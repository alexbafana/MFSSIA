// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMfssiaServiceManager

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

// ISignatureUtilsSignatureWithSaltAndExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithSaltAndExpiry struct {
	Signature []byte
	Salt      [32]byte
	Expiry    *big.Int
}

// ContractMfssiaServiceManagerMetaData contains all meta data concerning the ContractMfssiaServiceManager contract.
var ContractMfssiaServiceManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"_delegationManager\",\"type\":\"address\"},{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"contractIStakeRegistry\",\"name\":\"_stakeRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIMfssiaTaskManager\",\"name\":\"_mfssiaTaskManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"deregisterOperatorFromAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operatorAddr\",\"type\":\"address\"}],\"name\":\"freezeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getOperatorRestakedStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakeableStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mfssiaTaskManager\",\"outputs\":[{\"internalType\":\"contractIMfssiaTaskManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structISignatureUtils.SignatureWithSaltAndExpiry\",\"name\":\"operatorSignature\",\"type\":\"tuple\"}],\"name\":\"registerOperatorToAVS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"setMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620016cd380380620016cd83398101604081905262000035916200014f565b6001600160a01b0380851660a052808416608052821660c0528383836200005b62000074565b5050506001600160a01b031660e05250620001b7915050565b600054610100900460ff1615620000e15760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b60005460ff908116101562000134576000805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b03811681146200014c57600080fd5b50565b600080600080608085870312156200016657600080fd5b8451620001738162000136565b6020860151909450620001868162000136565b6040860151909350620001998162000136565b6060860151909250620001ac8162000136565b939692955090935050565b60805160a05160c05160e05161147a620002536000396000818161015701526106820152600081816103a1015281816104fd0152818161059401528181610af701528181610c7b0152610d1a0152600081816107500152818161081901526108ed0152600081816101cc0152818161025b015281816102db015281816107c50152818161089101528181610a350152610bd6015261147a6000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639926ee7d116100715780639926ee7d1461012c578063a364f4da1461013f578063a9fefd4014610152578063c4d66de814610179578063e481af9d1461018c578063f2fde38b1461019457600080fd5b806333cfb7b7146100ae57806338c8ee64146100d7578063715018a6146100ec578063750521f5146100f45780638da5cb5b14610107575b600080fd5b6100c16100bc366004610f8d565b6101a7565b6040516100ce9190610fb1565b60405180910390f35b6100ea6100e5366004610f8d565b610677565b005b6100ea61071d565b6100ea6101023660046110b3565b610731565b6033546001600160a01b03165b6040516001600160a01b0390911681526020016100ce565b6100ea61013a366004611104565b6107ba565b6100ea61014d366004610f8d565b610886565b6101147f000000000000000000000000000000000000000000000000000000000000000081565b6100ea610187366004610f8d565b61091c565b6100c1610a2f565b6100ea6101a2366004610f8d565b610df9565b6040516309aa152760e11b81526001600160a01b0382811660048301526060916000917f000000000000000000000000000000000000000000000000000000000000000016906313542a4e90602401602060405180830381865afa158015610213573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023791906111af565b60405163871ef04960e01b8152600481018290529091506000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063871ef04990602401602060405180830381865afa1580156102a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c691906111c8565b90506001600160c01b038116158061036057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610337573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035b91906111f1565b60ff16155b1561037c57505060408051600081526020810190915292915050565b6000610390826001600160c01b0316610e6f565b90506000805b8251811015610466577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ca5a5f58483815181106103e0576103e0611214565b01602001516040516001600160e01b031960e084901b16815260f89190911c6004820152602401602060405180830381865afa158015610424573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061044891906111af565b6104529083611240565b91508061045e81611258565b915050610396565b5060008167ffffffffffffffff81111561048257610482610ffe565b6040519080825280602002602001820160405280156104ab578160200160208202803683370190505b5090506000805b845181101561066a5760008582815181106104cf576104cf611214565b0160200151604051633ca5a5f560e01b815260f89190911c6004820181905291506000906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690633ca5a5f590602401602060405180830381865afa158015610544573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061056891906111af565b905060005b81811015610654576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa1580156105e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106069190611273565b6000015186868151811061061c5761061c611214565b6001600160a01b03909216602092830291909101909101528461063e81611258565b955050808061064c90611258565b91505061056d565b505050808061066290611258565b9150506104b2565b5090979650505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461071a5760405162461bcd60e51b815260206004820152603c60248201527f6f6e6c794d66737369615461736b4d616e616765723a206e6f742066726f6d2060448201527f6372656469626c65206d6673736961207461736b206d616e616765720000000060648201526084015b60405180910390fd5b50565b610725610ecc565b61072f6000610f26565b565b610739610ecc565b60405163a98fb35560e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063a98fb3559061078590849060040161133f565b600060405180830381600087803b15801561079f57600080fd5b505af11580156107b3573d6000803e3d6000fd5b5050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108025760405162461bcd60e51b815260040161071190611352565b604051639926ee7d60e01b81526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639926ee7d9061085090859085906004016113ca565b600060405180830381600087803b15801561086a57600080fd5b505af115801561087e573d6000803e3d6000fd5b505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108ce5760405162461bcd60e51b815260040161071190611352565b6040516351b27a6d60e11b81526001600160a01b0382811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063a364f4da90602401610785565b600054610100900460ff161580801561093c5750600054600160ff909116105b806109565750303b158015610956575060005460ff166001145b6109b95760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610711565b6000805460ff1916600117905580156109dc576000805461ff0019166101001790555b6109e582610f26565b8015610a2b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab591906111f1565b60ff16905080610ad357505060408051600081526020810190915290565b6000805b82811015610b8857604051633ca5a5f560e01b815260ff821660048201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610b46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b6a91906111af565b610b749083611240565b915080610b8081611258565b915050610ad7565b5060008167ffffffffffffffff811115610ba457610ba4610ffe565b604051908082528060200260200182016040528015610bcd578160200160208202803683370190505b5090506000805b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316639aa1653d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c32573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5691906111f1565b60ff16811015610def57604051633ca5a5f560e01b815260ff821660048201526000907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690633ca5a5f590602401602060405180830381865afa158015610cca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cee91906111af565b905060005b81811015610dda576040516356e4026d60e11b815260ff84166004820152602481018290527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063adc804da906044016040805180830381865afa158015610d68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d8c9190611273565b60000151858581518110610da257610da2611214565b6001600160a01b039092166020928302919091019091015283610dc481611258565b9450508080610dd290611258565b915050610cf3565b50508080610de790611258565b915050610bd4565b5090949350505050565b610e01610ecc565b6001600160a01b038116610e665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610711565b61071a81610f26565b60606000805b610100811015610ec5576001811b915083821615610eb557828160f81b604051602001610ea3929190611415565b60405160208183030381529060405292505b610ebe81611258565b9050610e75565b5050919050565b6033546001600160a01b0316331461072f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610711565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6001600160a01b038116811461071a57600080fd5b600060208284031215610f9f57600080fd5b8135610faa81610f78565b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610ff25783516001600160a01b031683529284019291840191600101610fcd565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6040516060810167ffffffffffffffff8111828210171561103757611037610ffe565b60405290565b600067ffffffffffffffff8084111561105857611058610ffe565b604051601f8501601f19908116603f0116810190828211818310171561108057611080610ffe565b8160405280935085815286868601111561109957600080fd5b858560208301376000602087830101525050509392505050565b6000602082840312156110c557600080fd5b813567ffffffffffffffff8111156110dc57600080fd5b8201601f810184136110ed57600080fd5b6110fc8482356020840161103d565b949350505050565b6000806040838503121561111757600080fd5b823561112281610f78565b9150602083013567ffffffffffffffff8082111561113f57600080fd5b908401906060828703121561115357600080fd5b61115b611014565b82358281111561116a57600080fd5b83019150601f8201871361117d57600080fd5b61118c8783356020850161103d565b815260208301356020820152604083013560408201528093505050509250929050565b6000602082840312156111c157600080fd5b5051919050565b6000602082840312156111da57600080fd5b81516001600160c01b0381168114610faa57600080fd5b60006020828403121561120357600080fd5b815160ff81168114610faa57600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082198211156112535761125361122a565b500190565b600060001982141561126c5761126c61122a565b5060010190565b60006040828403121561128557600080fd5b6040516040810181811067ffffffffffffffff821117156112a8576112a8610ffe565b60405282516112b681610f78565b815260208301516bffffffffffffffffffffffff811681146112d757600080fd5b60208201529392505050565b60005b838110156112fe5781810151838201526020016112e6565b8381111561130d576000848401525b50505050565b6000815180845261132b8160208601602086016112e3565b601f01601f19169290920160200192915050565b602081526000610faa6020830184611313565b60208082526052908201527f536572766963654d616e61676572426173652e6f6e6c7952656769737472794360408201527f6f6f7264696e61746f723a2063616c6c6572206973206e6f742074686520726560608201527133b4b9ba393c9031b7b7b93234b730ba37b960711b608082015260a00190565b60018060a01b03831681526040602082015260008251606060408401526113f460a0840182611313565b90506020840151606084015260408401516080840152809150509392505050565b600083516114278184602088016112e3565b6001600160f81b031993909316919092019081526001019291505056fea2646970667358221220ce236533051379db2b8247884457d6940a2a742712ce342261912b761ca6d24664736f6c634300080c0033",
}

// ContractMfssiaServiceManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMfssiaServiceManagerMetaData.ABI instead.
var ContractMfssiaServiceManagerABI = ContractMfssiaServiceManagerMetaData.ABI

// ContractMfssiaServiceManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMfssiaServiceManagerMetaData.Bin instead.
var ContractMfssiaServiceManagerBin = ContractMfssiaServiceManagerMetaData.Bin

// DeployContractMfssiaServiceManager deploys a new Ethereum contract, binding an instance of ContractMfssiaServiceManager to it.
func DeployContractMfssiaServiceManager(auth *bind.TransactOpts, backend bind.ContractBackend, _delegationManager common.Address, _registryCoordinator common.Address, _stakeRegistry common.Address, _mfssiaTaskManager common.Address) (common.Address, *types.Transaction, *ContractMfssiaServiceManager, error) {
	parsed, err := ContractMfssiaServiceManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractMfssiaServiceManagerBin), backend, _delegationManager, _registryCoordinator, _stakeRegistry, _mfssiaTaskManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractMfssiaServiceManager{ContractMfssiaServiceManagerCaller: ContractMfssiaServiceManagerCaller{contract: contract}, ContractMfssiaServiceManagerTransactor: ContractMfssiaServiceManagerTransactor{contract: contract}, ContractMfssiaServiceManagerFilterer: ContractMfssiaServiceManagerFilterer{contract: contract}}, nil
}

// ContractMfssiaServiceManager is an auto generated Go binding around an Ethereum contract.
type ContractMfssiaServiceManager struct {
	ContractMfssiaServiceManagerCaller     // Read-only binding to the contract
	ContractMfssiaServiceManagerTransactor // Write-only binding to the contract
	ContractMfssiaServiceManagerFilterer   // Log filterer for contract events
}

// ContractMfssiaServiceManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractMfssiaServiceManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaServiceManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractMfssiaServiceManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaServiceManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractMfssiaServiceManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaServiceManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractMfssiaServiceManagerSession struct {
	Contract     *ContractMfssiaServiceManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractMfssiaServiceManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractMfssiaServiceManagerCallerSession struct {
	Contract *ContractMfssiaServiceManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// ContractMfssiaServiceManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractMfssiaServiceManagerTransactorSession struct {
	Contract     *ContractMfssiaServiceManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// ContractMfssiaServiceManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractMfssiaServiceManagerRaw struct {
	Contract *ContractMfssiaServiceManager // Generic contract binding to access the raw methods on
}

// ContractMfssiaServiceManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractMfssiaServiceManagerCallerRaw struct {
	Contract *ContractMfssiaServiceManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractMfssiaServiceManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractMfssiaServiceManagerTransactorRaw struct {
	Contract *ContractMfssiaServiceManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractMfssiaServiceManager creates a new instance of ContractMfssiaServiceManager, bound to a specific deployed contract.
func NewContractMfssiaServiceManager(address common.Address, backend bind.ContractBackend) (*ContractMfssiaServiceManager, error) {
	contract, err := bindContractMfssiaServiceManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaServiceManager{ContractMfssiaServiceManagerCaller: ContractMfssiaServiceManagerCaller{contract: contract}, ContractMfssiaServiceManagerTransactor: ContractMfssiaServiceManagerTransactor{contract: contract}, ContractMfssiaServiceManagerFilterer: ContractMfssiaServiceManagerFilterer{contract: contract}}, nil
}

// NewContractMfssiaServiceManagerCaller creates a new read-only instance of ContractMfssiaServiceManager, bound to a specific deployed contract.
func NewContractMfssiaServiceManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractMfssiaServiceManagerCaller, error) {
	contract, err := bindContractMfssiaServiceManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaServiceManagerCaller{contract: contract}, nil
}

// NewContractMfssiaServiceManagerTransactor creates a new write-only instance of ContractMfssiaServiceManager, bound to a specific deployed contract.
func NewContractMfssiaServiceManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractMfssiaServiceManagerTransactor, error) {
	contract, err := bindContractMfssiaServiceManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaServiceManagerTransactor{contract: contract}, nil
}

// NewContractMfssiaServiceManagerFilterer creates a new log filterer instance of ContractMfssiaServiceManager, bound to a specific deployed contract.
func NewContractMfssiaServiceManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractMfssiaServiceManagerFilterer, error) {
	contract, err := bindContractMfssiaServiceManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaServiceManagerFilterer{contract: contract}, nil
}

// bindContractMfssiaServiceManager binds a generic wrapper to an already deployed contract.
func bindContractMfssiaServiceManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMfssiaServiceManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMfssiaServiceManager.Contract.ContractMfssiaServiceManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.ContractMfssiaServiceManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.ContractMfssiaServiceManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMfssiaServiceManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCaller) GetOperatorRestakedStrategies(opts *bind.CallOpts, operator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaServiceManager.contract.Call(opts, &out, "getOperatorRestakedStrategies", operator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractMfssiaServiceManager.CallOpts, operator)
}

// GetOperatorRestakedStrategies is a free data retrieval call binding the contract method 0x33cfb7b7.
//
// Solidity: function getOperatorRestakedStrategies(address operator) view returns(address[])
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCallerSession) GetOperatorRestakedStrategies(operator common.Address) ([]common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.GetOperatorRestakedStrategies(&_ContractMfssiaServiceManager.CallOpts, operator)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCaller) GetRestakeableStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaServiceManager.contract.Call(opts, &out, "getRestakeableStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.GetRestakeableStrategies(&_ContractMfssiaServiceManager.CallOpts)
}

// GetRestakeableStrategies is a free data retrieval call binding the contract method 0xe481af9d.
//
// Solidity: function getRestakeableStrategies() view returns(address[])
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCallerSession) GetRestakeableStrategies() ([]common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.GetRestakeableStrategies(&_ContractMfssiaServiceManager.CallOpts)
}

// MfssiaTaskManager is a free data retrieval call binding the contract method 0xa9fefd40.
//
// Solidity: function mfssiaTaskManager() view returns(address)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCaller) MfssiaTaskManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaServiceManager.contract.Call(opts, &out, "mfssiaTaskManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MfssiaTaskManager is a free data retrieval call binding the contract method 0xa9fefd40.
//
// Solidity: function mfssiaTaskManager() view returns(address)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) MfssiaTaskManager() (common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.MfssiaTaskManager(&_ContractMfssiaServiceManager.CallOpts)
}

// MfssiaTaskManager is a free data retrieval call binding the contract method 0xa9fefd40.
//
// Solidity: function mfssiaTaskManager() view returns(address)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCallerSession) MfssiaTaskManager() (common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.MfssiaTaskManager(&_ContractMfssiaServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaServiceManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) Owner() (common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.Owner(&_ContractMfssiaServiceManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerCallerSession) Owner() (common.Address, error) {
	return _ContractMfssiaServiceManager.Contract.Owner(&_ContractMfssiaServiceManager.CallOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.contract.Transact(opts, "deregisterOperatorFromAVS", operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractMfssiaServiceManager.TransactOpts, operator)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xa364f4da.
//
// Solidity: function deregisterOperatorFromAVS(address operator) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorSession) DeregisterOperatorFromAVS(operator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.DeregisterOperatorFromAVS(&_ContractMfssiaServiceManager.TransactOpts, operator)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactor) FreezeOperator(opts *bind.TransactOpts, operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.contract.Transact(opts, "freezeOperator", operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.FreezeOperator(&_ContractMfssiaServiceManager.TransactOpts, operatorAddr)
}

// FreezeOperator is a paid mutator transaction binding the contract method 0x38c8ee64.
//
// Solidity: function freezeOperator(address operatorAddr) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorSession) FreezeOperator(operatorAddr common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.FreezeOperator(&_ContractMfssiaServiceManager.TransactOpts, operatorAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.Initialize(&_ContractMfssiaServiceManager.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.Initialize(&_ContractMfssiaServiceManager.TransactOpts, initialOwner)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts, operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.contract.Transact(opts, "registerOperatorToAVS", operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.RegisterOperatorToAVS(&_ContractMfssiaServiceManager.TransactOpts, operator, operatorSignature)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0x9926ee7d.
//
// Solidity: function registerOperatorToAVS(address operator, (bytes,bytes32,uint256) operatorSignature) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorSession) RegisterOperatorToAVS(operator common.Address, operatorSignature ISignatureUtilsSignatureWithSaltAndExpiry) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.RegisterOperatorToAVS(&_ContractMfssiaServiceManager.TransactOpts, operator, operatorSignature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.RenounceOwnership(&_ContractMfssiaServiceManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.RenounceOwnership(&_ContractMfssiaServiceManager.TransactOpts)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactor) SetMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.contract.Transact(opts, "setMetadataURI", _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.SetMetadataURI(&_ContractMfssiaServiceManager.TransactOpts, _metadataURI)
}

// SetMetadataURI is a paid mutator transaction binding the contract method 0x750521f5.
//
// Solidity: function setMetadataURI(string _metadataURI) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorSession) SetMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.SetMetadataURI(&_ContractMfssiaServiceManager.TransactOpts, _metadataURI)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.TransferOwnership(&_ContractMfssiaServiceManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaServiceManager.Contract.TransferOwnership(&_ContractMfssiaServiceManager.TransactOpts, newOwner)
}

// ContractMfssiaServiceManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractMfssiaServiceManager contract.
type ContractMfssiaServiceManagerInitializedIterator struct {
	Event *ContractMfssiaServiceManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaServiceManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaServiceManagerInitialized)
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
		it.Event = new(ContractMfssiaServiceManagerInitialized)
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
func (it *ContractMfssiaServiceManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaServiceManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaServiceManagerInitialized represents a Initialized event raised by the ContractMfssiaServiceManager contract.
type ContractMfssiaServiceManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractMfssiaServiceManagerInitializedIterator, error) {

	logs, sub, err := _ContractMfssiaServiceManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaServiceManagerInitializedIterator{contract: _ContractMfssiaServiceManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractMfssiaServiceManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaServiceManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaServiceManagerInitialized)
				if err := _ContractMfssiaServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerFilterer) ParseInitialized(log types.Log) (*ContractMfssiaServiceManagerInitialized, error) {
	event := new(ContractMfssiaServiceManagerInitialized)
	if err := _ContractMfssiaServiceManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaServiceManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractMfssiaServiceManager contract.
type ContractMfssiaServiceManagerOwnershipTransferredIterator struct {
	Event *ContractMfssiaServiceManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaServiceManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaServiceManagerOwnershipTransferred)
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
		it.Event = new(ContractMfssiaServiceManagerOwnershipTransferred)
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
func (it *ContractMfssiaServiceManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaServiceManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaServiceManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractMfssiaServiceManager contract.
type ContractMfssiaServiceManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractMfssiaServiceManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMfssiaServiceManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaServiceManagerOwnershipTransferredIterator{contract: _ContractMfssiaServiceManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractMfssiaServiceManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMfssiaServiceManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaServiceManagerOwnershipTransferred)
				if err := _ContractMfssiaServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractMfssiaServiceManager *ContractMfssiaServiceManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractMfssiaServiceManagerOwnershipTransferred, error) {
	event := new(ContractMfssiaServiceManagerOwnershipTransferred)
	if err := _ContractMfssiaServiceManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
