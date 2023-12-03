package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("sepolia_client_wss"))
	if err != nil {
		fmt.Println("Wss connection: ", err)
		return
	}
	query := ethereum.FilterQuery{Addresses: []common.Address{common.HexToAddress("0x5D6AC6a7604bf5318DF2A31590B2dD650cebB68e")}}
	logs := make(chan types.Log)
	subs, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println("Filter log subs: ", err)
		return
	}
	for {
		select {
		case err := <-subs.Err():
			fmt.Println("Log subs: ", err)
		case vLog := <-logs:
			fmt.Println(vLog)
			abiFile, err := ioutil.ReadFile("/Users/dogukangundogan/techcareer_blockchain/backend/indexer/Indexer.abi")
			if err != nil {
				fmt.Println("Abi file read: ", err)
				return
			}
			contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
			if err != nil {
				fmt.Println("Abi file read: ", err)
				return
			}
			if vLog.Topics[0] == common.HexToHash("0x0d7fccda06d6eb51c23cbd16d49b9b3f3ebafb002dba1b074895cbb35d0e8130") { //Message event
				var messageEvent MessageEvent
				err := contractAbi.UnpackIntoInterface(&messageEvent, "MessageSent", vLog.Data)
				if err != nil {
					fmt.Println("Abi event unpack (message-event): ", err)
					return
				}
				fmt.Println(messageEvent)
			} else if vLog.Topics[0] == common.HexToHash("0xce9d3586d309892c73f4c3e05379ca39e09ec668a1f22820270142552bb560dd") { //Transfer event
				var transferEvent TransferEvent
				err := contractAbi.UnpackIntoInterface(&transferEvent, "AssetTransfer", vLog.Data)
				if err != nil {
					fmt.Println("Abi event unpack (transfer-event): ", err)
					return
				}
				fmt.Println(transferEvent)
			}
		}
	}
}
