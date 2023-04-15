// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

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
)

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// WhitelistingPaymasterMetaData contains all meta data concerning the WhitelistingPaymaster contract.
var WhitelistingPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sponsoredGas\",\"type\":\"uint256\"}],\"name\":\"AddressAddedToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COST_OF_POST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sponsoredGas\",\"type\":\"uint256\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"name\":\"setEntryPoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userDetails\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remainingGas\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WhitelistingPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use WhitelistingPaymasterMetaData.ABI instead.
var WhitelistingPaymasterABI = WhitelistingPaymasterMetaData.ABI

// WhitelistingPaymaster is an auto generated Go binding around an Ethereum contract.
type WhitelistingPaymaster struct {
	WhitelistingPaymasterCaller     // Read-only binding to the contract
	WhitelistingPaymasterTransactor // Write-only binding to the contract
	WhitelistingPaymasterFilterer   // Log filterer for contract events
}

// WhitelistingPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhitelistingPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistingPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhitelistingPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistingPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhitelistingPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhitelistingPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhitelistingPaymasterSession struct {
	Contract     *WhitelistingPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WhitelistingPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhitelistingPaymasterCallerSession struct {
	Contract *WhitelistingPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// WhitelistingPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhitelistingPaymasterTransactorSession struct {
	Contract     *WhitelistingPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// WhitelistingPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhitelistingPaymasterRaw struct {
	Contract *WhitelistingPaymaster // Generic contract binding to access the raw methods on
}

// WhitelistingPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhitelistingPaymasterCallerRaw struct {
	Contract *WhitelistingPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// WhitelistingPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhitelistingPaymasterTransactorRaw struct {
	Contract *WhitelistingPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhitelistingPaymaster creates a new instance of WhitelistingPaymaster, bound to a specific deployed contract.
func NewWhitelistingPaymaster(address common.Address, backend bind.ContractBackend) (*WhitelistingPaymaster, error) {
	contract, err := bindWhitelistingPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WhitelistingPaymaster{WhitelistingPaymasterCaller: WhitelistingPaymasterCaller{contract: contract}, WhitelistingPaymasterTransactor: WhitelistingPaymasterTransactor{contract: contract}, WhitelistingPaymasterFilterer: WhitelistingPaymasterFilterer{contract: contract}}, nil
}

// NewWhitelistingPaymasterCaller creates a new read-only instance of WhitelistingPaymaster, bound to a specific deployed contract.
func NewWhitelistingPaymasterCaller(address common.Address, caller bind.ContractCaller) (*WhitelistingPaymasterCaller, error) {
	contract, err := bindWhitelistingPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistingPaymasterCaller{contract: contract}, nil
}

// NewWhitelistingPaymasterTransactor creates a new write-only instance of WhitelistingPaymaster, bound to a specific deployed contract.
func NewWhitelistingPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*WhitelistingPaymasterTransactor, error) {
	contract, err := bindWhitelistingPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhitelistingPaymasterTransactor{contract: contract}, nil
}

// NewWhitelistingPaymasterFilterer creates a new log filterer instance of WhitelistingPaymaster, bound to a specific deployed contract.
func NewWhitelistingPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*WhitelistingPaymasterFilterer, error) {
	contract, err := bindWhitelistingPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhitelistingPaymasterFilterer{contract: contract}, nil
}

// bindWhitelistingPaymaster binds a generic wrapper to an already deployed contract.
func bindWhitelistingPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WhitelistingPaymasterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistingPaymaster *WhitelistingPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhitelistingPaymaster.Contract.WhitelistingPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistingPaymaster *WhitelistingPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.WhitelistingPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistingPaymaster *WhitelistingPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.WhitelistingPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WhitelistingPaymaster *WhitelistingPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WhitelistingPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.contract.Transact(opts, method, params...)
}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_WhitelistingPaymaster *WhitelistingPaymasterCaller) COSTOFPOST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WhitelistingPaymaster.contract.Call(opts, &out, "COST_OF_POST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) COSTOFPOST() (*big.Int, error) {
	return _WhitelistingPaymaster.Contract.COSTOFPOST(&_WhitelistingPaymaster.CallOpts)
}

// COSTOFPOST is a free data retrieval call binding the contract method 0x796d4371.
//
// Solidity: function COST_OF_POST() view returns(uint256)
func (_WhitelistingPaymaster *WhitelistingPaymasterCallerSession) COSTOFPOST() (*big.Int, error) {
	return _WhitelistingPaymaster.Contract.COSTOFPOST(&_WhitelistingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_WhitelistingPaymaster *WhitelistingPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhitelistingPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) EntryPoint() (common.Address, error) {
	return _WhitelistingPaymaster.Contract.EntryPoint(&_WhitelistingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_WhitelistingPaymaster *WhitelistingPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _WhitelistingPaymaster.Contract.EntryPoint(&_WhitelistingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_WhitelistingPaymaster *WhitelistingPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WhitelistingPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) GetDeposit() (*big.Int, error) {
	return _WhitelistingPaymaster.Contract.GetDeposit(&_WhitelistingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_WhitelistingPaymaster *WhitelistingPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _WhitelistingPaymaster.Contract.GetDeposit(&_WhitelistingPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WhitelistingPaymaster *WhitelistingPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WhitelistingPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) Owner() (common.Address, error) {
	return _WhitelistingPaymaster.Contract.Owner(&_WhitelistingPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WhitelistingPaymaster *WhitelistingPaymasterCallerSession) Owner() (common.Address, error) {
	return _WhitelistingPaymaster.Contract.Owner(&_WhitelistingPaymaster.CallOpts)
}

// UserDetails is a free data retrieval call binding the contract method 0x48dec2a7.
//
// Solidity: function userDetails(address ) view returns(bool isWhitelisted, uint256 remainingGas)
func (_WhitelistingPaymaster *WhitelistingPaymasterCaller) UserDetails(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsWhitelisted bool
	RemainingGas  *big.Int
}, error) {
	var out []interface{}
	err := _WhitelistingPaymaster.contract.Call(opts, &out, "userDetails", arg0)

	outstruct := new(struct {
		IsWhitelisted bool
		RemainingGas  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsWhitelisted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RemainingGas = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserDetails is a free data retrieval call binding the contract method 0x48dec2a7.
//
// Solidity: function userDetails(address ) view returns(bool isWhitelisted, uint256 remainingGas)
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) UserDetails(arg0 common.Address) (struct {
	IsWhitelisted bool
	RemainingGas  *big.Int
}, error) {
	return _WhitelistingPaymaster.Contract.UserDetails(&_WhitelistingPaymaster.CallOpts, arg0)
}

// UserDetails is a free data retrieval call binding the contract method 0x48dec2a7.
//
// Solidity: function userDetails(address ) view returns(bool isWhitelisted, uint256 remainingGas)
func (_WhitelistingPaymaster *WhitelistingPaymasterCallerSession) UserDetails(arg0 common.Address) (struct {
	IsWhitelisted bool
	RemainingGas  *big.Int
}, error) {
	return _WhitelistingPaymaster.Contract.UserDetails(&_WhitelistingPaymaster.CallOpts, arg0)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 maxCost) view returns(bytes context, uint256 deadline)
func (_WhitelistingPaymaster *WhitelistingPaymasterCaller) ValidatePaymasterUserOp(opts *bind.CallOpts, userOp UserOperation, arg1 [32]byte, maxCost *big.Int) (struct {
	Context  []byte
	Deadline *big.Int
}, error) {
	var out []interface{}
	err := _WhitelistingPaymaster.contract.Call(opts, &out, "validatePaymasterUserOp", userOp, arg1, maxCost)

	outstruct := new(struct {
		Context  []byte
		Deadline *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Context = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Deadline = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 maxCost) view returns(bytes context, uint256 deadline)
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) ValidatePaymasterUserOp(userOp UserOperation, arg1 [32]byte, maxCost *big.Int) (struct {
	Context  []byte
	Deadline *big.Int
}, error) {
	return _WhitelistingPaymaster.Contract.ValidatePaymasterUserOp(&_WhitelistingPaymaster.CallOpts, userOp, arg1, maxCost)
}

// ValidatePaymasterUserOp is a free data retrieval call binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 , uint256 maxCost) view returns(bytes context, uint256 deadline)
func (_WhitelistingPaymaster *WhitelistingPaymasterCallerSession) ValidatePaymasterUserOp(userOp UserOperation, arg1 [32]byte, maxCost *big.Int) (struct {
	Context  []byte
	Deadline *big.Int
}, error) {
	return _WhitelistingPaymaster.Contract.ValidatePaymasterUserOp(&_WhitelistingPaymaster.CallOpts, userOp, arg1, maxCost)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.AddStake(&_WhitelistingPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.AddStake(&_WhitelistingPaymaster.TransactOpts, unstakeDelaySec)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x214405fc.
//
// Solidity: function addToWhitelist(address addr, uint256 sponsoredGas) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) AddToWhitelist(opts *bind.TransactOpts, addr common.Address, sponsoredGas *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "addToWhitelist", addr, sponsoredGas)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x214405fc.
//
// Solidity: function addToWhitelist(address addr, uint256 sponsoredGas) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) AddToWhitelist(addr common.Address, sponsoredGas *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.AddToWhitelist(&_WhitelistingPaymaster.TransactOpts, addr, sponsoredGas)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x214405fc.
//
// Solidity: function addToWhitelist(address addr, uint256 sponsoredGas) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) AddToWhitelist(addr common.Address, sponsoredGas *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.AddToWhitelist(&_WhitelistingPaymaster.TransactOpts, addr, sponsoredGas)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) Deposit() (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.Deposit(&_WhitelistingPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.Deposit(&_WhitelistingPaymaster.TransactOpts)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.PostOp(&_WhitelistingPaymaster.TransactOpts, mode, context, actualGasCost)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.PostOp(&_WhitelistingPaymaster.TransactOpts, mode, context, actualGasCost)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.RenounceOwnership(&_WhitelistingPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.RenounceOwnership(&_WhitelistingPaymaster.TransactOpts)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) SetEntryPoint(opts *bind.TransactOpts, _entryPoint common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "setEntryPoint", _entryPoint)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) SetEntryPoint(_entryPoint common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.SetEntryPoint(&_WhitelistingPaymaster.TransactOpts, _entryPoint)
}

// SetEntryPoint is a paid mutator transaction binding the contract method 0x584465f2.
//
// Solidity: function setEntryPoint(address _entryPoint) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) SetEntryPoint(_entryPoint common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.SetEntryPoint(&_WhitelistingPaymaster.TransactOpts, _entryPoint)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.TransferOwnership(&_WhitelistingPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.TransferOwnership(&_WhitelistingPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.UnlockStake(&_WhitelistingPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.UnlockStake(&_WhitelistingPaymaster.TransactOpts)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.WithdrawStake(&_WhitelistingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.WithdrawStake(&_WhitelistingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.WithdrawTo(&_WhitelistingPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_WhitelistingPaymaster *WhitelistingPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WhitelistingPaymaster.Contract.WithdrawTo(&_WhitelistingPaymaster.TransactOpts, withdrawAddress, amount)
}

// WhitelistingPaymasterAddressAddedToWhitelistIterator is returned from FilterAddressAddedToWhitelist and is used to iterate over the raw logs and unpacked data for AddressAddedToWhitelist events raised by the WhitelistingPaymaster contract.
type WhitelistingPaymasterAddressAddedToWhitelistIterator struct {
	Event *WhitelistingPaymasterAddressAddedToWhitelist // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistingPaymasterAddressAddedToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistingPaymasterAddressAddedToWhitelist)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistingPaymasterAddressAddedToWhitelist)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistingPaymasterAddressAddedToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistingPaymasterAddressAddedToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistingPaymasterAddressAddedToWhitelist represents a AddressAddedToWhitelist event raised by the WhitelistingPaymaster contract.
type WhitelistingPaymasterAddressAddedToWhitelist struct {
	Addr         common.Address
	SponsoredGas *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddressAddedToWhitelist is a free log retrieval operation binding the contract event 0x3c144a3353f1bf91f2a96f3dc0c10980220384694d8304a8c68f2ba299eb0f8c.
//
// Solidity: event AddressAddedToWhitelist(address indexed addr, uint256 sponsoredGas)
func (_WhitelistingPaymaster *WhitelistingPaymasterFilterer) FilterAddressAddedToWhitelist(opts *bind.FilterOpts, addr []common.Address) (*WhitelistingPaymasterAddressAddedToWhitelistIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _WhitelistingPaymaster.contract.FilterLogs(opts, "AddressAddedToWhitelist", addrRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistingPaymasterAddressAddedToWhitelistIterator{contract: _WhitelistingPaymaster.contract, event: "AddressAddedToWhitelist", logs: logs, sub: sub}, nil
}

// WatchAddressAddedToWhitelist is a free log subscription operation binding the contract event 0x3c144a3353f1bf91f2a96f3dc0c10980220384694d8304a8c68f2ba299eb0f8c.
//
// Solidity: event AddressAddedToWhitelist(address indexed addr, uint256 sponsoredGas)
func (_WhitelistingPaymaster *WhitelistingPaymasterFilterer) WatchAddressAddedToWhitelist(opts *bind.WatchOpts, sink chan<- *WhitelistingPaymasterAddressAddedToWhitelist, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _WhitelistingPaymaster.contract.WatchLogs(opts, "AddressAddedToWhitelist", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistingPaymasterAddressAddedToWhitelist)
				if err := _WhitelistingPaymaster.contract.UnpackLog(event, "AddressAddedToWhitelist", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressAddedToWhitelist is a log parse operation binding the contract event 0x3c144a3353f1bf91f2a96f3dc0c10980220384694d8304a8c68f2ba299eb0f8c.
//
// Solidity: event AddressAddedToWhitelist(address indexed addr, uint256 sponsoredGas)
func (_WhitelistingPaymaster *WhitelistingPaymasterFilterer) ParseAddressAddedToWhitelist(log types.Log) (*WhitelistingPaymasterAddressAddedToWhitelist, error) {
	event := new(WhitelistingPaymasterAddressAddedToWhitelist)
	if err := _WhitelistingPaymaster.contract.UnpackLog(event, "AddressAddedToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhitelistingPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WhitelistingPaymaster contract.
type WhitelistingPaymasterOwnershipTransferredIterator struct {
	Event *WhitelistingPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WhitelistingPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhitelistingPaymasterOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WhitelistingPaymasterOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WhitelistingPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhitelistingPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhitelistingPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the WhitelistingPaymaster contract.
type WhitelistingPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WhitelistingPaymaster *WhitelistingPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WhitelistingPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WhitelistingPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WhitelistingPaymasterOwnershipTransferredIterator{contract: _WhitelistingPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WhitelistingPaymaster *WhitelistingPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WhitelistingPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WhitelistingPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhitelistingPaymasterOwnershipTransferred)
				if err := _WhitelistingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WhitelistingPaymaster *WhitelistingPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*WhitelistingPaymasterOwnershipTransferred, error) {
	event := new(WhitelistingPaymasterOwnershipTransferred)
	if err := _WhitelistingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
