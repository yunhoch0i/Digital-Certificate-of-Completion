// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package certificate_nft

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

// CertificateNFTMetaData contains all meta data concerning the CertificateNFT contract.
var CertificateNFTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_tracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ISSUER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attendanceTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAttendanceTracker\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completionThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasCertificate\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"issueCertificate\",\"inputs\":[{\"name\":\"memberId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ipfsCid\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"memberName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"memberTokenId\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CertificateIssued\",\"inputs\":[{\"name\":\"memberId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"memberName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162001c3238038062001c328339810160408190526200003391620001b5565b6040518060400160405280601081526020016f4153424720436572746966696361746560801b815250604051806040016040528060088152602001671054d091d0d1549560c21b815250815f90816200008d91906200028c565b5060016200009c82826200028c565b50620000ad91505f90503362000105565b50620000da7f114e74f6ea3bd819998f78687bfcb11b140da08e9b7d222fa9c1f1ba1f2aa1223362000105565b50600880546001600160a01b0319166001600160a01b03939093169290921790915560095562000358565b5f8281526007602090815260408083206001600160a01b038516845290915281205460ff16620001ac575f8381526007602090815260408083206001600160a01b03861684529091529020805460ff19166001179055620001633390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001620001af565b505f5b92915050565b5f8060408385031215620001c7575f80fd5b82516001600160a01b0381168114620001de575f80fd5b6020939093015192949293505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200021757607f821691505b6020821081036200023657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200028757805f5260205f20601f840160051c81016020851015620002635750805b601f840160051c820191505b8181101562000284575f81556001016200026f565b50505b505050565b81516001600160401b03811115620002a857620002a8620001ee565b620002c081620002b9845462000202565b846200023c565b602080601f831160018114620002f6575f8415620002de5750858301515b5f19600386901b1c1916600185901b17855562000350565b5f85815260208120601f198616915b82811015620003265788860151825594840194600190910190840162000305565b50858210156200034457878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b6118cc80620003665f395ff3fe608060405234801561000f575f80fd5b506004361061016d575f3560e01c806382aefa24116100d9578063b88d4fde11610093578063d547741f1161006e578063d547741f14610366578063dbb03d1614610379578063e985e9c51461038c578063f24bed041461039f575f80fd5b8063b88d4fde1461032d578063c312fda714610340578063c87b56dd14610353575f80fd5b806382aefa24146102af5780638b5447c6146102d657806391d14854146102f857806395d89b411461030b578063a217fddf14610313578063a22cb4651461031a575f80fd5b80632f2ff15d1161012a5780632f2ff15d1461023157806336568abe1461024457806342842e0e1461025757806353068c7e1461026a5780636352211e1461028957806370a082311461029c575f80fd5b806301ffc9a71461017157806306fdde0314610199578063081812fc146101ae578063095ea7b3146101d957806323b872dd146101ee578063248a9ca314610201575b5f80fd5b61018461017f3660046112f4565b6103a8565b60405190151581526020015b60405180910390f35b6101a16103b8565b604051610190919061135c565b6101c16101bc36600461136e565b610447565b6040516001600160a01b039091168152602001610190565b6101ec6101e73660046113a0565b61046e565b005b6101ec6101fc3660046113c8565b61047d565b61022361020f36600461136e565b5f9081526007602052604090206001015490565b604051908152602001610190565b6101ec61023f366004611401565b61050b565b6101ec610252366004611401565b61052f565b6101ec6102653660046113c8565b610567565b61022361027836600461136e565b600c6020525f908152604090205481565b6101c161029736600461136e565b610581565b6102236102aa36600461142b565b61058b565b6102237f114e74f6ea3bd819998f78687bfcb11b140da08e9b7d222fa9c1f1ba1f2aa12281565b6101846102e436600461136e565b600b6020525f908152604090205460ff1681565b610184610306366004611401565b6105d0565b6101a16105fa565b6102235f81565b6101ec610328366004611444565b610609565b6101ec61033b366004611491565b610614565b6008546101c1906001600160a01b031681565b6101a161036136600461136e565b61062c565b6101ec610374366004611401565b6106ad565b6101ec6103873660046115ab565b6106d1565b61018461039a36600461161f565b6108c3565b61022360095481565b5f6103b2826108f0565b92915050565b60605f80546103c690611647565b80601f01602080910402602001604051908101604052809291908181526020018280546103f290611647565b801561043d5780601f106104145761010080835404028352916020019161043d565b820191905f5260205f20905b81548152906001019060200180831161042057829003601f168201915b5050505050905090565b5f61045182610914565b505f828152600460205260409020546001600160a01b03166103b2565b61047982823361094c565b5050565b6001600160a01b0382166104ab57604051633250574960e11b81525f60048201526024015b60405180910390fd5b5f6104b7838333610959565b9050836001600160a01b0316816001600160a01b031614610505576040516364283d7b60e01b81526001600160a01b03808616600483015260248201849052821660448201526064016104a2565b50505050565b5f8281526007602052604090206001015461052581610a4b565b6105058383610a58565b6001600160a01b03811633146105585760405163334bd91960e11b815260040160405180910390fd5b6105628282610ae9565b505050565b61056283838360405180602001604052805f815250610614565b5f6103b282610914565b5f6001600160a01b0382166105b5576040516322718ad960e21b81525f60048201526024016104a2565b506001600160a01b03165f9081526003602052604090205490565b5f9182526007602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6060600180546103c690611647565b610479338383610b54565b61061f84848461047d565b6105053385858585610c1b565b606061063782610914565b505f61064d60408051602081019091525f815290565b90505f61065984610d43565b905081515f0361066a579392505050565b80511561069c57818160405160200161068492919061167f565b60405160208183030381529060405292505050919050565b6106a584610de2565b949350505050565b5f828152600760205260409020600101546106c781610a4b565b6105058383610ae9565b7f114e74f6ea3bd819998f78687bfcb11b140da08e9b7d222fa9c1f1ba1f2aa1226106fb81610a4b565b60095460085460405163eb27153360e01b8152600481018990526001600160a01b039091169063eb27153390602401602060405180830381865afa158015610745573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061076991906116ad565b10156107b75760405162461bcd60e51b815260206004820181905260248201527f434552543a2062656c6f7720636f6d706c6574696f6e207468726573686f6c6460448201526064016104a2565b5f868152600b602052604090205460ff161561080c5760405162461bcd60e51b815260206004820152601460248201527310d154950e88185b1c9958591e481a5cdcdd595960621b60448201526064016104a2565b5f600a5f815461081b906116c4565b9182905550905061082c3382610e53565b6108578187876040516020016108439291906116e8565b604051602081830303815290604052610eb4565b5f878152600b60209081526040808320805460ff19166001179055600c909152908190208290555187907f12b8ba46d8efd84182d36de4f91a8f0922f453ef50e9b302b291191863a14ac3906108b29084908890889061170a565b60405180910390a250505050505050565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b5f6001600160e01b03198216637965db0b60e01b14806103b257506103b282610f03565b5f818152600260205260408120546001600160a01b0316806103b257604051637e27328960e01b8152600481018490526024016104a2565b6105628383836001610f27565b5f828152600260205260408120546001600160a01b03908116908316156109855761098581848661102b565b6001600160a01b038116156109bf576109a05f855f80610f27565b6001600160a01b0381165f90815260036020526040902080545f190190555b6001600160a01b038516156109ed576001600160a01b0385165f908152600360205260409020805460010190555b5f8481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b610a55813361108f565b50565b5f610a6383836105d0565b610ae2575f8381526007602090815260408083206001600160a01b03861684529091529020805460ff19166001179055610a9a3390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45060016103b2565b505f6103b2565b5f610af483836105d0565b15610ae2575f8381526007602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45060016103b2565b6001600160a01b038316610b7d5760405163a9fbf51f60e01b81525f60048201526024016104a2565b6001600160a01b038216610baf57604051630b61174360e31b81526001600160a01b03831660048201526024016104a2565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b15610d3c57604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290610c5d90889088908790879060040161173f565b6020604051808303815f875af1925050508015610c97575060408051601f3d908101601f19168201909252610c949181019061177b565b60015b610cfe573d808015610cc4576040519150601f19603f3d011682016040523d82523d5f602084013e610cc9565b606091505b5080515f03610cf657604051633250574960e11b81526001600160a01b03851660048201526024016104a2565b805160208201fd5b6001600160e01b03198116630a85bd0160e11b14610d3a57604051633250574960e11b81526001600160a01b03851660048201526024016104a2565b505b5050505050565b5f818152600660205260409020805460609190610d5f90611647565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8b90611647565b8015610dd65780601f10610dad57610100808354040283529160200191610dd6565b820191905f5260205f20905b815481529060010190602001808311610db957829003601f168201915b50505050509050919050565b6060610ded82610914565b505f610e0360408051602081019091525f815290565b90505f815111610e215760405180602001604052805f815250610e4c565b80610e2b846110c8565b604051602001610e3c92919061167f565b6040516020818303038152906040525b9392505050565b6001600160a01b038216610e7c57604051633250574960e11b81525f60048201526024016104a2565b5f610e8883835f610959565b90506001600160a01b03811615610562576040516339e3563760e11b81525f60048201526024016104a2565b5f828152600660205260409020610ecb82826117da565b506040518281527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a15050565b5f6001600160e01b03198216632483248360e11b14806103b257506103b282611158565b8080610f3b57506001600160a01b03821615155b15610ffc575f610f4a84610914565b90506001600160a01b03831615801590610f765750826001600160a01b0316816001600160a01b031614155b8015610f895750610f8781846108c3565b155b15610fb25760405163a9fbf51f60e01b81526001600160a01b03841660048201526024016104a2565b8115610ffa5783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b50505f90815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b6110368383836111a7565b610562576001600160a01b03831661106457604051637e27328960e01b8152600481018290526024016104a2565b60405163177e802f60e01b81526001600160a01b0383166004820152602481018290526044016104a2565b61109982826105d0565b6104795760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016104a2565b60605f6110d483611208565b60010190505f8167ffffffffffffffff8111156110f3576110f361147d565b6040519080825280601f01601f19166020018201604052801561111d576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461112757509392505050565b5f6001600160e01b031982166380ac58cd60e01b148061118857506001600160e01b03198216635b5e139f60e01b145b806103b257506301ffc9a760e01b6001600160e01b03198316146103b2565b5f6001600160a01b038316158015906106a55750826001600160a01b0316846001600160a01b031614806111e057506111e084846108c3565b806106a55750505f908152600460205260409020546001600160a01b03908116911614919050565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106112465772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611272576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061129057662386f26fc10000830492506010015b6305f5e10083106112a8576305f5e100830492506008015b61271083106112bc57612710830492506004015b606483106112ce576064830492506002015b600a83106103b25760010192915050565b6001600160e01b031981168114610a55575f80fd5b5f60208284031215611304575f80fd5b8135610e4c816112df565b5f5b83811015611329578181015183820152602001611311565b50505f910152565b5f815180845261134881602086016020860161130f565b601f01601f19169290920160200192915050565b602081525f610e4c6020830184611331565b5f6020828403121561137e575f80fd5b5035919050565b80356001600160a01b038116811461139b575f80fd5b919050565b5f80604083850312156113b1575f80fd5b6113ba83611385565b946020939093013593505050565b5f805f606084860312156113da575f80fd5b6113e384611385565b92506113f160208501611385565b9150604084013590509250925092565b5f8060408385031215611412575f80fd5b8235915061142260208401611385565b90509250929050565b5f6020828403121561143b575f80fd5b610e4c82611385565b5f8060408385031215611455575f80fd5b61145e83611385565b915060208301358015158114611472575f80fd5b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b5f805f80608085870312156114a4575f80fd5b6114ad85611385565b93506114bb60208601611385565b925060408501359150606085013567ffffffffffffffff808211156114de575f80fd5b818701915087601f8301126114f1575f80fd5b8135818111156115035761150361147d565b604051601f8201601f19908116603f0116810190838211818310171561152b5761152b61147d565b816040528281528a6020848701011115611543575f80fd5b826020860160208301375f60208483010152809550505050505092959194509250565b5f8083601f840112611576575f80fd5b50813567ffffffffffffffff81111561158d575f80fd5b6020830191508360208285010111156115a4575f80fd5b9250929050565b5f805f805f606086880312156115bf575f80fd5b85359450602086013567ffffffffffffffff808211156115dd575f80fd5b6115e989838a01611566565b90965094506040880135915080821115611601575f80fd5b5061160e88828901611566565b969995985093965092949392505050565b5f8060408385031215611630575f80fd5b61163983611385565b915061142260208401611385565b600181811c9082168061165b57607f821691505b60208210810361167957634e487b7160e01b5f52602260045260245ffd5b50919050565b5f835161169081846020880161130f565b8351908301906116a481836020880161130f565b01949350505050565b5f602082840312156116bd575f80fd5b5051919050565b5f600182016116e157634e487b7160e01b5f52601160045260245ffd5b5060010190565b66697066733a2f2f60c81b8152818360078301375f9101600701908152919050565b83815260406020820152816040820152818360608301375f818301606090810191909152601f909201601f1916010192915050565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f9061177190830184611331565b9695505050505050565b5f6020828403121561178b575f80fd5b8151610e4c816112df565b601f82111561056257805f5260205f20601f840160051c810160208510156117bb5750805b601f840160051c820191505b81811015610d3c575f81556001016117c7565b815167ffffffffffffffff8111156117f4576117f461147d565b611808816118028454611647565b84611796565b602080601f83116001811461183b575f84156118245750858301515b5f19600386901b1c1916600185901b178555610d3a565b5f85815260208120601f198616915b828110156118695788860151825594840194600190910190840161184a565b508582101561188657878501515f19600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212200640c88a513311d88dc9322e2f8cc21f84d96d90bd7012c19171dba35791050f64736f6c63430008180033",
}

// CertificateNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use CertificateNFTMetaData.ABI instead.
var CertificateNFTABI = CertificateNFTMetaData.ABI

// CertificateNFTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CertificateNFTMetaData.Bin instead.
var CertificateNFTBin = CertificateNFTMetaData.Bin

// DeployCertificateNFT deploys a new Ethereum contract, binding an instance of CertificateNFT to it.
func DeployCertificateNFT(auth *bind.TransactOpts, backend bind.ContractBackend, _tracker common.Address, _threshold *big.Int) (common.Address, *types.Transaction, *CertificateNFT, error) {
	parsed, err := CertificateNFTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CertificateNFTBin), backend, _tracker, _threshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CertificateNFT{CertificateNFTCaller: CertificateNFTCaller{contract: contract}, CertificateNFTTransactor: CertificateNFTTransactor{contract: contract}, CertificateNFTFilterer: CertificateNFTFilterer{contract: contract}}, nil
}

// CertificateNFT is an auto generated Go binding around an Ethereum contract.
type CertificateNFT struct {
	CertificateNFTCaller     // Read-only binding to the contract
	CertificateNFTTransactor // Write-only binding to the contract
	CertificateNFTFilterer   // Log filterer for contract events
}

// CertificateNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertificateNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertificateNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertificateNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertificateNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertificateNFTSession struct {
	Contract     *CertificateNFT   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertificateNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertificateNFTCallerSession struct {
	Contract *CertificateNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CertificateNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertificateNFTTransactorSession struct {
	Contract     *CertificateNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CertificateNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertificateNFTRaw struct {
	Contract *CertificateNFT // Generic contract binding to access the raw methods on
}

// CertificateNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertificateNFTCallerRaw struct {
	Contract *CertificateNFTCaller // Generic read-only contract binding to access the raw methods on
}

// CertificateNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertificateNFTTransactorRaw struct {
	Contract *CertificateNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCertificateNFT creates a new instance of CertificateNFT, bound to a specific deployed contract.
func NewCertificateNFT(address common.Address, backend bind.ContractBackend) (*CertificateNFT, error) {
	contract, err := bindCertificateNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CertificateNFT{CertificateNFTCaller: CertificateNFTCaller{contract: contract}, CertificateNFTTransactor: CertificateNFTTransactor{contract: contract}, CertificateNFTFilterer: CertificateNFTFilterer{contract: contract}}, nil
}

// NewCertificateNFTCaller creates a new read-only instance of CertificateNFT, bound to a specific deployed contract.
func NewCertificateNFTCaller(address common.Address, caller bind.ContractCaller) (*CertificateNFTCaller, error) {
	contract, err := bindCertificateNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTCaller{contract: contract}, nil
}

// NewCertificateNFTTransactor creates a new write-only instance of CertificateNFT, bound to a specific deployed contract.
func NewCertificateNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*CertificateNFTTransactor, error) {
	contract, err := bindCertificateNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTTransactor{contract: contract}, nil
}

// NewCertificateNFTFilterer creates a new log filterer instance of CertificateNFT, bound to a specific deployed contract.
func NewCertificateNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*CertificateNFTFilterer, error) {
	contract, err := bindCertificateNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTFilterer{contract: contract}, nil
}

// bindCertificateNFT binds a generic wrapper to an already deployed contract.
func bindCertificateNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CertificateNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateNFT *CertificateNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateNFT.Contract.CertificateNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateNFT *CertificateNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateNFT.Contract.CertificateNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateNFT *CertificateNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateNFT.Contract.CertificateNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CertificateNFT *CertificateNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CertificateNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CertificateNFT *CertificateNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CertificateNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CertificateNFT *CertificateNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CertificateNFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CertificateNFT *CertificateNFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CertificateNFT *CertificateNFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CertificateNFT.Contract.DEFAULTADMINROLE(&_CertificateNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CertificateNFT *CertificateNFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CertificateNFT.Contract.DEFAULTADMINROLE(&_CertificateNFT.CallOpts)
}

// ISSUERROLE is a free data retrieval call binding the contract method 0x82aefa24.
//
// Solidity: function ISSUER_ROLE() view returns(bytes32)
func (_CertificateNFT *CertificateNFTCaller) ISSUERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "ISSUER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ISSUERROLE is a free data retrieval call binding the contract method 0x82aefa24.
//
// Solidity: function ISSUER_ROLE() view returns(bytes32)
func (_CertificateNFT *CertificateNFTSession) ISSUERROLE() ([32]byte, error) {
	return _CertificateNFT.Contract.ISSUERROLE(&_CertificateNFT.CallOpts)
}

// ISSUERROLE is a free data retrieval call binding the contract method 0x82aefa24.
//
// Solidity: function ISSUER_ROLE() view returns(bytes32)
func (_CertificateNFT *CertificateNFTCallerSession) ISSUERROLE() ([32]byte, error) {
	return _CertificateNFT.Contract.ISSUERROLE(&_CertificateNFT.CallOpts)
}

// AttendanceTracker is a free data retrieval call binding the contract method 0xc312fda7.
//
// Solidity: function attendanceTracker() view returns(address)
func (_CertificateNFT *CertificateNFTCaller) AttendanceTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "attendanceTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttendanceTracker is a free data retrieval call binding the contract method 0xc312fda7.
//
// Solidity: function attendanceTracker() view returns(address)
func (_CertificateNFT *CertificateNFTSession) AttendanceTracker() (common.Address, error) {
	return _CertificateNFT.Contract.AttendanceTracker(&_CertificateNFT.CallOpts)
}

// AttendanceTracker is a free data retrieval call binding the contract method 0xc312fda7.
//
// Solidity: function attendanceTracker() view returns(address)
func (_CertificateNFT *CertificateNFTCallerSession) AttendanceTracker() (common.Address, error) {
	return _CertificateNFT.Contract.AttendanceTracker(&_CertificateNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CertificateNFT *CertificateNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CertificateNFT *CertificateNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CertificateNFT.Contract.BalanceOf(&_CertificateNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CertificateNFT *CertificateNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CertificateNFT.Contract.BalanceOf(&_CertificateNFT.CallOpts, owner)
}

// CompletionThreshold is a free data retrieval call binding the contract method 0xf24bed04.
//
// Solidity: function completionThreshold() view returns(uint256)
func (_CertificateNFT *CertificateNFTCaller) CompletionThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "completionThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompletionThreshold is a free data retrieval call binding the contract method 0xf24bed04.
//
// Solidity: function completionThreshold() view returns(uint256)
func (_CertificateNFT *CertificateNFTSession) CompletionThreshold() (*big.Int, error) {
	return _CertificateNFT.Contract.CompletionThreshold(&_CertificateNFT.CallOpts)
}

// CompletionThreshold is a free data retrieval call binding the contract method 0xf24bed04.
//
// Solidity: function completionThreshold() view returns(uint256)
func (_CertificateNFT *CertificateNFTCallerSession) CompletionThreshold() (*big.Int, error) {
	return _CertificateNFT.Contract.CompletionThreshold(&_CertificateNFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CertificateNFT *CertificateNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CertificateNFT *CertificateNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CertificateNFT.Contract.GetApproved(&_CertificateNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_CertificateNFT *CertificateNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _CertificateNFT.Contract.GetApproved(&_CertificateNFT.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CertificateNFT *CertificateNFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CertificateNFT *CertificateNFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CertificateNFT.Contract.GetRoleAdmin(&_CertificateNFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CertificateNFT *CertificateNFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CertificateNFT.Contract.GetRoleAdmin(&_CertificateNFT.CallOpts, role)
}

// HasCertificate is a free data retrieval call binding the contract method 0x8b5447c6.
//
// Solidity: function hasCertificate(uint256 ) view returns(bool)
func (_CertificateNFT *CertificateNFTCaller) HasCertificate(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "hasCertificate", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCertificate is a free data retrieval call binding the contract method 0x8b5447c6.
//
// Solidity: function hasCertificate(uint256 ) view returns(bool)
func (_CertificateNFT *CertificateNFTSession) HasCertificate(arg0 *big.Int) (bool, error) {
	return _CertificateNFT.Contract.HasCertificate(&_CertificateNFT.CallOpts, arg0)
}

// HasCertificate is a free data retrieval call binding the contract method 0x8b5447c6.
//
// Solidity: function hasCertificate(uint256 ) view returns(bool)
func (_CertificateNFT *CertificateNFTCallerSession) HasCertificate(arg0 *big.Int) (bool, error) {
	return _CertificateNFT.Contract.HasCertificate(&_CertificateNFT.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CertificateNFT *CertificateNFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CertificateNFT *CertificateNFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CertificateNFT.Contract.HasRole(&_CertificateNFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CertificateNFT *CertificateNFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CertificateNFT.Contract.HasRole(&_CertificateNFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CertificateNFT *CertificateNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CertificateNFT *CertificateNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CertificateNFT.Contract.IsApprovedForAll(&_CertificateNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_CertificateNFT *CertificateNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _CertificateNFT.Contract.IsApprovedForAll(&_CertificateNFT.CallOpts, owner, operator)
}

// MemberTokenId is a free data retrieval call binding the contract method 0x53068c7e.
//
// Solidity: function memberTokenId(uint256 ) view returns(uint256)
func (_CertificateNFT *CertificateNFTCaller) MemberTokenId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "memberTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MemberTokenId is a free data retrieval call binding the contract method 0x53068c7e.
//
// Solidity: function memberTokenId(uint256 ) view returns(uint256)
func (_CertificateNFT *CertificateNFTSession) MemberTokenId(arg0 *big.Int) (*big.Int, error) {
	return _CertificateNFT.Contract.MemberTokenId(&_CertificateNFT.CallOpts, arg0)
}

// MemberTokenId is a free data retrieval call binding the contract method 0x53068c7e.
//
// Solidity: function memberTokenId(uint256 ) view returns(uint256)
func (_CertificateNFT *CertificateNFTCallerSession) MemberTokenId(arg0 *big.Int) (*big.Int, error) {
	return _CertificateNFT.Contract.MemberTokenId(&_CertificateNFT.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CertificateNFT *CertificateNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CertificateNFT *CertificateNFTSession) Name() (string, error) {
	return _CertificateNFT.Contract.Name(&_CertificateNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CertificateNFT *CertificateNFTCallerSession) Name() (string, error) {
	return _CertificateNFT.Contract.Name(&_CertificateNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CertificateNFT *CertificateNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CertificateNFT *CertificateNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CertificateNFT.Contract.OwnerOf(&_CertificateNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_CertificateNFT *CertificateNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _CertificateNFT.Contract.OwnerOf(&_CertificateNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CertificateNFT *CertificateNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CertificateNFT *CertificateNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CertificateNFT.Contract.SupportsInterface(&_CertificateNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CertificateNFT *CertificateNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CertificateNFT.Contract.SupportsInterface(&_CertificateNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CertificateNFT *CertificateNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CertificateNFT *CertificateNFTSession) Symbol() (string, error) {
	return _CertificateNFT.Contract.Symbol(&_CertificateNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CertificateNFT *CertificateNFTCallerSession) Symbol() (string, error) {
	return _CertificateNFT.Contract.Symbol(&_CertificateNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CertificateNFT *CertificateNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _CertificateNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CertificateNFT *CertificateNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CertificateNFT.Contract.TokenURI(&_CertificateNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_CertificateNFT *CertificateNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _CertificateNFT.Contract.TokenURI(&_CertificateNFT.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.Contract.Approve(&_CertificateNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.Contract.Approve(&_CertificateNFT.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CertificateNFT *CertificateNFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CertificateNFT *CertificateNFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CertificateNFT.Contract.GrantRole(&_CertificateNFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CertificateNFT.Contract.GrantRole(&_CertificateNFT.TransactOpts, role, account)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xdbb03d16.
//
// Solidity: function issueCertificate(uint256 memberId, string ipfsCid, string memberName) returns()
func (_CertificateNFT *CertificateNFTTransactor) IssueCertificate(opts *bind.TransactOpts, memberId *big.Int, ipfsCid string, memberName string) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "issueCertificate", memberId, ipfsCid, memberName)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xdbb03d16.
//
// Solidity: function issueCertificate(uint256 memberId, string ipfsCid, string memberName) returns()
func (_CertificateNFT *CertificateNFTSession) IssueCertificate(memberId *big.Int, ipfsCid string, memberName string) (*types.Transaction, error) {
	return _CertificateNFT.Contract.IssueCertificate(&_CertificateNFT.TransactOpts, memberId, ipfsCid, memberName)
}

// IssueCertificate is a paid mutator transaction binding the contract method 0xdbb03d16.
//
// Solidity: function issueCertificate(uint256 memberId, string ipfsCid, string memberName) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) IssueCertificate(memberId *big.Int, ipfsCid string, memberName string) (*types.Transaction, error) {
	return _CertificateNFT.Contract.IssueCertificate(&_CertificateNFT.TransactOpts, memberId, ipfsCid, memberName)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CertificateNFT *CertificateNFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CertificateNFT *CertificateNFTSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CertificateNFT.Contract.RenounceRole(&_CertificateNFT.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CertificateNFT.Contract.RenounceRole(&_CertificateNFT.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CertificateNFT *CertificateNFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CertificateNFT *CertificateNFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CertificateNFT.Contract.RevokeRole(&_CertificateNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CertificateNFT.Contract.RevokeRole(&_CertificateNFT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.Contract.SafeTransferFrom(&_CertificateNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.Contract.SafeTransferFrom(&_CertificateNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CertificateNFT *CertificateNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CertificateNFT *CertificateNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CertificateNFT.Contract.SafeTransferFrom0(&_CertificateNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _CertificateNFT.Contract.SafeTransferFrom0(&_CertificateNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CertificateNFT *CertificateNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CertificateNFT *CertificateNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CertificateNFT.Contract.SetApprovalForAll(&_CertificateNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _CertificateNFT.Contract.SetApprovalForAll(&_CertificateNFT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.Contract.TransferFrom(&_CertificateNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_CertificateNFT *CertificateNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CertificateNFT.Contract.TransferFrom(&_CertificateNFT.TransactOpts, from, to, tokenId)
}

// CertificateNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CertificateNFT contract.
type CertificateNFTApprovalIterator struct {
	Event *CertificateNFTApproval // Event containing the contract specifics and raw log

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
func (it *CertificateNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTApproval)
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
		it.Event = new(CertificateNFTApproval)
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
func (it *CertificateNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTApproval represents a Approval event raised by the CertificateNFT contract.
type CertificateNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CertificateNFT *CertificateNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CertificateNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTApprovalIterator{contract: _CertificateNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CertificateNFT *CertificateNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CertificateNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTApproval)
				if err := _CertificateNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_CertificateNFT *CertificateNFTFilterer) ParseApproval(log types.Log) (*CertificateNFTApproval, error) {
	event := new(CertificateNFTApproval)
	if err := _CertificateNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the CertificateNFT contract.
type CertificateNFTApprovalForAllIterator struct {
	Event *CertificateNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CertificateNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTApprovalForAll)
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
		it.Event = new(CertificateNFTApprovalForAll)
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
func (it *CertificateNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTApprovalForAll represents a ApprovalForAll event raised by the CertificateNFT contract.
type CertificateNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CertificateNFT *CertificateNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CertificateNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTApprovalForAllIterator{contract: _CertificateNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CertificateNFT *CertificateNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CertificateNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTApprovalForAll)
				if err := _CertificateNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_CertificateNFT *CertificateNFTFilterer) ParseApprovalForAll(log types.Log) (*CertificateNFTApprovalForAll, error) {
	event := new(CertificateNFTApprovalForAll)
	if err := _CertificateNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the CertificateNFT contract.
type CertificateNFTBatchMetadataUpdateIterator struct {
	Event *CertificateNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *CertificateNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTBatchMetadataUpdate)
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
		it.Event = new(CertificateNFTBatchMetadataUpdate)
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
func (it *CertificateNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the CertificateNFT contract.
type CertificateNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_CertificateNFT *CertificateNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*CertificateNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &CertificateNFTBatchMetadataUpdateIterator{contract: _CertificateNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_CertificateNFT *CertificateNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *CertificateNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTBatchMetadataUpdate)
				if err := _CertificateNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_CertificateNFT *CertificateNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*CertificateNFTBatchMetadataUpdate, error) {
	event := new(CertificateNFTBatchMetadataUpdate)
	if err := _CertificateNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTCertificateIssuedIterator is returned from FilterCertificateIssued and is used to iterate over the raw logs and unpacked data for CertificateIssued events raised by the CertificateNFT contract.
type CertificateNFTCertificateIssuedIterator struct {
	Event *CertificateNFTCertificateIssued // Event containing the contract specifics and raw log

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
func (it *CertificateNFTCertificateIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTCertificateIssued)
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
		it.Event = new(CertificateNFTCertificateIssued)
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
func (it *CertificateNFTCertificateIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTCertificateIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTCertificateIssued represents a CertificateIssued event raised by the CertificateNFT contract.
type CertificateNFTCertificateIssued struct {
	MemberId   *big.Int
	TokenId    *big.Int
	MemberName string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCertificateIssued is a free log retrieval operation binding the contract event 0x12b8ba46d8efd84182d36de4f91a8f0922f453ef50e9b302b291191863a14ac3.
//
// Solidity: event CertificateIssued(uint256 indexed memberId, uint256 tokenId, string memberName)
func (_CertificateNFT *CertificateNFTFilterer) FilterCertificateIssued(opts *bind.FilterOpts, memberId []*big.Int) (*CertificateNFTCertificateIssuedIterator, error) {

	var memberIdRule []interface{}
	for _, memberIdItem := range memberId {
		memberIdRule = append(memberIdRule, memberIdItem)
	}

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "CertificateIssued", memberIdRule)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTCertificateIssuedIterator{contract: _CertificateNFT.contract, event: "CertificateIssued", logs: logs, sub: sub}, nil
}

// WatchCertificateIssued is a free log subscription operation binding the contract event 0x12b8ba46d8efd84182d36de4f91a8f0922f453ef50e9b302b291191863a14ac3.
//
// Solidity: event CertificateIssued(uint256 indexed memberId, uint256 tokenId, string memberName)
func (_CertificateNFT *CertificateNFTFilterer) WatchCertificateIssued(opts *bind.WatchOpts, sink chan<- *CertificateNFTCertificateIssued, memberId []*big.Int) (event.Subscription, error) {

	var memberIdRule []interface{}
	for _, memberIdItem := range memberId {
		memberIdRule = append(memberIdRule, memberIdItem)
	}

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "CertificateIssued", memberIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTCertificateIssued)
				if err := _CertificateNFT.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
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

// ParseCertificateIssued is a log parse operation binding the contract event 0x12b8ba46d8efd84182d36de4f91a8f0922f453ef50e9b302b291191863a14ac3.
//
// Solidity: event CertificateIssued(uint256 indexed memberId, uint256 tokenId, string memberName)
func (_CertificateNFT *CertificateNFTFilterer) ParseCertificateIssued(log types.Log) (*CertificateNFTCertificateIssued, error) {
	event := new(CertificateNFTCertificateIssued)
	if err := _CertificateNFT.contract.UnpackLog(event, "CertificateIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the CertificateNFT contract.
type CertificateNFTMetadataUpdateIterator struct {
	Event *CertificateNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *CertificateNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTMetadataUpdate)
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
		it.Event = new(CertificateNFTMetadataUpdate)
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
func (it *CertificateNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTMetadataUpdate represents a MetadataUpdate event raised by the CertificateNFT contract.
type CertificateNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_CertificateNFT *CertificateNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*CertificateNFTMetadataUpdateIterator, error) {

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &CertificateNFTMetadataUpdateIterator{contract: _CertificateNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_CertificateNFT *CertificateNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *CertificateNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTMetadataUpdate)
				if err := _CertificateNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_CertificateNFT *CertificateNFTFilterer) ParseMetadataUpdate(log types.Log) (*CertificateNFTMetadataUpdate, error) {
	event := new(CertificateNFTMetadataUpdate)
	if err := _CertificateNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CertificateNFT contract.
type CertificateNFTRoleAdminChangedIterator struct {
	Event *CertificateNFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CertificateNFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTRoleAdminChanged)
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
		it.Event = new(CertificateNFTRoleAdminChanged)
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
func (it *CertificateNFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTRoleAdminChanged represents a RoleAdminChanged event raised by the CertificateNFT contract.
type CertificateNFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CertificateNFT *CertificateNFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CertificateNFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTRoleAdminChangedIterator{contract: _CertificateNFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CertificateNFT *CertificateNFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CertificateNFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTRoleAdminChanged)
				if err := _CertificateNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CertificateNFT *CertificateNFTFilterer) ParseRoleAdminChanged(log types.Log) (*CertificateNFTRoleAdminChanged, error) {
	event := new(CertificateNFTRoleAdminChanged)
	if err := _CertificateNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CertificateNFT contract.
type CertificateNFTRoleGrantedIterator struct {
	Event *CertificateNFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *CertificateNFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTRoleGranted)
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
		it.Event = new(CertificateNFTRoleGranted)
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
func (it *CertificateNFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTRoleGranted represents a RoleGranted event raised by the CertificateNFT contract.
type CertificateNFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CertificateNFT *CertificateNFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CertificateNFTRoleGrantedIterator, error) {

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

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTRoleGrantedIterator{contract: _CertificateNFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CertificateNFT *CertificateNFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CertificateNFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTRoleGranted)
				if err := _CertificateNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CertificateNFT *CertificateNFTFilterer) ParseRoleGranted(log types.Log) (*CertificateNFTRoleGranted, error) {
	event := new(CertificateNFTRoleGranted)
	if err := _CertificateNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CertificateNFT contract.
type CertificateNFTRoleRevokedIterator struct {
	Event *CertificateNFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CertificateNFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTRoleRevoked)
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
		it.Event = new(CertificateNFTRoleRevoked)
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
func (it *CertificateNFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTRoleRevoked represents a RoleRevoked event raised by the CertificateNFT contract.
type CertificateNFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CertificateNFT *CertificateNFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CertificateNFTRoleRevokedIterator, error) {

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

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTRoleRevokedIterator{contract: _CertificateNFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CertificateNFT *CertificateNFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CertificateNFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTRoleRevoked)
				if err := _CertificateNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CertificateNFT *CertificateNFTFilterer) ParseRoleRevoked(log types.Log) (*CertificateNFTRoleRevoked, error) {
	event := new(CertificateNFTRoleRevoked)
	if err := _CertificateNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CertificateNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CertificateNFT contract.
type CertificateNFTTransferIterator struct {
	Event *CertificateNFTTransfer // Event containing the contract specifics and raw log

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
func (it *CertificateNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertificateNFTTransfer)
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
		it.Event = new(CertificateNFTTransfer)
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
func (it *CertificateNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertificateNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertificateNFTTransfer represents a Transfer event raised by the CertificateNFT contract.
type CertificateNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CertificateNFT *CertificateNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CertificateNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CertificateNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CertificateNFTTransferIterator{contract: _CertificateNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CertificateNFT *CertificateNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CertificateNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _CertificateNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertificateNFTTransfer)
				if err := _CertificateNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_CertificateNFT *CertificateNFTFilterer) ParseTransfer(log types.Log) (*CertificateNFTTransfer, error) {
	event := new(CertificateNFTTransfer)
	if err := _CertificateNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
