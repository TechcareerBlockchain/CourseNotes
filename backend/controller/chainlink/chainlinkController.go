package chainlinkController

import (
	contracts "backend/contracts/aggregatorV3"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math"
	"math/big"
	"net/http"
	"os"
)

func GetPriceOf() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody ChainlinkPriceRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			fmt.Println("Request Body: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		client, err := ethclient.Dial(os.Getenv("eth_http"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		//token address => erc20 contract
		//priceFeedAddress => chainlik aggregator address
		erc20Instance, err := contracts.NewErc20Caller(common.HexToAddress(requestBody.TokenAddress), client)
		if err != nil {
			fmt.Println("Erc20 contract: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		decimals, err := erc20Instance.Decimals(nil)
		if err != nil {
			fmt.Println("Erc20 contract: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		fmt.Println("Decimal :", decimals)
		priceFeedInstance, err := contracts.NewAggregatorV3InterfaceCaller(common.HexToAddress(requestBody.PriceFeedAddress), client)
		if err != nil {
			fmt.Println("Aggregator contract: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		latestRoundData, err := priceFeedInstance.LatestRoundData(nil)
		fmt.Println("Latest round data :", latestRoundData)
		decimalsFromChainlinkOracle, err := priceFeedInstance.Decimals(nil)
		fmt.Println("Chainlink Price oracle Decimals data :", decimalsFromChainlinkOracle)
		oracleAnswer := big.NewFloat(0).SetInt(latestRoundData.Answer)
		divisor := big.NewFloat(math.Pow(10, float64(decimalsFromChainlinkOracle))) //18 => 10^18
		price, _ := big.NewFloat(0).Quo(oracleAnswer, divisor).Float64()            // 17571 / 10^18 => 17571 * 10^-18
		c.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": price})
	}
}
