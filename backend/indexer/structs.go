package main

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type TransferEvent struct {
	Sender common.Address `json:"sender"`
	Value  *big.Int       `json:"value"`
}

type MessageEvent struct {
	Sender  common.Address `json:"sender"`
	Message string         `json:"value"`
}
