package controllerSmartContract

import (
	"backend/contractcaller"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"os"
	"strconv"
)

func GetBalanceHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		client, err := ethclient.Dial(os.Getenv("sepolia_client_http"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		balance, err := getBalance(client, c.Param("walletAddress"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		//c.Request.Header.Get("walletAddress") - c.GetHeader("walletAddress")
		fmt.Println(balance)
		c.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": balance})
	}
}

func PostAdditionFunctionHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody AdditionRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			fmt.Println("Request Body: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		client, err := ethclient.Dial(os.Getenv("sepolia_client_http"))
		if err != nil {
			fmt.Println("Client error: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		} //0x2297aD0b4444Cde97345b689648Ab9dB1b52111c
		err, contractInstance, _, _, _ := contractcaller.CreateFunctionRequirementsForMathContract(client, c.Param("smartContractAddress"), os.Getenv("private_key"))
		variable1, err := strconv.Atoi(requestBody.V1)
		if err != nil {
			fmt.Println("Number1 conv: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		variable2, err := strconv.Atoi(requestBody.V2)
		if err != nil {
			fmt.Println("Number2 conv: ", err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		responseFromSmartContract, err := contractInstance.Addition(&bind.CallOpts{}, big.NewInt(int64(variable1)), big.NewInt(int64(variable2)))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"status": false, "data": err})
			return
		}
		fmt.Println(responseFromSmartContract)
		c.JSON(http.StatusOK, map[string]interface{}{"status": true, "data": responseFromSmartContract})
	}
}
