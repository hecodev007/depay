// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PayABI is the input ABI used to generate the binding from.
const PayABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"merchant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"product\",\"type\":\"uint256\"}],\"name\":\"CancelSubScribe\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MerchantWithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"merchant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"PayOrderEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"merchant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"product\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"SubScribe\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"merchant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"SubScribePay\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_point\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"addMerchant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_point\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"addMerchantBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_product\",\"type\":\"uint256\"}],\"name\":\"cancelSubScribe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getMerchantBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getMerchantInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"}],\"name\":\"getTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"merchantInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"merchantId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"merchantWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_orderid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_merchantId\",\"type\":\"uint256\"}],\"name\":\"payOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_point\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_denominator\",\"type\":\"uint256\"}],\"name\":\"setMerchant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenERC20\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"power\",\"type\":\"bool\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subScribe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"usdtAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_product\",\"type\":\"uint256\"}],\"name\":\"subScribePay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_merchant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_product\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"subScribes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Pay is an auto generated Go binding around an Ethereum contract.
type Pay struct {
	PayCaller     // Read-only binding to the contract
	PayTransactor // Write-only binding to the contract
	PayFilterer   // Log filterer for contract events
}

// PayCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaySession struct {
	Contract     *Pay              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayCallerSession struct {
	Contract *PayCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayTransactorSession struct {
	Contract     *PayTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayRaw struct {
	Contract *Pay // Generic contract binding to access the raw methods on
}

// PayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayCallerRaw struct {
	Contract *PayCaller // Generic read-only contract binding to access the raw methods on
}

// PayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayTransactorRaw struct {
	Contract *PayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPay creates a new instance of Pay, bound to a specific deployed contract.
func NewPay(address common.Address, backend bind.ContractBackend) (*Pay, error) {
	contract, err := bindPay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pay{PayCaller: PayCaller{contract: contract}, PayTransactor: PayTransactor{contract: contract}, PayFilterer: PayFilterer{contract: contract}}, nil
}

// NewPayCaller creates a new read-only instance of Pay, bound to a specific deployed contract.
func NewPayCaller(address common.Address, caller bind.ContractCaller) (*PayCaller, error) {
	contract, err := bindPay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayCaller{contract: contract}, nil
}

// NewPayTransactor creates a new write-only instance of Pay, bound to a specific deployed contract.
func NewPayTransactor(address common.Address, transactor bind.ContractTransactor) (*PayTransactor, error) {
	contract, err := bindPay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayTransactor{contract: contract}, nil
}

// NewPayFilterer creates a new log filterer instance of Pay, bound to a specific deployed contract.
func NewPayFilterer(address common.Address, filterer bind.ContractFilterer) (*PayFilterer, error) {
	contract, err := bindPay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayFilterer{contract: contract}, nil
}

// bindPay binds a generic wrapper to an already deployed contract.
func bindPay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pay *PayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pay.Contract.PayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pay *PayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.Contract.PayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pay *PayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pay.Contract.PayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pay *PayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pay *PayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pay *PayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pay.Contract.contract.Transact(opts, method, params...)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Pay *PayCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Pay *PaySession) FACTORY() (common.Address, error) {
	return _Pay.Contract.FACTORY(&_Pay.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_Pay *PayCallerSession) FACTORY() (common.Address, error) {
	return _Pay.Contract.FACTORY(&_Pay.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Pay *PayCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Pay *PaySession) WETH() (common.Address, error) {
	return _Pay.Contract.WETH(&_Pay.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Pay *PayCallerSession) WETH() (common.Address, error) {
	return _Pay.Contract.WETH(&_Pay.CallOpts)
}

// Adminer is a free data retrieval call binding the contract method 0x57dc5d9d.
//
// Solidity: function adminer() view returns(address)
func (_Pay *PayCaller) Adminer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "adminer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Adminer is a free data retrieval call binding the contract method 0x57dc5d9d.
//
// Solidity: function adminer() view returns(address)
func (_Pay *PaySession) Adminer() (common.Address, error) {
	return _Pay.Contract.Adminer(&_Pay.CallOpts)
}

// Adminer is a free data retrieval call binding the contract method 0x57dc5d9d.
//
// Solidity: function adminer() view returns(address)
func (_Pay *PayCallerSession) Adminer() (common.Address, error) {
	return _Pay.Contract.Adminer(&_Pay.CallOpts)
}

// GetMerchantBalance is a free data retrieval call binding the contract method 0x1628cfa4.
//
// Solidity: function getMerchantBalance(address _merchant, address token) view returns(uint256)
func (_Pay *PayCaller) GetMerchantBalance(opts *bind.CallOpts, _merchant common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "getMerchantBalance", _merchant, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerchantBalance is a free data retrieval call binding the contract method 0x1628cfa4.
//
// Solidity: function getMerchantBalance(address _merchant, address token) view returns(uint256)
func (_Pay *PaySession) GetMerchantBalance(_merchant common.Address, token common.Address) (*big.Int, error) {
	return _Pay.Contract.GetMerchantBalance(&_Pay.CallOpts, _merchant, token)
}

// GetMerchantBalance is a free data retrieval call binding the contract method 0x1628cfa4.
//
// Solidity: function getMerchantBalance(address _merchant, address token) view returns(uint256)
func (_Pay *PayCallerSession) GetMerchantBalance(_merchant common.Address, token common.Address) (*big.Int, error) {
	return _Pay.Contract.GetMerchantBalance(&_Pay.CallOpts, _merchant, token)
}

// GetMerchantInfo is a free data retrieval call binding the contract method 0x6862724a.
//
// Solidity: function getMerchantInfo(address _merchant, address token) view returns(uint256, uint256, uint256)
func (_Pay *PayCaller) GetMerchantInfo(opts *bind.CallOpts, _merchant common.Address, token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "getMerchantInfo", _merchant, token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetMerchantInfo is a free data retrieval call binding the contract method 0x6862724a.
//
// Solidity: function getMerchantInfo(address _merchant, address token) view returns(uint256, uint256, uint256)
func (_Pay *PaySession) GetMerchantInfo(_merchant common.Address, token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Pay.Contract.GetMerchantInfo(&_Pay.CallOpts, _merchant, token)
}

// GetMerchantInfo is a free data retrieval call binding the contract method 0x6862724a.
//
// Solidity: function getMerchantInfo(address _merchant, address token) view returns(uint256, uint256, uint256)
func (_Pay *PayCallerSession) GetMerchantInfo(_merchant common.Address, token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Pay.Contract.GetMerchantInfo(&_Pay.CallOpts, _merchant, token)
}

// GetTokenAmount is a free data retrieval call binding the contract method 0x115ece4c.
//
// Solidity: function getTokenAmount(address _token, uint256 usdtAmount) view returns(uint256)
func (_Pay *PayCaller) GetTokenAmount(opts *bind.CallOpts, _token common.Address, usdtAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "getTokenAmount", _token, usdtAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmount is a free data retrieval call binding the contract method 0x115ece4c.
//
// Solidity: function getTokenAmount(address _token, uint256 usdtAmount) view returns(uint256)
func (_Pay *PaySession) GetTokenAmount(_token common.Address, usdtAmount *big.Int) (*big.Int, error) {
	return _Pay.Contract.GetTokenAmount(&_Pay.CallOpts, _token, usdtAmount)
}

// GetTokenAmount is a free data retrieval call binding the contract method 0x115ece4c.
//
// Solidity: function getTokenAmount(address _token, uint256 usdtAmount) view returns(uint256)
func (_Pay *PayCallerSession) GetTokenAmount(_token common.Address, usdtAmount *big.Int) (*big.Int, error) {
	return _Pay.Contract.GetTokenAmount(&_Pay.CallOpts, _token, usdtAmount)
}

// MerchantInfo is a free data retrieval call binding the contract method 0x8972e5b5.
//
// Solidity: function merchantInfo(address , address ) view returns(uint256 merchantId, uint256 feePoint, uint256 feeDenominator, uint256 balance)
func (_Pay *PayCaller) MerchantInfo(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	MerchantId     *big.Int
	FeePoint       *big.Int
	FeeDenominator *big.Int
	Balance        *big.Int
}, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "merchantInfo", arg0, arg1)

	outstruct := new(struct {
		MerchantId     *big.Int
		FeePoint       *big.Int
		FeeDenominator *big.Int
		Balance        *big.Int
	})

	outstruct.MerchantId = out[0].(*big.Int)
	outstruct.FeePoint = out[1].(*big.Int)
	outstruct.FeeDenominator = out[2].(*big.Int)
	outstruct.Balance = out[3].(*big.Int)

	return *outstruct, err

}

// MerchantInfo is a free data retrieval call binding the contract method 0x8972e5b5.
//
// Solidity: function merchantInfo(address , address ) view returns(uint256 merchantId, uint256 feePoint, uint256 feeDenominator, uint256 balance)
func (_Pay *PaySession) MerchantInfo(arg0 common.Address, arg1 common.Address) (struct {
	MerchantId     *big.Int
	FeePoint       *big.Int
	FeeDenominator *big.Int
	Balance        *big.Int
}, error) {
	return _Pay.Contract.MerchantInfo(&_Pay.CallOpts, arg0, arg1)
}

// MerchantInfo is a free data retrieval call binding the contract method 0x8972e5b5.
//
// Solidity: function merchantInfo(address , address ) view returns(uint256 merchantId, uint256 feePoint, uint256 feeDenominator, uint256 balance)
func (_Pay *PayCallerSession) MerchantInfo(arg0 common.Address, arg1 common.Address) (struct {
	MerchantId     *big.Int
	FeePoint       *big.Int
	FeeDenominator *big.Int
	Balance        *big.Int
}, error) {
	return _Pay.Contract.MerchantInfo(&_Pay.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pay *PayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pay *PaySession) Owner() (common.Address, error) {
	return _Pay.Contract.Owner(&_Pay.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pay *PayCallerSession) Owner() (common.Address, error) {
	return _Pay.Contract.Owner(&_Pay.CallOpts)
}

// SubScribe is a free data retrieval call binding the contract method 0xde0ca2ac.
//
// Solidity: function subScribe(address , address , uint256 ) view returns(uint256)
func (_Pay *PayCaller) SubScribe(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "subScribe", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubScribe is a free data retrieval call binding the contract method 0xde0ca2ac.
//
// Solidity: function subScribe(address , address , uint256 ) view returns(uint256)
func (_Pay *PaySession) SubScribe(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Pay.Contract.SubScribe(&_Pay.CallOpts, arg0, arg1, arg2)
}

// SubScribe is a free data retrieval call binding the contract method 0xde0ca2ac.
//
// Solidity: function subScribe(address , address , uint256 ) view returns(uint256)
func (_Pay *PayCallerSession) SubScribe(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*big.Int, error) {
	return _Pay.Contract.SubScribe(&_Pay.CallOpts, arg0, arg1, arg2)
}

// TokenERC20 is a free data retrieval call binding the contract method 0x9d143e8e.
//
// Solidity: function tokenERC20() view returns(address)
func (_Pay *PayCaller) TokenERC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "tokenERC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenERC20 is a free data retrieval call binding the contract method 0x9d143e8e.
//
// Solidity: function tokenERC20() view returns(address)
func (_Pay *PaySession) TokenERC20() (common.Address, error) {
	return _Pay.Contract.TokenERC20(&_Pay.CallOpts)
}

// TokenERC20 is a free data retrieval call binding the contract method 0x9d143e8e.
//
// Solidity: function tokenERC20() view returns(address)
func (_Pay *PayCallerSession) TokenERC20() (common.Address, error) {
	return _Pay.Contract.TokenERC20(&_Pay.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Pay *PayCaller) UsdtToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pay.contract.Call(opts, &out, "usdtToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Pay *PaySession) UsdtToken() (common.Address, error) {
	return _Pay.Contract.UsdtToken(&_Pay.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Pay *PayCallerSession) UsdtToken() (common.Address, error) {
	return _Pay.Contract.UsdtToken(&_Pay.CallOpts)
}

// AddMerchant is a paid mutator transaction binding the contract method 0x28619bc3.
//
// Solidity: function addMerchant(address _merchant, uint256 _id, uint256 _point, uint256 _denominator) returns()
func (_Pay *PayTransactor) AddMerchant(opts *bind.TransactOpts, _merchant common.Address, _id *big.Int, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "addMerchant", _merchant, _id, _point, _denominator)
}

// AddMerchant is a paid mutator transaction binding the contract method 0x28619bc3.
//
// Solidity: function addMerchant(address _merchant, uint256 _id, uint256 _point, uint256 _denominator) returns()
func (_Pay *PaySession) AddMerchant(_merchant common.Address, _id *big.Int, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.AddMerchant(&_Pay.TransactOpts, _merchant, _id, _point, _denominator)
}

// AddMerchant is a paid mutator transaction binding the contract method 0x28619bc3.
//
// Solidity: function addMerchant(address _merchant, uint256 _id, uint256 _point, uint256 _denominator) returns()
func (_Pay *PayTransactorSession) AddMerchant(_merchant common.Address, _id *big.Int, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.AddMerchant(&_Pay.TransactOpts, _merchant, _id, _point, _denominator)
}

// AddMerchantBase is a paid mutator transaction binding the contract method 0x4b671f97.
//
// Solidity: function addMerchantBase(address _merchant, uint256 _id, address _token, uint256 _point, uint256 _denominator) returns()
func (_Pay *PayTransactor) AddMerchantBase(opts *bind.TransactOpts, _merchant common.Address, _id *big.Int, _token common.Address, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "addMerchantBase", _merchant, _id, _token, _point, _denominator)
}

// AddMerchantBase is a paid mutator transaction binding the contract method 0x4b671f97.
//
// Solidity: function addMerchantBase(address _merchant, uint256 _id, address _token, uint256 _point, uint256 _denominator) returns()
func (_Pay *PaySession) AddMerchantBase(_merchant common.Address, _id *big.Int, _token common.Address, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.AddMerchantBase(&_Pay.TransactOpts, _merchant, _id, _token, _point, _denominator)
}

// AddMerchantBase is a paid mutator transaction binding the contract method 0x4b671f97.
//
// Solidity: function addMerchantBase(address _merchant, uint256 _id, address _token, uint256 _point, uint256 _denominator) returns()
func (_Pay *PayTransactorSession) AddMerchantBase(_merchant common.Address, _id *big.Int, _token common.Address, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.AddMerchantBase(&_Pay.TransactOpts, _merchant, _id, _token, _point, _denominator)
}

// CancelSubScribe is a paid mutator transaction binding the contract method 0x070e5f84.
//
// Solidity: function cancelSubScribe(address _merchant, uint256 _product) returns()
func (_Pay *PayTransactor) CancelSubScribe(opts *bind.TransactOpts, _merchant common.Address, _product *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "cancelSubScribe", _merchant, _product)
}

// CancelSubScribe is a paid mutator transaction binding the contract method 0x070e5f84.
//
// Solidity: function cancelSubScribe(address _merchant, uint256 _product) returns()
func (_Pay *PaySession) CancelSubScribe(_merchant common.Address, _product *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.CancelSubScribe(&_Pay.TransactOpts, _merchant, _product)
}

// CancelSubScribe is a paid mutator transaction binding the contract method 0x070e5f84.
//
// Solidity: function cancelSubScribe(address _merchant, uint256 _product) returns()
func (_Pay *PayTransactorSession) CancelSubScribe(_merchant common.Address, _product *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.CancelSubScribe(&_Pay.TransactOpts, _merchant, _product)
}

// MerchantWithdraw is a paid mutator transaction binding the contract method 0xb42a1c70.
//
// Solidity: function merchantWithdraw(address token) returns()
func (_Pay *PayTransactor) MerchantWithdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "merchantWithdraw", token)
}

// MerchantWithdraw is a paid mutator transaction binding the contract method 0xb42a1c70.
//
// Solidity: function merchantWithdraw(address token) returns()
func (_Pay *PaySession) MerchantWithdraw(token common.Address) (*types.Transaction, error) {
	return _Pay.Contract.MerchantWithdraw(&_Pay.TransactOpts, token)
}

// MerchantWithdraw is a paid mutator transaction binding the contract method 0xb42a1c70.
//
// Solidity: function merchantWithdraw(address token) returns()
func (_Pay *PayTransactorSession) MerchantWithdraw(token common.Address) (*types.Transaction, error) {
	return _Pay.Contract.MerchantWithdraw(&_Pay.TransactOpts, token)
}

// PayOrder is a paid mutator transaction binding the contract method 0xec62213b.
//
// Solidity: function payOrder(uint256 usdtAmount, uint256 tokenAmount, uint256 _orderid, address _token, address _merchant, uint256 _merchantId) payable returns()
func (_Pay *PayTransactor) PayOrder(opts *bind.TransactOpts, usdtAmount *big.Int, tokenAmount *big.Int, _orderid *big.Int, _token common.Address, _merchant common.Address, _merchantId *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "payOrder", usdtAmount, tokenAmount, _orderid, _token, _merchant, _merchantId)
}

// PayOrder is a paid mutator transaction binding the contract method 0xec62213b.
//
// Solidity: function payOrder(uint256 usdtAmount, uint256 tokenAmount, uint256 _orderid, address _token, address _merchant, uint256 _merchantId) payable returns()
func (_Pay *PaySession) PayOrder(usdtAmount *big.Int, tokenAmount *big.Int, _orderid *big.Int, _token common.Address, _merchant common.Address, _merchantId *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.PayOrder(&_Pay.TransactOpts, usdtAmount, tokenAmount, _orderid, _token, _merchant, _merchantId)
}

// PayOrder is a paid mutator transaction binding the contract method 0xec62213b.
//
// Solidity: function payOrder(uint256 usdtAmount, uint256 tokenAmount, uint256 _orderid, address _token, address _merchant, uint256 _merchantId) payable returns()
func (_Pay *PayTransactorSession) PayOrder(usdtAmount *big.Int, tokenAmount *big.Int, _orderid *big.Int, _token common.Address, _merchant common.Address, _merchantId *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.PayOrder(&_Pay.TransactOpts, usdtAmount, tokenAmount, _orderid, _token, _merchant, _merchantId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pay *PayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pay *PaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Pay.Contract.RenounceOwnership(&_Pay.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pay *PayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pay.Contract.RenounceOwnership(&_Pay.TransactOpts)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Pay *PayTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Pay *PaySession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Pay.Contract.SetAdmin(&_Pay.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Pay *PayTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Pay.Contract.SetAdmin(&_Pay.TransactOpts, _admin)
}

// SetMerchant is a paid mutator transaction binding the contract method 0x781dd33d.
//
// Solidity: function setMerchant(address _merchant, uint256 _point, uint256 _denominator) returns()
func (_Pay *PayTransactor) SetMerchant(opts *bind.TransactOpts, _merchant common.Address, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "setMerchant", _merchant, _point, _denominator)
}

// SetMerchant is a paid mutator transaction binding the contract method 0x781dd33d.
//
// Solidity: function setMerchant(address _merchant, uint256 _point, uint256 _denominator) returns()
func (_Pay *PaySession) SetMerchant(_merchant common.Address, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.SetMerchant(&_Pay.TransactOpts, _merchant, _point, _denominator)
}

// SetMerchant is a paid mutator transaction binding the contract method 0x781dd33d.
//
// Solidity: function setMerchant(address _merchant, uint256 _point, uint256 _denominator) returns()
func (_Pay *PayTransactorSession) SetMerchant(_merchant common.Address, _point *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.SetMerchant(&_Pay.TransactOpts, _merchant, _point, _denominator)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address _tokenERC20, bool power) returns()
func (_Pay *PayTransactor) SetToken(opts *bind.TransactOpts, _tokenERC20 common.Address, power bool) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "setToken", _tokenERC20, power)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address _tokenERC20, bool power) returns()
func (_Pay *PaySession) SetToken(_tokenERC20 common.Address, power bool) (*types.Transaction, error) {
	return _Pay.Contract.SetToken(&_Pay.TransactOpts, _tokenERC20, power)
}

// SetToken is a paid mutator transaction binding the contract method 0x3816a292.
//
// Solidity: function setToken(address _tokenERC20, bool power) returns()
func (_Pay *PayTransactorSession) SetToken(_tokenERC20 common.Address, power bool) (*types.Transaction, error) {
	return _Pay.Contract.SetToken(&_Pay.TransactOpts, _tokenERC20, power)
}

// SubScribePay is a paid mutator transaction binding the contract method 0xe5d0bef1.
//
// Solidity: function subScribePay(address _merchant, address _user, address _token, uint256 usdtAmount, uint256 _product) returns()
func (_Pay *PayTransactor) SubScribePay(opts *bind.TransactOpts, _merchant common.Address, _user common.Address, _token common.Address, usdtAmount *big.Int, _product *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "subScribePay", _merchant, _user, _token, usdtAmount, _product)
}

// SubScribePay is a paid mutator transaction binding the contract method 0xe5d0bef1.
//
// Solidity: function subScribePay(address _merchant, address _user, address _token, uint256 usdtAmount, uint256 _product) returns()
func (_Pay *PaySession) SubScribePay(_merchant common.Address, _user common.Address, _token common.Address, usdtAmount *big.Int, _product *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.SubScribePay(&_Pay.TransactOpts, _merchant, _user, _token, usdtAmount, _product)
}

// SubScribePay is a paid mutator transaction binding the contract method 0xe5d0bef1.
//
// Solidity: function subScribePay(address _merchant, address _user, address _token, uint256 usdtAmount, uint256 _product) returns()
func (_Pay *PayTransactorSession) SubScribePay(_merchant common.Address, _user common.Address, _token common.Address, usdtAmount *big.Int, _product *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.SubScribePay(&_Pay.TransactOpts, _merchant, _user, _token, usdtAmount, _product)
}

// SubScribes is a paid mutator transaction binding the contract method 0x0adb6dac.
//
// Solidity: function subScribes(address _merchant, uint256 _product, uint256 _period) returns()
func (_Pay *PayTransactor) SubScribes(opts *bind.TransactOpts, _merchant common.Address, _product *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "subScribes", _merchant, _product, _period)
}

// SubScribes is a paid mutator transaction binding the contract method 0x0adb6dac.
//
// Solidity: function subScribes(address _merchant, uint256 _product, uint256 _period) returns()
func (_Pay *PaySession) SubScribes(_merchant common.Address, _product *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.SubScribes(&_Pay.TransactOpts, _merchant, _product, _period)
}

// SubScribes is a paid mutator transaction binding the contract method 0x0adb6dac.
//
// Solidity: function subScribes(address _merchant, uint256 _product, uint256 _period) returns()
func (_Pay *PayTransactorSession) SubScribes(_merchant common.Address, _product *big.Int, _period *big.Int) (*types.Transaction, error) {
	return _Pay.Contract.SubScribes(&_Pay.TransactOpts, _merchant, _product, _period)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pay *PayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pay.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pay *PaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pay.Contract.TransferOwnership(&_Pay.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pay *PayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pay.Contract.TransferOwnership(&_Pay.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pay *PayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pay.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pay *PaySession) Receive() (*types.Transaction, error) {
	return _Pay.Contract.Receive(&_Pay.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pay *PayTransactorSession) Receive() (*types.Transaction, error) {
	return _Pay.Contract.Receive(&_Pay.TransactOpts)
}

// PayCancelSubScribeIterator is returned from FilterCancelSubScribe and is used to iterate over the raw logs and unpacked data for CancelSubScribe events raised by the Pay contract.
type PayCancelSubScribeIterator struct {
	Event *PayCancelSubScribe // Event containing the contract specifics and raw log

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
func (it *PayCancelSubScribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayCancelSubScribe)
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
		it.Event = new(PayCancelSubScribe)
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
func (it *PayCancelSubScribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayCancelSubScribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayCancelSubScribe represents a CancelSubScribe event raised by the Pay contract.
type PayCancelSubScribe struct {
	User     common.Address
	Merchant common.Address
	Product  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCancelSubScribe is a free log retrieval operation binding the contract event 0x56adf33f94be28a032c02bd196e107cc9f407ef76c6749e9da1eabae8646ac25.
//
// Solidity: event CancelSubScribe(address user, address merchant, uint256 product)
func (_Pay *PayFilterer) FilterCancelSubScribe(opts *bind.FilterOpts) (*PayCancelSubScribeIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "CancelSubScribe")
	if err != nil {
		return nil, err
	}
	return &PayCancelSubScribeIterator{contract: _Pay.contract, event: "CancelSubScribe", logs: logs, sub: sub}, nil
}

// WatchCancelSubScribe is a free log subscription operation binding the contract event 0x56adf33f94be28a032c02bd196e107cc9f407ef76c6749e9da1eabae8646ac25.
//
// Solidity: event CancelSubScribe(address user, address merchant, uint256 product)
func (_Pay *PayFilterer) WatchCancelSubScribe(opts *bind.WatchOpts, sink chan<- *PayCancelSubScribe) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "CancelSubScribe")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayCancelSubScribe)
				if err := _Pay.contract.UnpackLog(event, "CancelSubScribe", log); err != nil {
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

// ParseCancelSubScribe is a log parse operation binding the contract event 0x56adf33f94be28a032c02bd196e107cc9f407ef76c6749e9da1eabae8646ac25.
//
// Solidity: event CancelSubScribe(address user, address merchant, uint256 product)
func (_Pay *PayFilterer) ParseCancelSubScribe(log types.Log) (*PayCancelSubScribe, error) {
	event := new(PayCancelSubScribe)
	if err := _Pay.contract.UnpackLog(event, "CancelSubScribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayMerchantWithdrawEventIterator is returned from FilterMerchantWithdrawEvent and is used to iterate over the raw logs and unpacked data for MerchantWithdrawEvent events raised by the Pay contract.
type PayMerchantWithdrawEventIterator struct {
	Event *PayMerchantWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *PayMerchantWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayMerchantWithdrawEvent)
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
		it.Event = new(PayMerchantWithdrawEvent)
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
func (it *PayMerchantWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayMerchantWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayMerchantWithdrawEvent represents a MerchantWithdrawEvent event raised by the Pay contract.
type PayMerchantWithdrawEvent struct {
	User   common.Address
	Usdt   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMerchantWithdrawEvent is a free log retrieval operation binding the contract event 0x0de91e3a3f6e247268a4925d673990d2f12dce87116d9df5abf65a4852ee923f.
//
// Solidity: event MerchantWithdrawEvent(address user, address _usdt, uint256 amount)
func (_Pay *PayFilterer) FilterMerchantWithdrawEvent(opts *bind.FilterOpts) (*PayMerchantWithdrawEventIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "MerchantWithdrawEvent")
	if err != nil {
		return nil, err
	}
	return &PayMerchantWithdrawEventIterator{contract: _Pay.contract, event: "MerchantWithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchMerchantWithdrawEvent is a free log subscription operation binding the contract event 0x0de91e3a3f6e247268a4925d673990d2f12dce87116d9df5abf65a4852ee923f.
//
// Solidity: event MerchantWithdrawEvent(address user, address _usdt, uint256 amount)
func (_Pay *PayFilterer) WatchMerchantWithdrawEvent(opts *bind.WatchOpts, sink chan<- *PayMerchantWithdrawEvent) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "MerchantWithdrawEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayMerchantWithdrawEvent)
				if err := _Pay.contract.UnpackLog(event, "MerchantWithdrawEvent", log); err != nil {
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

// ParseMerchantWithdrawEvent is a log parse operation binding the contract event 0x0de91e3a3f6e247268a4925d673990d2f12dce87116d9df5abf65a4852ee923f.
//
// Solidity: event MerchantWithdrawEvent(address user, address _usdt, uint256 amount)
func (_Pay *PayFilterer) ParseMerchantWithdrawEvent(log types.Log) (*PayMerchantWithdrawEvent, error) {
	event := new(PayMerchantWithdrawEvent)
	if err := _Pay.contract.UnpackLog(event, "MerchantWithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pay contract.
type PayOwnershipTransferredIterator struct {
	Event *PayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayOwnershipTransferred)
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
		it.Event = new(PayOwnershipTransferred)
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
func (it *PayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayOwnershipTransferred represents a OwnershipTransferred event raised by the Pay contract.
type PayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pay *PayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pay.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PayOwnershipTransferredIterator{contract: _Pay.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pay *PayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pay.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayOwnershipTransferred)
				if err := _Pay.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pay *PayFilterer) ParseOwnershipTransferred(log types.Log) (*PayOwnershipTransferred, error) {
	event := new(PayOwnershipTransferred)
	if err := _Pay.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PayPayOrderEventIterator is returned from FilterPayOrderEvent and is used to iterate over the raw logs and unpacked data for PayOrderEvent events raised by the Pay contract.
type PayPayOrderEventIterator struct {
	Event *PayPayOrderEvent // Event containing the contract specifics and raw log

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
func (it *PayPayOrderEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PayPayOrderEvent)
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
		it.Event = new(PayPayOrderEvent)
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
func (it *PayPayOrderEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PayPayOrderEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PayPayOrderEvent represents a PayOrderEvent event raised by the Pay contract.
type PayPayOrderEvent struct {
	OrderId     *big.Int
	User        common.Address
	Merchant    common.Address
	PayToken    common.Address
	TokenAmount *big.Int
	PayAmount   *big.Int
	SwapAmount  *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPayOrderEvent is a free log retrieval operation binding the contract event 0x753d1526f8fcb265ab0f908316c22eaec5eb9801f1ae3bf7038e5e915fd7f34f.
//
// Solidity: event PayOrderEvent(uint256 orderId, address user, address merchant, address payToken, uint256 tokenAmount, uint256 payAmount, uint256 swapAmount, uint256 fee)
func (_Pay *PayFilterer) FilterPayOrderEvent(opts *bind.FilterOpts) (*PayPayOrderEventIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "PayOrderEvent")
	if err != nil {
		return nil, err
	}
	return &PayPayOrderEventIterator{contract: _Pay.contract, event: "PayOrderEvent", logs: logs, sub: sub}, nil
}

// WatchPayOrderEvent is a free log subscription operation binding the contract event 0x753d1526f8fcb265ab0f908316c22eaec5eb9801f1ae3bf7038e5e915fd7f34f.
//
// Solidity: event PayOrderEvent(uint256 orderId, address user, address merchant, address payToken, uint256 tokenAmount, uint256 payAmount, uint256 swapAmount, uint256 fee)
func (_Pay *PayFilterer) WatchPayOrderEvent(opts *bind.WatchOpts, sink chan<- *PayPayOrderEvent) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "PayOrderEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PayPayOrderEvent)
				if err := _Pay.contract.UnpackLog(event, "PayOrderEvent", log); err != nil {
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

// ParsePayOrderEvent is a log parse operation binding the contract event 0x753d1526f8fcb265ab0f908316c22eaec5eb9801f1ae3bf7038e5e915fd7f34f.
//
// Solidity: event PayOrderEvent(uint256 orderId, address user, address merchant, address payToken, uint256 tokenAmount, uint256 payAmount, uint256 swapAmount, uint256 fee)
func (_Pay *PayFilterer) ParsePayOrderEvent(log types.Log) (*PayPayOrderEvent, error) {
	event := new(PayPayOrderEvent)
	if err := _Pay.contract.UnpackLog(event, "PayOrderEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaySubScribeIterator is returned from FilterSubScribe and is used to iterate over the raw logs and unpacked data for SubScribe events raised by the Pay contract.
type PaySubScribeIterator struct {
	Event *PaySubScribe // Event containing the contract specifics and raw log

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
func (it *PaySubScribeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaySubScribe)
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
		it.Event = new(PaySubScribe)
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
func (it *PaySubScribeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaySubScribeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaySubScribe represents a SubScribe event raised by the Pay contract.
type PaySubScribe struct {
	User     common.Address
	Merchant common.Address
	Product  *big.Int
	Period   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubScribe is a free log retrieval operation binding the contract event 0x54803366f6f5e96c62465b8e1e0b1dfdfe995c762a6658e0a3718cbc908f378f.
//
// Solidity: event SubScribe(address user, address merchant, uint256 product, uint256 _period)
func (_Pay *PayFilterer) FilterSubScribe(opts *bind.FilterOpts) (*PaySubScribeIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "SubScribe")
	if err != nil {
		return nil, err
	}
	return &PaySubScribeIterator{contract: _Pay.contract, event: "SubScribe", logs: logs, sub: sub}, nil
}

// WatchSubScribe is a free log subscription operation binding the contract event 0x54803366f6f5e96c62465b8e1e0b1dfdfe995c762a6658e0a3718cbc908f378f.
//
// Solidity: event SubScribe(address user, address merchant, uint256 product, uint256 _period)
func (_Pay *PayFilterer) WatchSubScribe(opts *bind.WatchOpts, sink chan<- *PaySubScribe) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "SubScribe")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaySubScribe)
				if err := _Pay.contract.UnpackLog(event, "SubScribe", log); err != nil {
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

// ParseSubScribe is a log parse operation binding the contract event 0x54803366f6f5e96c62465b8e1e0b1dfdfe995c762a6658e0a3718cbc908f378f.
//
// Solidity: event SubScribe(address user, address merchant, uint256 product, uint256 _period)
func (_Pay *PayFilterer) ParseSubScribe(log types.Log) (*PaySubScribe, error) {
	event := new(PaySubScribe)
	if err := _Pay.contract.UnpackLog(event, "SubScribe", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaySubScribePayIterator is returned from FilterSubScribePay and is used to iterate over the raw logs and unpacked data for SubScribePay events raised by the Pay contract.
type PaySubScribePayIterator struct {
	Event *PaySubScribePay // Event containing the contract specifics and raw log

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
func (it *PaySubScribePayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaySubScribePay)
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
		it.Event = new(PaySubScribePay)
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
func (it *PaySubScribePayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaySubScribePayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaySubScribePay represents a SubScribePay event raised by the Pay contract.
type PaySubScribePay struct {
	User     common.Address
	Merchant common.Address
	Token    common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubScribePay is a free log retrieval operation binding the contract event 0xaad8e4ecfd03e117888597719af4909304da88c03d25df02af62f1fa49a37e48.
//
// Solidity: event SubScribePay(address user, address merchant, address _token, uint256 amount, uint256 fee)
func (_Pay *PayFilterer) FilterSubScribePay(opts *bind.FilterOpts) (*PaySubScribePayIterator, error) {

	logs, sub, err := _Pay.contract.FilterLogs(opts, "SubScribePay")
	if err != nil {
		return nil, err
	}
	return &PaySubScribePayIterator{contract: _Pay.contract, event: "SubScribePay", logs: logs, sub: sub}, nil
}

// WatchSubScribePay is a free log subscription operation binding the contract event 0xaad8e4ecfd03e117888597719af4909304da88c03d25df02af62f1fa49a37e48.
//
// Solidity: event SubScribePay(address user, address merchant, address _token, uint256 amount, uint256 fee)
func (_Pay *PayFilterer) WatchSubScribePay(opts *bind.WatchOpts, sink chan<- *PaySubScribePay) (event.Subscription, error) {

	logs, sub, err := _Pay.contract.WatchLogs(opts, "SubScribePay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaySubScribePay)
				if err := _Pay.contract.UnpackLog(event, "SubScribePay", log); err != nil {
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

// ParseSubScribePay is a log parse operation binding the contract event 0xaad8e4ecfd03e117888597719af4909304da88c03d25df02af62f1fa49a37e48.
//
// Solidity: event SubScribePay(address user, address merchant, address _token, uint256 amount, uint256 fee)
func (_Pay *PayFilterer) ParseSubScribePay(log types.Log) (*PaySubScribePay, error) {
	event := new(PaySubScribePay)
	if err := _Pay.contract.UnpackLog(event, "SubScribePay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
