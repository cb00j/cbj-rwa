// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rwa

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

// CBJTokenMetaData contains all meta data concerning the CBJToken contract.
var CBJTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONFIGURE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MULTIPLIER_UPDATE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cbjCompliance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICBJCompliance\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cbjTokenPauseManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICBJTokenPauseManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharesByUnderlyingAmount\",\"inputs\":[{\"name\":\"underlyingAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnderlyingAmountByShares\",\"inputs\":[{\"name\":\"sharesAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_cbjCompliance\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenPauseManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"multiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"multiplierNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCompliance\",\"inputs\":[{\"name\":\"_cbjCompliance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setName\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSymbol\",\"inputs\":[{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenPauseManager\",\"inputs\":[{\"name\":\"_tokenPauseManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sharesOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferShares\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sharesAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMultiplier\",\"inputs\":[{\"name\":\"newMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ComplianceSet\",\"inputs\":[{\"name\":\"oldCompliance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newCompliance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MultiplierUpdated\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NameChanged\",\"inputs\":[{\"name\":\"oldName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SymbolChanged\",\"inputs\":[{\"name\":\"oldSymbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"newSymbol\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenPauseManagerSet\",\"inputs\":[{\"name\":\"oldTokenPauseManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newTokenPauseManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferShares\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BurnAmountExceedsBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BurnFromCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ComplianceCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintToCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenPauseManagerCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferAmountExceedsBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFromCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferToCannotBeZero\",\"inputs\":[]}]",
}

// CBJTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use CBJTokenMetaData.ABI instead.
var CBJTokenABI = CBJTokenMetaData.ABI

// CBJToken is an auto generated Go binding around an Ethereum contract.
type CBJToken struct {
	CBJTokenCaller     // Read-only binding to the contract
	CBJTokenTransactor // Write-only binding to the contract
	CBJTokenFilterer   // Log filterer for contract events
}

// CBJTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CBJTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CBJTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CBJTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CBJTokenSession struct {
	Contract     *CBJToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CBJTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CBJTokenCallerSession struct {
	Contract *CBJTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CBJTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CBJTokenTransactorSession struct {
	Contract     *CBJTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CBJTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CBJTokenRaw struct {
	Contract *CBJToken // Generic contract binding to access the raw methods on
}

// CBJTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CBJTokenCallerRaw struct {
	Contract *CBJTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CBJTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CBJTokenTransactorRaw struct {
	Contract *CBJTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCBJToken creates a new instance of CBJToken, bound to a specific deployed contract.
func NewCBJToken(address common.Address, backend bind.ContractBackend) (*CBJToken, error) {
	contract, err := bindCBJToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CBJToken{CBJTokenCaller: CBJTokenCaller{contract: contract}, CBJTokenTransactor: CBJTokenTransactor{contract: contract}, CBJTokenFilterer: CBJTokenFilterer{contract: contract}}, nil
}

// NewCBJTokenCaller creates a new read-only instance of CBJToken, bound to a specific deployed contract.
func NewCBJTokenCaller(address common.Address, caller bind.ContractCaller) (*CBJTokenCaller, error) {
	contract, err := bindCBJToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CBJTokenCaller{contract: contract}, nil
}

// NewCBJTokenTransactor creates a new write-only instance of CBJToken, bound to a specific deployed contract.
func NewCBJTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CBJTokenTransactor, error) {
	contract, err := bindCBJToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CBJTokenTransactor{contract: contract}, nil
}

// NewCBJTokenFilterer creates a new log filterer instance of CBJToken, bound to a specific deployed contract.
func NewCBJTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CBJTokenFilterer, error) {
	contract, err := bindCBJToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CBJTokenFilterer{contract: contract}, nil
}

// bindCBJToken binds a generic wrapper to an already deployed contract.
func bindCBJToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CBJTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBJToken *CBJTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CBJToken.Contract.CBJTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBJToken *CBJTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJToken.Contract.CBJTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBJToken *CBJTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBJToken.Contract.CBJTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBJToken *CBJTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CBJToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBJToken *CBJTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBJToken *CBJTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBJToken.Contract.contract.Transact(opts, method, params...)
}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCaller) BURNROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "BURN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenSession) BURNROLE() ([32]byte, error) {
	return _CBJToken.Contract.BURNROLE(&_CBJToken.CallOpts)
}

// BURNROLE is a free data retrieval call binding the contract method 0xb930908f.
//
// Solidity: function BURN_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCallerSession) BURNROLE() ([32]byte, error) {
	return _CBJToken.Contract.BURNROLE(&_CBJToken.CallOpts)
}

// CONFIGUREROLE is a free data retrieval call binding the contract method 0x2b41ceb5.
//
// Solidity: function CONFIGURE_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCaller) CONFIGUREROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "CONFIGURE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGUREROLE is a free data retrieval call binding the contract method 0x2b41ceb5.
//
// Solidity: function CONFIGURE_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenSession) CONFIGUREROLE() ([32]byte, error) {
	return _CBJToken.Contract.CONFIGUREROLE(&_CBJToken.CallOpts)
}

// CONFIGUREROLE is a free data retrieval call binding the contract method 0x2b41ceb5.
//
// Solidity: function CONFIGURE_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCallerSession) CONFIGUREROLE() ([32]byte, error) {
	return _CBJToken.Contract.CONFIGUREROLE(&_CBJToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CBJToken.Contract.DEFAULTADMINROLE(&_CBJToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CBJToken.Contract.DEFAULTADMINROLE(&_CBJToken.CallOpts)
}

// MINTROLE is a free data retrieval call binding the contract method 0xe9a9c850.
//
// Solidity: function MINT_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCaller) MINTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "MINT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTROLE is a free data retrieval call binding the contract method 0xe9a9c850.
//
// Solidity: function MINT_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenSession) MINTROLE() ([32]byte, error) {
	return _CBJToken.Contract.MINTROLE(&_CBJToken.CallOpts)
}

// MINTROLE is a free data retrieval call binding the contract method 0xe9a9c850.
//
// Solidity: function MINT_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCallerSession) MINTROLE() ([32]byte, error) {
	return _CBJToken.Contract.MINTROLE(&_CBJToken.CallOpts)
}

// MULTIPLIERUPDATEROLE is a free data retrieval call binding the contract method 0x8ec95900.
//
// Solidity: function MULTIPLIER_UPDATE_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCaller) MULTIPLIERUPDATEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "MULTIPLIER_UPDATE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MULTIPLIERUPDATEROLE is a free data retrieval call binding the contract method 0x8ec95900.
//
// Solidity: function MULTIPLIER_UPDATE_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenSession) MULTIPLIERUPDATEROLE() ([32]byte, error) {
	return _CBJToken.Contract.MULTIPLIERUPDATEROLE(&_CBJToken.CallOpts)
}

// MULTIPLIERUPDATEROLE is a free data retrieval call binding the contract method 0x8ec95900.
//
// Solidity: function MULTIPLIER_UPDATE_ROLE() view returns(bytes32)
func (_CBJToken *CBJTokenCallerSession) MULTIPLIERUPDATEROLE() ([32]byte, error) {
	return _CBJToken.Contract.MULTIPLIERUPDATEROLE(&_CBJToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CBJToken *CBJTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CBJToken *CBJTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CBJToken.Contract.Allowance(&_CBJToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CBJToken.Contract.Allowance(&_CBJToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CBJToken *CBJTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CBJToken *CBJTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CBJToken.Contract.BalanceOf(&_CBJToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CBJToken.Contract.BalanceOf(&_CBJToken.CallOpts, account)
}

// CbjCompliance is a free data retrieval call binding the contract method 0x253960a9.
//
// Solidity: function cbjCompliance() view returns(address)
func (_CBJToken *CBJTokenCaller) CbjCompliance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "cbjCompliance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CbjCompliance is a free data retrieval call binding the contract method 0x253960a9.
//
// Solidity: function cbjCompliance() view returns(address)
func (_CBJToken *CBJTokenSession) CbjCompliance() (common.Address, error) {
	return _CBJToken.Contract.CbjCompliance(&_CBJToken.CallOpts)
}

// CbjCompliance is a free data retrieval call binding the contract method 0x253960a9.
//
// Solidity: function cbjCompliance() view returns(address)
func (_CBJToken *CBJTokenCallerSession) CbjCompliance() (common.Address, error) {
	return _CBJToken.Contract.CbjCompliance(&_CBJToken.CallOpts)
}

// CbjTokenPauseManager is a free data retrieval call binding the contract method 0x64c0546b.
//
// Solidity: function cbjTokenPauseManager() view returns(address)
func (_CBJToken *CBJTokenCaller) CbjTokenPauseManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "cbjTokenPauseManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CbjTokenPauseManager is a free data retrieval call binding the contract method 0x64c0546b.
//
// Solidity: function cbjTokenPauseManager() view returns(address)
func (_CBJToken *CBJTokenSession) CbjTokenPauseManager() (common.Address, error) {
	return _CBJToken.Contract.CbjTokenPauseManager(&_CBJToken.CallOpts)
}

// CbjTokenPauseManager is a free data retrieval call binding the contract method 0x64c0546b.
//
// Solidity: function cbjTokenPauseManager() view returns(address)
func (_CBJToken *CBJTokenCallerSession) CbjTokenPauseManager() (common.Address, error) {
	return _CBJToken.Contract.CbjTokenPauseManager(&_CBJToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CBJToken *CBJTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CBJToken *CBJTokenSession) Decimals() (uint8, error) {
	return _CBJToken.Contract.Decimals(&_CBJToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CBJToken *CBJTokenCallerSession) Decimals() (uint8, error) {
	return _CBJToken.Contract.Decimals(&_CBJToken.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CBJToken *CBJTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CBJToken *CBJTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CBJToken.Contract.GetRoleAdmin(&_CBJToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CBJToken *CBJTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CBJToken.Contract.GetRoleAdmin(&_CBJToken.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CBJToken *CBJTokenCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CBJToken *CBJTokenSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CBJToken.Contract.GetRoleMember(&_CBJToken.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CBJToken *CBJTokenCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CBJToken.Contract.GetRoleMember(&_CBJToken.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CBJToken *CBJTokenCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CBJToken *CBJTokenSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CBJToken.Contract.GetRoleMemberCount(&_CBJToken.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CBJToken.Contract.GetRoleMemberCount(&_CBJToken.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CBJToken *CBJTokenCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CBJToken *CBJTokenSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CBJToken.Contract.GetRoleMembers(&_CBJToken.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CBJToken *CBJTokenCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CBJToken.Contract.GetRoleMembers(&_CBJToken.CallOpts, role)
}

// GetSharesByUnderlyingAmount is a free data retrieval call binding the contract method 0x44acb51b.
//
// Solidity: function getSharesByUnderlyingAmount(uint256 underlyingAmount) view returns(uint256)
func (_CBJToken *CBJTokenCaller) GetSharesByUnderlyingAmount(opts *bind.CallOpts, underlyingAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "getSharesByUnderlyingAmount", underlyingAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByUnderlyingAmount is a free data retrieval call binding the contract method 0x44acb51b.
//
// Solidity: function getSharesByUnderlyingAmount(uint256 underlyingAmount) view returns(uint256)
func (_CBJToken *CBJTokenSession) GetSharesByUnderlyingAmount(underlyingAmount *big.Int) (*big.Int, error) {
	return _CBJToken.Contract.GetSharesByUnderlyingAmount(&_CBJToken.CallOpts, underlyingAmount)
}

// GetSharesByUnderlyingAmount is a free data retrieval call binding the contract method 0x44acb51b.
//
// Solidity: function getSharesByUnderlyingAmount(uint256 underlyingAmount) view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) GetSharesByUnderlyingAmount(underlyingAmount *big.Int) (*big.Int, error) {
	return _CBJToken.Contract.GetSharesByUnderlyingAmount(&_CBJToken.CallOpts, underlyingAmount)
}

// GetUnderlyingAmountByShares is a free data retrieval call binding the contract method 0x944b511c.
//
// Solidity: function getUnderlyingAmountByShares(uint256 sharesAmount) view returns(uint256)
func (_CBJToken *CBJTokenCaller) GetUnderlyingAmountByShares(opts *bind.CallOpts, sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "getUnderlyingAmountByShares", sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingAmountByShares is a free data retrieval call binding the contract method 0x944b511c.
//
// Solidity: function getUnderlyingAmountByShares(uint256 sharesAmount) view returns(uint256)
func (_CBJToken *CBJTokenSession) GetUnderlyingAmountByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _CBJToken.Contract.GetUnderlyingAmountByShares(&_CBJToken.CallOpts, sharesAmount)
}

// GetUnderlyingAmountByShares is a free data retrieval call binding the contract method 0x944b511c.
//
// Solidity: function getUnderlyingAmountByShares(uint256 sharesAmount) view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) GetUnderlyingAmountByShares(sharesAmount *big.Int) (*big.Int, error) {
	return _CBJToken.Contract.GetUnderlyingAmountByShares(&_CBJToken.CallOpts, sharesAmount)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CBJToken *CBJTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CBJToken *CBJTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CBJToken.Contract.HasRole(&_CBJToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CBJToken *CBJTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CBJToken.Contract.HasRole(&_CBJToken.CallOpts, role, account)
}

// Multiplier is a free data retrieval call binding the contract method 0x1b3ed722.
//
// Solidity: function multiplier() view returns(uint256)
func (_CBJToken *CBJTokenCaller) Multiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "multiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiplier is a free data retrieval call binding the contract method 0x1b3ed722.
//
// Solidity: function multiplier() view returns(uint256)
func (_CBJToken *CBJTokenSession) Multiplier() (*big.Int, error) {
	return _CBJToken.Contract.Multiplier(&_CBJToken.CallOpts)
}

// Multiplier is a free data retrieval call binding the contract method 0x1b3ed722.
//
// Solidity: function multiplier() view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) Multiplier() (*big.Int, error) {
	return _CBJToken.Contract.Multiplier(&_CBJToken.CallOpts)
}

// MultiplierNonce is a free data retrieval call binding the contract method 0xf9e47896.
//
// Solidity: function multiplierNonce() view returns(uint256)
func (_CBJToken *CBJTokenCaller) MultiplierNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "multiplierNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierNonce is a free data retrieval call binding the contract method 0xf9e47896.
//
// Solidity: function multiplierNonce() view returns(uint256)
func (_CBJToken *CBJTokenSession) MultiplierNonce() (*big.Int, error) {
	return _CBJToken.Contract.MultiplierNonce(&_CBJToken.CallOpts)
}

// MultiplierNonce is a free data retrieval call binding the contract method 0xf9e47896.
//
// Solidity: function multiplierNonce() view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) MultiplierNonce() (*big.Int, error) {
	return _CBJToken.Contract.MultiplierNonce(&_CBJToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CBJToken *CBJTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CBJToken *CBJTokenSession) Name() (string, error) {
	return _CBJToken.Contract.Name(&_CBJToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CBJToken *CBJTokenCallerSession) Name() (string, error) {
	return _CBJToken.Contract.Name(&_CBJToken.CallOpts)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address account) view returns(uint256)
func (_CBJToken *CBJTokenCaller) SharesOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "sharesOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address account) view returns(uint256)
func (_CBJToken *CBJTokenSession) SharesOf(account common.Address) (*big.Int, error) {
	return _CBJToken.Contract.SharesOf(&_CBJToken.CallOpts, account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address account) view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) SharesOf(account common.Address) (*big.Int, error) {
	return _CBJToken.Contract.SharesOf(&_CBJToken.CallOpts, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CBJToken *CBJTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CBJToken *CBJTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CBJToken.Contract.SupportsInterface(&_CBJToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CBJToken *CBJTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CBJToken.Contract.SupportsInterface(&_CBJToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CBJToken *CBJTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CBJToken *CBJTokenSession) Symbol() (string, error) {
	return _CBJToken.Contract.Symbol(&_CBJToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CBJToken *CBJTokenCallerSession) Symbol() (string, error) {
	return _CBJToken.Contract.Symbol(&_CBJToken.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_CBJToken *CBJTokenCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_CBJToken *CBJTokenSession) TotalShares() (*big.Int, error) {
	return _CBJToken.Contract.TotalShares(&_CBJToken.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) TotalShares() (*big.Int, error) {
	return _CBJToken.Contract.TotalShares(&_CBJToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CBJToken *CBJTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CBJToken *CBJTokenSession) TotalSupply() (*big.Int, error) {
	return _CBJToken.Contract.TotalSupply(&_CBJToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CBJToken *CBJTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CBJToken.Contract.TotalSupply(&_CBJToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CBJToken *CBJTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CBJToken *CBJTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Approve(&_CBJToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_CBJToken *CBJTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Approve(&_CBJToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CBJToken *CBJTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CBJToken *CBJTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Burn(&_CBJToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_CBJToken *CBJTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Burn(&_CBJToken.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CBJToken *CBJTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CBJToken *CBJTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.GrantRole(&_CBJToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CBJToken *CBJTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.GrantRole(&_CBJToken.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string name_, string symbol_, address _cbjCompliance, address _tokenPauseManager) returns()
func (_CBJToken *CBJTokenTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, _cbjCompliance common.Address, _tokenPauseManager common.Address) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "initialize", name_, symbol_, _cbjCompliance, _tokenPauseManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string name_, string symbol_, address _cbjCompliance, address _tokenPauseManager) returns()
func (_CBJToken *CBJTokenSession) Initialize(name_ string, symbol_ string, _cbjCompliance common.Address, _tokenPauseManager common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.Initialize(&_CBJToken.TransactOpts, name_, symbol_, _cbjCompliance, _tokenPauseManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string name_, string symbol_, address _cbjCompliance, address _tokenPauseManager) returns()
func (_CBJToken *CBJTokenTransactorSession) Initialize(name_ string, symbol_ string, _cbjCompliance common.Address, _tokenPauseManager common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.Initialize(&_CBJToken.TransactOpts, name_, symbol_, _cbjCompliance, _tokenPauseManager)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CBJToken *CBJTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CBJToken *CBJTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Mint(&_CBJToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CBJToken *CBJTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Mint(&_CBJToken.TransactOpts, to, amount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CBJToken *CBJTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CBJToken *CBJTokenSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.RenounceRole(&_CBJToken.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CBJToken *CBJTokenTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.RenounceRole(&_CBJToken.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CBJToken *CBJTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CBJToken *CBJTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.RevokeRole(&_CBJToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CBJToken *CBJTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.RevokeRole(&_CBJToken.TransactOpts, role, account)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _cbjCompliance) returns()
func (_CBJToken *CBJTokenTransactor) SetCompliance(opts *bind.TransactOpts, _cbjCompliance common.Address) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "setCompliance", _cbjCompliance)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _cbjCompliance) returns()
func (_CBJToken *CBJTokenSession) SetCompliance(_cbjCompliance common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.SetCompliance(&_CBJToken.TransactOpts, _cbjCompliance)
}

// SetCompliance is a paid mutator transaction binding the contract method 0xf8981789.
//
// Solidity: function setCompliance(address _cbjCompliance) returns()
func (_CBJToken *CBJTokenTransactorSession) SetCompliance(_cbjCompliance common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.SetCompliance(&_CBJToken.TransactOpts, _cbjCompliance)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_CBJToken *CBJTokenTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_CBJToken *CBJTokenSession) SetName(name_ string) (*types.Transaction, error) {
	return _CBJToken.Contract.SetName(&_CBJToken.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name_) returns()
func (_CBJToken *CBJTokenTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _CBJToken.Contract.SetName(&_CBJToken.TransactOpts, name_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_CBJToken *CBJTokenTransactor) SetSymbol(opts *bind.TransactOpts, symbol_ string) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "setSymbol", symbol_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_CBJToken *CBJTokenSession) SetSymbol(symbol_ string) (*types.Transaction, error) {
	return _CBJToken.Contract.SetSymbol(&_CBJToken.TransactOpts, symbol_)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string symbol_) returns()
func (_CBJToken *CBJTokenTransactorSession) SetSymbol(symbol_ string) (*types.Transaction, error) {
	return _CBJToken.Contract.SetSymbol(&_CBJToken.TransactOpts, symbol_)
}

// SetTokenPauseManager is a paid mutator transaction binding the contract method 0x248609c6.
//
// Solidity: function setTokenPauseManager(address _tokenPauseManager) returns()
func (_CBJToken *CBJTokenTransactor) SetTokenPauseManager(opts *bind.TransactOpts, _tokenPauseManager common.Address) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "setTokenPauseManager", _tokenPauseManager)
}

// SetTokenPauseManager is a paid mutator transaction binding the contract method 0x248609c6.
//
// Solidity: function setTokenPauseManager(address _tokenPauseManager) returns()
func (_CBJToken *CBJTokenSession) SetTokenPauseManager(_tokenPauseManager common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.SetTokenPauseManager(&_CBJToken.TransactOpts, _tokenPauseManager)
}

// SetTokenPauseManager is a paid mutator transaction binding the contract method 0x248609c6.
//
// Solidity: function setTokenPauseManager(address _tokenPauseManager) returns()
func (_CBJToken *CBJTokenTransactorSession) SetTokenPauseManager(_tokenPauseManager common.Address) (*types.Transaction, error) {
	return _CBJToken.Contract.SetTokenPauseManager(&_CBJToken.TransactOpts, _tokenPauseManager)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CBJToken *CBJTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CBJToken *CBJTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Transfer(&_CBJToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_CBJToken *CBJTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.Transfer(&_CBJToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CBJToken *CBJTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CBJToken *CBJTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.TransferFrom(&_CBJToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_CBJToken *CBJTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.TransferFrom(&_CBJToken.TransactOpts, from, to, value)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address to, uint256 sharesAmount) returns(bool)
func (_CBJToken *CBJTokenTransactor) TransferShares(opts *bind.TransactOpts, to common.Address, sharesAmount *big.Int) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "transferShares", to, sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address to, uint256 sharesAmount) returns(bool)
func (_CBJToken *CBJTokenSession) TransferShares(to common.Address, sharesAmount *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.TransferShares(&_CBJToken.TransactOpts, to, sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address to, uint256 sharesAmount) returns(bool)
func (_CBJToken *CBJTokenTransactorSession) TransferShares(to common.Address, sharesAmount *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.TransferShares(&_CBJToken.TransactOpts, to, sharesAmount)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 newMultiplier) returns()
func (_CBJToken *CBJTokenTransactor) UpdateMultiplier(opts *bind.TransactOpts, newMultiplier *big.Int) (*types.Transaction, error) {
	return _CBJToken.contract.Transact(opts, "updateMultiplier", newMultiplier)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 newMultiplier) returns()
func (_CBJToken *CBJTokenSession) UpdateMultiplier(newMultiplier *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.UpdateMultiplier(&_CBJToken.TransactOpts, newMultiplier)
}

// UpdateMultiplier is a paid mutator transaction binding the contract method 0x5ffe6146.
//
// Solidity: function updateMultiplier(uint256 newMultiplier) returns()
func (_CBJToken *CBJTokenTransactorSession) UpdateMultiplier(newMultiplier *big.Int) (*types.Transaction, error) {
	return _CBJToken.Contract.UpdateMultiplier(&_CBJToken.TransactOpts, newMultiplier)
}

// CBJTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CBJToken contract.
type CBJTokenApprovalIterator struct {
	Event *CBJTokenApproval // Event containing the contract specifics and raw log

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
func (it *CBJTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenApproval)
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
		it.Event = new(CBJTokenApproval)
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
func (it *CBJTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenApproval represents a Approval event raised by the CBJToken contract.
type CBJTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CBJToken *CBJTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CBJTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenApprovalIterator{contract: _CBJToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CBJToken *CBJTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CBJTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenApproval)
				if err := _CBJToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CBJToken *CBJTokenFilterer) ParseApproval(log types.Log) (*CBJTokenApproval, error) {
	event := new(CBJTokenApproval)
	if err := _CBJToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenComplianceSetIterator is returned from FilterComplianceSet and is used to iterate over the raw logs and unpacked data for ComplianceSet events raised by the CBJToken contract.
type CBJTokenComplianceSetIterator struct {
	Event *CBJTokenComplianceSet // Event containing the contract specifics and raw log

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
func (it *CBJTokenComplianceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenComplianceSet)
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
		it.Event = new(CBJTokenComplianceSet)
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
func (it *CBJTokenComplianceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenComplianceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenComplianceSet represents a ComplianceSet event raised by the CBJToken contract.
type CBJTokenComplianceSet struct {
	OldCompliance common.Address
	NewCompliance common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterComplianceSet is a free log retrieval operation binding the contract event 0x12f44a0fc4b932332de2e987123a00d8e239e0c9785dc5c8d05f9a41ea7a42ef.
//
// Solidity: event ComplianceSet(address indexed oldCompliance, address indexed newCompliance)
func (_CBJToken *CBJTokenFilterer) FilterComplianceSet(opts *bind.FilterOpts, oldCompliance []common.Address, newCompliance []common.Address) (*CBJTokenComplianceSetIterator, error) {

	var oldComplianceRule []interface{}
	for _, oldComplianceItem := range oldCompliance {
		oldComplianceRule = append(oldComplianceRule, oldComplianceItem)
	}
	var newComplianceRule []interface{}
	for _, newComplianceItem := range newCompliance {
		newComplianceRule = append(newComplianceRule, newComplianceItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "ComplianceSet", oldComplianceRule, newComplianceRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenComplianceSetIterator{contract: _CBJToken.contract, event: "ComplianceSet", logs: logs, sub: sub}, nil
}

// WatchComplianceSet is a free log subscription operation binding the contract event 0x12f44a0fc4b932332de2e987123a00d8e239e0c9785dc5c8d05f9a41ea7a42ef.
//
// Solidity: event ComplianceSet(address indexed oldCompliance, address indexed newCompliance)
func (_CBJToken *CBJTokenFilterer) WatchComplianceSet(opts *bind.WatchOpts, sink chan<- *CBJTokenComplianceSet, oldCompliance []common.Address, newCompliance []common.Address) (event.Subscription, error) {

	var oldComplianceRule []interface{}
	for _, oldComplianceItem := range oldCompliance {
		oldComplianceRule = append(oldComplianceRule, oldComplianceItem)
	}
	var newComplianceRule []interface{}
	for _, newComplianceItem := range newCompliance {
		newComplianceRule = append(newComplianceRule, newComplianceItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "ComplianceSet", oldComplianceRule, newComplianceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenComplianceSet)
				if err := _CBJToken.contract.UnpackLog(event, "ComplianceSet", log); err != nil {
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

// ParseComplianceSet is a log parse operation binding the contract event 0x12f44a0fc4b932332de2e987123a00d8e239e0c9785dc5c8d05f9a41ea7a42ef.
//
// Solidity: event ComplianceSet(address indexed oldCompliance, address indexed newCompliance)
func (_CBJToken *CBJTokenFilterer) ParseComplianceSet(log types.Log) (*CBJTokenComplianceSet, error) {
	event := new(CBJTokenComplianceSet)
	if err := _CBJToken.contract.UnpackLog(event, "ComplianceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CBJToken contract.
type CBJTokenInitializedIterator struct {
	Event *CBJTokenInitialized // Event containing the contract specifics and raw log

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
func (it *CBJTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenInitialized)
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
		it.Event = new(CBJTokenInitialized)
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
func (it *CBJTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenInitialized represents a Initialized event raised by the CBJToken contract.
type CBJTokenInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CBJToken *CBJTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*CBJTokenInitializedIterator, error) {

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CBJTokenInitializedIterator{contract: _CBJToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CBJToken *CBJTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CBJTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenInitialized)
				if err := _CBJToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CBJToken *CBJTokenFilterer) ParseInitialized(log types.Log) (*CBJTokenInitialized, error) {
	event := new(CBJTokenInitialized)
	if err := _CBJToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenMultiplierUpdatedIterator is returned from FilterMultiplierUpdated and is used to iterate over the raw logs and unpacked data for MultiplierUpdated events raised by the CBJToken contract.
type CBJTokenMultiplierUpdatedIterator struct {
	Event *CBJTokenMultiplierUpdated // Event containing the contract specifics and raw log

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
func (it *CBJTokenMultiplierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenMultiplierUpdated)
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
		it.Event = new(CBJTokenMultiplierUpdated)
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
func (it *CBJTokenMultiplierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenMultiplierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenMultiplierUpdated represents a MultiplierUpdated event raised by the CBJToken contract.
type CBJTokenMultiplierUpdated struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMultiplierUpdated is a free log retrieval operation binding the contract event 0x4dbe4840d7465bd162f67814cea0b519567a2e0e578bcde61e7f4ced361e5a3d.
//
// Solidity: event MultiplierUpdated(uint256 value)
func (_CBJToken *CBJTokenFilterer) FilterMultiplierUpdated(opts *bind.FilterOpts) (*CBJTokenMultiplierUpdatedIterator, error) {

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "MultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return &CBJTokenMultiplierUpdatedIterator{contract: _CBJToken.contract, event: "MultiplierUpdated", logs: logs, sub: sub}, nil
}

// WatchMultiplierUpdated is a free log subscription operation binding the contract event 0x4dbe4840d7465bd162f67814cea0b519567a2e0e578bcde61e7f4ced361e5a3d.
//
// Solidity: event MultiplierUpdated(uint256 value)
func (_CBJToken *CBJTokenFilterer) WatchMultiplierUpdated(opts *bind.WatchOpts, sink chan<- *CBJTokenMultiplierUpdated) (event.Subscription, error) {

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "MultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenMultiplierUpdated)
				if err := _CBJToken.contract.UnpackLog(event, "MultiplierUpdated", log); err != nil {
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

// ParseMultiplierUpdated is a log parse operation binding the contract event 0x4dbe4840d7465bd162f67814cea0b519567a2e0e578bcde61e7f4ced361e5a3d.
//
// Solidity: event MultiplierUpdated(uint256 value)
func (_CBJToken *CBJTokenFilterer) ParseMultiplierUpdated(log types.Log) (*CBJTokenMultiplierUpdated, error) {
	event := new(CBJTokenMultiplierUpdated)
	if err := _CBJToken.contract.UnpackLog(event, "MultiplierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenNameChangedIterator is returned from FilterNameChanged and is used to iterate over the raw logs and unpacked data for NameChanged events raised by the CBJToken contract.
type CBJTokenNameChangedIterator struct {
	Event *CBJTokenNameChanged // Event containing the contract specifics and raw log

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
func (it *CBJTokenNameChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenNameChanged)
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
		it.Event = new(CBJTokenNameChanged)
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
func (it *CBJTokenNameChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenNameChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenNameChanged represents a NameChanged event raised by the CBJToken contract.
type CBJTokenNameChanged struct {
	OldName string
	NewName string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameChanged is a free log retrieval operation binding the contract event 0x6c20b91d1723b78732eba64ff11ebd7966a6e4af568a00fa4f6b72c20f58b02a.
//
// Solidity: event NameChanged(string oldName, string newName)
func (_CBJToken *CBJTokenFilterer) FilterNameChanged(opts *bind.FilterOpts) (*CBJTokenNameChangedIterator, error) {

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return &CBJTokenNameChangedIterator{contract: _CBJToken.contract, event: "NameChanged", logs: logs, sub: sub}, nil
}

// WatchNameChanged is a free log subscription operation binding the contract event 0x6c20b91d1723b78732eba64ff11ebd7966a6e4af568a00fa4f6b72c20f58b02a.
//
// Solidity: event NameChanged(string oldName, string newName)
func (_CBJToken *CBJTokenFilterer) WatchNameChanged(opts *bind.WatchOpts, sink chan<- *CBJTokenNameChanged) (event.Subscription, error) {

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "NameChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenNameChanged)
				if err := _CBJToken.contract.UnpackLog(event, "NameChanged", log); err != nil {
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

// ParseNameChanged is a log parse operation binding the contract event 0x6c20b91d1723b78732eba64ff11ebd7966a6e4af568a00fa4f6b72c20f58b02a.
//
// Solidity: event NameChanged(string oldName, string newName)
func (_CBJToken *CBJTokenFilterer) ParseNameChanged(log types.Log) (*CBJTokenNameChanged, error) {
	event := new(CBJTokenNameChanged)
	if err := _CBJToken.contract.UnpackLog(event, "NameChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CBJToken contract.
type CBJTokenRoleAdminChangedIterator struct {
	Event *CBJTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CBJTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenRoleAdminChanged)
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
		it.Event = new(CBJTokenRoleAdminChanged)
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
func (it *CBJTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenRoleAdminChanged represents a RoleAdminChanged event raised by the CBJToken contract.
type CBJTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CBJToken *CBJTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CBJTokenRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenRoleAdminChangedIterator{contract: _CBJToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CBJToken *CBJTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CBJTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenRoleAdminChanged)
				if err := _CBJToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CBJToken *CBJTokenFilterer) ParseRoleAdminChanged(log types.Log) (*CBJTokenRoleAdminChanged, error) {
	event := new(CBJTokenRoleAdminChanged)
	if err := _CBJToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CBJToken contract.
type CBJTokenRoleGrantedIterator struct {
	Event *CBJTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *CBJTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenRoleGranted)
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
		it.Event = new(CBJTokenRoleGranted)
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
func (it *CBJTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenRoleGranted represents a RoleGranted event raised by the CBJToken contract.
type CBJTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJToken *CBJTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CBJTokenRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenRoleGrantedIterator{contract: _CBJToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJToken *CBJTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CBJTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenRoleGranted)
				if err := _CBJToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJToken *CBJTokenFilterer) ParseRoleGranted(log types.Log) (*CBJTokenRoleGranted, error) {
	event := new(CBJTokenRoleGranted)
	if err := _CBJToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CBJToken contract.
type CBJTokenRoleRevokedIterator struct {
	Event *CBJTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CBJTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenRoleRevoked)
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
		it.Event = new(CBJTokenRoleRevoked)
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
func (it *CBJTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenRoleRevoked represents a RoleRevoked event raised by the CBJToken contract.
type CBJTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJToken *CBJTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CBJTokenRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenRoleRevokedIterator{contract: _CBJToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJToken *CBJTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CBJTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenRoleRevoked)
				if err := _CBJToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJToken *CBJTokenFilterer) ParseRoleRevoked(log types.Log) (*CBJTokenRoleRevoked, error) {
	event := new(CBJTokenRoleRevoked)
	if err := _CBJToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenSymbolChangedIterator is returned from FilterSymbolChanged and is used to iterate over the raw logs and unpacked data for SymbolChanged events raised by the CBJToken contract.
type CBJTokenSymbolChangedIterator struct {
	Event *CBJTokenSymbolChanged // Event containing the contract specifics and raw log

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
func (it *CBJTokenSymbolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenSymbolChanged)
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
		it.Event = new(CBJTokenSymbolChanged)
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
func (it *CBJTokenSymbolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenSymbolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenSymbolChanged represents a SymbolChanged event raised by the CBJToken contract.
type CBJTokenSymbolChanged struct {
	OldSymbol string
	NewSymbol string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSymbolChanged is a free log retrieval operation binding the contract event 0xd7ad744cc76ebad190995130eec8ba506b3605612d23b5b9cef8e27f14d138b4.
//
// Solidity: event SymbolChanged(string oldSymbol, string newSymbol)
func (_CBJToken *CBJTokenFilterer) FilterSymbolChanged(opts *bind.FilterOpts) (*CBJTokenSymbolChangedIterator, error) {

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "SymbolChanged")
	if err != nil {
		return nil, err
	}
	return &CBJTokenSymbolChangedIterator{contract: _CBJToken.contract, event: "SymbolChanged", logs: logs, sub: sub}, nil
}

// WatchSymbolChanged is a free log subscription operation binding the contract event 0xd7ad744cc76ebad190995130eec8ba506b3605612d23b5b9cef8e27f14d138b4.
//
// Solidity: event SymbolChanged(string oldSymbol, string newSymbol)
func (_CBJToken *CBJTokenFilterer) WatchSymbolChanged(opts *bind.WatchOpts, sink chan<- *CBJTokenSymbolChanged) (event.Subscription, error) {

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "SymbolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenSymbolChanged)
				if err := _CBJToken.contract.UnpackLog(event, "SymbolChanged", log); err != nil {
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

// ParseSymbolChanged is a log parse operation binding the contract event 0xd7ad744cc76ebad190995130eec8ba506b3605612d23b5b9cef8e27f14d138b4.
//
// Solidity: event SymbolChanged(string oldSymbol, string newSymbol)
func (_CBJToken *CBJTokenFilterer) ParseSymbolChanged(log types.Log) (*CBJTokenSymbolChanged, error) {
	event := new(CBJTokenSymbolChanged)
	if err := _CBJToken.contract.UnpackLog(event, "SymbolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenTokenPauseManagerSetIterator is returned from FilterTokenPauseManagerSet and is used to iterate over the raw logs and unpacked data for TokenPauseManagerSet events raised by the CBJToken contract.
type CBJTokenTokenPauseManagerSetIterator struct {
	Event *CBJTokenTokenPauseManagerSet // Event containing the contract specifics and raw log

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
func (it *CBJTokenTokenPauseManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenTokenPauseManagerSet)
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
		it.Event = new(CBJTokenTokenPauseManagerSet)
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
func (it *CBJTokenTokenPauseManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenTokenPauseManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenTokenPauseManagerSet represents a TokenPauseManagerSet event raised by the CBJToken contract.
type CBJTokenTokenPauseManagerSet struct {
	OldTokenPauseManager common.Address
	NewTokenPauseManager common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTokenPauseManagerSet is a free log retrieval operation binding the contract event 0xe9029de0b26006bd4600e17ec4210e5009422c541c7d0988176dcb60d0c0138e.
//
// Solidity: event TokenPauseManagerSet(address indexed oldTokenPauseManager, address indexed newTokenPauseManager)
func (_CBJToken *CBJTokenFilterer) FilterTokenPauseManagerSet(opts *bind.FilterOpts, oldTokenPauseManager []common.Address, newTokenPauseManager []common.Address) (*CBJTokenTokenPauseManagerSetIterator, error) {

	var oldTokenPauseManagerRule []interface{}
	for _, oldTokenPauseManagerItem := range oldTokenPauseManager {
		oldTokenPauseManagerRule = append(oldTokenPauseManagerRule, oldTokenPauseManagerItem)
	}
	var newTokenPauseManagerRule []interface{}
	for _, newTokenPauseManagerItem := range newTokenPauseManager {
		newTokenPauseManagerRule = append(newTokenPauseManagerRule, newTokenPauseManagerItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "TokenPauseManagerSet", oldTokenPauseManagerRule, newTokenPauseManagerRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenTokenPauseManagerSetIterator{contract: _CBJToken.contract, event: "TokenPauseManagerSet", logs: logs, sub: sub}, nil
}

// WatchTokenPauseManagerSet is a free log subscription operation binding the contract event 0xe9029de0b26006bd4600e17ec4210e5009422c541c7d0988176dcb60d0c0138e.
//
// Solidity: event TokenPauseManagerSet(address indexed oldTokenPauseManager, address indexed newTokenPauseManager)
func (_CBJToken *CBJTokenFilterer) WatchTokenPauseManagerSet(opts *bind.WatchOpts, sink chan<- *CBJTokenTokenPauseManagerSet, oldTokenPauseManager []common.Address, newTokenPauseManager []common.Address) (event.Subscription, error) {

	var oldTokenPauseManagerRule []interface{}
	for _, oldTokenPauseManagerItem := range oldTokenPauseManager {
		oldTokenPauseManagerRule = append(oldTokenPauseManagerRule, oldTokenPauseManagerItem)
	}
	var newTokenPauseManagerRule []interface{}
	for _, newTokenPauseManagerItem := range newTokenPauseManager {
		newTokenPauseManagerRule = append(newTokenPauseManagerRule, newTokenPauseManagerItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "TokenPauseManagerSet", oldTokenPauseManagerRule, newTokenPauseManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenTokenPauseManagerSet)
				if err := _CBJToken.contract.UnpackLog(event, "TokenPauseManagerSet", log); err != nil {
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

// ParseTokenPauseManagerSet is a log parse operation binding the contract event 0xe9029de0b26006bd4600e17ec4210e5009422c541c7d0988176dcb60d0c0138e.
//
// Solidity: event TokenPauseManagerSet(address indexed oldTokenPauseManager, address indexed newTokenPauseManager)
func (_CBJToken *CBJTokenFilterer) ParseTokenPauseManagerSet(log types.Log) (*CBJTokenTokenPauseManagerSet, error) {
	event := new(CBJTokenTokenPauseManagerSet)
	if err := _CBJToken.contract.UnpackLog(event, "TokenPauseManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CBJToken contract.
type CBJTokenTransferIterator struct {
	Event *CBJTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CBJTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenTransfer)
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
		it.Event = new(CBJTokenTransfer)
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
func (it *CBJTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenTransfer represents a Transfer event raised by the CBJToken contract.
type CBJTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CBJToken *CBJTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CBJTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenTransferIterator{contract: _CBJToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CBJToken *CBJTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CBJTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenTransfer)
				if err := _CBJToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CBJToken *CBJTokenFilterer) ParseTransfer(log types.Log) (*CBJTokenTransfer, error) {
	event := new(CBJTokenTransfer)
	if err := _CBJToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJTokenTransferSharesIterator is returned from FilterTransferShares and is used to iterate over the raw logs and unpacked data for TransferShares events raised by the CBJToken contract.
type CBJTokenTransferSharesIterator struct {
	Event *CBJTokenTransferShares // Event containing the contract specifics and raw log

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
func (it *CBJTokenTransferSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJTokenTransferShares)
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
		it.Event = new(CBJTokenTransferShares)
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
func (it *CBJTokenTransferSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJTokenTransferSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJTokenTransferShares represents a TransferShares event raised by the CBJToken contract.
type CBJTokenTransferShares struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransferShares is a free log retrieval operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 value)
func (_CBJToken *CBJTokenFilterer) FilterTransferShares(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CBJTokenTransferSharesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CBJToken.contract.FilterLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CBJTokenTransferSharesIterator{contract: _CBJToken.contract, event: "TransferShares", logs: logs, sub: sub}, nil
}

// WatchTransferShares is a free log subscription operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 value)
func (_CBJToken *CBJTokenFilterer) WatchTransferShares(opts *bind.WatchOpts, sink chan<- *CBJTokenTransferShares, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CBJToken.contract.WatchLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJTokenTransferShares)
				if err := _CBJToken.contract.UnpackLog(event, "TransferShares", log); err != nil {
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

// ParseTransferShares is a log parse operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 value)
func (_CBJToken *CBJTokenFilterer) ParseTransferShares(log types.Log) (*CBJTokenTransferShares, error) {
	event := new(CBJTokenTransferShares)
	if err := _CBJToken.contract.UnpackLog(event, "TransferShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
