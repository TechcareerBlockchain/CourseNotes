// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mathContract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MathContractMetaData contains all meta data concerning the MathContract contract.
var MathContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"}],\"name\":\"addition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506101a58061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c806354f363a31461002d575b5f80fd5b610047600480360381019061004291906100a9565b61005d565b60405161005491906100f6565b60405180910390f35b5f818361006a919061013c565b905092915050565b5f80fd5b5f819050919050565b61008881610076565b8114610092575f80fd5b50565b5f813590506100a38161007f565b92915050565b5f80604083850312156100bf576100be610072565b5b5f6100cc85828601610095565b92505060206100dd85828601610095565b9150509250929050565b6100f081610076565b82525050565b5f6020820190506101095f8301846100e7565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61014682610076565b915061015183610076565b92508282019050808211156101695761016861010f565b5b9291505056fea2646970667358221220f83e4ac40cc6d72ef7209deff13a98f6e558c135c9e0ee15227d246884452d9864736f6c63430008150033",
}

// MathContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MathContractMetaData.ABI instead.
var MathContractABI = MathContractMetaData.ABI

// MathContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathContractMetaData.Bin instead.
var MathContractBin = MathContractMetaData.Bin

// DeployMathContract deploys a new Ethereum contract, binding an instance of MathContract to it.
func DeployMathContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MathContract, error) {
	parsed, err := MathContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MathContract{MathContractCaller: MathContractCaller{contract: contract}, MathContractTransactor: MathContractTransactor{contract: contract}, MathContractFilterer: MathContractFilterer{contract: contract}}, nil
}

// MathContract is an auto generated Go binding around an Ethereum contract.
type MathContract struct {
	MathContractCaller     // Read-only binding to the contract
	MathContractTransactor // Write-only binding to the contract
	MathContractFilterer   // Log filterer for contract events
}

// MathContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathContractSession struct {
	Contract     *MathContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathContractCallerSession struct {
	Contract *MathContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MathContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathContractTransactorSession struct {
	Contract     *MathContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MathContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathContractRaw struct {
	Contract *MathContract // Generic contract binding to access the raw methods on
}

// MathContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathContractCallerRaw struct {
	Contract *MathContractCaller // Generic read-only contract binding to access the raw methods on
}

// MathContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathContractTransactorRaw struct {
	Contract *MathContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMathContract creates a new instance of MathContract, bound to a specific deployed contract.
func NewMathContract(address common.Address, backend bind.ContractBackend) (*MathContract, error) {
	contract, err := bindMathContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MathContract{MathContractCaller: MathContractCaller{contract: contract}, MathContractTransactor: MathContractTransactor{contract: contract}, MathContractFilterer: MathContractFilterer{contract: contract}}, nil
}

// NewMathContractCaller creates a new read-only instance of MathContract, bound to a specific deployed contract.
func NewMathContractCaller(address common.Address, caller bind.ContractCaller) (*MathContractCaller, error) {
	contract, err := bindMathContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathContractCaller{contract: contract}, nil
}

// NewMathContractTransactor creates a new write-only instance of MathContract, bound to a specific deployed contract.
func NewMathContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MathContractTransactor, error) {
	contract, err := bindMathContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathContractTransactor{contract: contract}, nil
}

// NewMathContractFilterer creates a new log filterer instance of MathContract, bound to a specific deployed contract.
func NewMathContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MathContractFilterer, error) {
	contract, err := bindMathContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathContractFilterer{contract: contract}, nil
}

// bindMathContract binds a generic wrapper to an already deployed contract.
func bindMathContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MathContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathContract *MathContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathContract.Contract.MathContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathContract *MathContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathContract.Contract.MathContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathContract *MathContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathContract.Contract.MathContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathContract *MathContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathContract *MathContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathContract *MathContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathContract.Contract.contract.Transact(opts, method, params...)
}

// Addition is a free data retrieval call binding the contract method 0x54f363a3.
//
// Solidity: function addition(uint256 v1, uint256 v2) pure returns(uint256)
func (_MathContract *MathContractCaller) Addition(opts *bind.CallOpts, v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MathContract.contract.Call(opts, &out, "addition", v1, v2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Addition is a free data retrieval call binding the contract method 0x54f363a3.
//
// Solidity: function addition(uint256 v1, uint256 v2) pure returns(uint256)
func (_MathContract *MathContractSession) Addition(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _MathContract.Contract.Addition(&_MathContract.CallOpts, v1, v2)
}

// Addition is a free data retrieval call binding the contract method 0x54f363a3.
//
// Solidity: function addition(uint256 v1, uint256 v2) pure returns(uint256)
func (_MathContract *MathContractCallerSession) Addition(v1 *big.Int, v2 *big.Int) (*big.Int, error) {
	return _MathContract.Contract.Addition(&_MathContract.CallOpts, v1, v2)
}
