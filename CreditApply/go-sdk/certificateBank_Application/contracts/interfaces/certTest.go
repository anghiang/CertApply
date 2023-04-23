// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// iBankCertificate is an auto generated low-level Go binding around an user-defined struct.
type IBankCertificate struct {
	Id             string
	OwnerAddr      common.Address
	Name           string
	IssueTime      uint64
	ValidityPeriod uint64
}

// iBankCredit is an auto generated low-level Go binding around an user-defined struct.
type IBankCredit struct {
	Score uint64
	CType uint8
}

// iBankUser is an auto generated low-level Go binding around an user-defined struct.
type IBankUser struct {
	Id    uint64
	Name  string
	Age   uint8
	Phone string
}


// CertTestABI is the input ABI used to generate the binding from.
const CertTestABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_uAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_issueTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_validityPeriod\",\"type\":\"uint64\"}],\"name\":\"addCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_uAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_score\",\"type\":\"uint8\"},{\"internalType\":\"enumiBank.CreditTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"certificateBank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"}],\"name\":\"getUserByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"internalType\":\"structiBank.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"}],\"name\":\"getUserCertificate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"issueTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"validityPeriod\",\"type\":\"uint64\"}],\"internalType\":\"structiBank.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"}],\"name\":\"getUserCredits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"enumiBank.CreditTypes\",\"name\":\"cType\",\"type\":\"uint8\"}],\"internalType\":\"structiBank.Credit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getmsgsender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCertificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"issueTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"validityPeriod\",\"type\":\"uint64\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCredit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"enumiBank.CreditTypes\",\"name\":\"cType\",\"type\":\"uint8\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userUser\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"stateMutability\":\"template\",\"type\":\"function\"}]"

// CertTest is an auto generated Go binding around a Solidity contract.
type CertTest struct {
	CertTestCaller     // Read-only binding to the contract
	CertTestTransactor // Write-only binding to the contract
	CertTestFilterer   // Log filterer for contract events
}

// CertTestCaller is an auto generated read-only Go binding around a Solidity contract.
type CertTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTestTransactor is an auto generated write-only Go binding around a Solidity contract.
type CertTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTestFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type CertTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTestSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type CertTestSession struct {
	Contract     *CertTest         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertTestCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type CertTestCallerSession struct {
	Contract *CertTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CertTestTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type CertTestTransactorSession struct {
	Contract     *CertTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CertTestRaw is an auto generated low-level Go binding around a Solidity contract.
type CertTestRaw struct {
	Contract *CertTest // Generic contract binding to access the raw methods on
}

// CertTestCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type CertTestCallerRaw struct {
	Contract *CertTestCaller // Generic read-only contract binding to access the raw methods on
}

// CertTestTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type CertTestTransactorRaw struct {
	Contract *CertTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertTest creates a new instance of CertTest, bound to a specific deployed contract.
func NewCertTest(address common.Address, backend bind.ContractBackend) (*CertTest, error) {
	contract, err := bindCertTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertTest{CertTestCaller: CertTestCaller{contract: contract}, CertTestTransactor: CertTestTransactor{contract: contract}, CertTestFilterer: CertTestFilterer{contract: contract}}, nil
}

// NewCertTestCaller creates a new read-only instance of CertTest, bound to a specific deployed contract.
func NewCertTestCaller(address common.Address, caller bind.ContractCaller) (*CertTestCaller, error) {
	contract, err := bindCertTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertTestCaller{contract: contract}, nil
}

// NewCertTestTransactor creates a new write-only instance of CertTest, bound to a specific deployed contract.
func NewCertTestTransactor(address common.Address, transactor bind.ContractTransactor) (*CertTestTransactor, error) {
	contract, err := bindCertTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertTestTransactor{contract: contract}, nil
}

// NewCertTestFilterer creates a new log filterer instance of CertTest, bound to a specific deployed contract.
func NewCertTestFilterer(address common.Address, filterer bind.ContractFilterer) (*CertTestFilterer, error) {
	contract, err := bindCertTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertTestFilterer{contract: contract}, nil
}

// bindCertTest binds a generic wrapper to an already deployed contract.
func bindCertTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertTest *CertTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertTest.Contract.CertTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertTest *CertTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.CertTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertTest *CertTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.CertTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertTest *CertTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CertTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertTest *CertTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertTest *CertTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.contract.Transact(opts, method, params...)
}

// CertificateBank is a free data retrieval call binding the contract method 0x0eafc810.
//
// Solidity: function certificateBank() constant returns(address)
func (_CertTest *CertTestCaller) CertificateBank(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CertTest.contract.Call(opts, out, "certificateBank")
	return *ret0, err
}

// CertificateBank is a free data retrieval call binding the contract method 0x0eafc810.
//
// Solidity: function certificateBank() constant returns(address)
func (_CertTest *CertTestSession) CertificateBank() (common.Address, error) {
	return _CertTest.Contract.CertificateBank(&_CertTest.CallOpts)
}

// CertificateBank is a free data retrieval call binding the contract method 0x0eafc810.
//
// Solidity: function certificateBank() constant returns(address)
func (_CertTest *CertTestCallerSession) CertificateBank() (common.Address, error) {
	return _CertTest.Contract.CertificateBank(&_CertTest.CallOpts)
}

// GetUserByAddress is a free data retrieval call binding the contract method 0x69c212f6.
//
// Solidity: function getUserByAddress(address uAddr) constant returns(iBankUser)
func (_CertTest *CertTestCaller) GetUserByAddress(opts *bind.CallOpts, uAddr common.Address) (IBankUser, error) {
	var (
		ret0 = new(IBankUser)
	)
	out := ret0
	err := _CertTest.contract.Call(opts, out, "getUserByAddress", uAddr)
	return *ret0, err
}

// GetUserByAddress is a free data retrieval call binding the contract method 0x69c212f6.
//
// Solidity: function getUserByAddress(address uAddr) constant returns(iBankUser)
func (_CertTest *CertTestSession) GetUserByAddress(uAddr common.Address) (IBankUser, error) {
	return _CertTest.Contract.GetUserByAddress(&_CertTest.CallOpts, uAddr)
}

// GetUserByAddress is a free data retrieval call binding the contract method 0x69c212f6.
//
// Solidity: function getUserByAddress(address uAddr) constant returns(iBankUser)
func (_CertTest *CertTestCallerSession) GetUserByAddress(uAddr common.Address) (IBankUser, error) {
	return _CertTest.Contract.GetUserByAddress(&_CertTest.CallOpts, uAddr)
}

// GetUserCertificate is a free data retrieval call binding the contract method 0x0bfdc7d7.
//
// Solidity: function getUserCertificate(address uAddr) constant returns([]iBankCertificate)
func (_CertTest *CertTestCaller) GetUserCertificate(opts *bind.CallOpts, uAddr common.Address) ([]IBankCertificate, error) {
	var (
		ret0 = new([]IBankCertificate)
	)
	out := ret0
	err := _CertTest.contract.Call(opts, out, "getUserCertificate", uAddr)
	return *ret0, err
}

// GetUserCertificate is a free data retrieval call binding the contract method 0x0bfdc7d7.
//
// Solidity: function getUserCertificate(address uAddr) constant returns([]iBankCertificate)
func (_CertTest *CertTestSession) GetUserCertificate(uAddr common.Address) ([]IBankCertificate, error) {
	return _CertTest.Contract.GetUserCertificate(&_CertTest.CallOpts, uAddr)
}

// GetUserCertificate is a free data retrieval call binding the contract method 0x0bfdc7d7.
//
// Solidity: function getUserCertificate(address uAddr) constant returns([]iBankCertificate)
func (_CertTest *CertTestCallerSession) GetUserCertificate(uAddr common.Address) ([]IBankCertificate, error) {
	return _CertTest.Contract.GetUserCertificate(&_CertTest.CallOpts, uAddr)
}

// GetUserCredits is a free data retrieval call binding the contract method 0x81b26dda.
//
// Solidity: function getUserCredits(address uAddr) constant returns([]iBankCredit)
func (_CertTest *CertTestCaller) GetUserCredits(opts *bind.CallOpts, uAddr common.Address) ([]IBankCredit, error) {
	var (
		ret0 = new([]IBankCredit)
	)
	out := ret0
	err := _CertTest.contract.Call(opts, out, "getUserCredits", uAddr)
	return *ret0, err
}

// GetUserCredits is a free data retrieval call binding the contract method 0x81b26dda.
//
// Solidity: function getUserCredits(address uAddr) constant returns([]iBankCredit)
func (_CertTest *CertTestSession) GetUserCredits(uAddr common.Address) ([]IBankCredit, error) {
	return _CertTest.Contract.GetUserCredits(&_CertTest.CallOpts, uAddr)
}

// GetUserCredits is a free data retrieval call binding the contract method 0x81b26dda.
//
// Solidity: function getUserCredits(address uAddr) constant returns([]iBankCredit)
func (_CertTest *CertTestCallerSession) GetUserCredits(uAddr common.Address) ([]IBankCredit, error) {
	return _CertTest.Contract.GetUserCredits(&_CertTest.CallOpts, uAddr)
}

// Getmsgsender is a free data retrieval call binding the contract method 0x3727c0b6.
//
// Solidity: function getmsgsender() constant returns(address)
func (_CertTest *CertTestCaller) Getmsgsender(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CertTest.contract.Call(opts, out, "getmsgsender")
	return *ret0, err
}

// Getmsgsender is a free data retrieval call binding the contract method 0x3727c0b6.
//
// Solidity: function getmsgsender() constant returns(address)
func (_CertTest *CertTestSession) Getmsgsender() (common.Address, error) {
	return _CertTest.Contract.Getmsgsender(&_CertTest.CallOpts)
}

// Getmsgsender is a free data retrieval call binding the contract method 0x3727c0b6.
//
// Solidity: function getmsgsender() constant returns(address)
func (_CertTest *CertTestCallerSession) Getmsgsender() (common.Address, error) {
	return _CertTest.Contract.Getmsgsender(&_CertTest.CallOpts)
}

// UserCertificates is a free data retrieval call binding the contract method 0x0175f889.
//
// Solidity: function userCertificates(address , uint256 ) constant returns(string id, address ownerAddr, string name, uint64 issueTime, uint64 validityPeriod)
func (_CertTest *CertTestCaller) UserCertificates(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Id             string
	OwnerAddr      common.Address
	Name           string
	IssueTime      uint64
	ValidityPeriod uint64
}, error) {
	ret := new(struct {
		Id             string
		OwnerAddr      common.Address
		Name           string
		IssueTime      uint64
		ValidityPeriod uint64
	})
	out := ret
	err := _CertTest.contract.Call(opts, out, "userCertificates", arg0, arg1)
	return *ret, err
}

// UserCertificates is a free data retrieval call binding the contract method 0x0175f889.
//
// Solidity: function userCertificates(address , uint256 ) constant returns(string id, address ownerAddr, string name, uint64 issueTime, uint64 validityPeriod)
func (_CertTest *CertTestSession) UserCertificates(arg0 common.Address, arg1 *big.Int) (struct {
	Id             string
	OwnerAddr      common.Address
	Name           string
	IssueTime      uint64
	ValidityPeriod uint64
}, error) {
	return _CertTest.Contract.UserCertificates(&_CertTest.CallOpts, arg0, arg1)
}

// UserCertificates is a free data retrieval call binding the contract method 0x0175f889.
//
// Solidity: function userCertificates(address , uint256 ) constant returns(string id, address ownerAddr, string name, uint64 issueTime, uint64 validityPeriod)
func (_CertTest *CertTestCallerSession) UserCertificates(arg0 common.Address, arg1 *big.Int) (struct {
	Id             string
	OwnerAddr      common.Address
	Name           string
	IssueTime      uint64
	ValidityPeriod uint64
}, error) {
	return _CertTest.Contract.UserCertificates(&_CertTest.CallOpts, arg0, arg1)
}

// UserCredit is a free data retrieval call binding the contract method 0xd426de72.
//
// Solidity: function userCredit(address , uint256 ) constant returns(uint64 score, uint8 cType)
func (_CertTest *CertTestCaller) UserCredit(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Score uint64
	CType uint8
}, error) {
	ret := new(struct {
		Score uint64
		CType uint8
	})
	out := ret
	err := _CertTest.contract.Call(opts, out, "userCredit", arg0, arg1)
	return *ret, err
}

// UserCredit is a free data retrieval call binding the contract method 0xd426de72.
//
// Solidity: function userCredit(address , uint256 ) constant returns(uint64 score, uint8 cType)
func (_CertTest *CertTestSession) UserCredit(arg0 common.Address, arg1 *big.Int) (struct {
	Score uint64
	CType uint8
}, error) {
	return _CertTest.Contract.UserCredit(&_CertTest.CallOpts, arg0, arg1)
}

// UserCredit is a free data retrieval call binding the contract method 0xd426de72.
//
// Solidity: function userCredit(address , uint256 ) constant returns(uint64 score, uint8 cType)
func (_CertTest *CertTestCallerSession) UserCredit(arg0 common.Address, arg1 *big.Int) (struct {
	Score uint64
	CType uint8
}, error) {
	return _CertTest.Contract.UserCredit(&_CertTest.CallOpts, arg0, arg1)
}

// UserUser is a free data retrieval call binding the contract method 0xabe21a8c.
//
// Solidity: function userUser(address ) constant returns(uint64 id, string name, uint8 age, string phone)
func (_CertTest *CertTestCaller) UserUser(opts *bind.CallOpts, arg0 common.Address) (struct {
	Id    uint64
	Name  string
	Age   uint8
	Phone string
}, error) {
	ret := new(struct {
		Id    uint64
		Name  string
		Age   uint8
		Phone string
	})
	out := ret
	err := _CertTest.contract.Call(opts, out, "userUser", arg0)
	return *ret, err
}

// UserUser is a free data retrieval call binding the contract method 0xabe21a8c.
//
// Solidity: function userUser(address ) constant returns(uint64 id, string name, uint8 age, string phone)
func (_CertTest *CertTestSession) UserUser(arg0 common.Address) (struct {
	Id    uint64
	Name  string
	Age   uint8
	Phone string
}, error) {
	return _CertTest.Contract.UserUser(&_CertTest.CallOpts, arg0)
}

// UserUser is a free data retrieval call binding the contract method 0xabe21a8c.
//
// Solidity: function userUser(address ) constant returns(uint64 id, string name, uint8 age, string phone)
func (_CertTest *CertTestCallerSession) UserUser(arg0 common.Address) (struct {
	Id    uint64
	Name  string
	Age   uint8
	Phone string
}, error) {
	return _CertTest.Contract.UserUser(&_CertTest.CallOpts, arg0)
}

// AddCertificate is a paid mutator transaction binding the contract method 0x91ee1839.
//
// Solidity: function addCertificate(address _uAddr, string _cid, string _cName, uint64 _issueTime, uint64 _validityPeriod) returns()
func (_CertTest *CertTestTransactor) AddCertificate(opts *bind.TransactOpts, _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.contract.Transact(opts, "addCertificate", _uAddr, _cid, _cName, _issueTime, _validityPeriod)
}

func (_CertTest *CertTestTransactor) AsyncAddCertificate(handler func(*types.Receipt, error), opts *bind.TransactOpts, _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, error) {
	return _CertTest.contract.AsyncTransact(opts, handler, "addCertificate", _uAddr, _cid, _cName, _issueTime, _validityPeriod)
}

// AddCertificate is a paid mutator transaction binding the contract method 0x91ee1839.
//
// Solidity: function addCertificate(address _uAddr, string _cid, string _cName, uint64 _issueTime, uint64 _validityPeriod) returns()
func (_CertTest *CertTestSession) AddCertificate(_uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.AddCertificate(&_CertTest.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
}

func (_CertTest *CertTestSession) AsyncAddCertificate(handler func(*types.Receipt, error), _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, error) {
	return _CertTest.Contract.AsyncAddCertificate(handler, &_CertTest.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
}

// AddCertificate is a paid mutator transaction binding the contract method 0x91ee1839.
//
// Solidity: function addCertificate(address _uAddr, string _cid, string _cName, uint64 _issueTime, uint64 _validityPeriod) returns()
func (_CertTest *CertTestTransactorSession) AddCertificate(_uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.AddCertificate(&_CertTest.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
}

func (_CertTest *CertTestTransactorSession) AsyncAddCertificate(handler func(*types.Receipt, error), _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, error) {
	return _CertTest.Contract.AsyncAddCertificate(handler, &_CertTest.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
}

// AddCredit is a paid mutator transaction binding the contract method 0x7587dc2b.
//
// Solidity: function addCredit(address _uAddr, uint8 _score, uint8 _type) returns()
func (_CertTest *CertTestTransactor) AddCredit(opts *bind.TransactOpts, _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.contract.Transact(opts, "addCredit", _uAddr, _score, _type)
}

func (_CertTest *CertTestTransactor) AsyncAddCredit(handler func(*types.Receipt, error), opts *bind.TransactOpts, _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, error) {
	return _CertTest.contract.AsyncTransact(opts, handler, "addCredit", _uAddr, _score, _type)
}

// AddCredit is a paid mutator transaction binding the contract method 0x7587dc2b.
//
// Solidity: function addCredit(address _uAddr, uint8 _score, uint8 _type) returns()
func (_CertTest *CertTestSession) AddCredit(_uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.AddCredit(&_CertTest.TransactOpts, _uAddr, _score, _type)
}

func (_CertTest *CertTestSession) AsyncAddCredit(handler func(*types.Receipt, error), _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, error) {
	return _CertTest.Contract.AsyncAddCredit(handler, &_CertTest.TransactOpts, _uAddr, _score, _type)
}

// AddCredit is a paid mutator transaction binding the contract method 0x7587dc2b.
//
// Solidity: function addCredit(address _uAddr, uint8 _score, uint8 _type) returns()
func (_CertTest *CertTestTransactorSession) AddCredit(_uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, *types.Receipt, error) {
	return _CertTest.Contract.AddCredit(&_CertTest.TransactOpts, _uAddr, _score, _type)
}

func (_CertTest *CertTestTransactorSession) AsyncAddCredit(handler func(*types.Receipt, error), _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, error) {
	return _CertTest.Contract.AsyncAddCredit(handler, &_CertTest.TransactOpts, _uAddr, _score, _type)
}
