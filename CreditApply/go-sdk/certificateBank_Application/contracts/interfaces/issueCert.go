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

// IssueCertABI is the input ABI used to generate the binding from.
const IssueCertABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"document\",\"type\":\"bytes32\"}],\"name\":\"certIssuedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"document\",\"type\":\"bytes32\"}],\"name\":\"certRevokedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_role\",\"type\":\"address\"}],\"name\":\"addRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"certIssued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"certRevoked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"isIssued\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"isRevoked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"template\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cert\",\"type\":\"bytes32\"}],\"name\":\"revoke\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IssueCert is an auto generated Go binding around a Solidity contract.
type IssueCert struct {
	IssueCertCaller     // Read-only binding to the contract
	IssueCertTransactor // Write-only binding to the contract
	IssueCertFilterer   // Log filterer for contract events
}

// IssueCertCaller is an auto generated read-only Go binding around a Solidity contract.
type IssueCertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssueCertTransactor is an auto generated write-only Go binding around a Solidity contract.
type IssueCertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssueCertFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IssueCertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IssueCertSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IssueCertSession struct {
	Contract     *IssueCert        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IssueCertCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IssueCertCallerSession struct {
	Contract *IssueCertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IssueCertTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IssueCertTransactorSession struct {
	Contract     *IssueCertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IssueCertRaw is an auto generated low-level Go binding around a Solidity contract.
type IssueCertRaw struct {
	Contract *IssueCert // Generic contract binding to access the raw methods on
}

// IssueCertCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IssueCertCallerRaw struct {
	Contract *IssueCertCaller // Generic read-only contract binding to access the raw methods on
}

// IssueCertTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IssueCertTransactorRaw struct {
	Contract *IssueCertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIssueCert creates a new instance of IssueCert, bound to a specific deployed contract.
func NewIssueCert(address common.Address, backend bind.ContractBackend) (*IssueCert, error) {
	contract, err := bindIssueCert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IssueCert{IssueCertCaller: IssueCertCaller{contract: contract}, IssueCertTransactor: IssueCertTransactor{contract: contract}, IssueCertFilterer: IssueCertFilterer{contract: contract}}, nil
}

// NewIssueCertCaller creates a new read-only instance of IssueCert, bound to a specific deployed contract.
func NewIssueCertCaller(address common.Address, caller bind.ContractCaller) (*IssueCertCaller, error) {
	contract, err := bindIssueCert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IssueCertCaller{contract: contract}, nil
}

// NewIssueCertTransactor creates a new write-only instance of IssueCert, bound to a specific deployed contract.
func NewIssueCertTransactor(address common.Address, transactor bind.ContractTransactor) (*IssueCertTransactor, error) {
	contract, err := bindIssueCert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IssueCertTransactor{contract: contract}, nil
}

// NewIssueCertFilterer creates a new log filterer instance of IssueCert, bound to a specific deployed contract.
func NewIssueCertFilterer(address common.Address, filterer bind.ContractFilterer) (*IssueCertFilterer, error) {
	contract, err := bindIssueCert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IssueCertFilterer{contract: contract}, nil
}

// bindIssueCert binds a generic wrapper to an already deployed contract.
func bindIssueCert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IssueCertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IssueCert *IssueCertRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IssueCert.Contract.IssueCertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IssueCert *IssueCertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.IssueCertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IssueCert *IssueCertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.IssueCertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IssueCert *IssueCertCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IssueCert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IssueCert *IssueCertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IssueCert *IssueCertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.contract.Transact(opts, method, params...)
}

// CertIssued is a free data retrieval call binding the contract method 0xbdb92387.
//
// Solidity: function certIssued(bytes32 ) constant returns(uint256)
func (_IssueCert *IssueCertCaller) CertIssued(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IssueCert.contract.Call(opts, out, "certIssued", arg0)
	return *ret0, err
}

// CertIssued is a free data retrieval call binding the contract method 0xbdb92387.
//
// Solidity: function certIssued(bytes32 ) constant returns(uint256)
func (_IssueCert *IssueCertSession) CertIssued(arg0 [32]byte) (*big.Int, error) {
	return _IssueCert.Contract.CertIssued(&_IssueCert.CallOpts, arg0)
}

// CertIssued is a free data retrieval call binding the contract method 0xbdb92387.
//
// Solidity: function certIssued(bytes32 ) constant returns(uint256)
func (_IssueCert *IssueCertCallerSession) CertIssued(arg0 [32]byte) (*big.Int, error) {
	return _IssueCert.Contract.CertIssued(&_IssueCert.CallOpts, arg0)
}

// CertRevoked is a free data retrieval call binding the contract method 0x1dc4e16f.
//
// Solidity: function certRevoked(bytes32 ) constant returns(uint256)
func (_IssueCert *IssueCertCaller) CertRevoked(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IssueCert.contract.Call(opts, out, "certRevoked", arg0)
	return *ret0, err
}

// CertRevoked is a free data retrieval call binding the contract method 0x1dc4e16f.
//
// Solidity: function certRevoked(bytes32 ) constant returns(uint256)
func (_IssueCert *IssueCertSession) CertRevoked(arg0 [32]byte) (*big.Int, error) {
	return _IssueCert.Contract.CertRevoked(&_IssueCert.CallOpts, arg0)
}

// CertRevoked is a free data retrieval call binding the contract method 0x1dc4e16f.
//
// Solidity: function certRevoked(bytes32 ) constant returns(uint256)
func (_IssueCert *IssueCertCallerSession) CertRevoked(arg0 [32]byte) (*big.Int, error) {
	return _IssueCert.Contract.CertRevoked(&_IssueCert.CallOpts, arg0)
}

// IsIssued is a free data retrieval call binding the contract method 0x163aa631.
//
// Solidity: function isIssued(bytes32 cert) constant returns(bool)
func (_IssueCert *IssueCertCaller) IsIssued(opts *bind.CallOpts, cert [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IssueCert.contract.Call(opts, out, "isIssued", cert)
	return *ret0, err
}

// IsIssued is a free data retrieval call binding the contract method 0x163aa631.
//
// Solidity: function isIssued(bytes32 cert) constant returns(bool)
func (_IssueCert *IssueCertSession) IsIssued(cert [32]byte) (bool, error) {
	return _IssueCert.Contract.IsIssued(&_IssueCert.CallOpts, cert)
}

// IsIssued is a free data retrieval call binding the contract method 0x163aa631.
//
// Solidity: function isIssued(bytes32 cert) constant returns(bool)
func (_IssueCert *IssueCertCallerSession) IsIssued(cert [32]byte) (bool, error) {
	return _IssueCert.Contract.IsIssued(&_IssueCert.CallOpts, cert)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 cert) constant returns(bool)
func (_IssueCert *IssueCertCaller) IsRevoked(opts *bind.CallOpts, cert [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IssueCert.contract.Call(opts, out, "isRevoked", cert)
	return *ret0, err
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 cert) constant returns(bool)
func (_IssueCert *IssueCertSession) IsRevoked(cert [32]byte) (bool, error) {
	return _IssueCert.Contract.IsRevoked(&_IssueCert.CallOpts, cert)
}

// IsRevoked is a free data retrieval call binding the contract method 0x4294857f.
//
// Solidity: function isRevoked(bytes32 cert) constant returns(bool)
func (_IssueCert *IssueCertCallerSession) IsRevoked(cert [32]byte) (bool, error) {
	return _IssueCert.Contract.IsRevoked(&_IssueCert.CallOpts, cert)
}

// IsRole is a free data retrieval call binding the contract method 0x7beeb945.
//
// Solidity: function isRole(address ) constant returns(bool)
func (_IssueCert *IssueCertCaller) IsRole(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IssueCert.contract.Call(opts, out, "isRole", arg0)
	return *ret0, err
}

// IsRole is a free data retrieval call binding the contract method 0x7beeb945.
//
// Solidity: function isRole(address ) constant returns(bool)
func (_IssueCert *IssueCertSession) IsRole(arg0 common.Address) (bool, error) {
	return _IssueCert.Contract.IsRole(&_IssueCert.CallOpts, arg0)
}

// IsRole is a free data retrieval call binding the contract method 0x7beeb945.
//
// Solidity: function isRole(address ) constant returns(bool)
func (_IssueCert *IssueCertCallerSession) IsRole(arg0 common.Address) (bool, error) {
	return _IssueCert.Contract.IsRole(&_IssueCert.CallOpts, arg0)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() constant returns(address)
func (_IssueCert *IssueCertCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IssueCert.contract.Call(opts, out, "master")
	return *ret0, err
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() constant returns(address)
func (_IssueCert *IssueCertSession) Master() (common.Address, error) {
	return _IssueCert.Contract.Master(&_IssueCert.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() constant returns(address)
func (_IssueCert *IssueCertCallerSession) Master() (common.Address, error) {
	return _IssueCert.Contract.Master(&_IssueCert.CallOpts)
}

// AddRole is a paid mutator transaction binding the contract method 0x3f877f72.
//
// Solidity: function addRole(address _role) returns()
func (_IssueCert *IssueCertTransactor) AddRole(opts *bind.TransactOpts, _role common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.contract.Transact(opts, "addRole", _role)
}

func (_IssueCert *IssueCertTransactor) AsyncAddRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, _role common.Address) (*types.Transaction, error) {
	return _IssueCert.contract.AsyncTransact(opts, handler, "addRole", _role)
}

// AddRole is a paid mutator transaction binding the contract method 0x3f877f72.
//
// Solidity: function addRole(address _role) returns()
func (_IssueCert *IssueCertSession) AddRole(_role common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.AddRole(&_IssueCert.TransactOpts, _role)
}

func (_IssueCert *IssueCertSession) AsyncAddRole(handler func(*types.Receipt, error), _role common.Address) (*types.Transaction, error) {
	return _IssueCert.Contract.AsyncAddRole(handler, &_IssueCert.TransactOpts, _role)
}

// AddRole is a paid mutator transaction binding the contract method 0x3f877f72.
//
// Solidity: function addRole(address _role) returns()
func (_IssueCert *IssueCertTransactorSession) AddRole(_role common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.AddRole(&_IssueCert.TransactOpts, _role)
}

func (_IssueCert *IssueCertTransactorSession) AsyncAddRole(handler func(*types.Receipt, error), _role common.Address) (*types.Transaction, error) {
	return _IssueCert.Contract.AsyncAddRole(handler, &_IssueCert.TransactOpts, _role)
}

// Issue is a paid mutator transaction binding the contract method 0x0f75e81f.
//
// Solidity: function issue(bytes32 cert) returns()
func (_IssueCert *IssueCertTransactor) Issue(opts *bind.TransactOpts, cert [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.contract.Transact(opts, "issue", cert)
}

func (_IssueCert *IssueCertTransactor) AsyncIssue(handler func(*types.Receipt, error), opts *bind.TransactOpts, cert [32]byte) (*types.Transaction, error) {
	return _IssueCert.contract.AsyncTransact(opts, handler, "issue", cert)
}

// Issue is a paid mutator transaction binding the contract method 0x0f75e81f.
//
// Solidity: function issue(bytes32 cert) returns()
func (_IssueCert *IssueCertSession) Issue(cert [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.Issue(&_IssueCert.TransactOpts, cert)
}

func (_IssueCert *IssueCertSession) AsyncIssue(handler func(*types.Receipt, error), cert [32]byte) (*types.Transaction, error) {
	return _IssueCert.Contract.AsyncIssue(handler, &_IssueCert.TransactOpts, cert)
}

// Issue is a paid mutator transaction binding the contract method 0x0f75e81f.
//
// Solidity: function issue(bytes32 cert) returns()
func (_IssueCert *IssueCertTransactorSession) Issue(cert [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.Issue(&_IssueCert.TransactOpts, cert)
}

func (_IssueCert *IssueCertTransactorSession) AsyncIssue(handler func(*types.Receipt, error), cert [32]byte) (*types.Transaction, error) {
	return _IssueCert.Contract.AsyncIssue(handler, &_IssueCert.TransactOpts, cert)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 cert) returns(bool)
func (_IssueCert *IssueCertTransactor) Revoke(opts *bind.TransactOpts, cert [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.contract.Transact(opts, "revoke", cert)
}

func (_IssueCert *IssueCertTransactor) AsyncRevoke(handler func(*types.Receipt, error), opts *bind.TransactOpts, cert [32]byte) (*types.Transaction, error) {
	return _IssueCert.contract.AsyncTransact(opts, handler, "revoke", cert)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 cert) returns(bool)
func (_IssueCert *IssueCertSession) Revoke(cert [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.Revoke(&_IssueCert.TransactOpts, cert)
}

func (_IssueCert *IssueCertSession) AsyncRevoke(handler func(*types.Receipt, error), cert [32]byte) (*types.Transaction, error) {
	return _IssueCert.Contract.AsyncRevoke(handler, &_IssueCert.TransactOpts, cert)
}

// Revoke is a paid mutator transaction binding the contract method 0xb75c7dc6.
//
// Solidity: function revoke(bytes32 cert) returns(bool)
func (_IssueCert *IssueCertTransactorSession) Revoke(cert [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IssueCert.Contract.Revoke(&_IssueCert.TransactOpts, cert)
}

func (_IssueCert *IssueCertTransactorSession) AsyncRevoke(handler func(*types.Receipt, error), cert [32]byte) (*types.Transaction, error) {
	return _IssueCert.Contract.AsyncRevoke(handler, &_IssueCert.TransactOpts, cert)
}

// IssueCertCertIssuedEvent represents a CertIssuedEvent event raised by the IssueCert contract.
type IssueCertCertIssuedEvent struct {
	Document [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchCertIssuedEvent is a free log subscription operation binding the contract event 0x9ad45bd7d1ae19105b76d621dada359f6c746bfc62b3e5eedba36df2d7710d56.
//
// Solidity: event certIssuedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertFilterer) WatchCertIssuedEvent(fromBlock *uint64, handler func(int, []types.Log), document [32]byte) error {
	return _IssueCert.contract.WatchLogs(fromBlock, handler, "certIssuedEvent", document)
}

func (_IssueCert *IssueCertFilterer) WatchAllCertIssuedEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _IssueCert.contract.WatchLogs(fromBlock, handler, "certIssuedEvent")
}

// ParseCertIssuedEvent is a log parse operation binding the contract event 0x9ad45bd7d1ae19105b76d621dada359f6c746bfc62b3e5eedba36df2d7710d56.
//
// Solidity: event certIssuedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertFilterer) ParseCertIssuedEvent(log types.Log) (*IssueCertCertIssuedEvent, error) {
	event := new(IssueCertCertIssuedEvent)
	if err := _IssueCert.contract.UnpackLog(event, "certIssuedEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCertIssuedEvent is a free log subscription operation binding the contract event 0x9ad45bd7d1ae19105b76d621dada359f6c746bfc62b3e5eedba36df2d7710d56.
//
// Solidity: event certIssuedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertSession) WatchCertIssuedEvent(fromBlock *uint64, handler func(int, []types.Log), document [32]byte) error {
	return _IssueCert.Contract.WatchCertIssuedEvent(fromBlock, handler, document)
}

func (_IssueCert *IssueCertSession) WatchAllCertIssuedEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _IssueCert.Contract.WatchAllCertIssuedEvent(fromBlock, handler)
}

// ParseCertIssuedEvent is a log parse operation binding the contract event 0x9ad45bd7d1ae19105b76d621dada359f6c746bfc62b3e5eedba36df2d7710d56.
//
// Solidity: event certIssuedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertSession) ParseCertIssuedEvent(log types.Log) (*IssueCertCertIssuedEvent, error) {
	return _IssueCert.Contract.ParseCertIssuedEvent(log)
}

// IssueCertCertRevokedEvent represents a CertRevokedEvent event raised by the IssueCert contract.
type IssueCertCertRevokedEvent struct {
	Document [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// WatchCertRevokedEvent is a free log subscription operation binding the contract event 0xadce10f389cd96cfbdfbd6bed6369d6f66d77cf374c89b4bceaee22d2fa04a07.
//
// Solidity: event certRevokedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertFilterer) WatchCertRevokedEvent(fromBlock *uint64, handler func(int, []types.Log), document [32]byte) error {
	return _IssueCert.contract.WatchLogs(fromBlock, handler, "certRevokedEvent", document)
}

func (_IssueCert *IssueCertFilterer) WatchAllCertRevokedEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _IssueCert.contract.WatchLogs(fromBlock, handler, "certRevokedEvent")
}

// ParseCertRevokedEvent is a log parse operation binding the contract event 0xadce10f389cd96cfbdfbd6bed6369d6f66d77cf374c89b4bceaee22d2fa04a07.
//
// Solidity: event certRevokedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertFilterer) ParseCertRevokedEvent(log types.Log) (*IssueCertCertRevokedEvent, error) {
	event := new(IssueCertCertRevokedEvent)
	if err := _IssueCert.contract.UnpackLog(event, "certRevokedEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchCertRevokedEvent is a free log subscription operation binding the contract event 0xadce10f389cd96cfbdfbd6bed6369d6f66d77cf374c89b4bceaee22d2fa04a07.
//
// Solidity: event certRevokedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertSession) WatchCertRevokedEvent(fromBlock *uint64, handler func(int, []types.Log), document [32]byte) error {
	return _IssueCert.Contract.WatchCertRevokedEvent(fromBlock, handler, document)
}

func (_IssueCert *IssueCertSession) WatchAllCertRevokedEvent(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _IssueCert.Contract.WatchAllCertRevokedEvent(fromBlock, handler)
}

// ParseCertRevokedEvent is a log parse operation binding the contract event 0xadce10f389cd96cfbdfbd6bed6369d6f66d77cf374c89b4bceaee22d2fa04a07.
//
// Solidity: event certRevokedEvent(bytes32 indexed document)
func (_IssueCert *IssueCertSession) ParseCertRevokedEvent(log types.Log) (*IssueCertCertRevokedEvent, error) {
	return _IssueCert.Contract.ParseCertRevokedEvent(log)
}
