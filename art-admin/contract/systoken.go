// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systoken

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
)

// SystokenMetaData contains all meta data concerning the Systoken contract.
var SystokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_initBaseURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_initNotRevealedUri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseExtension\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMintAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notRevealedUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revealed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseExtension\",\"type\":\"string\"}],\"name\":\"setBaseExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCost\",\"type\":\"uint256\"}],\"name\":\"setCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_notRevealedURI\",\"type\":\"string\"}],\"name\":\"setNotRevealedURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newmaxMintAmount\",\"type\":\"uint256\"}],\"name\":\"setmaxMintAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"walletOfOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60c06040526005608090815264173539b7b760d91b60a052600c9062000026908262000264565b5066b1a2bc2ec50000600d55612710600e556005600f556010805461ffff191690553480156200005557600080fd5b5060405162002749380380620027498339810160408190526200007891620003df565b8383600062000088838262000264565b50600162000097828262000264565b505050620000b4620000ae620000d460201b60201c565b620000d8565b620000bf826200012a565b620000ca8162000146565b5050505062000498565b3390565b600a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b620001346200015e565b600b62000142828262000264565b5050565b620001506200015e565b601162000142828262000264565b600a546001600160a01b03163314620001bd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640160405180910390fd5b565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620001ea57607f821691505b6020821081036200020b57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200025f57600081815260208120601f850160051c810160208610156200023a5750805b601f850160051c820191505b818110156200025b5782815560010162000246565b5050505b505050565b81516001600160401b03811115620002805762000280620001bf565b6200029881620002918454620001d5565b8462000211565b602080601f831160018114620002d05760008415620002b75750858301515b600019600386901b1c1916600185901b1785556200025b565b600085815260208120601f198616915b828110156200030157888601518255948401946001909101908401620002e0565b5085821015620003205787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600082601f8301126200034257600080fd5b81516001600160401b03808211156200035f576200035f620001bf565b604051601f8301601f19908116603f011681019082821181831017156200038a576200038a620001bf565b81604052838152602092508683858801011115620003a757600080fd5b600091505b83821015620003cb5785820183015181830184015290820190620003ac565b600093810190920192909252949350505050565b60008060008060808587031215620003f657600080fd5b84516001600160401b03808211156200040e57600080fd5b6200041c8883890162000330565b955060208701519150808211156200043357600080fd5b620004418883890162000330565b945060408701519150808211156200045857600080fd5b620004668883890162000330565b935060608701519150808211156200047d57600080fd5b506200048c8782880162000330565b91505092959194509250565b6122a180620004a86000396000f3fe60806040526004361061021a5760003560e01c80635c975abb11610123578063a475b5dd116100ab578063da3ef23f1161006f578063da3ef23f146105dd578063dd4efd5e146105fd578063e985e9c51461061d578063f2c4ce1e14610666578063f2fde38b1461068657600080fd5b8063a475b5dd1461055d578063b88d4fde14610572578063c668286214610592578063c87b56dd146105a7578063d5abeb01146105c757600080fd5b80637f00c7a6116100f25780637f00c7a6146104d75780638da5cb5b146104f757806395d89b4114610515578063a0712d681461052a578063a22cb4651461053d57600080fd5b80635c975abb146104685780636352211e1461048257806370a08231146104a2578063715018a6146104c257600080fd5b806323b872dd116101a6578063438b630011610175578063438b6300146103bc57806344a0d68a146103e95780634f6ccce714610409578063518302271461042957806355f804b31461044857600080fd5b806323b872dd146103545780632f745c59146103745780633ccfd60b1461039457806342842e0e1461039c57600080fd5b8063081c8c44116101ed578063081c8c44146102d0578063095ea7b3146102e557806313faede61461030557806318160ddd14610329578063239c70ae1461033e57600080fd5b806301ffc9a71461021f57806302329a291461025457806306fdde0314610276578063081812fc14610298575b600080fd5b34801561022b57600080fd5b5061023f61023a366004611b58565b6106a6565b60405190151581526020015b60405180910390f35b34801561026057600080fd5b5061027461026f366004611b8a565b6106d1565b005b34801561028257600080fd5b5061028b6106ec565b60405161024b9190611bf5565b3480156102a457600080fd5b506102b86102b3366004611c08565b61077e565b6040516001600160a01b03909116815260200161024b565b3480156102dc57600080fd5b5061028b6107a5565b3480156102f157600080fd5b50610274610300366004611c38565b610833565b34801561031157600080fd5b5061031b600d5481565b60405190815260200161024b565b34801561033557600080fd5b5060085461031b565b34801561034a57600080fd5b5061031b600f5481565b34801561036057600080fd5b5061027461036f366004611c62565b61094d565b34801561038057600080fd5b5061031b61038f366004611c38565b61097e565b610274610a14565b3480156103a857600080fd5b506102746103b7366004611c62565b610a90565b3480156103c857600080fd5b506103dc6103d7366004611c9e565b610aab565b60405161024b9190611cb9565b3480156103f557600080fd5b50610274610404366004611c08565b610b4d565b34801561041557600080fd5b5061031b610424366004611c08565b610b5a565b34801561043557600080fd5b5060105461023f90610100900460ff1681565b34801561045457600080fd5b50610274610463366004611d89565b610bed565b34801561047457600080fd5b5060105461023f9060ff1681565b34801561048e57600080fd5b506102b861049d366004611c08565b610c05565b3480156104ae57600080fd5b5061031b6104bd366004611c9e565b610c65565b3480156104ce57600080fd5b50610274610ceb565b3480156104e357600080fd5b506102746104f2366004611c08565b610cff565b34801561050357600080fd5b50600a546001600160a01b03166102b8565b34801561052157600080fd5b5061028b610d0c565b610274610538366004611c08565b610d1b565b34801561054957600080fd5b50610274610558366004611dd2565b610dc8565b34801561056957600080fd5b50610274610dd3565b34801561057e57600080fd5b5061027461058d366004611e05565b610dec565b34801561059e57600080fd5b5061028b610e24565b3480156105b357600080fd5b5061028b6105c2366004611c08565b610e31565b3480156105d357600080fd5b5061031b600e5481565b3480156105e957600080fd5b506102746105f8366004611d89565b610fb5565b34801561060957600080fd5b506102b8610618366004611c08565b610fc9565b34801561062957600080fd5b5061023f610638366004611e81565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b34801561067257600080fd5b50610274610681366004611d89565b610fd4565b34801561069257600080fd5b506102746106a1366004611c9e565b610fe8565b60006001600160e01b0319821663780e9d6360e01b14806106cb57506106cb8261105e565b92915050565b6106d96110ae565b6010805460ff1916911515919091179055565b6060600080546106fb90611eab565b80601f016020809104026020016040519081016040528092919081815260200182805461072790611eab565b80156107745780601f1061074957610100808354040283529160200191610774565b820191906000526020600020905b81548152906001019060200180831161075757829003601f168201915b5050505050905090565b600061078982611108565b506000908152600460205260409020546001600160a01b031690565b601180546107b290611eab565b80601f01602080910402602001604051908101604052809291908181526020018280546107de90611eab565b801561082b5780601f106108005761010080835404028352916020019161082b565b820191906000526020600020905b81548152906001019060200180831161080e57829003601f168201915b505050505081565b600061083e82610c05565b9050806001600160a01b0316836001600160a01b0316036108b05760405162461bcd60e51b815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e656044820152603960f91b60648201526084015b60405180910390fd5b336001600160a01b03821614806108cc57506108cc8133610638565b61093e5760405162461bcd60e51b815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c000060648201526084016108a7565b6109488383611167565b505050565b61095733826111d5565b6109735760405162461bcd60e51b81526004016108a790611ee5565b610948838383611254565b600061098983610c65565b82106109eb5760405162461bcd60e51b815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201526a74206f6620626f756e647360a81b60648201526084016108a7565b506001600160a01b03919091166000908152600660209081526040808320938352929052205490565b610a1c6110ae565b6000610a30600a546001600160a01b031690565b6001600160a01b03164760405160006040518083038185875af1925050503d8060008114610a7a576040519150601f19603f3d011682016040523d82523d6000602084013e610a7f565b606091505b5050905080610a8d57600080fd5b50565b61094883838360405180602001604052806000815250610dec565b60606000610ab883610c65565b905060008167ffffffffffffffff811115610ad557610ad5611cfd565b604051908082528060200260200182016040528015610afe578160200160208202803683370190505b50905060005b82811015610b4557610b16858261097e565b828281518110610b2857610b28611f33565b602090810291909101015280610b3d81611f5f565b915050610b04565b509392505050565b610b556110ae565b600d55565b6000610b6560085490565b8210610bc85760405162461bcd60e51b815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201526b7574206f6620626f756e647360a01b60648201526084016108a7565b60088281548110610bdb57610bdb611f33565b90600052602060002001549050919050565b610bf56110ae565b600b610c018282611fc6565b5050565b6000818152600260205260408120546001600160a01b0316806106cb5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016108a7565b60006001600160a01b038216610ccf5760405162461bcd60e51b815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f7420612076616044820152683634b21037bbb732b960b91b60648201526084016108a7565b506001600160a01b031660009081526003602052604090205490565b610cf36110ae565b610cfd60006113fb565b565b610d076110ae565b600f55565b6060600180546106fb90611eab565b6000610d2660085490565b60105490915060ff1615610d3957600080fd5b60008211610d4657600080fd5b600f54821115610d5557600080fd5b600e54610d628383612086565b1115610d6d57600080fd5b600a546001600160a01b03163314610d995781600d54610d8d9190612099565b341015610d9957600080fd5b60015b82811161094857610db633610db18385612086565b61144d565b80610dc081611f5f565b915050610d9c565b610c01338383611467565b610ddb6110ae565b6010805461ff001916610100179055565b610df633836111d5565b610e125760405162461bcd60e51b81526004016108a790611ee5565b610e1e84848484611535565b50505050565b600c80546107b290611eab565b6000818152600260205260409020546060906001600160a01b0316610eb05760405162461bcd60e51b815260206004820152602f60248201527f4552433732314d657461646174613a2055524920717565727920666f72206e6f60448201526e3732bc34b9ba32b73a103a37b5b2b760891b60648201526084016108a7565b601054610100900460ff161515600003610f565760118054610ed190611eab565b80601f0160208091040260200160405190810160405280929190818152602001828054610efd90611eab565b8015610f4a5780601f10610f1f57610100808354040283529160200191610f4a565b820191906000526020600020905b815481529060010190602001808311610f2d57829003601f168201915b50505050509050919050565b6000610f60611568565b90506000815111610f805760405180602001604052806000815250610fae565b80610f8a84611577565b600c604051602001610f9e939291906120b8565b6040516020818303038152906040525b9392505050565b610fbd6110ae565b600c610c018282611fc6565b60006106cb82610c05565b610fdc6110ae565b6011610c018282611fc6565b610ff06110ae565b6001600160a01b0381166110555760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016108a7565b610a8d816113fb565b60006001600160e01b031982166380ac58cd60e01b148061108f57506001600160e01b03198216635b5e139f60e01b145b806106cb57506301ffc9a760e01b6001600160e01b03198316146106cb565b600a546001600160a01b03163314610cfd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016108a7565b6000818152600260205260409020546001600160a01b0316610a8d5760405162461bcd60e51b8152602060048201526018602482015277115490cdcc8c4e881a5b9d985b1a59081d1bdad95b88125160421b60448201526064016108a7565b600081815260046020526040902080546001600160a01b0319166001600160a01b038416908117909155819061119c82610c05565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000806111e183610c05565b9050806001600160a01b0316846001600160a01b0316148061122857506001600160a01b0380821660009081526005602090815260408083209388168352929052205460ff165b8061124c5750836001600160a01b03166112418461077e565b6001600160a01b0316145b949350505050565b826001600160a01b031661126782610c05565b6001600160a01b0316146112cb5760405162461bcd60e51b815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201526437bbb732b960d91b60648201526084016108a7565b6001600160a01b03821661132d5760405162461bcd60e51b8152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f206164646044820152637265737360e01b60648201526084016108a7565b611338838383611678565b611343600082611167565b6001600160a01b038316600090815260036020526040812080546001929061136c908490612158565b90915550506001600160a01b038216600090815260036020526040812080546001929061139a908490612086565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600a80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b610c01828260405180602001604052806000815250611730565b816001600160a01b0316836001600160a01b0316036114c85760405162461bcd60e51b815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c65720000000000000060448201526064016108a7565b6001600160a01b03838116600081815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b611540848484611254565b61154c84848484611763565b610e1e5760405162461bcd60e51b81526004016108a79061216b565b6060600b80546106fb90611eab565b60608160000361159e5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156115c857806115b281611f5f565b91506115c19050600a836121d3565b91506115a2565b60008167ffffffffffffffff8111156115e3576115e3611cfd565b6040519080825280601f01601f19166020018201604052801561160d576020820181803683370190505b5090505b841561124c57611622600183612158565b915061162f600a866121e7565b61163a906030612086565b60f81b81838151811061164f5761164f611f33565b60200101906001600160f81b031916908160001a905350611671600a866121d3565b9450611611565b6001600160a01b0383166116d3576116ce81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b6116f6565b816001600160a01b0316836001600160a01b0316146116f6576116f68382611864565b6001600160a01b03821661170d5761094881611901565b826001600160a01b0316826001600160a01b0316146109485761094882826119b0565b61173a83836119f4565b6117476000848484611763565b6109485760405162461bcd60e51b81526004016108a79061216b565b60006001600160a01b0384163b1561185957604051630a85bd0160e11b81526001600160a01b0385169063150b7a02906117a79033908990889088906004016121fb565b6020604051808303816000875af19250505080156117e2575060408051601f3d908101601f191682019092526117df91810190612238565b60015b61183f573d808015611810576040519150601f19603f3d011682016040523d82523d6000602084013e611815565b606091505b5080516000036118375760405162461bcd60e51b81526004016108a79061216b565b805181602001fd5b6001600160e01b031916630a85bd0160e11b14905061124c565b506001949350505050565b6000600161187184610c65565b61187b9190612158565b6000838152600760205260409020549091508082146118ce576001600160a01b03841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b5060009182526007602090815260408084208490556001600160a01b039094168352600681528383209183525290812055565b60085460009061191390600190612158565b6000838152600960205260408120546008805493945090928490811061193b5761193b611f33565b90600052602060002001549050806008838154811061195c5761195c611f33565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061199457611994612255565b6001900381819060005260206000200160009055905550505050565b60006119bb83610c65565b6001600160a01b039093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b6001600160a01b038216611a4a5760405162461bcd60e51b815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f206164647265737360448201526064016108a7565b6000818152600260205260409020546001600160a01b031615611aaf5760405162461bcd60e51b815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060448201526064016108a7565b611abb60008383611678565b6001600160a01b0382166000908152600360205260408120805460019290611ae4908490612086565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b6001600160e01b031981168114610a8d57600080fd5b600060208284031215611b6a57600080fd5b8135610fae81611b42565b80358015158114611b8557600080fd5b919050565b600060208284031215611b9c57600080fd5b610fae82611b75565b60005b83811015611bc0578181015183820152602001611ba8565b50506000910152565b60008151808452611be1816020860160208601611ba5565b601f01601f19169290920160200192915050565b602081526000610fae6020830184611bc9565b600060208284031215611c1a57600080fd5b5035919050565b80356001600160a01b0381168114611b8557600080fd5b60008060408385031215611c4b57600080fd5b611c5483611c21565b946020939093013593505050565b600080600060608486031215611c7757600080fd5b611c8084611c21565b9250611c8e60208501611c21565b9150604084013590509250925092565b600060208284031215611cb057600080fd5b610fae82611c21565b6020808252825182820181905260009190848201906040850190845b81811015611cf157835183529284019291840191600101611cd5565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff80841115611d2e57611d2e611cfd565b604051601f8501601f19908116603f01168101908282118183101715611d5657611d56611cfd565b81604052809350858152868686011115611d6f57600080fd5b858560208301376000602087830101525050509392505050565b600060208284031215611d9b57600080fd5b813567ffffffffffffffff811115611db257600080fd5b8201601f81018413611dc357600080fd5b61124c84823560208401611d13565b60008060408385031215611de557600080fd5b611dee83611c21565b9150611dfc60208401611b75565b90509250929050565b60008060008060808587031215611e1b57600080fd5b611e2485611c21565b9350611e3260208601611c21565b925060408501359150606085013567ffffffffffffffff811115611e5557600080fd5b8501601f81018713611e6657600080fd5b611e7587823560208401611d13565b91505092959194509250565b60008060408385031215611e9457600080fd5b611e9d83611c21565b9150611dfc60208401611c21565b600181811c90821680611ebf57607f821691505b602082108103611edf57634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252602e908201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560408201526d1c881b9bdc88185c1c1c9bdd995960921b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611f7157611f71611f49565b5060010190565b601f82111561094857600081815260208120601f850160051c81016020861015611f9f5750805b601f850160051c820191505b81811015611fbe57828155600101611fab565b505050505050565b815167ffffffffffffffff811115611fe057611fe0611cfd565b611ff481611fee8454611eab565b84611f78565b602080601f83116001811461202957600084156120115750858301515b600019600386901b1c1916600185901b178555611fbe565b600085815260208120601f198616915b8281101561205857888601518255948401946001909101908401612039565b50858210156120765787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156106cb576106cb611f49565b60008160001904831182151516156120b3576120b3611f49565b500290565b6000845160206120cb8285838a01611ba5565b8551918401916120de8184848a01611ba5565b85549201916000906120ef81611eab565b60018281168015612107576001811461211c57612148565b60ff1984168752821515830287019450612148565b896000528560002060005b8481101561214057815489820152908301908701612127565b505082870194505b50929a9950505050505050505050565b818103818111156106cb576106cb611f49565b60208082526032908201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560408201527131b2b4bb32b91034b6b83632b6b2b73a32b960711b606082015260800190565b634e487b7160e01b600052601260045260246000fd5b6000826121e2576121e26121bd565b500490565b6000826121f6576121f66121bd565b500690565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061222e90830184611bc9565b9695505050505050565b60006020828403121561224a57600080fd5b8151610fae81611b42565b634e487b7160e01b600052603160045260246000fdfea2646970667358221220bf948797bbd2d96dc6cb692e04e680f60f866473f342c0c2ba4872c6f7d96e0364736f6c63430008100033",
}

// SystokenABI is the input ABI used to generate the binding from.
// Deprecated: Use SystokenMetaData.ABI instead.
var SystokenABI = SystokenMetaData.ABI

// SystokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystokenMetaData.Bin instead.
var SystokenBin = SystokenMetaData.Bin

// DeploySystoken deploys a new Ethereum contract, binding an instance of Systoken to it.
func DeploySystoken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _initBaseURI string, _initNotRevealedUri string) (common.Address, *types.Transaction, *Systoken, error) {
	parsed, err := SystokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystokenBin), backend, _name, _symbol, _initBaseURI, _initNotRevealedUri)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Systoken{SystokenCaller: SystokenCaller{contract: contract}, SystokenTransactor: SystokenTransactor{contract: contract}, SystokenFilterer: SystokenFilterer{contract: contract}}, nil
}

// Systoken is an auto generated Go binding around an Ethereum contract.
type Systoken struct {
	SystokenCaller     // Read-only binding to the contract
	SystokenTransactor // Write-only binding to the contract
	SystokenFilterer   // Log filterer for contract events
}

// SystokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystokenSession struct {
	Contract     *Systoken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystokenCallerSession struct {
	Contract *SystokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SystokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystokenTransactorSession struct {
	Contract     *SystokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SystokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystokenRaw struct {
	Contract *Systoken // Generic contract binding to access the raw methods on
}

// SystokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystokenCallerRaw struct {
	Contract *SystokenCaller // Generic read-only contract binding to access the raw methods on
}

// SystokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystokenTransactorRaw struct {
	Contract *SystokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystoken creates a new instance of Systoken, bound to a specific deployed contract.
func NewSystoken(address common.Address, backend bind.ContractBackend) (*Systoken, error) {
	contract, err := bindSystoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Systoken{SystokenCaller: SystokenCaller{contract: contract}, SystokenTransactor: SystokenTransactor{contract: contract}, SystokenFilterer: SystokenFilterer{contract: contract}}, nil
}

// NewSystokenCaller creates a new read-only instance of Systoken, bound to a specific deployed contract.
func NewSystokenCaller(address common.Address, caller bind.ContractCaller) (*SystokenCaller, error) {
	contract, err := bindSystoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystokenCaller{contract: contract}, nil
}

// NewSystokenTransactor creates a new write-only instance of Systoken, bound to a specific deployed contract.
func NewSystokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SystokenTransactor, error) {
	contract, err := bindSystoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystokenTransactor{contract: contract}, nil
}

// NewSystokenFilterer creates a new log filterer instance of Systoken, bound to a specific deployed contract.
func NewSystokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SystokenFilterer, error) {
	contract, err := bindSystoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystokenFilterer{contract: contract}, nil
}

// bindSystoken binds a generic wrapper to an already deployed contract.
func bindSystoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systoken *SystokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systoken.Contract.SystokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systoken *SystokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systoken.Contract.SystokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systoken *SystokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systoken.Contract.SystokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systoken *SystokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systoken *SystokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systoken *SystokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systoken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Systoken *SystokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Systoken *SystokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Systoken.Contract.BalanceOf(&_Systoken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Systoken *SystokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Systoken.Contract.BalanceOf(&_Systoken.CallOpts, owner)
}

// BaseExtension is a free data retrieval call binding the contract method 0xc6682862.
//
// Solidity: function baseExtension() view returns(string)
func (_Systoken *SystokenCaller) BaseExtension(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "baseExtension")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseExtension is a free data retrieval call binding the contract method 0xc6682862.
//
// Solidity: function baseExtension() view returns(string)
func (_Systoken *SystokenSession) BaseExtension() (string, error) {
	return _Systoken.Contract.BaseExtension(&_Systoken.CallOpts)
}

// BaseExtension is a free data retrieval call binding the contract method 0xc6682862.
//
// Solidity: function baseExtension() view returns(string)
func (_Systoken *SystokenCallerSession) BaseExtension() (string, error) {
	return _Systoken.Contract.BaseExtension(&_Systoken.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_Systoken *SystokenCaller) Cost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "cost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_Systoken *SystokenSession) Cost() (*big.Int, error) {
	return _Systoken.Contract.Cost(&_Systoken.CallOpts)
}

// Cost is a free data retrieval call binding the contract method 0x13faede6.
//
// Solidity: function cost() view returns(uint256)
func (_Systoken *SystokenCallerSession) Cost() (*big.Int, error) {
	return _Systoken.Contract.Cost(&_Systoken.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Systoken *SystokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Systoken *SystokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Systoken.Contract.GetApproved(&_Systoken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Systoken *SystokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Systoken.Contract.GetApproved(&_Systoken.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Systoken *SystokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Systoken *SystokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Systoken.Contract.IsApprovedForAll(&_Systoken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Systoken *SystokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Systoken.Contract.IsApprovedForAll(&_Systoken.CallOpts, owner, operator)
}

// MaxMintAmount is a free data retrieval call binding the contract method 0x239c70ae.
//
// Solidity: function maxMintAmount() view returns(uint256)
func (_Systoken *SystokenCaller) MaxMintAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "maxMintAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMintAmount is a free data retrieval call binding the contract method 0x239c70ae.
//
// Solidity: function maxMintAmount() view returns(uint256)
func (_Systoken *SystokenSession) MaxMintAmount() (*big.Int, error) {
	return _Systoken.Contract.MaxMintAmount(&_Systoken.CallOpts)
}

// MaxMintAmount is a free data retrieval call binding the contract method 0x239c70ae.
//
// Solidity: function maxMintAmount() view returns(uint256)
func (_Systoken *SystokenCallerSession) MaxMintAmount() (*big.Int, error) {
	return _Systoken.Contract.MaxMintAmount(&_Systoken.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Systoken *SystokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Systoken *SystokenSession) MaxSupply() (*big.Int, error) {
	return _Systoken.Contract.MaxSupply(&_Systoken.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Systoken *SystokenCallerSession) MaxSupply() (*big.Int, error) {
	return _Systoken.Contract.MaxSupply(&_Systoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Systoken *SystokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Systoken *SystokenSession) Name() (string, error) {
	return _Systoken.Contract.Name(&_Systoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Systoken *SystokenCallerSession) Name() (string, error) {
	return _Systoken.Contract.Name(&_Systoken.CallOpts)
}

// NotRevealedUri is a free data retrieval call binding the contract method 0x081c8c44.
//
// Solidity: function notRevealedUri() view returns(string)
func (_Systoken *SystokenCaller) NotRevealedUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "notRevealedUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NotRevealedUri is a free data retrieval call binding the contract method 0x081c8c44.
//
// Solidity: function notRevealedUri() view returns(string)
func (_Systoken *SystokenSession) NotRevealedUri() (string, error) {
	return _Systoken.Contract.NotRevealedUri(&_Systoken.CallOpts)
}

// NotRevealedUri is a free data retrieval call binding the contract method 0x081c8c44.
//
// Solidity: function notRevealedUri() view returns(string)
func (_Systoken *SystokenCallerSession) NotRevealedUri() (string, error) {
	return _Systoken.Contract.NotRevealedUri(&_Systoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Systoken *SystokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Systoken *SystokenSession) Owner() (common.Address, error) {
	return _Systoken.Contract.Owner(&_Systoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Systoken *SystokenCallerSession) Owner() (common.Address, error) {
	return _Systoken.Contract.Owner(&_Systoken.CallOpts)
}

// OwnerById is a free data retrieval call binding the contract method 0xdd4efd5e.
//
// Solidity: function ownerById(uint256 tokenId) view returns(address)
func (_Systoken *SystokenCaller) OwnerById(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "ownerById", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerById is a free data retrieval call binding the contract method 0xdd4efd5e.
//
// Solidity: function ownerById(uint256 tokenId) view returns(address)
func (_Systoken *SystokenSession) OwnerById(tokenId *big.Int) (common.Address, error) {
	return _Systoken.Contract.OwnerById(&_Systoken.CallOpts, tokenId)
}

// OwnerById is a free data retrieval call binding the contract method 0xdd4efd5e.
//
// Solidity: function ownerById(uint256 tokenId) view returns(address)
func (_Systoken *SystokenCallerSession) OwnerById(tokenId *big.Int) (common.Address, error) {
	return _Systoken.Contract.OwnerById(&_Systoken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Systoken *SystokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Systoken *SystokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Systoken.Contract.OwnerOf(&_Systoken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Systoken *SystokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Systoken.Contract.OwnerOf(&_Systoken.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Systoken *SystokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Systoken *SystokenSession) Paused() (bool, error) {
	return _Systoken.Contract.Paused(&_Systoken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Systoken *SystokenCallerSession) Paused() (bool, error) {
	return _Systoken.Contract.Paused(&_Systoken.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_Systoken *SystokenCaller) Revealed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "revealed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_Systoken *SystokenSession) Revealed() (bool, error) {
	return _Systoken.Contract.Revealed(&_Systoken.CallOpts)
}

// Revealed is a free data retrieval call binding the contract method 0x51830227.
//
// Solidity: function revealed() view returns(bool)
func (_Systoken *SystokenCallerSession) Revealed() (bool, error) {
	return _Systoken.Contract.Revealed(&_Systoken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Systoken *SystokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Systoken *SystokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Systoken.Contract.SupportsInterface(&_Systoken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Systoken *SystokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Systoken.Contract.SupportsInterface(&_Systoken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Systoken *SystokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Systoken *SystokenSession) Symbol() (string, error) {
	return _Systoken.Contract.Symbol(&_Systoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Systoken *SystokenCallerSession) Symbol() (string, error) {
	return _Systoken.Contract.Symbol(&_Systoken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Systoken *SystokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Systoken *SystokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Systoken.Contract.TokenByIndex(&_Systoken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Systoken *SystokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Systoken.Contract.TokenByIndex(&_Systoken.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Systoken *SystokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Systoken *SystokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Systoken.Contract.TokenOfOwnerByIndex(&_Systoken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Systoken *SystokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Systoken.Contract.TokenOfOwnerByIndex(&_Systoken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Systoken *SystokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Systoken *SystokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Systoken.Contract.TokenURI(&_Systoken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Systoken *SystokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Systoken.Contract.TokenURI(&_Systoken.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Systoken *SystokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Systoken *SystokenSession) TotalSupply() (*big.Int, error) {
	return _Systoken.Contract.TotalSupply(&_Systoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Systoken *SystokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Systoken.Contract.TotalSupply(&_Systoken.CallOpts)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_Systoken *SystokenCaller) WalletOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Systoken.contract.Call(opts, &out, "walletOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_Systoken *SystokenSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Systoken.Contract.WalletOfOwner(&_Systoken.CallOpts, _owner)
}

// WalletOfOwner is a free data retrieval call binding the contract method 0x438b6300.
//
// Solidity: function walletOfOwner(address _owner) view returns(uint256[])
func (_Systoken *SystokenCallerSession) WalletOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _Systoken.Contract.WalletOfOwner(&_Systoken.CallOpts, _owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Systoken *SystokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Systoken *SystokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.Approve(&_Systoken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Systoken *SystokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.Approve(&_Systoken.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _mintAmount) payable returns()
func (_Systoken *SystokenTransactor) Mint(opts *bind.TransactOpts, _mintAmount *big.Int) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "mint", _mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _mintAmount) payable returns()
func (_Systoken *SystokenSession) Mint(_mintAmount *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.Mint(&_Systoken.TransactOpts, _mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _mintAmount) payable returns()
func (_Systoken *SystokenTransactorSession) Mint(_mintAmount *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.Mint(&_Systoken.TransactOpts, _mintAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Systoken *SystokenTransactor) Pause(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "pause", _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Systoken *SystokenSession) Pause(_state bool) (*types.Transaction, error) {
	return _Systoken.Contract.Pause(&_Systoken.TransactOpts, _state)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _state) returns()
func (_Systoken *SystokenTransactorSession) Pause(_state bool) (*types.Transaction, error) {
	return _Systoken.Contract.Pause(&_Systoken.TransactOpts, _state)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Systoken *SystokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Systoken *SystokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Systoken.Contract.RenounceOwnership(&_Systoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Systoken *SystokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Systoken.Contract.RenounceOwnership(&_Systoken.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_Systoken *SystokenTransactor) Reveal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "reveal")
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_Systoken *SystokenSession) Reveal() (*types.Transaction, error) {
	return _Systoken.Contract.Reveal(&_Systoken.TransactOpts)
}

// Reveal is a paid mutator transaction binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() returns()
func (_Systoken *SystokenTransactorSession) Reveal() (*types.Transaction, error) {
	return _Systoken.Contract.Reveal(&_Systoken.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Systoken *SystokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Systoken *SystokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.SafeTransferFrom(&_Systoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Systoken *SystokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.SafeTransferFrom(&_Systoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Systoken *SystokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Systoken *SystokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Systoken.Contract.SafeTransferFrom0(&_Systoken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Systoken *SystokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Systoken.Contract.SafeTransferFrom0(&_Systoken.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Systoken *SystokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Systoken *SystokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Systoken.Contract.SetApprovalForAll(&_Systoken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Systoken *SystokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Systoken.Contract.SetApprovalForAll(&_Systoken.TransactOpts, operator, approved)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string _newBaseExtension) returns()
func (_Systoken *SystokenTransactor) SetBaseExtension(opts *bind.TransactOpts, _newBaseExtension string) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "setBaseExtension", _newBaseExtension)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string _newBaseExtension) returns()
func (_Systoken *SystokenSession) SetBaseExtension(_newBaseExtension string) (*types.Transaction, error) {
	return _Systoken.Contract.SetBaseExtension(&_Systoken.TransactOpts, _newBaseExtension)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string _newBaseExtension) returns()
func (_Systoken *SystokenTransactorSession) SetBaseExtension(_newBaseExtension string) (*types.Transaction, error) {
	return _Systoken.Contract.SetBaseExtension(&_Systoken.TransactOpts, _newBaseExtension)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Systoken *SystokenTransactor) SetBaseURI(opts *bind.TransactOpts, _newBaseURI string) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "setBaseURI", _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Systoken *SystokenSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Systoken.Contract.SetBaseURI(&_Systoken.TransactOpts, _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_Systoken *SystokenTransactorSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _Systoken.Contract.SetBaseURI(&_Systoken.TransactOpts, _newBaseURI)
}

// SetCost is a paid mutator transaction binding the contract method 0x44a0d68a.
//
// Solidity: function setCost(uint256 _newCost) returns()
func (_Systoken *SystokenTransactor) SetCost(opts *bind.TransactOpts, _newCost *big.Int) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "setCost", _newCost)
}

// SetCost is a paid mutator transaction binding the contract method 0x44a0d68a.
//
// Solidity: function setCost(uint256 _newCost) returns()
func (_Systoken *SystokenSession) SetCost(_newCost *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.SetCost(&_Systoken.TransactOpts, _newCost)
}

// SetCost is a paid mutator transaction binding the contract method 0x44a0d68a.
//
// Solidity: function setCost(uint256 _newCost) returns()
func (_Systoken *SystokenTransactorSession) SetCost(_newCost *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.SetCost(&_Systoken.TransactOpts, _newCost)
}

// SetNotRevealedURI is a paid mutator transaction binding the contract method 0xf2c4ce1e.
//
// Solidity: function setNotRevealedURI(string _notRevealedURI) returns()
func (_Systoken *SystokenTransactor) SetNotRevealedURI(opts *bind.TransactOpts, _notRevealedURI string) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "setNotRevealedURI", _notRevealedURI)
}

// SetNotRevealedURI is a paid mutator transaction binding the contract method 0xf2c4ce1e.
//
// Solidity: function setNotRevealedURI(string _notRevealedURI) returns()
func (_Systoken *SystokenSession) SetNotRevealedURI(_notRevealedURI string) (*types.Transaction, error) {
	return _Systoken.Contract.SetNotRevealedURI(&_Systoken.TransactOpts, _notRevealedURI)
}

// SetNotRevealedURI is a paid mutator transaction binding the contract method 0xf2c4ce1e.
//
// Solidity: function setNotRevealedURI(string _notRevealedURI) returns()
func (_Systoken *SystokenTransactorSession) SetNotRevealedURI(_notRevealedURI string) (*types.Transaction, error) {
	return _Systoken.Contract.SetNotRevealedURI(&_Systoken.TransactOpts, _notRevealedURI)
}

// SetmaxMintAmount is a paid mutator transaction binding the contract method 0x7f00c7a6.
//
// Solidity: function setmaxMintAmount(uint256 _newmaxMintAmount) returns()
func (_Systoken *SystokenTransactor) SetmaxMintAmount(opts *bind.TransactOpts, _newmaxMintAmount *big.Int) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "setmaxMintAmount", _newmaxMintAmount)
}

// SetmaxMintAmount is a paid mutator transaction binding the contract method 0x7f00c7a6.
//
// Solidity: function setmaxMintAmount(uint256 _newmaxMintAmount) returns()
func (_Systoken *SystokenSession) SetmaxMintAmount(_newmaxMintAmount *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.SetmaxMintAmount(&_Systoken.TransactOpts, _newmaxMintAmount)
}

// SetmaxMintAmount is a paid mutator transaction binding the contract method 0x7f00c7a6.
//
// Solidity: function setmaxMintAmount(uint256 _newmaxMintAmount) returns()
func (_Systoken *SystokenTransactorSession) SetmaxMintAmount(_newmaxMintAmount *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.SetmaxMintAmount(&_Systoken.TransactOpts, _newmaxMintAmount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Systoken *SystokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Systoken *SystokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.TransferFrom(&_Systoken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Systoken *SystokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Systoken.Contract.TransferFrom(&_Systoken.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Systoken *SystokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Systoken *SystokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Systoken.Contract.TransferOwnership(&_Systoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Systoken *SystokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Systoken.Contract.TransferOwnership(&_Systoken.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Systoken *SystokenTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systoken.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Systoken *SystokenSession) Withdraw() (*types.Transaction, error) {
	return _Systoken.Contract.Withdraw(&_Systoken.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Systoken *SystokenTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Systoken.Contract.Withdraw(&_Systoken.TransactOpts)
}

// SystokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Systoken contract.
type SystokenApprovalIterator struct {
	Event *SystokenApproval // Event containing the contract specifics and raw log

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
func (it *SystokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystokenApproval)
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
		it.Event = new(SystokenApproval)
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
func (it *SystokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystokenApproval represents a Approval event raised by the Systoken contract.
type SystokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Systoken *SystokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SystokenApprovalIterator, error) {

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

	logs, sub, err := _Systoken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystokenApprovalIterator{contract: _Systoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Systoken *SystokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SystokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Systoken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystokenApproval)
				if err := _Systoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Systoken *SystokenFilterer) ParseApproval(log types.Log) (*SystokenApproval, error) {
	event := new(SystokenApproval)
	if err := _Systoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Systoken contract.
type SystokenApprovalForAllIterator struct {
	Event *SystokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SystokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystokenApprovalForAll)
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
		it.Event = new(SystokenApprovalForAll)
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
func (it *SystokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystokenApprovalForAll represents a ApprovalForAll event raised by the Systoken contract.
type SystokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Systoken *SystokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SystokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Systoken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SystokenApprovalForAllIterator{contract: _Systoken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Systoken *SystokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SystokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Systoken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystokenApprovalForAll)
				if err := _Systoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Systoken *SystokenFilterer) ParseApprovalForAll(log types.Log) (*SystokenApprovalForAll, error) {
	event := new(SystokenApprovalForAll)
	if err := _Systoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Systoken contract.
type SystokenOwnershipTransferredIterator struct {
	Event *SystokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystokenOwnershipTransferred)
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
		it.Event = new(SystokenOwnershipTransferred)
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
func (it *SystokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystokenOwnershipTransferred represents a OwnershipTransferred event raised by the Systoken contract.
type SystokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Systoken *SystokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Systoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystokenOwnershipTransferredIterator{contract: _Systoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Systoken *SystokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Systoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystokenOwnershipTransferred)
				if err := _Systoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Systoken *SystokenFilterer) ParseOwnershipTransferred(log types.Log) (*SystokenOwnershipTransferred, error) {
	event := new(SystokenOwnershipTransferred)
	if err := _Systoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Systoken contract.
type SystokenTransferIterator struct {
	Event *SystokenTransfer // Event containing the contract specifics and raw log

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
func (it *SystokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystokenTransfer)
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
		it.Event = new(SystokenTransfer)
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
func (it *SystokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystokenTransfer represents a Transfer event raised by the Systoken contract.
type SystokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Systoken *SystokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SystokenTransferIterator, error) {

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

	logs, sub, err := _Systoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SystokenTransferIterator{contract: _Systoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Systoken *SystokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SystokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Systoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystokenTransfer)
				if err := _Systoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Systoken *SystokenFilterer) ParseTransfer(log types.Log) (*SystokenTransfer, error) {
	event := new(SystokenTransfer)
	if err := _Systoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
