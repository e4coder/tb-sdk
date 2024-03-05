package tbsdk_test

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	tbsdk "github.com/e4coder/tb-sdk"
	"github.com/ethereum/go-ethereum/common"
)

const TestContractCreationCode = "60806040526040516102da3803806102da83398181016040528101906100259190610103565b815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806001819055505050610141565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61009f82610076565b9050919050565b6100af81610095565b81146100b9575f80fd5b50565b5f815190506100ca816100a6565b92915050565b5f819050919050565b6100e2816100d0565b81146100ec575f80fd5b50565b5f815190506100fd816100d9565b92915050565b5f806040838503121561011957610118610072565b5b5f610126858286016100bc565b9250506020610137858286016100ef565b9150509250929050565b61018c8061014e5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806312065fe0146100435780638da5cb5b14610061578063c29855781461007f575b5f80fd5b61004b61009d565b60405161005891906100e5565b60405180910390f35b6100696100a4565b604051610076919061013d565b60405180910390f35b6100876100c7565b60405161009491906100e5565b60405180910390f35b5f47905090565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b5f819050919050565b6100df816100cd565b82525050565b5f6020820190506100f85f8301846100d6565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610127826100fe565b9050919050565b6101378161011d565b82525050565b5f6020820190506101505f83018461012e565b9291505056fea26469706673582212209160fc917f6c728557338ef43a745e4577c9e3fe7d4e94dd399cb4d508b5fdff64736f6c63430008180033"
const TestOwner = "0xB77e8533ce316AC7355191bf775F1Be1967650eD"
const TestAccountContractCreationCode = "60806040523480156200001157600080fd5b506040516200136e3803806200136e8339810160408190526200003491620000a6565b600280546001600160a01b03199081166001600160a01b039586161790915560018054821693851693909317909255600380548316918416919091179055600080549091169290911691909117905562000103565b80516001600160a01b0381168114620000a157600080fd5b919050565b60008060008060808587031215620000bd57600080fd5b620000c88562000089565b9350620000d86020860162000089565b9250620000e86040860162000089565b9150620000f86060860162000089565b905092959194509250565b61125b80620001136000396000f3fe6080604052600436106100eb5760003560e01c80638da5cb5b1161008a578063d6762c4a11610059578063d6762c4a14610314578063f23a6e6114610334578063f2fde38b1461037a578063f8a8fd6d1461039a57600080fd5b80638da5cb5b1461026c578063b0d691fe1461028c578063b61d27f6146102ac578063bc197c81146102cc57600080fd5b8063150b7a02116100c6578063150b7a02146101af57806318dfb3c7146101f45780633a871cdd14610214578063687cd9c11461023457600080fd5b806223de291461012f57806301ffc9a71461015657806306661abd1461018b57600080fd5b3661012a57604080513381523460208201527f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874910160405180910390a1005b600080fd5b34801561013b57600080fd5b5061015461014a366004610d6f565b5050505050505050565b005b34801561016257600080fd5b50610176610171366004610e1a565b6103af565b60405190151581526020015b60405180910390f35b34801561019757600080fd5b506101a160045481565b604051908152602001610182565b3480156101bb57600080fd5b506101db6101ca366004610e4b565b630a85bd0160e11b95945050505050565b6040516001600160e01b03199091168152602001610182565b34801561020057600080fd5b5061015461020f366004610eff565b610433565b34801561022057600080fd5b506101a161022f366004610f6b565b610531565b34801561024057600080fd5b50600354610254906001600160a01b031681565b6040516001600160a01b039091168152602001610182565b34801561027857600080fd5b50600054610254906001600160a01b031681565b34801561029857600080fd5b50600154610254906001600160a01b031681565b3480156102b857600080fd5b506101546102c7366004610fbf565b6105fc565b3480156102d857600080fd5b506101db6102e736600461100d565b7fbc197c810000000000000000000000000000000000000000000000000000000098975050505050505050565b34801561032057600080fd5b50600254610254906001600160a01b031681565b34801561034057600080fd5b506101db61034f3660046110a7565b7ff23a6e61000000000000000000000000000000000000000000000000000000009695505050505050565b34801561038657600080fd5b5061015461039536600461111f565b61064b565b3480156103a657600080fd5b506101546107ef565b60006001600160e01b03198216630a85bd0160e11b14806103f957506001600160e01b031982167f4e2312e000000000000000000000000000000000000000000000000000000000145b8061042d57506001600160e01b031982167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b61043b610806565b828114610483576040517f223ef21e00000000000000000000000000000000000000000000000000000000815260048101829052602481018490526044015b60405180910390fd5b60005b8381101561052a576105188585838181106104a3576104a361113a565b90506020020160208101906104b8919061111f565b60008585858181106104cc576104cc61113a565b90506020028101906104de9190611150565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061084e92505050565b8061052281611197565b915050610486565b5050505050565b600061053b610806565b600360009054906101000a90046001600160a01b03166001600160a01b031663bb25e1326040518163ffffffff1660e01b8152600401602060405180830381865afa15801561058e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b291906111be565b6105ea576040517f7fe4a4bc00000000000000000000000000000000000000000000000000000000815233600482015260240161047a565b6105f484846108be565b949350505050565b610604610806565b610645848484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061084e92505050565b50505050565b6001600160a01b038116158061066957506001600160a01b03811630145b8061068157506000546001600160a01b038281169116145b156106c3576040517fb20f76e30000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161047a565b6000546001600160a01b031633148015906106de5750333014155b801561076c57506002546040517fd78b27b10000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b039091169063d78b27b190602401602060405180830381865afa158015610746573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061076a91906111be565b155b1561078c57604051634a0bfec160e01b815233600482015260240161047a565b600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a350565b600480549060006107ff83611197565b9190505550565b6001546001600160a01b0316331480159061082c57506000546001600160a01b03163314155b1561084c57604051634a0bfec160e01b815233600482015260240161047a565b565b600080846001600160a01b0316848460405161086a91906111e0565b60006040518083038185875af1925050503d80600081146108a7576040519150601f19603f3d011682016040523d82523d6000602084013e6108ac565b606091505b50915091508161052a57805160208201fd5b600080610918836040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905061096861092b610140860186611150565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525085939250506109919050565b6000546001600160a01b0390811691161461098757600191505061042d565b5060009392505050565b60008060006109a085856109b5565b915091506109ad81610a23565b509392505050565b60008082516041036109eb5760208301516040840151606085015160001a6109df87828585610bdc565b94509450505050610a1c565b8251604003610a145760208301516040840151610a09868383610cc9565b935093505050610a1c565b506000905060025b9250929050565b6000816004811115610a3757610a3761120f565b03610a3f5750565b6001816004811115610a5357610a5361120f565b03610aa05760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161047a565b6002816004811115610ab457610ab461120f565b03610b015760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161047a565b6003816004811115610b1557610b1561120f565b03610b6d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161047a565b6004816004811115610b8157610b8161120f565b03610bd95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b606482015260840161047a565b50565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610c135750600090506003610cc0565b8460ff16601b14158015610c2b57508460ff16601c14155b15610c3c5750600090506004610cc0565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610c90573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116610cb957600060019250925050610cc0565b9150600090505b94509492505050565b6000807f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831660ff84901c601b01610d0387828885610bdc565b935093505050935093915050565b80356001600160a01b0381168114610d2857600080fd5b919050565b60008083601f840112610d3f57600080fd5b50813567ffffffffffffffff811115610d5757600080fd5b602083019150836020828501011115610a1c57600080fd5b60008060008060008060008060c0898b031215610d8b57600080fd5b610d9489610d11565b9750610da260208a01610d11565b9650610db060408a01610d11565b955060608901359450608089013567ffffffffffffffff80821115610dd457600080fd5b610de08c838d01610d2d565b909650945060a08b0135915080821115610df957600080fd5b50610e068b828c01610d2d565b999c989b5096995094979396929594505050565b600060208284031215610e2c57600080fd5b81356001600160e01b031981168114610e4457600080fd5b9392505050565b600080600080600060808688031215610e6357600080fd5b610e6c86610d11565b9450610e7a60208701610d11565b935060408601359250606086013567ffffffffffffffff811115610e9d57600080fd5b610ea988828901610d2d565b969995985093965092949392505050565b60008083601f840112610ecc57600080fd5b50813567ffffffffffffffff811115610ee457600080fd5b6020830191508360208260051b8501011115610a1c57600080fd5b60008060008060408587031215610f1557600080fd5b843567ffffffffffffffff80821115610f2d57600080fd5b610f3988838901610eba565b90965094506020870135915080821115610f5257600080fd5b50610f5f87828801610eba565b95989497509550505050565b600080600060608486031215610f8057600080fd5b833567ffffffffffffffff811115610f9757600080fd5b84016101608187031215610faa57600080fd5b95602085013595506040909401359392505050565b60008060008060608587031215610fd557600080fd5b610fde85610d11565b935060208501359250604085013567ffffffffffffffff81111561100157600080fd5b610f5f87828801610d2d565b60008060008060008060008060a0898b03121561102957600080fd5b61103289610d11565b975061104060208a01610d11565b9650604089013567ffffffffffffffff8082111561105d57600080fd5b6110698c838d01610eba565b909850965060608b013591508082111561108257600080fd5b61108e8c838d01610eba565b909650945060808b0135915080821115610df957600080fd5b60008060008060008060a087890312156110c057600080fd5b6110c987610d11565b95506110d760208801610d11565b94506040870135935060608701359250608087013567ffffffffffffffff81111561110157600080fd5b61110d89828a01610d2d565b979a9699509497509295939492505050565b60006020828403121561113157600080fd5b610e4482610d11565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261116757600080fd5b83018035915067ffffffffffffffff82111561118257600080fd5b602001915036819003821315610a1c57600080fd5b6000600182016111b757634e487b7160e01b600052601160045260246000fd5b5060010190565b6000602082840312156111d057600080fd5b81518015158114610e4457600080fd5b6000825160005b8181101561120157602081860181015185830152016111e7565b506000920191825250919050565b634e487b7160e01b600052602160045260246000fdfea26469706673582212203e2b85dbafffb29ff3e8021a205432dc067a80bdea97663bb3b27b4284cc924064736f6c63430008130033"

var Foo *big.Int = big.NewInt(99)

func TestComputeCounterFactualAddress(t *testing.T) {
	ownerBytes, _ := hex.DecodeString(strings.TrimPrefix(TestOwner, "0x"))
	fooBytes := Foo.Bytes()
	salt, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000001")
	sender := common.HexToAddress("0x9D7f74d0C41E726EC95884E0e97Fa6129e3b5E99").Bytes()
	creationCode, _ := hex.DecodeString(TestContractCreationCode)

	initCode := tbsdk.AbiEncodePacked(creationCode, tbsdk.AbiEncode(
		ownerBytes,
		fooBytes,
	))

	address := tbsdk.ComputeCounterfactualAddress(sender, salt, initCode)

	deployedAddress := "0x603119ed83c3070b4f77a332769618422f832aa3"

	if address != deployedAddress {
		t.Log(address)
		t.FailNow()
	}
}

func TestAccountContractCounterFactualAddress(t *testing.T) {
	owner, _ := hex.DecodeString("494E8f1c10bb14Bea02C2f16cFB33a84BC57ef74")
	factory, _ := hex.DecodeString("4BBa2E1c4856228c0572f7b64f14916E2F091391")
	ar, _ := hex.DecodeString("F235B58DC3b2169136A857B06aaedcE1aEC4c667")
	ep, _ := hex.DecodeString("8024A70A99d35FF24Cacc861e946945530ee96A3")
	salt, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000001")
	creationCode, _ := hex.DecodeString(TestAccountContractCreationCode)

	initCode := tbsdk.AbiEncodePacked(creationCode, tbsdk.AbiEncode(
		owner,
		ar,
		ep,
		factory,
	))

	address := tbsdk.ComputeCounterfactualAddress(factory, salt, initCode)

	// https://mumbai.polygonscan.com/address/0xd40aeab1d9e7c57523c2f5381f79c9738a73fe2d#internaltx
	deployedAddress := "0xd40aeab1d9e7c57523c2f5381f79c9738a73fe2d"

	if address != deployedAddress {
		t.Log(address)
		t.FailNow()
	}
}
