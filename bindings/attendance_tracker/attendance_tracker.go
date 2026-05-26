// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package attendance_tracker

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

// AttendanceTrackerMetaData contains all meta data concerning the AttendanceTracker contract.
var AttendanceTrackerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attendance\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintAttendance\",\"inputs\":[{\"name\":\"memberIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"eventId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AttendanceMinted\",\"inputs\":[{\"name\":\"memberId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"eventId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001a5f3361004b565b506100457f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a63361004b565b506100f4565b5f828152602081815260408083206001600160a01b038516845290915281205460ff166100eb575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556100a33390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016100ee565b505f5b92915050565b610677806101015f395ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c8063a217fddf11610063578063a217fddf14610132578063cafa063614610139578063d53913931461014c578063d547741f14610173578063eb27153314610186575f80fd5b806301ffc9a71461009f578063248a9ca3146100c75780632f2ff15d146100f757806336568abe1461010c57806391d148541461011f575b5f80fd5b6100b26100ad3660046104a3565b6101a5565b60405190151581526020015b60405180910390f35b6100e96100d53660046104d1565b5f9081526020819052604090206001015490565b6040519081526020016100be565b61010a6101053660046104e8565b6101db565b005b61010a61011a3660046104e8565b610205565b6100b261012d3660046104e8565b61023d565b6100e95f81565b61010a610147366004610521565b610265565b6100e97f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b61010a6101813660046104e8565b610339565b6100e96101943660046104d1565b60016020525f908152604090205481565b5f6001600160e01b03198216637965db0b60e01b14806101d557506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f828152602081905260409020600101546101f58161035d565b6101ff838361036a565b50505050565b6001600160a01b038116331461022e5760405163334bd91960e11b815260040160405180910390fd5b61023882826103f9565b505050565b5f918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a661028f8161035d565b5f5b848110156103315760015f8787848181106102ae576102ae6105db565b9050602002013581526020019081526020015f205f8154809291906102d2906105ef565b91905055508585828181106102e9576102e96105db565b905060200201357fb4462a03543ab44c07dc3716e3f35587b03f0878d48d5cce6bd114340d8658088585604051610321929190610613565b60405180910390a2600101610291565b505050505050565b5f828152602081905260409020600101546103538161035d565b6101ff83836103f9565b6103678133610462565b50565b5f610375838361023d565b6103f2575f838152602081815260408083206001600160a01b03861684529091529020805460ff191660011790556103aa3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016101d5565b505f6101d5565b5f610404838361023d565b156103f2575f838152602081815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016101d5565b61046c828261023d565b61049f5760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440160405180910390fd5b5050565b5f602082840312156104b3575f80fd5b81356001600160e01b0319811681146104ca575f80fd5b9392505050565b5f602082840312156104e1575f80fd5b5035919050565b5f80604083850312156104f9575f80fd5b8235915060208301356001600160a01b0381168114610516575f80fd5b809150509250929050565b5f805f8060408587031215610534575f80fd5b843567ffffffffffffffff8082111561054b575f80fd5b818701915087601f83011261055e575f80fd5b81358181111561056c575f80fd5b8860208260051b8501011115610580575f80fd5b60209283019650945090860135908082111561059a575f80fd5b818701915087601f8301126105ad575f80fd5b8135818111156105bb575f80fd5b8860208285010111156105cc575f80fd5b95989497505060200194505050565b634e487b7160e01b5f52603260045260245ffd5b5f6001820161060c57634e487b7160e01b5f52601160045260245ffd5b5060010190565b60208152816020820152818360408301375f818301604090810191909152601f909201601f1916010191905056fea26469706673582212207eec9d464426b6b2881fa8a4253ac5670e1b4cd94784a950d105125b30d63fa164736f6c63430008180033",
}

// AttendanceTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use AttendanceTrackerMetaData.ABI instead.
var AttendanceTrackerABI = AttendanceTrackerMetaData.ABI

// AttendanceTrackerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttendanceTrackerMetaData.Bin instead.
var AttendanceTrackerBin = AttendanceTrackerMetaData.Bin

// DeployAttendanceTracker deploys a new Ethereum contract, binding an instance of AttendanceTracker to it.
func DeployAttendanceTracker(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttendanceTracker, error) {
	parsed, err := AttendanceTrackerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttendanceTrackerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttendanceTracker{AttendanceTrackerCaller: AttendanceTrackerCaller{contract: contract}, AttendanceTrackerTransactor: AttendanceTrackerTransactor{contract: contract}, AttendanceTrackerFilterer: AttendanceTrackerFilterer{contract: contract}}, nil
}

// AttendanceTracker is an auto generated Go binding around an Ethereum contract.
type AttendanceTracker struct {
	AttendanceTrackerCaller     // Read-only binding to the contract
	AttendanceTrackerTransactor // Write-only binding to the contract
	AttendanceTrackerFilterer   // Log filterer for contract events
}

// AttendanceTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttendanceTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttendanceTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttendanceTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttendanceTrackerSession struct {
	Contract     *AttendanceTracker // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AttendanceTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttendanceTrackerCallerSession struct {
	Contract *AttendanceTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AttendanceTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttendanceTrackerTransactorSession struct {
	Contract     *AttendanceTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AttendanceTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttendanceTrackerRaw struct {
	Contract *AttendanceTracker // Generic contract binding to access the raw methods on
}

// AttendanceTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttendanceTrackerCallerRaw struct {
	Contract *AttendanceTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// AttendanceTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttendanceTrackerTransactorRaw struct {
	Contract *AttendanceTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttendanceTracker creates a new instance of AttendanceTracker, bound to a specific deployed contract.
func NewAttendanceTracker(address common.Address, backend bind.ContractBackend) (*AttendanceTracker, error) {
	contract, err := bindAttendanceTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttendanceTracker{AttendanceTrackerCaller: AttendanceTrackerCaller{contract: contract}, AttendanceTrackerTransactor: AttendanceTrackerTransactor{contract: contract}, AttendanceTrackerFilterer: AttendanceTrackerFilterer{contract: contract}}, nil
}

// NewAttendanceTrackerCaller creates a new read-only instance of AttendanceTracker, bound to a specific deployed contract.
func NewAttendanceTrackerCaller(address common.Address, caller bind.ContractCaller) (*AttendanceTrackerCaller, error) {
	contract, err := bindAttendanceTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceTrackerCaller{contract: contract}, nil
}

// NewAttendanceTrackerTransactor creates a new write-only instance of AttendanceTracker, bound to a specific deployed contract.
func NewAttendanceTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*AttendanceTrackerTransactor, error) {
	contract, err := bindAttendanceTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceTrackerTransactor{contract: contract}, nil
}

// NewAttendanceTrackerFilterer creates a new log filterer instance of AttendanceTracker, bound to a specific deployed contract.
func NewAttendanceTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*AttendanceTrackerFilterer, error) {
	contract, err := bindAttendanceTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttendanceTrackerFilterer{contract: contract}, nil
}

// bindAttendanceTracker binds a generic wrapper to an already deployed contract.
func bindAttendanceTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttendanceTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttendanceTracker *AttendanceTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttendanceTracker.Contract.AttendanceTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttendanceTracker *AttendanceTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.AttendanceTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttendanceTracker *AttendanceTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.AttendanceTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttendanceTracker *AttendanceTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttendanceTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttendanceTracker *AttendanceTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttendanceTracker *AttendanceTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AttendanceTracker.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AttendanceTracker.Contract.DEFAULTADMINROLE(&_AttendanceTracker.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AttendanceTracker.Contract.DEFAULTADMINROLE(&_AttendanceTracker.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AttendanceTracker.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerSession) MINTERROLE() ([32]byte, error) {
	return _AttendanceTracker.Contract.MINTERROLE(&_AttendanceTracker.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerCallerSession) MINTERROLE() ([32]byte, error) {
	return _AttendanceTracker.Contract.MINTERROLE(&_AttendanceTracker.CallOpts)
}

// Attendance is a free data retrieval call binding the contract method 0xeb271533.
//
// Solidity: function attendance(uint256 ) view returns(uint256)
func (_AttendanceTracker *AttendanceTrackerCaller) Attendance(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AttendanceTracker.contract.Call(opts, &out, "attendance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Attendance is a free data retrieval call binding the contract method 0xeb271533.
//
// Solidity: function attendance(uint256 ) view returns(uint256)
func (_AttendanceTracker *AttendanceTrackerSession) Attendance(arg0 *big.Int) (*big.Int, error) {
	return _AttendanceTracker.Contract.Attendance(&_AttendanceTracker.CallOpts, arg0)
}

// Attendance is a free data retrieval call binding the contract method 0xeb271533.
//
// Solidity: function attendance(uint256 ) view returns(uint256)
func (_AttendanceTracker *AttendanceTrackerCallerSession) Attendance(arg0 *big.Int) (*big.Int, error) {
	return _AttendanceTracker.Contract.Attendance(&_AttendanceTracker.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AttendanceTracker.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AttendanceTracker.Contract.GetRoleAdmin(&_AttendanceTracker.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AttendanceTracker *AttendanceTrackerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AttendanceTracker.Contract.GetRoleAdmin(&_AttendanceTracker.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AttendanceTracker *AttendanceTrackerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AttendanceTracker.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AttendanceTracker *AttendanceTrackerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AttendanceTracker.Contract.HasRole(&_AttendanceTracker.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AttendanceTracker *AttendanceTrackerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AttendanceTracker.Contract.HasRole(&_AttendanceTracker.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AttendanceTracker *AttendanceTrackerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AttendanceTracker.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AttendanceTracker *AttendanceTrackerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AttendanceTracker.Contract.SupportsInterface(&_AttendanceTracker.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AttendanceTracker *AttendanceTrackerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AttendanceTracker.Contract.SupportsInterface(&_AttendanceTracker.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AttendanceTracker *AttendanceTrackerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AttendanceTracker *AttendanceTrackerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.GrantRole(&_AttendanceTracker.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AttendanceTracker *AttendanceTrackerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.GrantRole(&_AttendanceTracker.TransactOpts, role, account)
}

// MintAttendance is a paid mutator transaction binding the contract method 0xcafa0636.
//
// Solidity: function mintAttendance(uint256[] memberIds, string eventId) returns()
func (_AttendanceTracker *AttendanceTrackerTransactor) MintAttendance(opts *bind.TransactOpts, memberIds []*big.Int, eventId string) (*types.Transaction, error) {
	return _AttendanceTracker.contract.Transact(opts, "mintAttendance", memberIds, eventId)
}

// MintAttendance is a paid mutator transaction binding the contract method 0xcafa0636.
//
// Solidity: function mintAttendance(uint256[] memberIds, string eventId) returns()
func (_AttendanceTracker *AttendanceTrackerSession) MintAttendance(memberIds []*big.Int, eventId string) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.MintAttendance(&_AttendanceTracker.TransactOpts, memberIds, eventId)
}

// MintAttendance is a paid mutator transaction binding the contract method 0xcafa0636.
//
// Solidity: function mintAttendance(uint256[] memberIds, string eventId) returns()
func (_AttendanceTracker *AttendanceTrackerTransactorSession) MintAttendance(memberIds []*big.Int, eventId string) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.MintAttendance(&_AttendanceTracker.TransactOpts, memberIds, eventId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AttendanceTracker *AttendanceTrackerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AttendanceTracker *AttendanceTrackerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.RenounceRole(&_AttendanceTracker.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AttendanceTracker *AttendanceTrackerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.RenounceRole(&_AttendanceTracker.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AttendanceTracker *AttendanceTrackerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AttendanceTracker *AttendanceTrackerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.RevokeRole(&_AttendanceTracker.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AttendanceTracker *AttendanceTrackerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceTracker.Contract.RevokeRole(&_AttendanceTracker.TransactOpts, role, account)
}

// AttendanceTrackerAttendanceMintedIterator is returned from FilterAttendanceMinted and is used to iterate over the raw logs and unpacked data for AttendanceMinted events raised by the AttendanceTracker contract.
type AttendanceTrackerAttendanceMintedIterator struct {
	Event *AttendanceTrackerAttendanceMinted // Event containing the contract specifics and raw log

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
func (it *AttendanceTrackerAttendanceMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTrackerAttendanceMinted)
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
		it.Event = new(AttendanceTrackerAttendanceMinted)
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
func (it *AttendanceTrackerAttendanceMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTrackerAttendanceMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTrackerAttendanceMinted represents a AttendanceMinted event raised by the AttendanceTracker contract.
type AttendanceTrackerAttendanceMinted struct {
	MemberId *big.Int
	EventId  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttendanceMinted is a free log retrieval operation binding the contract event 0xb4462a03543ab44c07dc3716e3f35587b03f0878d48d5cce6bd114340d865808.
//
// Solidity: event AttendanceMinted(uint256 indexed memberId, string eventId)
func (_AttendanceTracker *AttendanceTrackerFilterer) FilterAttendanceMinted(opts *bind.FilterOpts, memberId []*big.Int) (*AttendanceTrackerAttendanceMintedIterator, error) {

	var memberIdRule []interface{}
	for _, memberIdItem := range memberId {
		memberIdRule = append(memberIdRule, memberIdItem)
	}

	logs, sub, err := _AttendanceTracker.contract.FilterLogs(opts, "AttendanceMinted", memberIdRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTrackerAttendanceMintedIterator{contract: _AttendanceTracker.contract, event: "AttendanceMinted", logs: logs, sub: sub}, nil
}

// WatchAttendanceMinted is a free log subscription operation binding the contract event 0xb4462a03543ab44c07dc3716e3f35587b03f0878d48d5cce6bd114340d865808.
//
// Solidity: event AttendanceMinted(uint256 indexed memberId, string eventId)
func (_AttendanceTracker *AttendanceTrackerFilterer) WatchAttendanceMinted(opts *bind.WatchOpts, sink chan<- *AttendanceTrackerAttendanceMinted, memberId []*big.Int) (event.Subscription, error) {

	var memberIdRule []interface{}
	for _, memberIdItem := range memberId {
		memberIdRule = append(memberIdRule, memberIdItem)
	}

	logs, sub, err := _AttendanceTracker.contract.WatchLogs(opts, "AttendanceMinted", memberIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTrackerAttendanceMinted)
				if err := _AttendanceTracker.contract.UnpackLog(event, "AttendanceMinted", log); err != nil {
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

// ParseAttendanceMinted is a log parse operation binding the contract event 0xb4462a03543ab44c07dc3716e3f35587b03f0878d48d5cce6bd114340d865808.
//
// Solidity: event AttendanceMinted(uint256 indexed memberId, string eventId)
func (_AttendanceTracker *AttendanceTrackerFilterer) ParseAttendanceMinted(log types.Log) (*AttendanceTrackerAttendanceMinted, error) {
	event := new(AttendanceTrackerAttendanceMinted)
	if err := _AttendanceTracker.contract.UnpackLog(event, "AttendanceMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTrackerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AttendanceTracker contract.
type AttendanceTrackerRoleAdminChangedIterator struct {
	Event *AttendanceTrackerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AttendanceTrackerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTrackerRoleAdminChanged)
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
		it.Event = new(AttendanceTrackerRoleAdminChanged)
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
func (it *AttendanceTrackerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTrackerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTrackerRoleAdminChanged represents a RoleAdminChanged event raised by the AttendanceTracker contract.
type AttendanceTrackerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AttendanceTracker *AttendanceTrackerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AttendanceTrackerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AttendanceTracker.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTrackerRoleAdminChangedIterator{contract: _AttendanceTracker.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AttendanceTracker *AttendanceTrackerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AttendanceTrackerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AttendanceTracker.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTrackerRoleAdminChanged)
				if err := _AttendanceTracker.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AttendanceTracker *AttendanceTrackerFilterer) ParseRoleAdminChanged(log types.Log) (*AttendanceTrackerRoleAdminChanged, error) {
	event := new(AttendanceTrackerRoleAdminChanged)
	if err := _AttendanceTracker.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTrackerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AttendanceTracker contract.
type AttendanceTrackerRoleGrantedIterator struct {
	Event *AttendanceTrackerRoleGranted // Event containing the contract specifics and raw log

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
func (it *AttendanceTrackerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTrackerRoleGranted)
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
		it.Event = new(AttendanceTrackerRoleGranted)
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
func (it *AttendanceTrackerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTrackerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTrackerRoleGranted represents a RoleGranted event raised by the AttendanceTracker contract.
type AttendanceTrackerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceTracker *AttendanceTrackerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AttendanceTrackerRoleGrantedIterator, error) {

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

	logs, sub, err := _AttendanceTracker.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTrackerRoleGrantedIterator{contract: _AttendanceTracker.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceTracker *AttendanceTrackerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AttendanceTrackerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AttendanceTracker.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTrackerRoleGranted)
				if err := _AttendanceTracker.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AttendanceTracker *AttendanceTrackerFilterer) ParseRoleGranted(log types.Log) (*AttendanceTrackerRoleGranted, error) {
	event := new(AttendanceTrackerRoleGranted)
	if err := _AttendanceTracker.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTrackerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AttendanceTracker contract.
type AttendanceTrackerRoleRevokedIterator struct {
	Event *AttendanceTrackerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AttendanceTrackerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTrackerRoleRevoked)
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
		it.Event = new(AttendanceTrackerRoleRevoked)
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
func (it *AttendanceTrackerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTrackerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTrackerRoleRevoked represents a RoleRevoked event raised by the AttendanceTracker contract.
type AttendanceTrackerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceTracker *AttendanceTrackerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AttendanceTrackerRoleRevokedIterator, error) {

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

	logs, sub, err := _AttendanceTracker.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTrackerRoleRevokedIterator{contract: _AttendanceTracker.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceTracker *AttendanceTrackerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AttendanceTrackerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AttendanceTracker.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTrackerRoleRevoked)
				if err := _AttendanceTracker.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AttendanceTracker *AttendanceTrackerFilterer) ParseRoleRevoked(log types.Log) (*AttendanceTrackerRoleRevoked, error) {
	event := new(AttendanceTrackerRoleRevoked)
	if err := _AttendanceTracker.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
