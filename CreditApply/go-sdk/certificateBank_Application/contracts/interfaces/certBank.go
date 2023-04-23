// // Code generated - DO NOT EDIT.
// // This file is a generated binding and any manual changes will be lost.
package interfaces

//
//import (
//	"math/big"
//	"strings"
//
//	"github.com/FISCO-BCOS/go-sdk/abi"
//	"github.com/FISCO-BCOS/go-sdk/abi/bind"
//	"github.com/FISCO-BCOS/go-sdk/core/types"
//	ethereum "github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/common"
//)
//
//// Reference imports to suppress errors if they are not otherwise used.
//var (
//	_ = big.NewInt
//	_ = strings.NewReader
//	_ = ethereum.NotFound
//	_ = abi.U256
//	_ = bind.Bind
//	_ = common.Big1
//	_ = types.BloomLookup
//)
//
//// iBankCertificate is an auto generated low-level Go binding around an user-defined struct.
//type iBankCertificate struct {
//	Id             string
//	OwnerAddr      common.Address
//	Name           string
//	IssueTime      uint64
//	ValidityPeriod uint64
//}
//
//// iBankCredit is an auto generated low-level Go binding around an user-defined struct.
//type iBankCredit struct {
//	Score uint64
//	CType uint8
//}
//
//// iBankUser is an auto generated low-level Go binding around an user-defined struct.
//type iBankUser struct {
//	Id    uint64
//	Name  string
//	Age   uint8
//	Phone string
//}
//
//// CertBankABI is the input ABI used to generate the binding from.
//const CertBankABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_uAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cName\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_issueTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_validityPeriod\",\"type\":\"uint64\"}],\"name\":\"addCertificate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_uAddr\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_score\",\"type\":\"uint8\"},{\"internalType\":\"enumiBank.CreditTypes\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"addCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"certificateBank\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"}],\"name\":\"getUserByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"internalType\":\"structiBank.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"}],\"name\":\"getUserCertificate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"issueTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"validityPeriod\",\"type\":\"uint64\"}],\"internalType\":\"structiBank.Certificate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"uAddr\",\"type\":\"address\"}],\"name\":\"getUserCredits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"enumiBank.CreditTypes\",\"name\":\"cType\",\"type\":\"uint8\"}],\"internalType\":\"structiBank.Credit[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCertificates\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"issueTime\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"validityPeriod\",\"type\":\"uint64\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userCredit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"enumiBank.CreditTypes\",\"name\":\"cType\",\"type\":\"uint8\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userUser\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"age\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"}],\"stateMutability\":\"template\",\"type\":\"function\"}]"
//
//// CertBank is an auto generated Go binding around a Solidity contract.
//type CertBank struct {
//	CertBankCaller     // Read-only binding to the contract
//	CertBankTransactor // Write-only binding to the contract
//	CertBankFilterer   // Log filterer for contract events
//}
//
//// CertBankCaller is an auto generated read-only Go binding around a Solidity contract.
//type CertBankCaller struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// CertBankTransactor is an auto generated write-only Go binding around a Solidity contract.
//type CertBankTransactor struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// CertBankFilterer is an auto generated log filtering Go binding around a Solidity contract events.
//type CertBankFilterer struct {
//	contract *bind.BoundContract // Generic contract wrapper for the low level calls
//}
//
//// CertBankSession is an auto generated Go binding around a Solidity contract,
//// with pre-set call and transact options.
//type CertBankSession struct {
//	Contract     *CertBank         // Generic contract binding to set the session for
//	CallOpts     bind.CallOpts     // Call options to use throughout this session
//	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
//}
//
//// CertBankCallerSession is an auto generated read-only Go binding around a Solidity contract,
//// with pre-set call options.
//type CertBankCallerSession struct {
//	Contract *CertBankCaller // Generic contract caller binding to set the session for
//	CallOpts bind.CallOpts   // Call options to use throughout this session
//}
//
//// CertBankTransactorSession is an auto generated write-only Go binding around a Solidity contract,
//// with pre-set transact options.
//type CertBankTransactorSession struct {
//	Contract     *CertBankTransactor // Generic contract transactor binding to set the session for
//	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
//}
//
//// CertBankRaw is an auto generated low-level Go binding around a Solidity contract.
//type CertBankRaw struct {
//	Contract *CertBank // Generic contract binding to access the raw methods on
//}
//
//// CertBankCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
//type CertBankCallerRaw struct {
//	Contract *CertBankCaller // Generic read-only contract binding to access the raw methods on
//}
//
//// CertBankTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
//type CertBankTransactorRaw struct {
//	Contract *CertBankTransactor // Generic write-only contract binding to access the raw methods on
//}
//
//// NewCertBank creates a new instance of CertBank, bound to a specific deployed contract.
//func NewCertBank(address common.Address, backend bind.ContractBackend) (*CertBank, error) {
//	contract, err := bindCertBank(address, backend, backend, backend)
//	if err != nil {
//		return nil, err
//	}
//	return &CertBank{CertBankCaller: CertBankCaller{contract: contract}, CertBankTransactor: CertBankTransactor{contract: contract}, CertBankFilterer: CertBankFilterer{contract: contract}}, nil
//}
//
//// NewCertBankCaller creates a new read-only instance of CertBank, bound to a specific deployed contract.
//func NewCertBankCaller(address common.Address, caller bind.ContractCaller) (*CertBankCaller, error) {
//	contract, err := bindCertBank(address, caller, nil, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &CertBankCaller{contract: contract}, nil
//}
//
//// NewCertBankTransactor creates a new write-only instance of CertBank, bound to a specific deployed contract.
//func NewCertBankTransactor(address common.Address, transactor bind.ContractTransactor) (*CertBankTransactor, error) {
//	contract, err := bindCertBank(address, nil, transactor, nil)
//	if err != nil {
//		return nil, err
//	}
//	return &CertBankTransactor{contract: contract}, nil
//}
//
//// NewCertBankFilterer creates a new log filterer instance of CertBank, bound to a specific deployed contract.
//func NewCertBankFilterer(address common.Address, filterer bind.ContractFilterer) (*CertBankFilterer, error) {
//	contract, err := bindCertBank(address, nil, nil, filterer)
//	if err != nil {
//		return nil, err
//	}
//	return &CertBankFilterer{contract: contract}, nil
//}
//
//// bindCertBank binds a generic wrapper to an already deployed contract.
//func bindCertBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
//	parsed, err := abi.JSON(strings.NewReader(CertBankABI))
//	if err != nil {
//		return nil, err
//	}
//	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
//}
//
//// Call invokes the (constant) contract method with params as input values and
//// sets the output to result. The result type might be a single field for simple
//// returns, a slice of interfaces for anonymous returns and a struct for named
//// returns.
//func (_CertBank *CertBankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
//	return _CertBank.Contract.CertBankCaller.contract.Call(opts, result, method, params...)
//}
//
//// Transfer initiates a plain transaction to move funds to the contract, calling
//// its default method if one is available.
//func (_CertBank *CertBankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.CertBankTransactor.contract.Transfer(opts)
//}
//
//// Transact invokes the (paid) contract method with params as input values.
//func (_CertBank *CertBankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.CertBankTransactor.contract.Transact(opts, method, params...)
//}
//
//// Call invokes the (constant) contract method with params as input values and
//// sets the output to result. The result type might be a single field for simple
//// returns, a slice of interfaces for anonymous returns and a struct for named
//// returns.
//func (_CertBank *CertBankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
//	return _CertBank.Contract.contract.Call(opts, result, method, params...)
//}
//
//// Transfer initiates a plain transaction to move funds to the contract, calling
//// its default method if one is available.
//func (_CertBank *CertBankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.contract.Transfer(opts)
//}
//
//// Transact invokes the (paid) contract method with params as input values.
//func (_CertBank *CertBankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.contract.Transact(opts, method, params...)
//}
//
//// CertificateBank is a free data retrieval call binding the contract method 0x0eafc810.
////
//// Solidity: function certificateBank() constant returns(address)
//func (_CertBank *CertBankCaller) CertificateBank(opts *bind.CallOpts) (common.Address, error) {
//	var (
//		ret0 = new(common.Address)
//	)
//	out := ret0
//	err := _CertBank.contract.Call(opts, out, "certificateBank")
//	return *ret0, err
//}
//
//// CertificateBank is a free data retrieval call binding the contract method 0x0eafc810.
////
//// Solidity: function certificateBank() constant returns(address)
//func (_CertBank *CertBankSession) CertificateBank() (common.Address, error) {
//	return _CertBank.Contract.CertificateBank(&_CertBank.CallOpts)
//}
//
//// CertificateBank is a free data retrieval call binding the contract method 0x0eafc810.
////
//// Solidity: function certificateBank() constant returns(address)
//func (_CertBank *CertBankCallerSession) CertificateBank() (common.Address, error) {
//	return _CertBank.Contract.CertificateBank(&_CertBank.CallOpts)
//}
//
//// GetUserByAddress is a free data retrieval call binding the contract method 0x69c212f6.
////
//// Solidity: function getUserByAddress(address uAddr) constant returns(iBankUser)
//func (_CertBank *CertBankCaller) GetUserByAddress(opts *bind.CallOpts, uAddr common.Address) (iBankUser, error) {
//	var (
//		ret0 = new(iBankUser)
//	)
//	out := ret0
//	err := _CertBank.contract.Call(opts, out, "getUserByAddress", uAddr)
//	return *ret0, err
//}
//
//// GetUserByAddress is a free data retrieval call binding the contract method 0x69c212f6.
////
//// Solidity: function getUserByAddress(address uAddr) constant returns(iBankUser)
//func (_CertBank *CertBankSession) GetUserByAddress(uAddr common.Address) (iBankUser, error) {
//	return _CertBank.Contract.GetUserByAddress(&_CertBank.CallOpts, uAddr)
//}
//
//// GetUserByAddress is a free data retrieval call binding the contract method 0x69c212f6.
////
//// Solidity: function getUserByAddress(address uAddr) constant returns(iBankUser)
//func (_CertBank *CertBankCallerSession) GetUserByAddress(uAddr common.Address) (iBankUser, error) {
//	return _CertBank.Contract.GetUserByAddress(&_CertBank.CallOpts, uAddr)
//}
//
//// GetUserCertificate is a free data retrieval call binding the contract method 0x0bfdc7d7.
////
//// Solidity: function getUserCertificate(address uAddr) constant returns([]iBankCertificate)
//func (_CertBank *CertBankCaller) GetUserCertificate(opts *bind.CallOpts, uAddr common.Address) ([]iBankCertificate, error) {
//	var (
//		ret0 = new([]iBankCertificate)
//	)
//	out := ret0
//	err := _CertBank.contract.Call(opts, out, "getUserCertificate", uAddr)
//	return *ret0, err
//}
//
//// GetUserCertificate is a free data retrieval call binding the contract method 0x0bfdc7d7.
////
//// Solidity: function getUserCertificate(address uAddr) constant returns([]iBankCertificate)
//func (_CertBank *CertBankSession) GetUserCertificate(uAddr common.Address) ([]iBankCertificate, error) {
//	return _CertBank.Contract.GetUserCertificate(&_CertBank.CallOpts, uAddr)
//}
//
//// GetUserCertificate is a free data retrieval call binding the contract method 0x0bfdc7d7.
////
//// Solidity: function getUserCertificate(address uAddr) constant returns([]iBankCertificate)
//func (_CertBank *CertBankCallerSession) GetUserCertificate(uAddr common.Address) ([]iBankCertificate, error) {
//	return _CertBank.Contract.GetUserCertificate(&_CertBank.CallOpts, uAddr)
//}
//
//// GetUserCredits is a free data retrieval call binding the contract method 0x81b26dda.
////
//// Solidity: function getUserCredits(address uAddr) constant returns([]iBankCredit)
//func (_CertBank *CertBankCaller) GetUserCredits(opts *bind.CallOpts, uAddr common.Address) ([]iBankCredit, error) {
//	var (
//		ret0 = new([]iBankCredit)
//	)
//	out := ret0
//	err := _CertBank.contract.Call(opts, out, "getUserCredits", uAddr)
//	return *ret0, err
//}
//
//// GetUserCredits is a free data retrieval call binding the contract method 0x81b26dda.
////
//// Solidity: function getUserCredits(address uAddr) constant returns([]iBankCredit)
//func (_CertBank *CertBankSession) GetUserCredits(uAddr common.Address) ([]iBankCredit, error) {
//	return _CertBank.Contract.GetUserCredits(&_CertBank.CallOpts, uAddr)
//}
//
//// GetUserCredits is a free data retrieval call binding the contract method 0x81b26dda.
////
//// Solidity: function getUserCredits(address uAddr) constant returns([]iBankCredit)
//func (_CertBank *CertBankCallerSession) GetUserCredits(uAddr common.Address) ([]iBankCredit, error) {
//	return _CertBank.Contract.GetUserCredits(&_CertBank.CallOpts, uAddr)
//}
//
//// UserCertificates is a free data retrieval call binding the contract method 0x0175f889.
////
//// Solidity: function userCertificates(address , uint256 ) constant returns(string id, address ownerAddr, string name, uint64 issueTime, uint64 validityPeriod)
//func (_CertBank *CertBankCaller) UserCertificates(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
//	Id             string
//	OwnerAddr      common.Address
//	Name           string
//	IssueTime      uint64
//	ValidityPeriod uint64
//}, error) {
//	ret := new(struct {
//		Id             string
//		OwnerAddr      common.Address
//		Name           string
//		IssueTime      uint64
//		ValidityPeriod uint64
//	})
//	out := ret
//	err := _CertBank.contract.Call(opts, out, "userCertificates", arg0, arg1)
//	return *ret, err
//}
//
//// UserCertificates is a free data retrieval call binding the contract method 0x0175f889.
////
//// Solidity: function userCertificates(address , uint256 ) constant returns(string id, address ownerAddr, string name, uint64 issueTime, uint64 validityPeriod)
//func (_CertBank *CertBankSession) UserCertificates(arg0 common.Address, arg1 *big.Int) (struct {
//	Id             string
//	OwnerAddr      common.Address
//	Name           string
//	IssueTime      uint64
//	ValidityPeriod uint64
//}, error) {
//	return _CertBank.Contract.UserCertificates(&_CertBank.CallOpts, arg0, arg1)
//}
//
//// UserCertificates is a free data retrieval call binding the contract method 0x0175f889.
////
//// Solidity: function userCertificates(address , uint256 ) constant returns(string id, address ownerAddr, string name, uint64 issueTime, uint64 validityPeriod)
//func (_CertBank *CertBankCallerSession) UserCertificates(arg0 common.Address, arg1 *big.Int) (struct {
//	Id             string
//	OwnerAddr      common.Address
//	Name           string
//	IssueTime      uint64
//	ValidityPeriod uint64
//}, error) {
//	return _CertBank.Contract.UserCertificates(&_CertBank.CallOpts, arg0, arg1)
//}
//
//// UserCredit is a free data retrieval call binding the contract method 0xd426de72.
////
//// Solidity: function userCredit(address , uint256 ) constant returns(uint64 score, uint8 cType)
//func (_CertBank *CertBankCaller) UserCredit(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
//	Score uint64
//	CType uint8
//}, error) {
//	ret := new(struct {
//		Score uint64
//		CType uint8
//	})
//	out := ret
//	err := _CertBank.contract.Call(opts, out, "userCredit", arg0, arg1)
//	return *ret, err
//}
//
//// UserCredit is a free data retrieval call binding the contract method 0xd426de72.
////
//// Solidity: function userCredit(address , uint256 ) constant returns(uint64 score, uint8 cType)
//func (_CertBank *CertBankSession) UserCredit(arg0 common.Address, arg1 *big.Int) (struct {
//	Score uint64
//	CType uint8
//}, error) {
//	return _CertBank.Contract.UserCredit(&_CertBank.CallOpts, arg0, arg1)
//}
//
//// UserCredit is a free data retrieval call binding the contract method 0xd426de72.
////
//// Solidity: function userCredit(address , uint256 ) constant returns(uint64 score, uint8 cType)
//func (_CertBank *CertBankCallerSession) UserCredit(arg0 common.Address, arg1 *big.Int) (struct {
//	Score uint64
//	CType uint8
//}, error) {
//	return _CertBank.Contract.UserCredit(&_CertBank.CallOpts, arg0, arg1)
//}
//
//// UserUser is a free data retrieval call binding the contract method 0xabe21a8c.
////
//// Solidity: function userUser(address ) constant returns(uint64 id, string name, uint8 age, string phone)
//func (_CertBank *CertBankCaller) UserUser(opts *bind.CallOpts, arg0 common.Address) (struct {
//	Id    uint64
//	Name  string
//	Age   uint8
//	Phone string
//}, error) {
//	ret := new(struct {
//		Id    uint64
//		Name  string
//		Age   uint8
//		Phone string
//	})
//	out := ret
//	err := _CertBank.contract.Call(opts, out, "userUser", arg0)
//	return *ret, err
//}
//
//// UserUser is a free data retrieval call binding the contract method 0xabe21a8c.
////
//// Solidity: function userUser(address ) constant returns(uint64 id, string name, uint8 age, string phone)
//func (_CertBank *CertBankSession) UserUser(arg0 common.Address) (struct {
//	Id    uint64
//	Name  string
//	Age   uint8
//	Phone string
//}, error) {
//	return _CertBank.Contract.UserUser(&_CertBank.CallOpts, arg0)
//}
//
//// UserUser is a free data retrieval call binding the contract method 0xabe21a8c.
////
//// Solidity: function userUser(address ) constant returns(uint64 id, string name, uint8 age, string phone)
//func (_CertBank *CertBankCallerSession) UserUser(arg0 common.Address) (struct {
//	Id    uint64
//	Name  string
//	Age   uint8
//	Phone string
//}, error) {
//	return _CertBank.Contract.UserUser(&_CertBank.CallOpts, arg0)
//}
//
//// AddCertificate is a paid mutator transaction binding the contract method 0x91ee1839.
////
//// Solidity: function addCertificate(address _uAddr, string _cid, string _cName, uint64 _issueTime, uint64 _validityPeriod) returns()
//func (_CertBank *CertBankTransactor) AddCertificate(opts *bind.TransactOpts, _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.contract.Transact(opts, "addCertificate", _uAddr, _cid, _cName, _issueTime, _validityPeriod)
//}
//
//func (_CertBank *CertBankTransactor) AsyncAddCertificate(handler func(*types.Receipt, error), opts *bind.TransactOpts, _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, error) {
//	return _CertBank.contract.AsyncTransact(opts, handler, "addCertificate", _uAddr, _cid, _cName, _issueTime, _validityPeriod)
//}
//
//// AddCertificate is a paid mutator transaction binding the contract method 0x91ee1839.
////
//// Solidity: function addCertificate(address _uAddr, string _cid, string _cName, uint64 _issueTime, uint64 _validityPeriod) returns()
//func (_CertBank *CertBankSession) AddCertificate(_uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.AddCertificate(&_CertBank.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
//}
//
//func (_CertBank *CertBankSession) AsyncAddCertificate(handler func(*types.Receipt, error), _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, error) {
//	return _CertBank.Contract.AsyncAddCertificate(handler, &_CertBank.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
//}
//
//// AddCertificate is a paid mutator transaction binding the contract method 0x91ee1839.
////
//// Solidity: function addCertificate(address _uAddr, string _cid, string _cName, uint64 _issueTime, uint64 _validityPeriod) returns()
//func (_CertBank *CertBankTransactorSession) AddCertificate(_uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.AddCertificate(&_CertBank.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
//}
//
//func (_CertBank *CertBankTransactorSession) AsyncAddCertificate(handler func(*types.Receipt, error), _uAddr common.Address, _cid string, _cName string, _issueTime uint64, _validityPeriod uint64) (*types.Transaction, error) {
//	return _CertBank.Contract.AsyncAddCertificate(handler, &_CertBank.TransactOpts, _uAddr, _cid, _cName, _issueTime, _validityPeriod)
//}
//
//// AddCredit is a paid mutator transaction binding the contract method 0x7587dc2b.
////
//// Solidity: function addCredit(address _uAddr, uint8 _score, uint8 _type) returns()
//func (_CertBank *CertBankTransactor) AddCredit(opts *bind.TransactOpts, _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.contract.Transact(opts, "addCredit", _uAddr, _score, _type)
//}
//
//func (_CertBank *CertBankTransactor) AsyncAddCredit(handler func(*types.Receipt, error), opts *bind.TransactOpts, _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, error) {
//	return _CertBank.contract.AsyncTransact(opts, handler, "addCredit", _uAddr, _score, _type)
//}
//
//// AddCredit is a paid mutator transaction binding the contract method 0x7587dc2b.
////
//// Solidity: function addCredit(address _uAddr, uint8 _score, uint8 _type) returns()
//func (_CertBank *CertBankSession) AddCredit(_uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.AddCredit(&_CertBank.TransactOpts, _uAddr, _score, _type)
//}
//
//func (_CertBank *CertBankSession) AsyncAddCredit(handler func(*types.Receipt, error), _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, error) {
//	return _CertBank.Contract.AsyncAddCredit(handler, &_CertBank.TransactOpts, _uAddr, _score, _type)
//}
//
//// AddCredit is a paid mutator transaction binding the contract method 0x7587dc2b.
////
//// Solidity: function addCredit(address _uAddr, uint8 _score, uint8 _type) returns()
//func (_CertBank *CertBankTransactorSession) AddCredit(_uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, *types.Receipt, error) {
//	return _CertBank.Contract.AddCredit(&_CertBank.TransactOpts, _uAddr, _score, _type)
//}
//
//func (_CertBank *CertBankTransactorSession) AsyncAddCredit(handler func(*types.Receipt, error), _uAddr common.Address, _score uint8, _type uint8) (*types.Transaction, error) {
//	return _CertBank.Contract.AsyncAddCredit(handler, &_CertBank.TransactOpts, _uAddr, _score, _type)
//}
