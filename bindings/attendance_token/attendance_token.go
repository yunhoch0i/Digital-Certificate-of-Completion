// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package attendance_token

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

// AttendanceTokenMetaData contains all meta data concerning the AttendanceToken contract.
var AttendanceTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintAttendance\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"eventId\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AttendanceMinted\",\"inputs\":[{\"name\":\"member\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"eventId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801562000010575f80fd5b506040518060400160405280601581526020017f4153424720417474656e64616e636520546f6b656e00000000000000000000008152506040518060400160405280600481526020017f415342470000000000000000000000000000000000000000000000000000000081525081600390816200008e9190620004bc565b508060049081620000a09190620004bc565b505050620000b75f801b33620000f160201b60201c565b50620000ea7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633620000f160201b60201c565b50620005a0565b5f620001048383620001ed60201b60201c565b620001e357600160055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055506200017f6200025160201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050620001e7565b5f90505b92915050565b5f60055f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f33905090565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620002d457607f821691505b602082108103620002ea57620002e96200028f565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200034e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000311565b6200035a868362000311565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620003a46200039e620003988462000372565b6200037b565b62000372565b9050919050565b5f819050919050565b620003bf8362000384565b620003d7620003ce82620003ab565b8484546200031d565b825550505050565b5f90565b620003ed620003df565b620003fa818484620003b4565b505050565b5b818110156200042157620004155f82620003e3565b60018101905062000400565b5050565b601f82111562000470576200043a81620002f0565b620004458462000302565b8101602085101562000455578190505b6200046d620004648562000302565b830182620003ff565b50505b505050565b5f82821c905092915050565b5f620004925f198460080262000475565b1980831691505092915050565b5f620004ac838362000481565b9150826002028217905092915050565b620004c78262000258565b67ffffffffffffffff811115620004e357620004e262000262565b5b620004ef8254620002bc565b620004fc82828562000425565b5f60209050601f83116001811462000532575f84156200051d578287015190505b6200052985826200049f565b86555062000598565b601f1984166200054286620002f0565b5f5b828110156200056b5784890151825560018201915060208501945060208101905062000544565b868310156200058b578489015162000587601f89168262000481565b8355505b6001600288020188555050505b505050505050565b61190e80620005ae5f395ff3fe608060405234801561000f575f80fd5b5060043610610114575f3560e01c806336568abe116100a0578063a217fddf1161006f578063a217fddf14610304578063a9059cbb14610322578063d539139314610352578063d547741f14610370578063dd62ed3e1461038c57610114565b806336568abe1461026a57806370a082311461028657806391d14854146102b657806395d89b41146102e657610114565b806323b872dd116100e757806323b872dd146101b4578063248a9ca3146101e45780632ed522fa146102145780632f2ff15d14610230578063313ce5671461024c57610114565b806301ffc9a71461011857806306fdde0314610148578063095ea7b31461016657806318160ddd14610196575b5f80fd5b610132600480360381019061012d9190611247565b6103bc565b60405161013f919061128c565b60405180910390f35b610150610435565b60405161015d919061132f565b60405180910390f35b610180600480360381019061017b91906113dc565b6104c5565b60405161018d919061128c565b60405180910390f35b61019e6104e7565b6040516101ab9190611429565b60405180910390f35b6101ce60048036038101906101c99190611442565b6104f0565b6040516101db919061128c565b60405180910390f35b6101fe60048036038101906101f991906114c5565b61051e565b60405161020b91906114ff565b60405180910390f35b61022e60048036038101906102299190611579565b61053b565b005b61024a600480360381019061024591906115ea565b6105c8565b005b6102546105ea565b6040516102619190611643565b60405180910390f35b610284600480360381019061027f91906115ea565b6105ee565b005b6102a0600480360381019061029b919061165c565b610669565b6040516102ad9190611429565b60405180910390f35b6102d060048036038101906102cb91906115ea565b6106ae565b6040516102dd919061128c565b60405180910390f35b6102ee610712565b6040516102fb919061132f565b60405180910390f35b61030c6107a2565b60405161031991906114ff565b60405180910390f35b61033c600480360381019061033791906113dc565b6107a8565b604051610349919061128c565b60405180910390f35b61035a6107ca565b60405161036791906114ff565b60405180910390f35b61038a600480360381019061038591906115ea565b6107ee565b005b6103a660048036038101906103a19190611687565b610810565b6040516103b39190611429565b60405180910390f35b5f7f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061042e575061042d82610892565b5b9050919050565b606060038054610444906116f2565b80601f0160208091040260200160405190810160405280929190818152602001828054610470906116f2565b80156104bb5780601f10610492576101008083540402835291602001916104bb565b820191905f5260205f20905b81548152906001019060200180831161049e57829003601f168201915b5050505050905090565b5f806104cf6108fb565b90506104dc818585610902565b600191505092915050565b5f600254905090565b5f806104fa6108fb565b9050610507858285610914565b6105128585856109a7565b60019150509392505050565b5f60055f8381526020019081526020015f20600101549050919050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a661056581610a97565b61056f8585610aab565b8473ffffffffffffffffffffffffffffffffffffffff167f9cc58cd9de4c207bc33a561f774552ecef14dde35f27dc990b8bc56059f5a5b28585856040516105b99392919061175c565b60405180910390a25050505050565b6105d18261051e565b6105da81610a97565b6105e48383610b2a565b50505050565b5f90565b6105f66108fb565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461065a576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106648282610c14565b505050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f60055f8481526020019081526020015f205f015f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b606060048054610721906116f2565b80601f016020809104026020016040519081016040528092919081815260200182805461074d906116f2565b80156107985780601f1061076f57610100808354040283529160200191610798565b820191905f5260205f20905b81548152906001019060200180831161077b57829003601f168201915b5050505050905090565b5f801b81565b5f806107b26108fb565b90506107bf8185856109a7565b600191505092915050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b6107f78261051e565b61080081610a97565b61080a8383610c14565b50505050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f33905090565b61090f8383836001610cfe565b505050565b5f61091f8484610810565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156109a15781811015610992578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016109899392919061179b565b60405180910390fd5b6109a084848484035f610cfe565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a17575f6040517f96c6fd1e000000000000000000000000000000000000000000000000000000008152600401610a0e91906117d0565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a87575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610a7e91906117d0565b60405180910390fd5b610a92838383610ecd565b505050565b610aa881610aa36108fb565b610f80565b50565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b1b575f6040517fec442f05000000000000000000000000000000000000000000000000000000008152600401610b1291906117d0565b60405180910390fd5b610b265f8383610ecd565b5050565b5f610b3583836106ae565b610c0a57600160055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610ba76108fb565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019050610c0e565b5f90505b92915050565b5f610c1f83836106ae565b15610cf4575f60055f8581526020019081526020015f205f015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550610c916108fb565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16847ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a460019050610cf8565b5f90505b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610d6e575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610d6591906117d0565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610dde575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610dd591906117d0565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610ec7578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610ebe9190611429565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015610f305750610f2e7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6846106ae565b155b15610f70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6790611833565b60405180910390fd5b610f7b838383610fd1565b505050565b610f8a82826106ae565b610fcd5780826040517fe2517d3f000000000000000000000000000000000000000000000000000000008152600401610fc4929190611851565b60405180910390fd5b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611021578060025f82825461101591906118a5565b925050819055506110ef565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050818110156110aa578381836040517fe450d38c0000000000000000000000000000000000000000000000000000000081526004016110a19392919061179b565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611136578060025f8282540392505081905550611180565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111dd9190611429565b60405180910390a3505050565b5f80fd5b5f80fd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611226816111f2565b8114611230575f80fd5b50565b5f813590506112418161121d565b92915050565b5f6020828403121561125c5761125b6111ea565b5b5f61126984828501611233565b91505092915050565b5f8115159050919050565b61128681611272565b82525050565b5f60208201905061129f5f83018461127d565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156112dc5780820151818401526020810190506112c1565b5f8484015250505050565b5f601f19601f8301169050919050565b5f611301826112a5565b61130b81856112af565b935061131b8185602086016112bf565b611324816112e7565b840191505092915050565b5f6020820190508181035f83015261134781846112f7565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6113788261134f565b9050919050565b6113888161136e565b8114611392575f80fd5b50565b5f813590506113a38161137f565b92915050565b5f819050919050565b6113bb816113a9565b81146113c5575f80fd5b50565b5f813590506113d6816113b2565b92915050565b5f80604083850312156113f2576113f16111ea565b5b5f6113ff85828601611395565b9250506020611410858286016113c8565b9150509250929050565b611423816113a9565b82525050565b5f60208201905061143c5f83018461141a565b92915050565b5f805f60608486031215611459576114586111ea565b5b5f61146686828701611395565b935050602061147786828701611395565b9250506040611488868287016113c8565b9150509250925092565b5f819050919050565b6114a481611492565b81146114ae575f80fd5b50565b5f813590506114bf8161149b565b92915050565b5f602082840312156114da576114d96111ea565b5b5f6114e7848285016114b1565b91505092915050565b6114f981611492565b82525050565b5f6020820190506115125f8301846114f0565b92915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261153957611538611518565b5b8235905067ffffffffffffffff8111156115565761155561151c565b5b60208301915083600182028301111561157257611571611520565b5b9250929050565b5f805f8060608587031215611591576115906111ea565b5b5f61159e87828801611395565b94505060206115af878288016113c8565b935050604085013567ffffffffffffffff8111156115d0576115cf6111ee565b5b6115dc87828801611524565b925092505092959194509250565b5f8060408385031215611600576115ff6111ea565b5b5f61160d858286016114b1565b925050602061161e85828601611395565b9150509250929050565b5f60ff82169050919050565b61163d81611628565b82525050565b5f6020820190506116565f830184611634565b92915050565b5f60208284031215611671576116706111ea565b5b5f61167e84828501611395565b91505092915050565b5f806040838503121561169d5761169c6111ea565b5b5f6116aa85828601611395565b92505060206116bb85828601611395565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061170957607f821691505b60208210810361171c5761171b6116c5565b5b50919050565b828183375f83830152505050565b5f61173b83856112af565b9350611748838584611722565b611751836112e7565b840190509392505050565b5f60408201905061176f5f83018661141a565b8181036020830152611782818486611730565b9050949350505050565b6117958161136e565b82525050565b5f6060820190506117ae5f83018661178c565b6117bb602083018561141a565b6117c8604083018461141a565b949350505050565b5f6020820190506117e35f83018461178c565b92915050565b7f41544b3a207472616e736665722064697361626c6564000000000000000000005f82015250565b5f61181d6016836112af565b9150611828826117e9565b602082019050919050565b5f6020820190508181035f83015261184a81611811565b9050919050565b5f6040820190506118645f83018561178c565b61187160208301846114f0565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6118af826113a9565b91506118ba836113a9565b92508282019050808211156118d2576118d1611878565b5b9291505056fea26469706673582212206d573a207ec4b25abdf696dbce1997aeb2a9496d0d423d1034c4ea7ce9a77e2f64736f6c63430008180033",
}

// AttendanceTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use AttendanceTokenMetaData.ABI instead.
var AttendanceTokenABI = AttendanceTokenMetaData.ABI

// AttendanceTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttendanceTokenMetaData.Bin instead.
var AttendanceTokenBin = AttendanceTokenMetaData.Bin

// DeployAttendanceToken deploys a new Ethereum contract, binding an instance of AttendanceToken to it.
func DeployAttendanceToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttendanceToken, error) {
	parsed, err := AttendanceTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttendanceTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttendanceToken{AttendanceTokenCaller: AttendanceTokenCaller{contract: contract}, AttendanceTokenTransactor: AttendanceTokenTransactor{contract: contract}, AttendanceTokenFilterer: AttendanceTokenFilterer{contract: contract}}, nil
}

// AttendanceToken is an auto generated Go binding around an Ethereum contract.
type AttendanceToken struct {
	AttendanceTokenCaller     // Read-only binding to the contract
	AttendanceTokenTransactor // Write-only binding to the contract
	AttendanceTokenFilterer   // Log filterer for contract events
}

// AttendanceTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttendanceTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttendanceTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttendanceTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttendanceTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttendanceTokenSession struct {
	Contract     *AttendanceToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttendanceTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttendanceTokenCallerSession struct {
	Contract *AttendanceTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AttendanceTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttendanceTokenTransactorSession struct {
	Contract     *AttendanceTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AttendanceTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttendanceTokenRaw struct {
	Contract *AttendanceToken // Generic contract binding to access the raw methods on
}

// AttendanceTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttendanceTokenCallerRaw struct {
	Contract *AttendanceTokenCaller // Generic read-only contract binding to access the raw methods on
}

// AttendanceTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttendanceTokenTransactorRaw struct {
	Contract *AttendanceTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttendanceToken creates a new instance of AttendanceToken, bound to a specific deployed contract.
func NewAttendanceToken(address common.Address, backend bind.ContractBackend) (*AttendanceToken, error) {
	contract, err := bindAttendanceToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttendanceToken{AttendanceTokenCaller: AttendanceTokenCaller{contract: contract}, AttendanceTokenTransactor: AttendanceTokenTransactor{contract: contract}, AttendanceTokenFilterer: AttendanceTokenFilterer{contract: contract}}, nil
}

// NewAttendanceTokenCaller creates a new read-only instance of AttendanceToken, bound to a specific deployed contract.
func NewAttendanceTokenCaller(address common.Address, caller bind.ContractCaller) (*AttendanceTokenCaller, error) {
	contract, err := bindAttendanceToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenCaller{contract: contract}, nil
}

// NewAttendanceTokenTransactor creates a new write-only instance of AttendanceToken, bound to a specific deployed contract.
func NewAttendanceTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AttendanceTokenTransactor, error) {
	contract, err := bindAttendanceToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenTransactor{contract: contract}, nil
}

// NewAttendanceTokenFilterer creates a new log filterer instance of AttendanceToken, bound to a specific deployed contract.
func NewAttendanceTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AttendanceTokenFilterer, error) {
	contract, err := bindAttendanceToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenFilterer{contract: contract}, nil
}

// bindAttendanceToken binds a generic wrapper to an already deployed contract.
func bindAttendanceToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttendanceTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttendanceToken *AttendanceTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttendanceToken.Contract.AttendanceTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttendanceToken *AttendanceTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttendanceToken.Contract.AttendanceTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttendanceToken *AttendanceTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttendanceToken.Contract.AttendanceTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttendanceToken *AttendanceTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttendanceToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttendanceToken *AttendanceTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttendanceToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttendanceToken *AttendanceTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttendanceToken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AttendanceToken *AttendanceTokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AttendanceToken *AttendanceTokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AttendanceToken.Contract.DEFAULTADMINROLE(&_AttendanceToken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AttendanceToken *AttendanceTokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AttendanceToken.Contract.DEFAULTADMINROLE(&_AttendanceToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AttendanceToken *AttendanceTokenCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AttendanceToken *AttendanceTokenSession) MINTERROLE() ([32]byte, error) {
	return _AttendanceToken.Contract.MINTERROLE(&_AttendanceToken.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_AttendanceToken *AttendanceTokenCallerSession) MINTERROLE() ([32]byte, error) {
	return _AttendanceToken.Contract.MINTERROLE(&_AttendanceToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AttendanceToken *AttendanceTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AttendanceToken *AttendanceTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AttendanceToken.Contract.Allowance(&_AttendanceToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AttendanceToken *AttendanceTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AttendanceToken.Contract.Allowance(&_AttendanceToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AttendanceToken *AttendanceTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AttendanceToken *AttendanceTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AttendanceToken.Contract.BalanceOf(&_AttendanceToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AttendanceToken *AttendanceTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AttendanceToken.Contract.BalanceOf(&_AttendanceToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_AttendanceToken *AttendanceTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_AttendanceToken *AttendanceTokenSession) Decimals() (uint8, error) {
	return _AttendanceToken.Contract.Decimals(&_AttendanceToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_AttendanceToken *AttendanceTokenCallerSession) Decimals() (uint8, error) {
	return _AttendanceToken.Contract.Decimals(&_AttendanceToken.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AttendanceToken *AttendanceTokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AttendanceToken *AttendanceTokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AttendanceToken.Contract.GetRoleAdmin(&_AttendanceToken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AttendanceToken *AttendanceTokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AttendanceToken.Contract.GetRoleAdmin(&_AttendanceToken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AttendanceToken *AttendanceTokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AttendanceToken *AttendanceTokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AttendanceToken.Contract.HasRole(&_AttendanceToken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AttendanceToken *AttendanceTokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AttendanceToken.Contract.HasRole(&_AttendanceToken.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AttendanceToken *AttendanceTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AttendanceToken *AttendanceTokenSession) Name() (string, error) {
	return _AttendanceToken.Contract.Name(&_AttendanceToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AttendanceToken *AttendanceTokenCallerSession) Name() (string, error) {
	return _AttendanceToken.Contract.Name(&_AttendanceToken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AttendanceToken *AttendanceTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AttendanceToken *AttendanceTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AttendanceToken.Contract.SupportsInterface(&_AttendanceToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AttendanceToken *AttendanceTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AttendanceToken.Contract.SupportsInterface(&_AttendanceToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AttendanceToken *AttendanceTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AttendanceToken *AttendanceTokenSession) Symbol() (string, error) {
	return _AttendanceToken.Contract.Symbol(&_AttendanceToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AttendanceToken *AttendanceTokenCallerSession) Symbol() (string, error) {
	return _AttendanceToken.Contract.Symbol(&_AttendanceToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AttendanceToken *AttendanceTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttendanceToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AttendanceToken *AttendanceTokenSession) TotalSupply() (*big.Int, error) {
	return _AttendanceToken.Contract.TotalSupply(&_AttendanceToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AttendanceToken *AttendanceTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _AttendanceToken.Contract.TotalSupply(&_AttendanceToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.Contract.Approve(&_AttendanceToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.Contract.Approve(&_AttendanceToken.TransactOpts, spender, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AttendanceToken *AttendanceTokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceToken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AttendanceToken *AttendanceTokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceToken.Contract.GrantRole(&_AttendanceToken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AttendanceToken *AttendanceTokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceToken.Contract.GrantRole(&_AttendanceToken.TransactOpts, role, account)
}

// MintAttendance is a paid mutator transaction binding the contract method 0x2ed522fa.
//
// Solidity: function mintAttendance(address member, uint256 amount, string eventId) returns()
func (_AttendanceToken *AttendanceTokenTransactor) MintAttendance(opts *bind.TransactOpts, member common.Address, amount *big.Int, eventId string) (*types.Transaction, error) {
	return _AttendanceToken.contract.Transact(opts, "mintAttendance", member, amount, eventId)
}

// MintAttendance is a paid mutator transaction binding the contract method 0x2ed522fa.
//
// Solidity: function mintAttendance(address member, uint256 amount, string eventId) returns()
func (_AttendanceToken *AttendanceTokenSession) MintAttendance(member common.Address, amount *big.Int, eventId string) (*types.Transaction, error) {
	return _AttendanceToken.Contract.MintAttendance(&_AttendanceToken.TransactOpts, member, amount, eventId)
}

// MintAttendance is a paid mutator transaction binding the contract method 0x2ed522fa.
//
// Solidity: function mintAttendance(address member, uint256 amount, string eventId) returns()
func (_AttendanceToken *AttendanceTokenTransactorSession) MintAttendance(member common.Address, amount *big.Int, eventId string) (*types.Transaction, error) {
	return _AttendanceToken.Contract.MintAttendance(&_AttendanceToken.TransactOpts, member, amount, eventId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AttendanceToken *AttendanceTokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AttendanceToken.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AttendanceToken *AttendanceTokenSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AttendanceToken.Contract.RenounceRole(&_AttendanceToken.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AttendanceToken *AttendanceTokenTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AttendanceToken.Contract.RenounceRole(&_AttendanceToken.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AttendanceToken *AttendanceTokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceToken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AttendanceToken *AttendanceTokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceToken.Contract.RevokeRole(&_AttendanceToken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AttendanceToken *AttendanceTokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AttendanceToken.Contract.RevokeRole(&_AttendanceToken.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.Contract.Transfer(&_AttendanceToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.Contract.Transfer(&_AttendanceToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.Contract.TransferFrom(&_AttendanceToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AttendanceToken *AttendanceTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AttendanceToken.Contract.TransferFrom(&_AttendanceToken.TransactOpts, from, to, value)
}

// AttendanceTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AttendanceToken contract.
type AttendanceTokenApprovalIterator struct {
	Event *AttendanceTokenApproval // Event containing the contract specifics and raw log

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
func (it *AttendanceTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTokenApproval)
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
		it.Event = new(AttendanceTokenApproval)
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
func (it *AttendanceTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTokenApproval represents a Approval event raised by the AttendanceToken contract.
type AttendanceTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AttendanceToken *AttendanceTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AttendanceTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AttendanceToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenApprovalIterator{contract: _AttendanceToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AttendanceToken *AttendanceTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AttendanceTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AttendanceToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTokenApproval)
				if err := _AttendanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AttendanceToken *AttendanceTokenFilterer) ParseApproval(log types.Log) (*AttendanceTokenApproval, error) {
	event := new(AttendanceTokenApproval)
	if err := _AttendanceToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTokenAttendanceMintedIterator is returned from FilterAttendanceMinted and is used to iterate over the raw logs and unpacked data for AttendanceMinted events raised by the AttendanceToken contract.
type AttendanceTokenAttendanceMintedIterator struct {
	Event *AttendanceTokenAttendanceMinted // Event containing the contract specifics and raw log

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
func (it *AttendanceTokenAttendanceMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTokenAttendanceMinted)
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
		it.Event = new(AttendanceTokenAttendanceMinted)
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
func (it *AttendanceTokenAttendanceMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTokenAttendanceMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTokenAttendanceMinted represents a AttendanceMinted event raised by the AttendanceToken contract.
type AttendanceTokenAttendanceMinted struct {
	Member  common.Address
	Amount  *big.Int
	EventId string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAttendanceMinted is a free log retrieval operation binding the contract event 0x9cc58cd9de4c207bc33a561f774552ecef14dde35f27dc990b8bc56059f5a5b2.
//
// Solidity: event AttendanceMinted(address indexed member, uint256 amount, string eventId)
func (_AttendanceToken *AttendanceTokenFilterer) FilterAttendanceMinted(opts *bind.FilterOpts, member []common.Address) (*AttendanceTokenAttendanceMintedIterator, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _AttendanceToken.contract.FilterLogs(opts, "AttendanceMinted", memberRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenAttendanceMintedIterator{contract: _AttendanceToken.contract, event: "AttendanceMinted", logs: logs, sub: sub}, nil
}

// WatchAttendanceMinted is a free log subscription operation binding the contract event 0x9cc58cd9de4c207bc33a561f774552ecef14dde35f27dc990b8bc56059f5a5b2.
//
// Solidity: event AttendanceMinted(address indexed member, uint256 amount, string eventId)
func (_AttendanceToken *AttendanceTokenFilterer) WatchAttendanceMinted(opts *bind.WatchOpts, sink chan<- *AttendanceTokenAttendanceMinted, member []common.Address) (event.Subscription, error) {

	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _AttendanceToken.contract.WatchLogs(opts, "AttendanceMinted", memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTokenAttendanceMinted)
				if err := _AttendanceToken.contract.UnpackLog(event, "AttendanceMinted", log); err != nil {
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

// ParseAttendanceMinted is a log parse operation binding the contract event 0x9cc58cd9de4c207bc33a561f774552ecef14dde35f27dc990b8bc56059f5a5b2.
//
// Solidity: event AttendanceMinted(address indexed member, uint256 amount, string eventId)
func (_AttendanceToken *AttendanceTokenFilterer) ParseAttendanceMinted(log types.Log) (*AttendanceTokenAttendanceMinted, error) {
	event := new(AttendanceTokenAttendanceMinted)
	if err := _AttendanceToken.contract.UnpackLog(event, "AttendanceMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AttendanceToken contract.
type AttendanceTokenRoleAdminChangedIterator struct {
	Event *AttendanceTokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AttendanceTokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTokenRoleAdminChanged)
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
		it.Event = new(AttendanceTokenRoleAdminChanged)
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
func (it *AttendanceTokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTokenRoleAdminChanged represents a RoleAdminChanged event raised by the AttendanceToken contract.
type AttendanceTokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AttendanceToken *AttendanceTokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AttendanceTokenRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AttendanceToken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenRoleAdminChangedIterator{contract: _AttendanceToken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AttendanceToken *AttendanceTokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AttendanceTokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AttendanceToken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTokenRoleAdminChanged)
				if err := _AttendanceToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AttendanceToken *AttendanceTokenFilterer) ParseRoleAdminChanged(log types.Log) (*AttendanceTokenRoleAdminChanged, error) {
	event := new(AttendanceTokenRoleAdminChanged)
	if err := _AttendanceToken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AttendanceToken contract.
type AttendanceTokenRoleGrantedIterator struct {
	Event *AttendanceTokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *AttendanceTokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTokenRoleGranted)
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
		it.Event = new(AttendanceTokenRoleGranted)
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
func (it *AttendanceTokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTokenRoleGranted represents a RoleGranted event raised by the AttendanceToken contract.
type AttendanceTokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceToken *AttendanceTokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AttendanceTokenRoleGrantedIterator, error) {

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

	logs, sub, err := _AttendanceToken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenRoleGrantedIterator{contract: _AttendanceToken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceToken *AttendanceTokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AttendanceTokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AttendanceToken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTokenRoleGranted)
				if err := _AttendanceToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AttendanceToken *AttendanceTokenFilterer) ParseRoleGranted(log types.Log) (*AttendanceTokenRoleGranted, error) {
	event := new(AttendanceTokenRoleGranted)
	if err := _AttendanceToken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AttendanceToken contract.
type AttendanceTokenRoleRevokedIterator struct {
	Event *AttendanceTokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AttendanceTokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTokenRoleRevoked)
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
		it.Event = new(AttendanceTokenRoleRevoked)
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
func (it *AttendanceTokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTokenRoleRevoked represents a RoleRevoked event raised by the AttendanceToken contract.
type AttendanceTokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceToken *AttendanceTokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AttendanceTokenRoleRevokedIterator, error) {

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

	logs, sub, err := _AttendanceToken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenRoleRevokedIterator{contract: _AttendanceToken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AttendanceToken *AttendanceTokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AttendanceTokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AttendanceToken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTokenRoleRevoked)
				if err := _AttendanceToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AttendanceToken *AttendanceTokenFilterer) ParseRoleRevoked(log types.Log) (*AttendanceTokenRoleRevoked, error) {
	event := new(AttendanceTokenRoleRevoked)
	if err := _AttendanceToken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttendanceTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AttendanceToken contract.
type AttendanceTokenTransferIterator struct {
	Event *AttendanceTokenTransfer // Event containing the contract specifics and raw log

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
func (it *AttendanceTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttendanceTokenTransfer)
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
		it.Event = new(AttendanceTokenTransfer)
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
func (it *AttendanceTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttendanceTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttendanceTokenTransfer represents a Transfer event raised by the AttendanceToken contract.
type AttendanceTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AttendanceToken *AttendanceTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AttendanceTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AttendanceToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AttendanceTokenTransferIterator{contract: _AttendanceToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AttendanceToken *AttendanceTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AttendanceTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AttendanceToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttendanceTokenTransfer)
				if err := _AttendanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AttendanceToken *AttendanceTokenFilterer) ParseTransfer(log types.Log) (*AttendanceTokenTransfer, error) {
	event := new(AttendanceTokenTransfer)
	if err := _AttendanceToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
