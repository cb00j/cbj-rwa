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

// ICBJGatewayDepositOperation is an auto generated low-level Go binding around an user-defined struct.
type ICBJGatewayDepositOperation struct {
	User                 common.Address
	UsdcAmount           *big.Int
	PendingCbjUSDCAmount *big.Int
	Status               uint8
	Timestamp            *big.Int
}

// ICBJGatewayWithdrawalOperation is an auto generated low-level Go binding around an user-defined struct.
type ICBJGatewayWithdrawalOperation struct {
	User              common.Address
	CbjUSDCAmount     *big.Int
	PendingUSDCAmount *big.Int
	Status            uint8
	Timestamp         *big.Int
}

// CBJGatewayMetaData contains all meta data concerning the CBJGateway contract.
var CBJGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_cbjUSDC\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CBJ_USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONFIGURE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PENDING_CBJ_USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PENDING_USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROCESSOR_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositOperations\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pendingCbjUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumICBJGateway.OperationStatus\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositsArePaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositOperation\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICBJGateway.DepositOperation\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pendingCbjUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumICBJGateway.OperationStatus\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawalOperation\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICBJGateway.WithdrawalOperation\",\"components\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cbjUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pendingUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumICBJGateway.OperationStatus\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_cbjUSDC\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_guardian\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minimumDepositAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minimumWithdrawalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minimumDepositAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumWithdrawalAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseDeposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseWithdrawals\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processDeposit\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"cbjUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processWithdrawal\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"usdcAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumDepositAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinimumWithdrawalAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpauseDeposits\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseWithdrawals\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"cbjUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalOperations\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"cbjUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pendingUSDCAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumICBJGateway.OperationStatus\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalsArePaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DepositProcessed\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cbjUSDCAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositsPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositsUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumDepositAmountSet\",\"inputs\":[{\"name\":\"oldAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinimumWithdrawalAmountSet\",\"inputs\":[{\"name\":\"oldAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingDeposit\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"usdcAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pendingCbjUSDCAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PendingWithdraw\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cbjUSDCAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pendingUSDCAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalProcessed\",\"inputs\":[{\"name\":\"operationId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"usdcAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalsPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalsUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountCannotBeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositAmountTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositsArePaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperationId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOperationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAuthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperationAlreadyProcessed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WithdrawalAmountTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalsArePaused\",\"inputs\":[]}]",
}

// CBJGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use CBJGatewayMetaData.ABI instead.
var CBJGatewayABI = CBJGatewayMetaData.ABI

// CBJGateway is an auto generated Go binding around an Ethereum contract.
type CBJGateway struct {
	CBJGatewayCaller     // Read-only binding to the contract
	CBJGatewayTransactor // Write-only binding to the contract
	CBJGatewayFilterer   // Log filterer for contract events
}

// CBJGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type CBJGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CBJGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CBJGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CBJGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CBJGatewaySession struct {
	Contract     *CBJGateway       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CBJGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CBJGatewayCallerSession struct {
	Contract *CBJGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CBJGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CBJGatewayTransactorSession struct {
	Contract     *CBJGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CBJGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type CBJGatewayRaw struct {
	Contract *CBJGateway // Generic contract binding to access the raw methods on
}

// CBJGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CBJGatewayCallerRaw struct {
	Contract *CBJGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// CBJGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CBJGatewayTransactorRaw struct {
	Contract *CBJGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCBJGateway creates a new instance of CBJGateway, bound to a specific deployed contract.
func NewCBJGateway(address common.Address, backend bind.ContractBackend) (*CBJGateway, error) {
	contract, err := bindCBJGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CBJGateway{CBJGatewayCaller: CBJGatewayCaller{contract: contract}, CBJGatewayTransactor: CBJGatewayTransactor{contract: contract}, CBJGatewayFilterer: CBJGatewayFilterer{contract: contract}}, nil
}

// NewCBJGatewayCaller creates a new read-only instance of CBJGateway, bound to a specific deployed contract.
func NewCBJGatewayCaller(address common.Address, caller bind.ContractCaller) (*CBJGatewayCaller, error) {
	contract, err := bindCBJGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayCaller{contract: contract}, nil
}

// NewCBJGatewayTransactor creates a new write-only instance of CBJGateway, bound to a specific deployed contract.
func NewCBJGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*CBJGatewayTransactor, error) {
	contract, err := bindCBJGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayTransactor{contract: contract}, nil
}

// NewCBJGatewayFilterer creates a new log filterer instance of CBJGateway, bound to a specific deployed contract.
func NewCBJGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*CBJGatewayFilterer, error) {
	contract, err := bindCBJGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayFilterer{contract: contract}, nil
}

// bindCBJGateway binds a generic wrapper to an already deployed contract.
func bindCBJGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CBJGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBJGateway *CBJGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CBJGateway.Contract.CBJGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBJGateway *CBJGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJGateway.Contract.CBJGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBJGateway *CBJGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBJGateway.Contract.CBJGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CBJGateway *CBJGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CBJGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CBJGateway *CBJGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CBJGateway *CBJGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CBJGateway.Contract.contract.Transact(opts, method, params...)
}

// CBJUSDC is a free data retrieval call binding the contract method 0x30de2a96.
//
// Solidity: function CBJ_USDC() view returns(address)
func (_CBJGateway *CBJGatewayCaller) CBJUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "CBJ_USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CBJUSDC is a free data retrieval call binding the contract method 0x30de2a96.
//
// Solidity: function CBJ_USDC() view returns(address)
func (_CBJGateway *CBJGatewaySession) CBJUSDC() (common.Address, error) {
	return _CBJGateway.Contract.CBJUSDC(&_CBJGateway.CallOpts)
}

// CBJUSDC is a free data retrieval call binding the contract method 0x30de2a96.
//
// Solidity: function CBJ_USDC() view returns(address)
func (_CBJGateway *CBJGatewayCallerSession) CBJUSDC() (common.Address, error) {
	return _CBJGateway.Contract.CBJUSDC(&_CBJGateway.CallOpts)
}

// CONFIGUREROLE is a free data retrieval call binding the contract method 0x2b41ceb5.
//
// Solidity: function CONFIGURE_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCaller) CONFIGUREROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "CONFIGURE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGUREROLE is a free data retrieval call binding the contract method 0x2b41ceb5.
//
// Solidity: function CONFIGURE_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewaySession) CONFIGUREROLE() ([32]byte, error) {
	return _CBJGateway.Contract.CONFIGUREROLE(&_CBJGateway.CallOpts)
}

// CONFIGUREROLE is a free data retrieval call binding the contract method 0x2b41ceb5.
//
// Solidity: function CONFIGURE_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCallerSession) CONFIGUREROLE() ([32]byte, error) {
	return _CBJGateway.Contract.CONFIGUREROLE(&_CBJGateway.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CBJGateway.Contract.DEFAULTADMINROLE(&_CBJGateway.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CBJGateway.Contract.DEFAULTADMINROLE(&_CBJGateway.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewaySession) PAUSEROLE() ([32]byte, error) {
	return _CBJGateway.Contract.PAUSEROLE(&_CBJGateway.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCallerSession) PAUSEROLE() ([32]byte, error) {
	return _CBJGateway.Contract.PAUSEROLE(&_CBJGateway.CallOpts)
}

// PENDINGCBJUSDC is a free data retrieval call binding the contract method 0x5c24c752.
//
// Solidity: function PENDING_CBJ_USDC() view returns(address)
func (_CBJGateway *CBJGatewayCaller) PENDINGCBJUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "PENDING_CBJ_USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PENDINGCBJUSDC is a free data retrieval call binding the contract method 0x5c24c752.
//
// Solidity: function PENDING_CBJ_USDC() view returns(address)
func (_CBJGateway *CBJGatewaySession) PENDINGCBJUSDC() (common.Address, error) {
	return _CBJGateway.Contract.PENDINGCBJUSDC(&_CBJGateway.CallOpts)
}

// PENDINGCBJUSDC is a free data retrieval call binding the contract method 0x5c24c752.
//
// Solidity: function PENDING_CBJ_USDC() view returns(address)
func (_CBJGateway *CBJGatewayCallerSession) PENDINGCBJUSDC() (common.Address, error) {
	return _CBJGateway.Contract.PENDINGCBJUSDC(&_CBJGateway.CallOpts)
}

// PENDINGUSDC is a free data retrieval call binding the contract method 0xe9bda9a0.
//
// Solidity: function PENDING_USDC() view returns(address)
func (_CBJGateway *CBJGatewayCaller) PENDINGUSDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "PENDING_USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PENDINGUSDC is a free data retrieval call binding the contract method 0xe9bda9a0.
//
// Solidity: function PENDING_USDC() view returns(address)
func (_CBJGateway *CBJGatewaySession) PENDINGUSDC() (common.Address, error) {
	return _CBJGateway.Contract.PENDINGUSDC(&_CBJGateway.CallOpts)
}

// PENDINGUSDC is a free data retrieval call binding the contract method 0xe9bda9a0.
//
// Solidity: function PENDING_USDC() view returns(address)
func (_CBJGateway *CBJGatewayCallerSession) PENDINGUSDC() (common.Address, error) {
	return _CBJGateway.Contract.PENDINGUSDC(&_CBJGateway.CallOpts)
}

// PROCESSORROLE is a free data retrieval call binding the contract method 0x8222bdb2.
//
// Solidity: function PROCESSOR_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCaller) PROCESSORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "PROCESSOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROCESSORROLE is a free data retrieval call binding the contract method 0x8222bdb2.
//
// Solidity: function PROCESSOR_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewaySession) PROCESSORROLE() ([32]byte, error) {
	return _CBJGateway.Contract.PROCESSORROLE(&_CBJGateway.CallOpts)
}

// PROCESSORROLE is a free data retrieval call binding the contract method 0x8222bdb2.
//
// Solidity: function PROCESSOR_ROLE() view returns(bytes32)
func (_CBJGateway *CBJGatewayCallerSession) PROCESSORROLE() ([32]byte, error) {
	return _CBJGateway.Contract.PROCESSORROLE(&_CBJGateway.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_CBJGateway *CBJGatewayCaller) USDC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "USDC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_CBJGateway *CBJGatewaySession) USDC() (common.Address, error) {
	return _CBJGateway.Contract.USDC(&_CBJGateway.CallOpts)
}

// USDC is a free data retrieval call binding the contract method 0x89a30271.
//
// Solidity: function USDC() view returns(address)
func (_CBJGateway *CBJGatewayCallerSession) USDC() (common.Address, error) {
	return _CBJGateway.Contract.USDC(&_CBJGateway.CallOpts)
}

// DepositOperations is a free data retrieval call binding the contract method 0x6778a330.
//
// Solidity: function depositOperations(bytes32 ) view returns(address user, uint256 usdcAmount, uint256 pendingCbjUSDCAmount, uint8 status, uint256 timestamp)
func (_CBJGateway *CBJGatewayCaller) DepositOperations(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User                 common.Address
	UsdcAmount           *big.Int
	PendingCbjUSDCAmount *big.Int
	Status               uint8
	Timestamp            *big.Int
}, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "depositOperations", arg0)

	outstruct := new(struct {
		User                 common.Address
		UsdcAmount           *big.Int
		PendingCbjUSDCAmount *big.Int
		Status               uint8
		Timestamp            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.UsdcAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingCbjUSDCAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositOperations is a free data retrieval call binding the contract method 0x6778a330.
//
// Solidity: function depositOperations(bytes32 ) view returns(address user, uint256 usdcAmount, uint256 pendingCbjUSDCAmount, uint8 status, uint256 timestamp)
func (_CBJGateway *CBJGatewaySession) DepositOperations(arg0 [32]byte) (struct {
	User                 common.Address
	UsdcAmount           *big.Int
	PendingCbjUSDCAmount *big.Int
	Status               uint8
	Timestamp            *big.Int
}, error) {
	return _CBJGateway.Contract.DepositOperations(&_CBJGateway.CallOpts, arg0)
}

// DepositOperations is a free data retrieval call binding the contract method 0x6778a330.
//
// Solidity: function depositOperations(bytes32 ) view returns(address user, uint256 usdcAmount, uint256 pendingCbjUSDCAmount, uint8 status, uint256 timestamp)
func (_CBJGateway *CBJGatewayCallerSession) DepositOperations(arg0 [32]byte) (struct {
	User                 common.Address
	UsdcAmount           *big.Int
	PendingCbjUSDCAmount *big.Int
	Status               uint8
	Timestamp            *big.Int
}, error) {
	return _CBJGateway.Contract.DepositOperations(&_CBJGateway.CallOpts, arg0)
}

// DepositsArePaused is a free data retrieval call binding the contract method 0x248dac38.
//
// Solidity: function depositsArePaused() view returns(bool)
func (_CBJGateway *CBJGatewayCaller) DepositsArePaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "depositsArePaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositsArePaused is a free data retrieval call binding the contract method 0x248dac38.
//
// Solidity: function depositsArePaused() view returns(bool)
func (_CBJGateway *CBJGatewaySession) DepositsArePaused() (bool, error) {
	return _CBJGateway.Contract.DepositsArePaused(&_CBJGateway.CallOpts)
}

// DepositsArePaused is a free data retrieval call binding the contract method 0x248dac38.
//
// Solidity: function depositsArePaused() view returns(bool)
func (_CBJGateway *CBJGatewayCallerSession) DepositsArePaused() (bool, error) {
	return _CBJGateway.Contract.DepositsArePaused(&_CBJGateway.CallOpts)
}

// GetDepositOperation is a free data retrieval call binding the contract method 0xdbdc02b3.
//
// Solidity: function getDepositOperation(bytes32 operationId) view returns((address,uint256,uint256,uint8,uint256))
func (_CBJGateway *CBJGatewayCaller) GetDepositOperation(opts *bind.CallOpts, operationId [32]byte) (ICBJGatewayDepositOperation, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "getDepositOperation", operationId)

	if err != nil {
		return *new(ICBJGatewayDepositOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(ICBJGatewayDepositOperation)).(*ICBJGatewayDepositOperation)

	return out0, err

}

// GetDepositOperation is a free data retrieval call binding the contract method 0xdbdc02b3.
//
// Solidity: function getDepositOperation(bytes32 operationId) view returns((address,uint256,uint256,uint8,uint256))
func (_CBJGateway *CBJGatewaySession) GetDepositOperation(operationId [32]byte) (ICBJGatewayDepositOperation, error) {
	return _CBJGateway.Contract.GetDepositOperation(&_CBJGateway.CallOpts, operationId)
}

// GetDepositOperation is a free data retrieval call binding the contract method 0xdbdc02b3.
//
// Solidity: function getDepositOperation(bytes32 operationId) view returns((address,uint256,uint256,uint8,uint256))
func (_CBJGateway *CBJGatewayCallerSession) GetDepositOperation(operationId [32]byte) (ICBJGatewayDepositOperation, error) {
	return _CBJGateway.Contract.GetDepositOperation(&_CBJGateway.CallOpts, operationId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CBJGateway *CBJGatewayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CBJGateway *CBJGatewaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CBJGateway.Contract.GetRoleAdmin(&_CBJGateway.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CBJGateway *CBJGatewayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CBJGateway.Contract.GetRoleAdmin(&_CBJGateway.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CBJGateway *CBJGatewayCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CBJGateway *CBJGatewaySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CBJGateway.Contract.GetRoleMember(&_CBJGateway.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CBJGateway *CBJGatewayCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CBJGateway.Contract.GetRoleMember(&_CBJGateway.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CBJGateway *CBJGatewayCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CBJGateway *CBJGatewaySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CBJGateway.Contract.GetRoleMemberCount(&_CBJGateway.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CBJGateway *CBJGatewayCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CBJGateway.Contract.GetRoleMemberCount(&_CBJGateway.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CBJGateway *CBJGatewayCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CBJGateway *CBJGatewaySession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CBJGateway.Contract.GetRoleMembers(&_CBJGateway.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_CBJGateway *CBJGatewayCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _CBJGateway.Contract.GetRoleMembers(&_CBJGateway.CallOpts, role)
}

// GetWithdrawalOperation is a free data retrieval call binding the contract method 0xada9805f.
//
// Solidity: function getWithdrawalOperation(bytes32 operationId) view returns((address,uint256,uint256,uint8,uint256))
func (_CBJGateway *CBJGatewayCaller) GetWithdrawalOperation(opts *bind.CallOpts, operationId [32]byte) (ICBJGatewayWithdrawalOperation, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "getWithdrawalOperation", operationId)

	if err != nil {
		return *new(ICBJGatewayWithdrawalOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(ICBJGatewayWithdrawalOperation)).(*ICBJGatewayWithdrawalOperation)

	return out0, err

}

// GetWithdrawalOperation is a free data retrieval call binding the contract method 0xada9805f.
//
// Solidity: function getWithdrawalOperation(bytes32 operationId) view returns((address,uint256,uint256,uint8,uint256))
func (_CBJGateway *CBJGatewaySession) GetWithdrawalOperation(operationId [32]byte) (ICBJGatewayWithdrawalOperation, error) {
	return _CBJGateway.Contract.GetWithdrawalOperation(&_CBJGateway.CallOpts, operationId)
}

// GetWithdrawalOperation is a free data retrieval call binding the contract method 0xada9805f.
//
// Solidity: function getWithdrawalOperation(bytes32 operationId) view returns((address,uint256,uint256,uint8,uint256))
func (_CBJGateway *CBJGatewayCallerSession) GetWithdrawalOperation(operationId [32]byte) (ICBJGatewayWithdrawalOperation, error) {
	return _CBJGateway.Contract.GetWithdrawalOperation(&_CBJGateway.CallOpts, operationId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CBJGateway *CBJGatewayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CBJGateway *CBJGatewaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CBJGateway.Contract.HasRole(&_CBJGateway.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CBJGateway *CBJGatewayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CBJGateway.Contract.HasRole(&_CBJGateway.CallOpts, role, account)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_CBJGateway *CBJGatewayCaller) MinimumDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "minimumDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_CBJGateway *CBJGatewaySession) MinimumDepositAmount() (*big.Int, error) {
	return _CBJGateway.Contract.MinimumDepositAmount(&_CBJGateway.CallOpts)
}

// MinimumDepositAmount is a free data retrieval call binding the contract method 0x080c279a.
//
// Solidity: function minimumDepositAmount() view returns(uint256)
func (_CBJGateway *CBJGatewayCallerSession) MinimumDepositAmount() (*big.Int, error) {
	return _CBJGateway.Contract.MinimumDepositAmount(&_CBJGateway.CallOpts)
}

// MinimumWithdrawalAmount is a free data retrieval call binding the contract method 0x2b180646.
//
// Solidity: function minimumWithdrawalAmount() view returns(uint256)
func (_CBJGateway *CBJGatewayCaller) MinimumWithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "minimumWithdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumWithdrawalAmount is a free data retrieval call binding the contract method 0x2b180646.
//
// Solidity: function minimumWithdrawalAmount() view returns(uint256)
func (_CBJGateway *CBJGatewaySession) MinimumWithdrawalAmount() (*big.Int, error) {
	return _CBJGateway.Contract.MinimumWithdrawalAmount(&_CBJGateway.CallOpts)
}

// MinimumWithdrawalAmount is a free data retrieval call binding the contract method 0x2b180646.
//
// Solidity: function minimumWithdrawalAmount() view returns(uint256)
func (_CBJGateway *CBJGatewayCallerSession) MinimumWithdrawalAmount() (*big.Int, error) {
	return _CBJGateway.Contract.MinimumWithdrawalAmount(&_CBJGateway.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CBJGateway *CBJGatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CBJGateway *CBJGatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CBJGateway.Contract.SupportsInterface(&_CBJGateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CBJGateway *CBJGatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CBJGateway.Contract.SupportsInterface(&_CBJGateway.CallOpts, interfaceId)
}

// WithdrawalOperations is a free data retrieval call binding the contract method 0xf85440cc.
//
// Solidity: function withdrawalOperations(bytes32 ) view returns(address user, uint256 cbjUSDCAmount, uint256 pendingUSDCAmount, uint8 status, uint256 timestamp)
func (_CBJGateway *CBJGatewayCaller) WithdrawalOperations(opts *bind.CallOpts, arg0 [32]byte) (struct {
	User              common.Address
	CbjUSDCAmount     *big.Int
	PendingUSDCAmount *big.Int
	Status            uint8
	Timestamp         *big.Int
}, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "withdrawalOperations", arg0)

	outstruct := new(struct {
		User              common.Address
		CbjUSDCAmount     *big.Int
		PendingUSDCAmount *big.Int
		Status            uint8
		Timestamp         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CbjUSDCAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PendingUSDCAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawalOperations is a free data retrieval call binding the contract method 0xf85440cc.
//
// Solidity: function withdrawalOperations(bytes32 ) view returns(address user, uint256 cbjUSDCAmount, uint256 pendingUSDCAmount, uint8 status, uint256 timestamp)
func (_CBJGateway *CBJGatewaySession) WithdrawalOperations(arg0 [32]byte) (struct {
	User              common.Address
	CbjUSDCAmount     *big.Int
	PendingUSDCAmount *big.Int
	Status            uint8
	Timestamp         *big.Int
}, error) {
	return _CBJGateway.Contract.WithdrawalOperations(&_CBJGateway.CallOpts, arg0)
}

// WithdrawalOperations is a free data retrieval call binding the contract method 0xf85440cc.
//
// Solidity: function withdrawalOperations(bytes32 ) view returns(address user, uint256 cbjUSDCAmount, uint256 pendingUSDCAmount, uint8 status, uint256 timestamp)
func (_CBJGateway *CBJGatewayCallerSession) WithdrawalOperations(arg0 [32]byte) (struct {
	User              common.Address
	CbjUSDCAmount     *big.Int
	PendingUSDCAmount *big.Int
	Status            uint8
	Timestamp         *big.Int
}, error) {
	return _CBJGateway.Contract.WithdrawalOperations(&_CBJGateway.CallOpts, arg0)
}

// WithdrawalsArePaused is a free data retrieval call binding the contract method 0xca077176.
//
// Solidity: function withdrawalsArePaused() view returns(bool)
func (_CBJGateway *CBJGatewayCaller) WithdrawalsArePaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CBJGateway.contract.Call(opts, &out, "withdrawalsArePaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalsArePaused is a free data retrieval call binding the contract method 0xca077176.
//
// Solidity: function withdrawalsArePaused() view returns(bool)
func (_CBJGateway *CBJGatewaySession) WithdrawalsArePaused() (bool, error) {
	return _CBJGateway.Contract.WithdrawalsArePaused(&_CBJGateway.CallOpts)
}

// WithdrawalsArePaused is a free data retrieval call binding the contract method 0xca077176.
//
// Solidity: function withdrawalsArePaused() view returns(bool)
func (_CBJGateway *CBJGatewayCallerSession) WithdrawalsArePaused() (bool, error) {
	return _CBJGateway.Contract.WithdrawalsArePaused(&_CBJGateway.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(bytes32 operationId)
func (_CBJGateway *CBJGatewayTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(bytes32 operationId)
func (_CBJGateway *CBJGatewaySession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.Deposit(&_CBJGateway.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns(bytes32 operationId)
func (_CBJGateway *CBJGatewayTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.Deposit(&_CBJGateway.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CBJGateway *CBJGatewayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CBJGateway *CBJGatewaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJGateway.Contract.GrantRole(&_CBJGateway.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CBJGateway *CBJGatewayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJGateway.Contract.GrantRole(&_CBJGateway.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address _usdc, address _cbjUSDC, address _guardian, uint256 _minimumDepositAmount, uint256 _minimumWithdrawalAmount) returns()
func (_CBJGateway *CBJGatewayTransactor) Initialize(opts *bind.TransactOpts, _usdc common.Address, _cbjUSDC common.Address, _guardian common.Address, _minimumDepositAmount *big.Int, _minimumWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "initialize", _usdc, _cbjUSDC, _guardian, _minimumDepositAmount, _minimumWithdrawalAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address _usdc, address _cbjUSDC, address _guardian, uint256 _minimumDepositAmount, uint256 _minimumWithdrawalAmount) returns()
func (_CBJGateway *CBJGatewaySession) Initialize(_usdc common.Address, _cbjUSDC common.Address, _guardian common.Address, _minimumDepositAmount *big.Int, _minimumWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.Initialize(&_CBJGateway.TransactOpts, _usdc, _cbjUSDC, _guardian, _minimumDepositAmount, _minimumWithdrawalAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address _usdc, address _cbjUSDC, address _guardian, uint256 _minimumDepositAmount, uint256 _minimumWithdrawalAmount) returns()
func (_CBJGateway *CBJGatewayTransactorSession) Initialize(_usdc common.Address, _cbjUSDC common.Address, _guardian common.Address, _minimumDepositAmount *big.Int, _minimumWithdrawalAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.Initialize(&_CBJGateway.TransactOpts, _usdc, _cbjUSDC, _guardian, _minimumDepositAmount, _minimumWithdrawalAmount)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x02191980.
//
// Solidity: function pauseDeposits() returns()
func (_CBJGateway *CBJGatewayTransactor) PauseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "pauseDeposits")
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x02191980.
//
// Solidity: function pauseDeposits() returns()
func (_CBJGateway *CBJGatewaySession) PauseDeposits() (*types.Transaction, error) {
	return _CBJGateway.Contract.PauseDeposits(&_CBJGateway.TransactOpts)
}

// PauseDeposits is a paid mutator transaction binding the contract method 0x02191980.
//
// Solidity: function pauseDeposits() returns()
func (_CBJGateway *CBJGatewayTransactorSession) PauseDeposits() (*types.Transaction, error) {
	return _CBJGateway.Contract.PauseDeposits(&_CBJGateway.TransactOpts)
}

// PauseWithdrawals is a paid mutator transaction binding the contract method 0x56bb54a7.
//
// Solidity: function pauseWithdrawals() returns()
func (_CBJGateway *CBJGatewayTransactor) PauseWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "pauseWithdrawals")
}

// PauseWithdrawals is a paid mutator transaction binding the contract method 0x56bb54a7.
//
// Solidity: function pauseWithdrawals() returns()
func (_CBJGateway *CBJGatewaySession) PauseWithdrawals() (*types.Transaction, error) {
	return _CBJGateway.Contract.PauseWithdrawals(&_CBJGateway.TransactOpts)
}

// PauseWithdrawals is a paid mutator transaction binding the contract method 0x56bb54a7.
//
// Solidity: function pauseWithdrawals() returns()
func (_CBJGateway *CBJGatewayTransactorSession) PauseWithdrawals() (*types.Transaction, error) {
	return _CBJGateway.Contract.PauseWithdrawals(&_CBJGateway.TransactOpts)
}

// ProcessDeposit is a paid mutator transaction binding the contract method 0x161bd1fb.
//
// Solidity: function processDeposit(bytes32 operationId, uint256 cbjUSDCAmount) returns()
func (_CBJGateway *CBJGatewayTransactor) ProcessDeposit(opts *bind.TransactOpts, operationId [32]byte, cbjUSDCAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "processDeposit", operationId, cbjUSDCAmount)
}

// ProcessDeposit is a paid mutator transaction binding the contract method 0x161bd1fb.
//
// Solidity: function processDeposit(bytes32 operationId, uint256 cbjUSDCAmount) returns()
func (_CBJGateway *CBJGatewaySession) ProcessDeposit(operationId [32]byte, cbjUSDCAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.ProcessDeposit(&_CBJGateway.TransactOpts, operationId, cbjUSDCAmount)
}

// ProcessDeposit is a paid mutator transaction binding the contract method 0x161bd1fb.
//
// Solidity: function processDeposit(bytes32 operationId, uint256 cbjUSDCAmount) returns()
func (_CBJGateway *CBJGatewayTransactorSession) ProcessDeposit(operationId [32]byte, cbjUSDCAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.ProcessDeposit(&_CBJGateway.TransactOpts, operationId, cbjUSDCAmount)
}

// ProcessWithdrawal is a paid mutator transaction binding the contract method 0x99db98eb.
//
// Solidity: function processWithdrawal(bytes32 operationId, uint256 usdcAmount) returns()
func (_CBJGateway *CBJGatewayTransactor) ProcessWithdrawal(opts *bind.TransactOpts, operationId [32]byte, usdcAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "processWithdrawal", operationId, usdcAmount)
}

// ProcessWithdrawal is a paid mutator transaction binding the contract method 0x99db98eb.
//
// Solidity: function processWithdrawal(bytes32 operationId, uint256 usdcAmount) returns()
func (_CBJGateway *CBJGatewaySession) ProcessWithdrawal(operationId [32]byte, usdcAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.ProcessWithdrawal(&_CBJGateway.TransactOpts, operationId, usdcAmount)
}

// ProcessWithdrawal is a paid mutator transaction binding the contract method 0x99db98eb.
//
// Solidity: function processWithdrawal(bytes32 operationId, uint256 usdcAmount) returns()
func (_CBJGateway *CBJGatewayTransactorSession) ProcessWithdrawal(operationId [32]byte, usdcAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.ProcessWithdrawal(&_CBJGateway.TransactOpts, operationId, usdcAmount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CBJGateway *CBJGatewayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CBJGateway *CBJGatewaySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CBJGateway.Contract.RenounceRole(&_CBJGateway.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CBJGateway *CBJGatewayTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CBJGateway.Contract.RenounceRole(&_CBJGateway.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CBJGateway *CBJGatewayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CBJGateway *CBJGatewaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJGateway.Contract.RevokeRole(&_CBJGateway.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CBJGateway *CBJGatewayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CBJGateway.Contract.RevokeRole(&_CBJGateway.TransactOpts, role, account)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 amount) returns()
func (_CBJGateway *CBJGatewayTransactor) SetMinimumDepositAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "setMinimumDepositAmount", amount)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 amount) returns()
func (_CBJGateway *CBJGatewaySession) SetMinimumDepositAmount(amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.SetMinimumDepositAmount(&_CBJGateway.TransactOpts, amount)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 amount) returns()
func (_CBJGateway *CBJGatewayTransactorSession) SetMinimumDepositAmount(amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.SetMinimumDepositAmount(&_CBJGateway.TransactOpts, amount)
}

// SetMinimumWithdrawalAmount is a paid mutator transaction binding the contract method 0x3620d373.
//
// Solidity: function setMinimumWithdrawalAmount(uint256 amount) returns()
func (_CBJGateway *CBJGatewayTransactor) SetMinimumWithdrawalAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "setMinimumWithdrawalAmount", amount)
}

// SetMinimumWithdrawalAmount is a paid mutator transaction binding the contract method 0x3620d373.
//
// Solidity: function setMinimumWithdrawalAmount(uint256 amount) returns()
func (_CBJGateway *CBJGatewaySession) SetMinimumWithdrawalAmount(amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.SetMinimumWithdrawalAmount(&_CBJGateway.TransactOpts, amount)
}

// SetMinimumWithdrawalAmount is a paid mutator transaction binding the contract method 0x3620d373.
//
// Solidity: function setMinimumWithdrawalAmount(uint256 amount) returns()
func (_CBJGateway *CBJGatewayTransactorSession) SetMinimumWithdrawalAmount(amount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.SetMinimumWithdrawalAmount(&_CBJGateway.TransactOpts, amount)
}

// UnpauseDeposits is a paid mutator transaction binding the contract method 0x63d8882a.
//
// Solidity: function unpauseDeposits() returns()
func (_CBJGateway *CBJGatewayTransactor) UnpauseDeposits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "unpauseDeposits")
}

// UnpauseDeposits is a paid mutator transaction binding the contract method 0x63d8882a.
//
// Solidity: function unpauseDeposits() returns()
func (_CBJGateway *CBJGatewaySession) UnpauseDeposits() (*types.Transaction, error) {
	return _CBJGateway.Contract.UnpauseDeposits(&_CBJGateway.TransactOpts)
}

// UnpauseDeposits is a paid mutator transaction binding the contract method 0x63d8882a.
//
// Solidity: function unpauseDeposits() returns()
func (_CBJGateway *CBJGatewayTransactorSession) UnpauseDeposits() (*types.Transaction, error) {
	return _CBJGateway.Contract.UnpauseDeposits(&_CBJGateway.TransactOpts)
}

// UnpauseWithdrawals is a paid mutator transaction binding the contract method 0xe4c4be58.
//
// Solidity: function unpauseWithdrawals() returns()
func (_CBJGateway *CBJGatewayTransactor) UnpauseWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "unpauseWithdrawals")
}

// UnpauseWithdrawals is a paid mutator transaction binding the contract method 0xe4c4be58.
//
// Solidity: function unpauseWithdrawals() returns()
func (_CBJGateway *CBJGatewaySession) UnpauseWithdrawals() (*types.Transaction, error) {
	return _CBJGateway.Contract.UnpauseWithdrawals(&_CBJGateway.TransactOpts)
}

// UnpauseWithdrawals is a paid mutator transaction binding the contract method 0xe4c4be58.
//
// Solidity: function unpauseWithdrawals() returns()
func (_CBJGateway *CBJGatewayTransactorSession) UnpauseWithdrawals() (*types.Transaction, error) {
	return _CBJGateway.Contract.UnpauseWithdrawals(&_CBJGateway.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 cbjUSDCAmount) returns(bytes32 operationId)
func (_CBJGateway *CBJGatewayTransactor) Withdraw(opts *bind.TransactOpts, cbjUSDCAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.contract.Transact(opts, "withdraw", cbjUSDCAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 cbjUSDCAmount) returns(bytes32 operationId)
func (_CBJGateway *CBJGatewaySession) Withdraw(cbjUSDCAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.Withdraw(&_CBJGateway.TransactOpts, cbjUSDCAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 cbjUSDCAmount) returns(bytes32 operationId)
func (_CBJGateway *CBJGatewayTransactorSession) Withdraw(cbjUSDCAmount *big.Int) (*types.Transaction, error) {
	return _CBJGateway.Contract.Withdraw(&_CBJGateway.TransactOpts, cbjUSDCAmount)
}

// CBJGatewayDepositProcessedIterator is returned from FilterDepositProcessed and is used to iterate over the raw logs and unpacked data for DepositProcessed events raised by the CBJGateway contract.
type CBJGatewayDepositProcessedIterator struct {
	Event *CBJGatewayDepositProcessed // Event containing the contract specifics and raw log

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
func (it *CBJGatewayDepositProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayDepositProcessed)
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
		it.Event = new(CBJGatewayDepositProcessed)
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
func (it *CBJGatewayDepositProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayDepositProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayDepositProcessed represents a DepositProcessed event raised by the CBJGateway contract.
type CBJGatewayDepositProcessed struct {
	OperationId   [32]byte
	User          common.Address
	CbjUSDCAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDepositProcessed is a free log retrieval operation binding the contract event 0x7d77e3f1866839ef93f5eda23990fa4368bf7b0af583ae0f64ecae4da14a4d61.
//
// Solidity: event DepositProcessed(bytes32 indexed operationId, address indexed user, uint256 cbjUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) FilterDepositProcessed(opts *bind.FilterOpts, operationId [][32]byte, user []common.Address) (*CBJGatewayDepositProcessedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "DepositProcessed", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayDepositProcessedIterator{contract: _CBJGateway.contract, event: "DepositProcessed", logs: logs, sub: sub}, nil
}

// WatchDepositProcessed is a free log subscription operation binding the contract event 0x7d77e3f1866839ef93f5eda23990fa4368bf7b0af583ae0f64ecae4da14a4d61.
//
// Solidity: event DepositProcessed(bytes32 indexed operationId, address indexed user, uint256 cbjUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) WatchDepositProcessed(opts *bind.WatchOpts, sink chan<- *CBJGatewayDepositProcessed, operationId [][32]byte, user []common.Address) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "DepositProcessed", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayDepositProcessed)
				if err := _CBJGateway.contract.UnpackLog(event, "DepositProcessed", log); err != nil {
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

// ParseDepositProcessed is a log parse operation binding the contract event 0x7d77e3f1866839ef93f5eda23990fa4368bf7b0af583ae0f64ecae4da14a4d61.
//
// Solidity: event DepositProcessed(bytes32 indexed operationId, address indexed user, uint256 cbjUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) ParseDepositProcessed(log types.Log) (*CBJGatewayDepositProcessed, error) {
	event := new(CBJGatewayDepositProcessed)
	if err := _CBJGateway.contract.UnpackLog(event, "DepositProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayDepositsPausedIterator is returned from FilterDepositsPaused and is used to iterate over the raw logs and unpacked data for DepositsPaused events raised by the CBJGateway contract.
type CBJGatewayDepositsPausedIterator struct {
	Event *CBJGatewayDepositsPaused // Event containing the contract specifics and raw log

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
func (it *CBJGatewayDepositsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayDepositsPaused)
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
		it.Event = new(CBJGatewayDepositsPaused)
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
func (it *CBJGatewayDepositsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayDepositsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayDepositsPaused represents a DepositsPaused event raised by the CBJGateway contract.
type CBJGatewayDepositsPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositsPaused is a free log retrieval operation binding the contract event 0xdeeb69430b7153361c25d630947115165636e6a723fa8daea4b0de34b3247459.
//
// Solidity: event DepositsPaused()
func (_CBJGateway *CBJGatewayFilterer) FilterDepositsPaused(opts *bind.FilterOpts) (*CBJGatewayDepositsPausedIterator, error) {

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "DepositsPaused")
	if err != nil {
		return nil, err
	}
	return &CBJGatewayDepositsPausedIterator{contract: _CBJGateway.contract, event: "DepositsPaused", logs: logs, sub: sub}, nil
}

// WatchDepositsPaused is a free log subscription operation binding the contract event 0xdeeb69430b7153361c25d630947115165636e6a723fa8daea4b0de34b3247459.
//
// Solidity: event DepositsPaused()
func (_CBJGateway *CBJGatewayFilterer) WatchDepositsPaused(opts *bind.WatchOpts, sink chan<- *CBJGatewayDepositsPaused) (event.Subscription, error) {

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "DepositsPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayDepositsPaused)
				if err := _CBJGateway.contract.UnpackLog(event, "DepositsPaused", log); err != nil {
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

// ParseDepositsPaused is a log parse operation binding the contract event 0xdeeb69430b7153361c25d630947115165636e6a723fa8daea4b0de34b3247459.
//
// Solidity: event DepositsPaused()
func (_CBJGateway *CBJGatewayFilterer) ParseDepositsPaused(log types.Log) (*CBJGatewayDepositsPaused, error) {
	event := new(CBJGatewayDepositsPaused)
	if err := _CBJGateway.contract.UnpackLog(event, "DepositsPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayDepositsUnpausedIterator is returned from FilterDepositsUnpaused and is used to iterate over the raw logs and unpacked data for DepositsUnpaused events raised by the CBJGateway contract.
type CBJGatewayDepositsUnpausedIterator struct {
	Event *CBJGatewayDepositsUnpaused // Event containing the contract specifics and raw log

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
func (it *CBJGatewayDepositsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayDepositsUnpaused)
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
		it.Event = new(CBJGatewayDepositsUnpaused)
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
func (it *CBJGatewayDepositsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayDepositsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayDepositsUnpaused represents a DepositsUnpaused event raised by the CBJGateway contract.
type CBJGatewayDepositsUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDepositsUnpaused is a free log retrieval operation binding the contract event 0x823084e804e36d8971e8b86749b6b0ace7b9f87ed272bef910c1e72d123eeb48.
//
// Solidity: event DepositsUnpaused()
func (_CBJGateway *CBJGatewayFilterer) FilterDepositsUnpaused(opts *bind.FilterOpts) (*CBJGatewayDepositsUnpausedIterator, error) {

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "DepositsUnpaused")
	if err != nil {
		return nil, err
	}
	return &CBJGatewayDepositsUnpausedIterator{contract: _CBJGateway.contract, event: "DepositsUnpaused", logs: logs, sub: sub}, nil
}

// WatchDepositsUnpaused is a free log subscription operation binding the contract event 0x823084e804e36d8971e8b86749b6b0ace7b9f87ed272bef910c1e72d123eeb48.
//
// Solidity: event DepositsUnpaused()
func (_CBJGateway *CBJGatewayFilterer) WatchDepositsUnpaused(opts *bind.WatchOpts, sink chan<- *CBJGatewayDepositsUnpaused) (event.Subscription, error) {

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "DepositsUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayDepositsUnpaused)
				if err := _CBJGateway.contract.UnpackLog(event, "DepositsUnpaused", log); err != nil {
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

// ParseDepositsUnpaused is a log parse operation binding the contract event 0x823084e804e36d8971e8b86749b6b0ace7b9f87ed272bef910c1e72d123eeb48.
//
// Solidity: event DepositsUnpaused()
func (_CBJGateway *CBJGatewayFilterer) ParseDepositsUnpaused(log types.Log) (*CBJGatewayDepositsUnpaused, error) {
	event := new(CBJGatewayDepositsUnpaused)
	if err := _CBJGateway.contract.UnpackLog(event, "DepositsUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CBJGateway contract.
type CBJGatewayInitializedIterator struct {
	Event *CBJGatewayInitialized // Event containing the contract specifics and raw log

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
func (it *CBJGatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayInitialized)
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
		it.Event = new(CBJGatewayInitialized)
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
func (it *CBJGatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayInitialized represents a Initialized event raised by the CBJGateway contract.
type CBJGatewayInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CBJGateway *CBJGatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*CBJGatewayInitializedIterator, error) {

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CBJGatewayInitializedIterator{contract: _CBJGateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CBJGateway *CBJGatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CBJGatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayInitialized)
				if err := _CBJGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CBJGateway *CBJGatewayFilterer) ParseInitialized(log types.Log) (*CBJGatewayInitialized, error) {
	event := new(CBJGatewayInitialized)
	if err := _CBJGateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayMinimumDepositAmountSetIterator is returned from FilterMinimumDepositAmountSet and is used to iterate over the raw logs and unpacked data for MinimumDepositAmountSet events raised by the CBJGateway contract.
type CBJGatewayMinimumDepositAmountSetIterator struct {
	Event *CBJGatewayMinimumDepositAmountSet // Event containing the contract specifics and raw log

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
func (it *CBJGatewayMinimumDepositAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayMinimumDepositAmountSet)
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
		it.Event = new(CBJGatewayMinimumDepositAmountSet)
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
func (it *CBJGatewayMinimumDepositAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayMinimumDepositAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayMinimumDepositAmountSet represents a MinimumDepositAmountSet event raised by the CBJGateway contract.
type CBJGatewayMinimumDepositAmountSet struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinimumDepositAmountSet is a free log retrieval operation binding the contract event 0xe6e25add7363f8f8a40cbea9810d3115a33703b10972ef759104219b00657436.
//
// Solidity: event MinimumDepositAmountSet(uint256 indexed oldAmount, uint256 indexed newAmount)
func (_CBJGateway *CBJGatewayFilterer) FilterMinimumDepositAmountSet(opts *bind.FilterOpts, oldAmount []*big.Int, newAmount []*big.Int) (*CBJGatewayMinimumDepositAmountSetIterator, error) {

	var oldAmountRule []interface{}
	for _, oldAmountItem := range oldAmount {
		oldAmountRule = append(oldAmountRule, oldAmountItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "MinimumDepositAmountSet", oldAmountRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayMinimumDepositAmountSetIterator{contract: _CBJGateway.contract, event: "MinimumDepositAmountSet", logs: logs, sub: sub}, nil
}

// WatchMinimumDepositAmountSet is a free log subscription operation binding the contract event 0xe6e25add7363f8f8a40cbea9810d3115a33703b10972ef759104219b00657436.
//
// Solidity: event MinimumDepositAmountSet(uint256 indexed oldAmount, uint256 indexed newAmount)
func (_CBJGateway *CBJGatewayFilterer) WatchMinimumDepositAmountSet(opts *bind.WatchOpts, sink chan<- *CBJGatewayMinimumDepositAmountSet, oldAmount []*big.Int, newAmount []*big.Int) (event.Subscription, error) {

	var oldAmountRule []interface{}
	for _, oldAmountItem := range oldAmount {
		oldAmountRule = append(oldAmountRule, oldAmountItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "MinimumDepositAmountSet", oldAmountRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayMinimumDepositAmountSet)
				if err := _CBJGateway.contract.UnpackLog(event, "MinimumDepositAmountSet", log); err != nil {
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

// ParseMinimumDepositAmountSet is a log parse operation binding the contract event 0xe6e25add7363f8f8a40cbea9810d3115a33703b10972ef759104219b00657436.
//
// Solidity: event MinimumDepositAmountSet(uint256 indexed oldAmount, uint256 indexed newAmount)
func (_CBJGateway *CBJGatewayFilterer) ParseMinimumDepositAmountSet(log types.Log) (*CBJGatewayMinimumDepositAmountSet, error) {
	event := new(CBJGatewayMinimumDepositAmountSet)
	if err := _CBJGateway.contract.UnpackLog(event, "MinimumDepositAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayMinimumWithdrawalAmountSetIterator is returned from FilterMinimumWithdrawalAmountSet and is used to iterate over the raw logs and unpacked data for MinimumWithdrawalAmountSet events raised by the CBJGateway contract.
type CBJGatewayMinimumWithdrawalAmountSetIterator struct {
	Event *CBJGatewayMinimumWithdrawalAmountSet // Event containing the contract specifics and raw log

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
func (it *CBJGatewayMinimumWithdrawalAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayMinimumWithdrawalAmountSet)
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
		it.Event = new(CBJGatewayMinimumWithdrawalAmountSet)
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
func (it *CBJGatewayMinimumWithdrawalAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayMinimumWithdrawalAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayMinimumWithdrawalAmountSet represents a MinimumWithdrawalAmountSet event raised by the CBJGateway contract.
type CBJGatewayMinimumWithdrawalAmountSet struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinimumWithdrawalAmountSet is a free log retrieval operation binding the contract event 0x4aa0fb18e66f6e36ba65fa3a8c4808fce6914782d5b67cb0adae132b3fd1ad9f.
//
// Solidity: event MinimumWithdrawalAmountSet(uint256 indexed oldAmount, uint256 indexed newAmount)
func (_CBJGateway *CBJGatewayFilterer) FilterMinimumWithdrawalAmountSet(opts *bind.FilterOpts, oldAmount []*big.Int, newAmount []*big.Int) (*CBJGatewayMinimumWithdrawalAmountSetIterator, error) {

	var oldAmountRule []interface{}
	for _, oldAmountItem := range oldAmount {
		oldAmountRule = append(oldAmountRule, oldAmountItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "MinimumWithdrawalAmountSet", oldAmountRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayMinimumWithdrawalAmountSetIterator{contract: _CBJGateway.contract, event: "MinimumWithdrawalAmountSet", logs: logs, sub: sub}, nil
}

// WatchMinimumWithdrawalAmountSet is a free log subscription operation binding the contract event 0x4aa0fb18e66f6e36ba65fa3a8c4808fce6914782d5b67cb0adae132b3fd1ad9f.
//
// Solidity: event MinimumWithdrawalAmountSet(uint256 indexed oldAmount, uint256 indexed newAmount)
func (_CBJGateway *CBJGatewayFilterer) WatchMinimumWithdrawalAmountSet(opts *bind.WatchOpts, sink chan<- *CBJGatewayMinimumWithdrawalAmountSet, oldAmount []*big.Int, newAmount []*big.Int) (event.Subscription, error) {

	var oldAmountRule []interface{}
	for _, oldAmountItem := range oldAmount {
		oldAmountRule = append(oldAmountRule, oldAmountItem)
	}
	var newAmountRule []interface{}
	for _, newAmountItem := range newAmount {
		newAmountRule = append(newAmountRule, newAmountItem)
	}

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "MinimumWithdrawalAmountSet", oldAmountRule, newAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayMinimumWithdrawalAmountSet)
				if err := _CBJGateway.contract.UnpackLog(event, "MinimumWithdrawalAmountSet", log); err != nil {
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

// ParseMinimumWithdrawalAmountSet is a log parse operation binding the contract event 0x4aa0fb18e66f6e36ba65fa3a8c4808fce6914782d5b67cb0adae132b3fd1ad9f.
//
// Solidity: event MinimumWithdrawalAmountSet(uint256 indexed oldAmount, uint256 indexed newAmount)
func (_CBJGateway *CBJGatewayFilterer) ParseMinimumWithdrawalAmountSet(log types.Log) (*CBJGatewayMinimumWithdrawalAmountSet, error) {
	event := new(CBJGatewayMinimumWithdrawalAmountSet)
	if err := _CBJGateway.contract.UnpackLog(event, "MinimumWithdrawalAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayPendingDepositIterator is returned from FilterPendingDeposit and is used to iterate over the raw logs and unpacked data for PendingDeposit events raised by the CBJGateway contract.
type CBJGatewayPendingDepositIterator struct {
	Event *CBJGatewayPendingDeposit // Event containing the contract specifics and raw log

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
func (it *CBJGatewayPendingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayPendingDeposit)
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
		it.Event = new(CBJGatewayPendingDeposit)
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
func (it *CBJGatewayPendingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayPendingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayPendingDeposit represents a PendingDeposit event raised by the CBJGateway contract.
type CBJGatewayPendingDeposit struct {
	OperationId          [32]byte
	User                 common.Address
	UsdcAmount           *big.Int
	PendingCbjUSDCAmount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPendingDeposit is a free log retrieval operation binding the contract event 0xed1ef86cea829cd3f9a3e1dc8c5e00d75e12c3d56d7b0615af8d6293e2d3f1fb.
//
// Solidity: event PendingDeposit(bytes32 indexed operationId, address indexed user, uint256 usdcAmount, uint256 pendingCbjUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) FilterPendingDeposit(opts *bind.FilterOpts, operationId [][32]byte, user []common.Address) (*CBJGatewayPendingDepositIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "PendingDeposit", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayPendingDepositIterator{contract: _CBJGateway.contract, event: "PendingDeposit", logs: logs, sub: sub}, nil
}

// WatchPendingDeposit is a free log subscription operation binding the contract event 0xed1ef86cea829cd3f9a3e1dc8c5e00d75e12c3d56d7b0615af8d6293e2d3f1fb.
//
// Solidity: event PendingDeposit(bytes32 indexed operationId, address indexed user, uint256 usdcAmount, uint256 pendingCbjUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) WatchPendingDeposit(opts *bind.WatchOpts, sink chan<- *CBJGatewayPendingDeposit, operationId [][32]byte, user []common.Address) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "PendingDeposit", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayPendingDeposit)
				if err := _CBJGateway.contract.UnpackLog(event, "PendingDeposit", log); err != nil {
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

// ParsePendingDeposit is a log parse operation binding the contract event 0xed1ef86cea829cd3f9a3e1dc8c5e00d75e12c3d56d7b0615af8d6293e2d3f1fb.
//
// Solidity: event PendingDeposit(bytes32 indexed operationId, address indexed user, uint256 usdcAmount, uint256 pendingCbjUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) ParsePendingDeposit(log types.Log) (*CBJGatewayPendingDeposit, error) {
	event := new(CBJGatewayPendingDeposit)
	if err := _CBJGateway.contract.UnpackLog(event, "PendingDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayPendingWithdrawIterator is returned from FilterPendingWithdraw and is used to iterate over the raw logs and unpacked data for PendingWithdraw events raised by the CBJGateway contract.
type CBJGatewayPendingWithdrawIterator struct {
	Event *CBJGatewayPendingWithdraw // Event containing the contract specifics and raw log

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
func (it *CBJGatewayPendingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayPendingWithdraw)
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
		it.Event = new(CBJGatewayPendingWithdraw)
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
func (it *CBJGatewayPendingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayPendingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayPendingWithdraw represents a PendingWithdraw event raised by the CBJGateway contract.
type CBJGatewayPendingWithdraw struct {
	OperationId       [32]byte
	User              common.Address
	CbjUSDCAmount     *big.Int
	PendingUSDCAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPendingWithdraw is a free log retrieval operation binding the contract event 0x8b65ad5f6693cce066a491cc8e5e1055aa66f16ab75f392ca2f1825be5c47204.
//
// Solidity: event PendingWithdraw(bytes32 indexed operationId, address indexed user, uint256 cbjUSDCAmount, uint256 pendingUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) FilterPendingWithdraw(opts *bind.FilterOpts, operationId [][32]byte, user []common.Address) (*CBJGatewayPendingWithdrawIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "PendingWithdraw", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayPendingWithdrawIterator{contract: _CBJGateway.contract, event: "PendingWithdraw", logs: logs, sub: sub}, nil
}

// WatchPendingWithdraw is a free log subscription operation binding the contract event 0x8b65ad5f6693cce066a491cc8e5e1055aa66f16ab75f392ca2f1825be5c47204.
//
// Solidity: event PendingWithdraw(bytes32 indexed operationId, address indexed user, uint256 cbjUSDCAmount, uint256 pendingUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) WatchPendingWithdraw(opts *bind.WatchOpts, sink chan<- *CBJGatewayPendingWithdraw, operationId [][32]byte, user []common.Address) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "PendingWithdraw", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayPendingWithdraw)
				if err := _CBJGateway.contract.UnpackLog(event, "PendingWithdraw", log); err != nil {
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

// ParsePendingWithdraw is a log parse operation binding the contract event 0x8b65ad5f6693cce066a491cc8e5e1055aa66f16ab75f392ca2f1825be5c47204.
//
// Solidity: event PendingWithdraw(bytes32 indexed operationId, address indexed user, uint256 cbjUSDCAmount, uint256 pendingUSDCAmount)
func (_CBJGateway *CBJGatewayFilterer) ParsePendingWithdraw(log types.Log) (*CBJGatewayPendingWithdraw, error) {
	event := new(CBJGatewayPendingWithdraw)
	if err := _CBJGateway.contract.UnpackLog(event, "PendingWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CBJGateway contract.
type CBJGatewayRoleAdminChangedIterator struct {
	Event *CBJGatewayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CBJGatewayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayRoleAdminChanged)
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
		it.Event = new(CBJGatewayRoleAdminChanged)
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
func (it *CBJGatewayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayRoleAdminChanged represents a RoleAdminChanged event raised by the CBJGateway contract.
type CBJGatewayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CBJGateway *CBJGatewayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CBJGatewayRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayRoleAdminChangedIterator{contract: _CBJGateway.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CBJGateway *CBJGatewayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CBJGatewayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayRoleAdminChanged)
				if err := _CBJGateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CBJGateway *CBJGatewayFilterer) ParseRoleAdminChanged(log types.Log) (*CBJGatewayRoleAdminChanged, error) {
	event := new(CBJGatewayRoleAdminChanged)
	if err := _CBJGateway.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CBJGateway contract.
type CBJGatewayRoleGrantedIterator struct {
	Event *CBJGatewayRoleGranted // Event containing the contract specifics and raw log

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
func (it *CBJGatewayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayRoleGranted)
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
		it.Event = new(CBJGatewayRoleGranted)
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
func (it *CBJGatewayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayRoleGranted represents a RoleGranted event raised by the CBJGateway contract.
type CBJGatewayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJGateway *CBJGatewayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CBJGatewayRoleGrantedIterator, error) {

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

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayRoleGrantedIterator{contract: _CBJGateway.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJGateway *CBJGatewayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CBJGatewayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayRoleGranted)
				if err := _CBJGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CBJGateway *CBJGatewayFilterer) ParseRoleGranted(log types.Log) (*CBJGatewayRoleGranted, error) {
	event := new(CBJGatewayRoleGranted)
	if err := _CBJGateway.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CBJGateway contract.
type CBJGatewayRoleRevokedIterator struct {
	Event *CBJGatewayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CBJGatewayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayRoleRevoked)
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
		it.Event = new(CBJGatewayRoleRevoked)
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
func (it *CBJGatewayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayRoleRevoked represents a RoleRevoked event raised by the CBJGateway contract.
type CBJGatewayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJGateway *CBJGatewayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CBJGatewayRoleRevokedIterator, error) {

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

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayRoleRevokedIterator{contract: _CBJGateway.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CBJGateway *CBJGatewayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CBJGatewayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayRoleRevoked)
				if err := _CBJGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CBJGateway *CBJGatewayFilterer) ParseRoleRevoked(log types.Log) (*CBJGatewayRoleRevoked, error) {
	event := new(CBJGatewayRoleRevoked)
	if err := _CBJGateway.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayWithdrawalProcessedIterator is returned from FilterWithdrawalProcessed and is used to iterate over the raw logs and unpacked data for WithdrawalProcessed events raised by the CBJGateway contract.
type CBJGatewayWithdrawalProcessedIterator struct {
	Event *CBJGatewayWithdrawalProcessed // Event containing the contract specifics and raw log

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
func (it *CBJGatewayWithdrawalProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayWithdrawalProcessed)
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
		it.Event = new(CBJGatewayWithdrawalProcessed)
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
func (it *CBJGatewayWithdrawalProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayWithdrawalProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayWithdrawalProcessed represents a WithdrawalProcessed event raised by the CBJGateway contract.
type CBJGatewayWithdrawalProcessed struct {
	OperationId [32]byte
	User        common.Address
	UsdcAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalProcessed is a free log retrieval operation binding the contract event 0x7bc99555a3551f7af5c9766d73b50823d8cc5709a24fe9a20c0d6a42483c4556.
//
// Solidity: event WithdrawalProcessed(bytes32 indexed operationId, address indexed user, uint256 usdcAmount)
func (_CBJGateway *CBJGatewayFilterer) FilterWithdrawalProcessed(opts *bind.FilterOpts, operationId [][32]byte, user []common.Address) (*CBJGatewayWithdrawalProcessedIterator, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "WithdrawalProcessed", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &CBJGatewayWithdrawalProcessedIterator{contract: _CBJGateway.contract, event: "WithdrawalProcessed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalProcessed is a free log subscription operation binding the contract event 0x7bc99555a3551f7af5c9766d73b50823d8cc5709a24fe9a20c0d6a42483c4556.
//
// Solidity: event WithdrawalProcessed(bytes32 indexed operationId, address indexed user, uint256 usdcAmount)
func (_CBJGateway *CBJGatewayFilterer) WatchWithdrawalProcessed(opts *bind.WatchOpts, sink chan<- *CBJGatewayWithdrawalProcessed, operationId [][32]byte, user []common.Address) (event.Subscription, error) {

	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "WithdrawalProcessed", operationIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayWithdrawalProcessed)
				if err := _CBJGateway.contract.UnpackLog(event, "WithdrawalProcessed", log); err != nil {
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

// ParseWithdrawalProcessed is a log parse operation binding the contract event 0x7bc99555a3551f7af5c9766d73b50823d8cc5709a24fe9a20c0d6a42483c4556.
//
// Solidity: event WithdrawalProcessed(bytes32 indexed operationId, address indexed user, uint256 usdcAmount)
func (_CBJGateway *CBJGatewayFilterer) ParseWithdrawalProcessed(log types.Log) (*CBJGatewayWithdrawalProcessed, error) {
	event := new(CBJGatewayWithdrawalProcessed)
	if err := _CBJGateway.contract.UnpackLog(event, "WithdrawalProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayWithdrawalsPausedIterator is returned from FilterWithdrawalsPaused and is used to iterate over the raw logs and unpacked data for WithdrawalsPaused events raised by the CBJGateway contract.
type CBJGatewayWithdrawalsPausedIterator struct {
	Event *CBJGatewayWithdrawalsPaused // Event containing the contract specifics and raw log

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
func (it *CBJGatewayWithdrawalsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayWithdrawalsPaused)
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
		it.Event = new(CBJGatewayWithdrawalsPaused)
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
func (it *CBJGatewayWithdrawalsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayWithdrawalsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayWithdrawalsPaused represents a WithdrawalsPaused event raised by the CBJGateway contract.
type CBJGatewayWithdrawalsPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsPaused is a free log retrieval operation binding the contract event 0x6022a9e759c95aad593773b7a47586ff34cddc74d34ea6361f64c5bac98cf294.
//
// Solidity: event WithdrawalsPaused()
func (_CBJGateway *CBJGatewayFilterer) FilterWithdrawalsPaused(opts *bind.FilterOpts) (*CBJGatewayWithdrawalsPausedIterator, error) {

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "WithdrawalsPaused")
	if err != nil {
		return nil, err
	}
	return &CBJGatewayWithdrawalsPausedIterator{contract: _CBJGateway.contract, event: "WithdrawalsPaused", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsPaused is a free log subscription operation binding the contract event 0x6022a9e759c95aad593773b7a47586ff34cddc74d34ea6361f64c5bac98cf294.
//
// Solidity: event WithdrawalsPaused()
func (_CBJGateway *CBJGatewayFilterer) WatchWithdrawalsPaused(opts *bind.WatchOpts, sink chan<- *CBJGatewayWithdrawalsPaused) (event.Subscription, error) {

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "WithdrawalsPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayWithdrawalsPaused)
				if err := _CBJGateway.contract.UnpackLog(event, "WithdrawalsPaused", log); err != nil {
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

// ParseWithdrawalsPaused is a log parse operation binding the contract event 0x6022a9e759c95aad593773b7a47586ff34cddc74d34ea6361f64c5bac98cf294.
//
// Solidity: event WithdrawalsPaused()
func (_CBJGateway *CBJGatewayFilterer) ParseWithdrawalsPaused(log types.Log) (*CBJGatewayWithdrawalsPaused, error) {
	event := new(CBJGatewayWithdrawalsPaused)
	if err := _CBJGateway.contract.UnpackLog(event, "WithdrawalsPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CBJGatewayWithdrawalsUnpausedIterator is returned from FilterWithdrawalsUnpaused and is used to iterate over the raw logs and unpacked data for WithdrawalsUnpaused events raised by the CBJGateway contract.
type CBJGatewayWithdrawalsUnpausedIterator struct {
	Event *CBJGatewayWithdrawalsUnpaused // Event containing the contract specifics and raw log

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
func (it *CBJGatewayWithdrawalsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CBJGatewayWithdrawalsUnpaused)
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
		it.Event = new(CBJGatewayWithdrawalsUnpaused)
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
func (it *CBJGatewayWithdrawalsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CBJGatewayWithdrawalsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CBJGatewayWithdrawalsUnpaused represents a WithdrawalsUnpaused event raised by the CBJGateway contract.
type CBJGatewayWithdrawalsUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsUnpaused is a free log retrieval operation binding the contract event 0x73f109ff397323e276e38b27fa4bf2a3d7bc07fb71be9fccbcd541e39140ea3f.
//
// Solidity: event WithdrawalsUnpaused()
func (_CBJGateway *CBJGatewayFilterer) FilterWithdrawalsUnpaused(opts *bind.FilterOpts) (*CBJGatewayWithdrawalsUnpausedIterator, error) {

	logs, sub, err := _CBJGateway.contract.FilterLogs(opts, "WithdrawalsUnpaused")
	if err != nil {
		return nil, err
	}
	return &CBJGatewayWithdrawalsUnpausedIterator{contract: _CBJGateway.contract, event: "WithdrawalsUnpaused", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsUnpaused is a free log subscription operation binding the contract event 0x73f109ff397323e276e38b27fa4bf2a3d7bc07fb71be9fccbcd541e39140ea3f.
//
// Solidity: event WithdrawalsUnpaused()
func (_CBJGateway *CBJGatewayFilterer) WatchWithdrawalsUnpaused(opts *bind.WatchOpts, sink chan<- *CBJGatewayWithdrawalsUnpaused) (event.Subscription, error) {

	logs, sub, err := _CBJGateway.contract.WatchLogs(opts, "WithdrawalsUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CBJGatewayWithdrawalsUnpaused)
				if err := _CBJGateway.contract.UnpackLog(event, "WithdrawalsUnpaused", log); err != nil {
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

// ParseWithdrawalsUnpaused is a log parse operation binding the contract event 0x73f109ff397323e276e38b27fa4bf2a3d7bc07fb71be9fccbcd541e39140ea3f.
//
// Solidity: event WithdrawalsUnpaused()
func (_CBJGateway *CBJGatewayFilterer) ParseWithdrawalsUnpaused(log types.Log) (*CBJGatewayWithdrawalsUnpaused, error) {
	event := new(CBJGatewayWithdrawalsUnpaused)
	if err := _CBJGateway.contract.UnpackLog(event, "WithdrawalsUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
