// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Description string
	Completed   bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"TaskAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"addTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskId\",\"type\":\"uint256\"}],\"name\":\"completeTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllTasks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610e698061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80631d65e77e14610059578063672385621461008a578063779900b4146100a65780638d977672146100c4578063e1e29558146100f5575b5f5ffd5b610073600480360381019061006e9190610628565b610111565b6040516100819291906106dd565b60405180910390f35b6100a4600480360381019061009f9190610837565b610243565b005b6100ae610301565b6040516100bb91906109ca565b60405180910390f35b6100de60048036038101906100d99190610628565b610405565b6040516100ec9291906106dd565b60405180910390f35b61010f600480360381019061010a9190610628565b6104c5565b005b60605f5f80549050831061015a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015190610a34565b60405180910390fd5b5f5f848154811061016e5761016d610a52565b5b905f5260205f2090600202016040518060400160405290815f8201805461019490610aac565b80601f01602080910402602001604051908101604052809291908181526020018280546101c090610aac565b801561020b5780601f106101e25761010080835404028352916020019161020b565b820191905f5260205f20905b8154815290600101906020018083116101ee57829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250509050805f015181602001519250925050915091565b5f60405180604001604052808381526020015f1515815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f0190816102949190610c7c565b506020820151816001015f6101000a81548160ff021916908315150217905550505060015f805490506102c79190610d78565b7f2ee7ce8bac091b03a0c391c07ee88a246bdd72bdc6b1ef1ce4d99220595d32c5826040516102f69190610dab565b60405180910390a250565b60605f805480602002602001604051908101604052809291908181526020015f905b828210156103fc578382905f5260205f2090600202016040518060400160405290815f8201805461035390610aac565b80601f016020809104026020016040519081016040528092919081815260200182805461037f90610aac565b80156103ca5780601f106103a1576101008083540402835291602001916103ca565b820191905f5260205f20905b8154815290600101906020018083116103ad57829003601f168201915b50505050508152602001600182015f9054906101000a900460ff16151515158152505081526020019060010190610323565b50505050905090565b5f8181548110610413575f80fd5b905f5260205f2090600202015f91509050805f01805461043290610aac565b80601f016020809104026020016040519081016040528092919081815260200182805461045e90610aac565b80156104a95780601f10610480576101008083540402835291602001916104a9565b820191905f5260205f20905b81548152906001019060200180831161048c57829003601f168201915b505050505090806001015f9054906101000a900460ff16905082565b5f80549050811061050b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050290610a34565b60405180910390fd5b5f818154811061051e5761051d610a52565b5b905f5260205f2090600202016001015f9054906101000a900460ff161561057a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057190610e15565b60405180910390fd5b60015f828154811061058f5761058e610a52565b5b905f5260205f2090600202016001015f6101000a81548160ff021916908315150217905550807fc1fa7142cfb933d0855114a3bffde296944c30b727f297c14d0db4d553d3a5c760405160405180910390a250565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610607816105f5565b8114610611575f5ffd5b50565b5f81359050610622816105fe565b92915050565b5f6020828403121561063d5761063c6105ed565b5b5f61064a84828501610614565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61069582610653565b61069f818561065d565b93506106af81856020860161066d565b6106b88161067b565b840191505092915050565b5f8115159050919050565b6106d7816106c3565b82525050565b5f6040820190508181035f8301526106f5818561068b565b905061070460208301846106ce565b9392505050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6107498261067b565b810181811067ffffffffffffffff8211171561076857610767610713565b5b80604052505050565b5f61077a6105e4565b90506107868282610740565b919050565b5f67ffffffffffffffff8211156107a5576107a4610713565b5b6107ae8261067b565b9050602081019050919050565b828183375f83830152505050565b5f6107db6107d68461078b565b610771565b9050828152602081018484840111156107f7576107f661070f565b5b6108028482856107bb565b509392505050565b5f82601f83011261081e5761081d61070b565b5b813561082e8482602086016107c9565b91505092915050565b5f6020828403121561084c5761084b6105ed565b5b5f82013567ffffffffffffffff811115610869576108686105f1565b5b6108758482850161080a565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f6108c182610653565b6108cb81856108a7565b93506108db81856020860161066d565b6108e48161067b565b840191505092915050565b6108f8816106c3565b82525050565b5f604083015f8301518482035f86015261091882826108b7565b915050602083015161092d60208601826108ef565b508091505092915050565b5f61094383836108fe565b905092915050565b5f602082019050919050565b5f6109618261087e565b61096b8185610888565b93508360208202850161097d85610898565b805f5b858110156109b857848403895281516109998582610938565b94506109a48361094b565b925060208a01995050600181019050610980565b50829750879550505050505092915050565b5f6020820190508181035f8301526109e28184610957565b905092915050565b7f5461736b20646f6573206e6f74206578697374000000000000000000000000005f82015250565b5f610a1e60138361065d565b9150610a29826109ea565b602082019050919050565b5f6020820190508181035f830152610a4b81610a12565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610ac357607f821691505b602082108103610ad657610ad5610a7f565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610b387fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610afd565b610b428683610afd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610b7d610b78610b73846105f5565b610b5a565b6105f5565b9050919050565b5f819050919050565b610b9683610b63565b610baa610ba282610b84565b848454610b09565b825550505050565b5f5f905090565b610bc1610bb2565b610bcc818484610b8d565b505050565b5b81811015610bef57610be45f82610bb9565b600181019050610bd2565b5050565b601f821115610c3457610c0581610adc565b610c0e84610aee565b81016020851015610c1d578190505b610c31610c2985610aee565b830182610bd1565b50505b505050565b5f82821c905092915050565b5f610c545f1984600802610c39565b1980831691505092915050565b5f610c6c8383610c45565b9150826002028217905092915050565b610c8582610653565b67ffffffffffffffff811115610c9e57610c9d610713565b5b610ca88254610aac565b610cb3828285610bf3565b5f60209050601f831160018114610ce4575f8415610cd2578287015190505b610cdc8582610c61565b865550610d43565b601f198416610cf286610adc565b5f5b82811015610d1957848901518255600182019150602085019450602081019050610cf4565b86831015610d365784890151610d32601f891682610c45565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610d82826105f5565b9150610d8d836105f5565b9250828203905081811115610da557610da4610d4b565b5b92915050565b5f6020820190508181035f830152610dc3818461068b565b905092915050565b7f5461736b20616c726561647920636f6d706c65746564000000000000000000005f82015250565b5f610dff60168361065d565b9150610e0a82610dcb565b602082019050919050565b5f6020820190508181035f830152610e2c81610df3565b905091905056fea26469706673582212204765628d24a226b37493b501bfcac025b4ad5bca8104b9b07f3b24e31315263264736f6c634300081e0033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() view returns((string,bool)[])
func (_Todo *TodoCaller) GetAllTasks(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getAllTasks")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() view returns((string,bool)[])
func (_Todo *TodoSession) GetAllTasks() ([]TodoTask, error) {
	return _Todo.Contract.GetAllTasks(&_Todo.CallOpts)
}

// GetAllTasks is a free data retrieval call binding the contract method 0x779900b4.
//
// Solidity: function getAllTasks() view returns((string,bool)[])
func (_Todo *TodoCallerSession) GetAllTasks() ([]TodoTask, error) {
	return _Todo.Contract.GetAllTasks(&_Todo.CallOpts)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskId) view returns(string, bool)
func (_Todo *TodoCaller) GetTask(opts *bind.CallOpts, _taskId *big.Int) (string, bool, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "getTask", _taskId)

	if err != nil {
		return *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskId) view returns(string, bool)
func (_Todo *TodoSession) GetTask(_taskId *big.Int) (string, bool, error) {
	return _Todo.Contract.GetTask(&_Todo.CallOpts, _taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskId) view returns(string, bool)
func (_Todo *TodoCallerSession) GetTask(_taskId *big.Int) (string, bool, error) {
	return _Todo.Contract.GetTask(&_Todo.CallOpts, _taskId)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string description, bool completed)
func (_Todo *TodoCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description string
	Completed   bool
}, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Description string
		Completed   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Description = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Completed = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string description, bool completed)
func (_Todo *TodoSession) Tasks(arg0 *big.Int) (struct {
	Description string
	Completed   bool
}, error) {
	return _Todo.Contract.Tasks(&_Todo.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string description, bool completed)
func (_Todo *TodoCallerSession) Tasks(arg0 *big.Int) (struct {
	Description string
	Completed   bool
}, error) {
	return _Todo.Contract.Tasks(&_Todo.CallOpts, arg0)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _description) returns()
func (_Todo *TodoTransactor) AddTask(opts *bind.TransactOpts, _description string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "addTask", _description)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _description) returns()
func (_Todo *TodoSession) AddTask(_description string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _description)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _description) returns()
func (_Todo *TodoTransactorSession) AddTask(_description string) (*types.Transaction, error) {
	return _Todo.Contract.AddTask(&_Todo.TransactOpts, _description)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 _taskId) returns()
func (_Todo *TodoTransactor) CompleteTask(opts *bind.TransactOpts, _taskId *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "completeTask", _taskId)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 _taskId) returns()
func (_Todo *TodoSession) CompleteTask(_taskId *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.CompleteTask(&_Todo.TransactOpts, _taskId)
}

// CompleteTask is a paid mutator transaction binding the contract method 0xe1e29558.
//
// Solidity: function completeTask(uint256 _taskId) returns()
func (_Todo *TodoTransactorSession) CompleteTask(_taskId *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.CompleteTask(&_Todo.TransactOpts, _taskId)
}

// TodoTaskAddedIterator is returned from FilterTaskAdded and is used to iterate over the raw logs and unpacked data for TaskAdded events raised by the Todo contract.
type TodoTaskAddedIterator struct {
	Event *TodoTaskAdded // Event containing the contract specifics and raw log

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
func (it *TodoTaskAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodoTaskAdded)
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
		it.Event = new(TodoTaskAdded)
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
func (it *TodoTaskAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodoTaskAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodoTaskAdded represents a TaskAdded event raised by the Todo contract.
type TodoTaskAdded struct {
	TaskId      *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskAdded is a free log retrieval operation binding the contract event 0x2ee7ce8bac091b03a0c391c07ee88a246bdd72bdc6b1ef1ce4d99220595d32c5.
//
// Solidity: event TaskAdded(uint256 indexed taskId, string description)
func (_Todo *TodoFilterer) FilterTaskAdded(opts *bind.FilterOpts, taskId []*big.Int) (*TodoTaskAddedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _Todo.contract.FilterLogs(opts, "TaskAdded", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &TodoTaskAddedIterator{contract: _Todo.contract, event: "TaskAdded", logs: logs, sub: sub}, nil
}

// WatchTaskAdded is a free log subscription operation binding the contract event 0x2ee7ce8bac091b03a0c391c07ee88a246bdd72bdc6b1ef1ce4d99220595d32c5.
//
// Solidity: event TaskAdded(uint256 indexed taskId, string description)
func (_Todo *TodoFilterer) WatchTaskAdded(opts *bind.WatchOpts, sink chan<- *TodoTaskAdded, taskId []*big.Int) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _Todo.contract.WatchLogs(opts, "TaskAdded", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodoTaskAdded)
				if err := _Todo.contract.UnpackLog(event, "TaskAdded", log); err != nil {
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

// ParseTaskAdded is a log parse operation binding the contract event 0x2ee7ce8bac091b03a0c391c07ee88a246bdd72bdc6b1ef1ce4d99220595d32c5.
//
// Solidity: event TaskAdded(uint256 indexed taskId, string description)
func (_Todo *TodoFilterer) ParseTaskAdded(log types.Log) (*TodoTaskAdded, error) {
	event := new(TodoTaskAdded)
	if err := _Todo.contract.UnpackLog(event, "TaskAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TodoTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the Todo contract.
type TodoTaskCompletedIterator struct {
	Event *TodoTaskCompleted // Event containing the contract specifics and raw log

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
func (it *TodoTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodoTaskCompleted)
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
		it.Event = new(TodoTaskCompleted)
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
func (it *TodoTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodoTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodoTaskCompleted represents a TaskCompleted event raised by the Todo contract.
type TodoTaskCompleted struct {
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0xc1fa7142cfb933d0855114a3bffde296944c30b727f297c14d0db4d553d3a5c7.
//
// Solidity: event TaskCompleted(uint256 indexed taskId)
func (_Todo *TodoFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskId []*big.Int) (*TodoTaskCompletedIterator, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _Todo.contract.FilterLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return &TodoTaskCompletedIterator{contract: _Todo.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0xc1fa7142cfb933d0855114a3bffde296944c30b727f297c14d0db4d553d3a5c7.
//
// Solidity: event TaskCompleted(uint256 indexed taskId)
func (_Todo *TodoFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *TodoTaskCompleted, taskId []*big.Int) (event.Subscription, error) {

	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}

	logs, sub, err := _Todo.contract.WatchLogs(opts, "TaskCompleted", taskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodoTaskCompleted)
				if err := _Todo.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0xc1fa7142cfb933d0855114a3bffde296944c30b727f297c14d0db4d553d3a5c7.
//
// Solidity: event TaskCompleted(uint256 indexed taskId)
func (_Todo *TodoFilterer) ParseTaskCompleted(log types.Log) (*TodoTaskCompleted, error) {
	event := new(TodoTaskCompleted)
	if err := _Todo.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
