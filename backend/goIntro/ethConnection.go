package main

import (
	"backend/contractcaller"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
)

func main() {

	client, err := ethclient.Dial(os.Getenv("sepolia_client_http"))
	if err != nil {
		fmt.Println(err)
		return
	}
	balance, err := getBalance(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(balance)
	err, contractInstance, _, _, _ := contractcaller.CreateFunctionRequirementsForMathContract(client, "0x2297aD0b4444Cde97345b689648Ab9dB1b52111c", os.Getenv("private_key"))
	responseFromSmartContract, err := contractInstance.Addition(&bind.CallOpts{}, big.NewInt(15), big.NewInt(25))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(responseFromSmartContract)
}

func getBalance(client *ethclient.Client) (float64, error) {
	walletAddress := common.HexToAddress("0x6097f22127E2EF98C2cF31335Bede16D742f6890") //common.Address
	balance, err := client.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	balanceAsFloat, _ := balance.Float64()
	return balanceAsFloat * 1e19, nil
}
