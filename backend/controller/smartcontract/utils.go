package controllerSmartContract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getBalance(client *ethclient.Client, walletAddressAsString string) (float64, error) {
	walletAddress := common.HexToAddress(walletAddressAsString) //common.Address
	balance, err := client.BalanceAt(context.Background(), walletAddress, nil)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	balanceAsFloat, _ := balance.Float64()
	return balanceAsFloat * 1e19, nil
}
