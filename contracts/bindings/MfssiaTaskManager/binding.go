// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractMfssiaTaskManager

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSSignatureCheckerNonSignerStakesAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerNonSignerStakesAndSignature struct {
	NonSignerQuorumBitmapIndices []uint32
	NonSignerPubkeys             []BN254G1Point
	QuorumApks                   []BN254G1Point
	ApkG2                        BN254G2Point
	Sigma                        BN254G1Point
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// IBLSSignatureCheckerQuorumStakeTotals is an auto generated low-level Go binding around an user-defined struct.
type IBLSSignatureCheckerQuorumStakeTotals struct {
	SignedStakeForQuorum []*big.Int
	TotalStakeForQuorum  []*big.Int
}

// IMfssiaTaskManagerTask is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTask struct {
	TaskCreatedBlock          uint32
	QuorumNumbers             []byte
	System1Value              string
	System2Value              string
	DkgValue                  string
	QuorumThresholdPercentage uint32
}

// IMfssiaTaskManagerTaskResponse is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTaskResponse struct {
	ReferenceTaskIndex uint32
	FailedSystem       string
}

// IMfssiaTaskManagerTaskResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IMfssiaTaskManagerTaskResponseMetadata struct {
	TaskResponsedBlock uint32
	HashOfNonSigners   [32]byte
}

// OperatorStateRetrieverCheckSignaturesIndices is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverCheckSignaturesIndices struct {
	NonSignerQuorumBitmapIndices []uint32
	QuorumApkIndices             []uint32
	TotalStakeIndices            []uint32
	NonSignerStakeIndices        [][]uint32
}

// OperatorStateRetrieverOperator is an auto generated low-level Go binding around an user-defined struct.
type OperatorStateRetrieverOperator struct {
	OperatorId [32]byte
	Stake      *big.Int
}

// ContractMfssiaTaskManagerMetaData contains all meta data concerning the ContractMfssiaTaskManager contract.
var ContractMfssiaTaskManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"_registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_taskResponseWindowBlock\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"system1Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"system2Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dkgValue\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"}],\"name\":\"NewTaskCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"StaleStakesForbiddenUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedSuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"}],\"name\":\"TaskChallengedUnsuccessfully\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"taskIndex\",\"type\":\"uint32\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"failedSystem\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskResponsedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIMfssiaTaskManager.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"}],\"name\":\"TaskResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TASK_CHALLENGE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TASK_RESPONSE_WINDOW_BLOCK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"aggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allTaskHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"allTaskResponses\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blsApkRegistry\",\"outputs\":[{\"internalType\":\"contractIBLSApkRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"checkSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96[]\",\"name\":\"signedStakeForQuorum\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96[]\",\"name\":\"totalStakeForQuorum\",\"type\":\"uint96[]\"}],\"internalType\":\"structIBLSSignatureChecker.QuorumStakeTotals\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"system1Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"system2Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dkgValue\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"}],\"name\":\"createNewTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegation\",\"outputs\":[{\"internalType\":\"contractIDelegationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"referenceBlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"nonSignerOperatorIds\",\"type\":\"bytes32[]\"}],\"name\":\"getCheckSignaturesIndices\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structOperatorStateRetriever.CheckSignaturesIndices\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"registryCoordinator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"}],\"name\":\"getOperatorState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"operatorId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"stake\",\"type\":\"uint96\"}],\"internalType\":\"structOperatorStateRetriever.Operator[][]\",\"name\":\"\",\"type\":\"tuple[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTaskResponseWindowBlock\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_aggregator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_generator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTaskNum\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"system1Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"system2Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dkgValue\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"internalType\":\"structIMfssiaTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"failedSystem\",\"type\":\"string\"}],\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskResponsedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\"}],\"internalType\":\"structIMfssiaTaskManager.TaskResponseMetadata\",\"name\":\"taskResponseMetadata\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"pubkeysOfNonSigningOperators\",\"type\":\"tuple[]\"}],\"name\":\"raiseAndResolveChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryCoordinator\",\"outputs\":[{\"internalType\":\"contractIRegistryCoordinator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"taskCreatedBlock\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"quorumNumbers\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"system1Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"system2Value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dkgValue\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\"}],\"internalType\":\"structIMfssiaTaskManager.Task\",\"name\":\"task\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"referenceTaskIndex\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"failedSystem\",\"type\":\"string\"}],\"internalType\":\"structIMfssiaTaskManager.TaskResponse\",\"name\":\"taskResponse\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32[]\",\"name\":\"nonSignerQuorumBitmapIndices\",\"type\":\"uint32[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point[]\",\"name\":\"quorumApks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"},{\"internalType\":\"uint32[]\",\"name\":\"quorumApkIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[]\",\"name\":\"totalStakeIndices\",\"type\":\"uint32[]\"},{\"internalType\":\"uint32[][]\",\"name\":\"nonSignerStakeIndices\",\"type\":\"uint32[][]\"}],\"internalType\":\"structIBLSSignatureChecker.NonSignerStakesAndSignature\",\"name\":\"nonSignerStakesAndSignature\",\"type\":\"tuple\"}],\"name\":\"respondToTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setStaleStakesForbidden\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegistry\",\"outputs\":[{\"internalType\":\"contractIStakeRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staleStakesForbidden\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taskNumber\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"taskSuccesfullyChallenged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"msgHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"apk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"X\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"Y\",\"type\":\"uint256[2]\"}],\"internalType\":\"structBN254.G2Point\",\"name\":\"apkG2\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"X\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Y\",\"type\":\"uint256\"}],\"internalType\":\"structBN254.G1Point\",\"name\":\"sigma\",\"type\":\"tuple\"}],\"name\":\"trySignatureAndApkVerification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pairingSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"siganatureIsValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5060405162005b8738038062005b878339810160408190526200003591620001f7565b81806001600160a01b03166080816001600160a01b031681525050806001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200008f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000b591906200023e565b6001600160a01b031660a0816001600160a01b031681525050806001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200010d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200013391906200023e565b6001600160a01b031660c0816001600160a01b03168152505060a0516001600160a01b031663df5cf7236040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200018d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001b391906200023e565b6001600160a01b031660e052506097805460ff1916600117905563ffffffff16610100525062000265565b6001600160a01b0381168114620001f457600080fd5b50565b600080604083850312156200020b57600080fd5b82516200021881620001de565b602084015190925063ffffffff811681146200023357600080fd5b809150509250929050565b6000602082840312156200025157600080fd5b81516200025e81620001de565b9392505050565b60805160a05160c05160e05161010051615890620002f7600039600081816102720152818161057b01526128a70152600081816105440152611bff0152600081816103e001528181611de90152612f4f01526000818161040701528181611fbf015261218101526000818161042e01528181610ccd015281816118dd01528181611a750152611ca301526158906000f3fe608060405234801561001057600080fd5b50600436106102115760003560e01c80636efb463611610125578063b31c72c6116100ad578063f2fde38b1161007c578063f2fde38b14610566578063f5c9899d14610579578063f63c5bab1461059f578063f8c8765e146105a7578063fabc1cbc146105ba57600080fd5b8063b31c72c6146104fe578063b98d090814610511578063cefdc1d41461051e578063df5cf7231461053f57600080fd5b8063886f1195116100f4578063886f11951461049f5780638b00ce7c146104b25780638da5cb5b146104c7578063a11c1a5d146104d8578063aedde731146104eb57600080fd5b80636efb463614610450578063715018a61461047157806372d18e8d146104795780637afa1eed1461048c57600080fd5b8063416c7e5e116101a85780635c975abb116101775780635c975abb146103b05780635decc3f5146103b85780635df45946146103db57806368304835146104025780636d14a9871461042957600080fd5b8063416c7e5e146103425780634f739f7414610355578063595c6a67146103755780635ac86ab71461037d57600080fd5b8063245a7bfc116101e4578063245a7bfc146102a95780632cb223d5146102d45780632d89f6fc146103025780633563b0d11461032257600080fd5b806310d67a2f14610216578063136439dd1461022b578063171f1d5b1461023e5780631ad431891461026d575b600080fd5b6102296102243660046142cd565b6105cd565b005b6102296102393660046142ea565b610689565b61025161024c366004614468565b6107c8565b6040805192151583529015156020830152015b60405180910390f35b6102947f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610264565b609b546102bc906001600160a01b031681565b6040516001600160a01b039091168152602001610264565b6102f46102e23660046144d6565b60996020526000908152604090205481565b604051908152602001610264565b6102f46103103660046144d6565b60986020526000908152604090205481565b6103356103303660046144f3565b610952565b604051610264919061463a565b610229610350366004614662565b610ccb565b6103686103633660046146c7565b610e40565b60405161026491906147cb565b6102296114c4565b6103a061038b366004614895565b606654600160ff9092169190911b9081161490565b6040519015158152602001610264565b6066546102f4565b6103a06103c63660046144d6565b609a6020526000908152604090205460ff1681565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b61046361045e366004614b4b565b61158b565b604051610264929190614c0b565b610229612436565b609754610100900463ffffffff16610294565b609c546102bc906001600160a01b031681565b6065546102bc906001600160a01b031681565b60975461029490610100900463ffffffff1681565b6033546001600160a01b03166102bc565b6102296104e6366004614c54565b61244a565b6102296104f9366004614d58565b6126b9565b61022961050c366004614ddf565b612b35565b6097546103a09060ff1681565b61053161052c366004614e78565b613084565b604051610264929190614eba565b6102bc7f000000000000000000000000000000000000000000000000000000000000000081565b6102296105743660046142cd565b613216565b7f0000000000000000000000000000000000000000000000000000000000000000610294565b610294606481565b6102296105b5366004614edb565b61328c565b6102296105c83660046142ea565b6133dd565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610620573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106449190614f37565b6001600160a01b0316336001600160a01b03161461067d5760405162461bcd60e51b815260040161067490614f54565b60405180910390fd5b61068681613539565b50565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa1580156106d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f59190614f9e565b6107115760405162461bcd60e51b815260040161067490614fbb565b6066548181161461078a5760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e70617573653a20696e76616c696420617474656d70742060448201527f746f20756e70617573652066756e6374696f6e616c69747900000000000000006064820152608401610674565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d906020015b60405180910390a250565b60008060007f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018787600001518860200151886000015160006002811061081057610810615003565b60200201518951600160200201518a6020015160006002811061083557610835615003565b60200201518b6020015160016002811061085157610851615003565b602090810291909101518c518d8301516040516108ae9a99989796959401988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012060001c6108d19190615019565b90506109446108ea6108e38884613630565b86906136c7565b6108f261375b565b61093a61092b85610925604080518082018252600080825260209182015281518083019092526001825260029082015290565b90613630565b6109348c61381b565b906136c7565b886201d4c06138ab565b909890975095505050505050565b60606000846001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610994573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109b89190614f37565b90506000856001600160a01b0316639e9923c26040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109fa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1e9190614f37565b9050600085516001600160401b03811115610a3b57610a3b614303565b604051908082528060200260200182016040528015610a6e57816020015b6060815260200190600190039081610a595790505b50905060005b8651811015610cc0576000878281518110610a9157610a91615003565b0160200151604051638902624560e01b815260f89190911c6004820181905263ffffffff8916602483015291506000906001600160a01b03861690638902624590604401600060405180830381865afa158015610af2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b1a919081019061503b565b905080516001600160401b03811115610b3557610b35614303565b604051908082528060200260200182016040528015610b7a57816020015b6040805180820190915260008082526020820152815260200190600190039081610b535790505b50848481518110610b8d57610b8d615003565b602002602001018190525060005b8151811015610caa576000828281518110610bb857610bb8615003565b602090810291909101810151604080518082018252828152905163fa28c62760e01b81526004810183905260ff8816602482015263ffffffff8e166044820152919350918201906001600160a01b038b169063fa28c62790606401602060405180830381865afa158015610c30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5491906150cb565b6001600160601b0316815250868681518110610c7257610c72615003565b60200260200101518381518110610c8b57610c8b615003565b6020026020010181905250508080610ca29061510a565b915050610b9b565b5050508080610cb89061510a565b915050610a74565b509695505050505050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d4d9190614f37565b6001600160a01b0316336001600160a01b031614610df95760405162461bcd60e51b815260206004820152605c60248201527f424c535369676e6174757265436865636b65722e6f6e6c79436f6f7264696e6160448201527f746f724f776e65723a2063616c6c6572206973206e6f7420746865206f776e6560648201527f72206f6620746865207265676973747279436f6f7264696e61746f7200000000608482015260a401610674565b6097805460ff19168215159081179091556040519081527f40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc9060200160405180910390a150565b610e6b6040518060800160405280606081526020016060815260200160608152602001606081525090565b6000876001600160a01b031663683048356040518163ffffffff1660e01b8152600401602060405180830381865afa158015610eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ecf9190614f37565b9050610efc6040518060800160405280606081526020016060815260200160608152602001606081525090565b6040516361c8a12f60e11b81526001600160a01b038a169063c391425e90610f2c908b9089908990600401615125565b600060405180830381865afa158015610f49573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f71919081019061516f565b81526040516340e03a8160e11b81526001600160a01b038316906381c0750290610fa3908b908b908b90600401615226565b600060405180830381865afa158015610fc0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fe8919081019061516f565b6040820152856001600160401b0381111561100557611005614303565b60405190808252806020026020018201604052801561103857816020015b60608152602001906001900390816110235790505b50606082015260005b60ff81168711156113d5576000856001600160401b0381111561106657611066614303565b60405190808252806020026020018201604052801561108f578160200160208202803683370190505b5083606001518360ff16815181106110a9576110a9615003565b602002602001018190525060005b868110156112d55760008c6001600160a01b03166304ec63518a8a858181106110e2576110e2615003565b905060200201358e8860000151868151811061110057611100615003565b60200260200101516040518463ffffffff1660e01b815260040161113d9392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa15801561115a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061117e919061524f565b90508a8a8560ff1681811061119557611195615003565b6001600160c01b03841692013560f81c9190911c6001908116141590506112c257856001600160a01b031663dd9846b98a8a858181106111d7576111d7615003565b905060200201358d8d8860ff168181106111f3576111f3615003565b6040516001600160e01b031960e087901b1681526004810194909452919091013560f81c60248301525063ffffffff8f166044820152606401602060405180830381865afa158015611249573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126d9190615278565b85606001518560ff168151811061128657611286615003565b6020026020010151848151811061129f5761129f615003565b63ffffffff90921660209283029190910190910152826112be8161510a565b9350505b50806112cd8161510a565b9150506110b7565b506000816001600160401b038111156112f0576112f0614303565b604051908082528060200260200182016040528015611319578160200160208202803683370190505b50905060005b8281101561139a5784606001518460ff168151811061134057611340615003565b6020026020010151818151811061135957611359615003565b602002602001015182828151811061137357611373615003565b63ffffffff90921660209283029190910190910152806113928161510a565b91505061131f565b508084606001518460ff16815181106113b5576113b5615003565b6020026020010181905250505080806113cd90615295565b915050611041565b506000896001600160a01b0316635df459466040518163ffffffff1660e01b8152600401602060405180830381865afa158015611416573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061143a9190614f37565b60405163354952a360e21b81529091506001600160a01b0382169063d5254a8c9061146d908b908b908e906004016152b5565b600060405180830381865afa15801561148a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526114b2919081019061516f565b60208301525098975050505050505050565b60655460405163237dfb4760e11b81523360048201526001600160a01b03909116906346fbf68e90602401602060405180830381865afa15801561150c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115309190614f9e565b61154c5760405162461bcd60e51b815260040161067490614fbb565b600019606681905560405190815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2565b604080518082018252606080825260208201529082015151600090851480156115b8575060a08301515185145b80156115c8575060c08301515185145b80156115d8575060e08301515185145b6116425760405162461bcd60e51b8152602060048201526041602482015260008051602061583b83398151915260448201527f7265733a20696e7075742071756f72756d206c656e677468206d69736d6174636064820152600d60fb1b608482015260a401610674565b825151602084015151146116ba5760405162461bcd60e51b81526020600482015260446024820181905260008051602061583b833981519152908201527f7265733a20696e707574206e6f6e7369676e6572206c656e677468206d69736d6064820152630c2e8c6d60e31b608482015260a401610674565b4363ffffffff168463ffffffff16111561172a5760405162461bcd60e51b815260206004820152603c602482015260008051602061583b83398151915260448201527f7265733a20696e76616c6964207265666572656e636520626c6f636b000000006064820152608401610674565b6040805180820182526000808252602080830191909152825180840190935260608084529083015290866001600160401b0381111561176b5761176b614303565b604051908082528060200260200182016040528015611794578160200160208202803683370190505b506020820152866001600160401b038111156117b2576117b2614303565b6040519080825280602002602001820160405280156117db578160200160208202803683370190505b50815260408051808201909152606080825260208201528560200151516001600160401b0381111561180f5761180f614303565b604051908082528060200260200182016040528015611838578160200160208202803683370190505b5081526020860151516001600160401b0381111561185857611858614303565b604051908082528060200260200182016040528015611881578160200160208202803683370190505b50816020018190525060006119538a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505060408051639aa1653d60e01b815290516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169350639aa1653d925060048083019260209291908290030181865afa15801561192a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061194e91906152df565b613acf565b905060005b876020015151811015611bee5761199d8860200151828151811061197e5761197e615003565b6020026020010151805160009081526020918201519091526040902090565b836020015182815181106119b3576119b3615003565b60209081029190910101528015611a735760208301516119d46001836152fc565b815181106119e4576119e4615003565b602002602001015160001c83602001518281518110611a0557611a05615003565b602002602001015160001c11611a73576040805162461bcd60e51b815260206004820152602481019190915260008051602061583b83398151915260448201527f7265733a206e6f6e5369676e65725075626b657973206e6f7420736f727465646064820152608401610674565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166304ec635184602001518381518110611ab857611ab8615003565b60200260200101518b8b600001518581518110611ad757611ad7615003565b60200260200101516040518463ffffffff1660e01b8152600401611b149392919092835263ffffffff918216602084015216604082015260600190565b602060405180830381865afa158015611b31573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b55919061524f565b6001600160c01b031683600001518281518110611b7457611b74615003565b602002602001018181525050611bda6108e3611bae8486600001518581518110611ba057611ba0615003565b602002602001015116613b8a565b8a602001518481518110611bc457611bc4615003565b6020026020010151613bb590919063ffffffff16565b945080611be68161510a565b915050611958565b5050611bf983613c99565b925060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166350f73e7c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c5b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c7f9190615313565b60975490915060ff1660005b8a811015612305578115611de7578963ffffffff16837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663249a0c428f8f86818110611ce257611ce2615003565b60405160e085901b6001600160e01b031916815292013560f81c600483015250602401602060405180830381865afa158015611d22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d469190615313565b611d50919061532c565b1015611de75760405162461bcd60e51b8152602060048201526066602482015260008051602061583b83398151915260448201527f7265733a205374616b6552656769737472792075706461746573206d7573742060648201527f62652077697468696e207769746864726177616c44656c6179426c6f636b732060848201526577696e646f7760d01b60a482015260c401610674565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166368bccaac8d8d84818110611e2857611e28615003565b9050013560f81c60f81b60f81c8c8c60a001518581518110611e4c57611e4c615003565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa158015611ea8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ecc9190615344565b6001600160401b031916611eef8a60400151838151811061197e5761197e615003565b67ffffffffffffffff191614611f8b5760405162461bcd60e51b8152602060048201526061602482015260008051602061583b83398151915260448201527f7265733a2071756f72756d41706b206861736820696e2073746f72616765206460648201527f6f6573206e6f74206d617463682070726f76696465642071756f72756d2061706084820152606b60f81b60a482015260c401610674565b611fbb89604001518281518110611fa457611fa4615003565b6020026020010151876136c790919063ffffffff16565b95507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c8294c568d8d84818110611ffe57611ffe615003565b9050013560f81c60f81b60f81c8c8c60c00151858151811061202257612022615003565b60209081029190910101516040516001600160e01b031960e086901b16815260ff909316600484015263ffffffff9182166024840152166044820152606401602060405180830381865afa15801561207e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120a291906150cb565b856020015182815181106120b8576120b8615003565b6001600160601b039092166020928302919091018201528501518051829081106120e4576120e4615003565b60200260200101518560000151828151811061210257612102615003565b60200260200101906001600160601b031690816001600160601b0316815250506000805b8a60200151518110156122f05761217a8660000151828151811061214c5761214c615003565b60200260200101518f8f8681811061216657612166615003565b600192013560f81c9290921c811614919050565b156122de577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f2be94ae8f8f868181106121c0576121c0615003565b9050013560f81c60f81b60f81c8e896020015185815181106121e4576121e4615003565b60200260200101518f60e00151888151811061220257612202615003565b6020026020010151878151811061221b5761221b615003565b60209081029190910101516040516001600160e01b031960e087901b16815260ff909416600485015263ffffffff92831660248501526044840191909152166064820152608401602060405180830381865afa15801561227f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122a391906150cb565b87518051859081106122b7576122b7615003565b602002602001018181516122cb919061536f565b6001600160601b03169052506001909101905b806122e88161510a565b915050612126565b505080806122fd9061510a565b915050611c8b565b50505060008061231f8c868a606001518b608001516107c8565b91509150816123905760405162461bcd60e51b8152602060048201526043602482015260008051602061583b83398151915260448201527f7265733a2070616972696e6720707265636f6d70696c652063616c6c206661696064820152621b195960ea1b608482015260a401610674565b806123f15760405162461bcd60e51b8152602060048201526039602482015260008051602061583b83398151915260448201527f7265733a207369676e617475726520697320696e76616c6964000000000000006064820152608401610674565b5050600087826020015160405160200161240c929190615397565b60408051808303601f190181529190528051602090910120929b929a509198505050505050505050565b61243e613d34565b6124486000613d8e565b565b609c546001600160a01b031633146124ae5760405162461bcd60e51b815260206004820152602160248201527f5461736b2067656e657261746f72206d757374206265207468652063616c6c656044820152603960f91b6064820152608401610674565b6124f36040518060c00160405280600063ffffffff16815260200160608152602001606081526020016060815260200160608152602001600063ffffffff1681525090565b89898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060408083019190915280516020601f8a018190048102820181019092528881529089908990819084018382808284376000920191909152505050506060820152604080516020601f8801819004810282018101909252868152908790879081908401838280828437600092019190915250505050608082015263ffffffff4381168252841660a0820152604080516020601f85018190048102820181019092528381529084908490819084018382808284376000920191909152505050506020808301919091526040516125ff9183910161543b565b60408051601f1981840301815282825280516020918201206097805463ffffffff610100918290048116600090815260989095529490932091909155540416907f850fa919096ea51a0971805dcc782e21d54b81f21a8e3c76e9adfac51b6786499061266c90849061543b565b60405180910390a260975461268d90610100900463ffffffff1660016154db565b609760016101000a81548163ffffffff021916908363ffffffff16021790555050505050505050505050565b609b546001600160a01b031633146127135760405162461bcd60e51b815260206004820152601d60248201527f41676772656761746f72206d757374206265207468652063616c6c65720000006044820152606401610674565b600061272260208501856144d6565b90503660006127346020870187615503565b9092509050600061274b60c0880160a089016144d6565b90506098600061275e60208901896144d6565b63ffffffff1663ffffffff168152602001908152602001600020548760405160200161278a919061558e565b60405160208183030381529060405280519060200120146128135760405162461bcd60e51b815260206004820152603d60248201527f737570706c696564207461736b20646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610674565b600060998161282560208a018a6144d6565b63ffffffff1663ffffffff16815260200190815260200160002054146128a25760405162461bcd60e51b815260206004820152602c60248201527f41676772656761746f722068617320616c726561647920726573706f6e64656460448201526b20746f20746865207461736b60a01b6064820152608401610674565b6128cc7f0000000000000000000000000000000000000000000000000000000000000000856154db565b63ffffffff164363ffffffff16111561293d5760405162461bcd60e51b815260206004820152602d60248201527f41676772656761746f722068617320726573706f6e64656420746f207468652060448201526c7461736b20746f6f206c61746560981b6064820152608401610674565b60008660405160200161295091906156a3565b6040516020818303038152906040528051906020012090506000806129788387878a8c61158b565b9150915060005b85811015612a77578460ff16836020015182815181106129a1576129a1615003565b60200260200101516129b391906156b6565b6001600160601b03166064846000015183815181106129d4576129d4615003565b60200260200101516001600160601b03166129ef91906156e5565b1015612a65576040805162461bcd60e51b81526020600482015260248101919091527f5369676e61746f7269657320646f206e6f74206f776e206174206c656173742060448201527f7468726573686f6c642070657263656e74616765206f6620612071756f72756d6064820152608401610674565b80612a6f8161510a565b91505061297f565b5060408051808201825263ffffffff43168152602080820184905291519091612aa4918c91849101615704565b60405160208183030381529060405280519060200120609960008c6000016020810190612ad191906144d6565b63ffffffff1663ffffffff168152602001908152602001600020819055507fd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb548a82604051612b20929190615704565b60405180910390a15050505050505050505050565b6000612b4460208501856144d6565b63ffffffff8116600090815260996020526040902054909150612bb35760405162461bcd60e51b815260206004820152602160248201527f5461736b206861736e2774206265656e20726573706f6e64656420746f2079656044820152601d60fa1b6064820152608401610674565b8383604051602001612bc6929190615737565b60408051601f19818403018152918152815160209283012063ffffffff84166000908152609990935291205414612c655760405162461bcd60e51b815260206004820152603d60248201527f5461736b20726573706f6e736520646f6573206e6f74206d617463682074686560448201527f206f6e65207265636f7264656420696e2074686520636f6e74726163740000006064820152608401610674565b63ffffffff81166000908152609a602052604090205460ff1615612cfd5760405162461bcd60e51b815260206004820152604360248201527f54686520726573706f6e736520746f2074686973207461736b2068617320616c60448201527f7265616479206265656e206368616c6c656e676564207375636365737366756c606482015262363c9760e91b608482015260a401610674565b6064612d0c60208501856144d6565b612d1691906154db565b63ffffffff164363ffffffff161115612d975760405162461bcd60e51b815260206004820152603760248201527f546865206368616c6c656e676520706572696f6420666f72207468697320746160448201527f736b2068617320616c726561647920657870697265642e0000000000000000006064820152608401610674565b600082516001600160401b03811115612db257612db2614303565b604051908082528060200260200182016040528015612ddb578160200160208202803683370190505b50905060005b8351811015612e2e57612dff84828151811061197e5761197e615003565b828281518110612e1157612e11615003565b602090810291909101015280612e268161510a565b915050612de1565b506000612e3e60208801886144d6565b82604051602001612e50929190615397565b60405160208183030381529060405280519060200120905084602001358114612efa5760405162461bcd60e51b815260206004820152605060248201527f546865207075626b657973206f66206e6f6e2d7369676e696e67206f7065726160448201527f746f727320737570706c69656420627920746865206368616c6c656e6765722060648201526f30b932903737ba1031b7b93932b1ba1760811b608482015260a401610674565b600084516001600160401b03811115612f1557612f15614303565b604051908082528060200260200182016040528015612f3e578160200160208202803683370190505b50905060005b8551811015613031577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663e8bb9ae6858381518110612f8e57612f8e615003565b60200260200101516040518263ffffffff1660e01b8152600401612fb491815260200190565b602060405180830381865afa158015612fd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ff59190614f37565b82828151811061300757613007615003565b6001600160a01b0390921660209283029190910190910152806130298161510a565b915050612f44565b5063ffffffff84166000818152609a6020526040808220805460ff19166001179055513392917fc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec91a35050505050505050565b60408051600180825281830190925260009160609183916020808301908036833701905050905084816000815181106130bf576130bf615003565b60209081029190910101526040516361c8a12f60e11b81526000906001600160a01b0388169063c391425e906130fb9088908690600401615775565b600060405180830381865afa158015613118573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052613140919081019061516f565b60008151811061315257613152615003565b60209081029190910101516040516304ec635160e01b81526004810188905263ffffffff87811660248301529091166044820181905291506000906001600160a01b038916906304ec635190606401602060405180830381865afa1580156131be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131e2919061524f565b6001600160c01b0316905060006131f882613de0565b9050816132068a838a610952565b9550955050505050935093915050565b61321e613d34565b6001600160a01b0381166132835760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610674565b61068681613d8e565b600054610100900460ff16158080156132ac5750600054600160ff909116105b806132c65750303b1580156132c6575060005460ff166001145b6133295760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610674565b6000805460ff19166001179055801561334c576000805461ff0019166101001790555b613357856000613e3d565b61336084613d8e565b609b80546001600160a01b038086166001600160a01b031992831617909255609c80549285169290911691909117905580156133d6576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b606560009054906101000a90046001600160a01b03166001600160a01b031663eab66d7a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613430573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134549190614f37565b6001600160a01b0316336001600160a01b0316146134845760405162461bcd60e51b815260040161067490614f54565b6066541981196066541916146135025760405162461bcd60e51b815260206004820152603860248201527f5061757361626c652e756e70617573653a20696e76616c696420617474656d7060448201527f7420746f2070617573652066756e6374696f6e616c69747900000000000000006064820152608401610674565b606681905560405181815233907f3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c906020016107bd565b6001600160a01b0381166135c75760405162461bcd60e51b815260206004820152604960248201527f5061757361626c652e5f73657450617573657252656769737472793a206e657760448201527f50617573657252656769737472792063616e6e6f7420626520746865207a65726064820152686f206164647265737360b81b608482015260a401610674565b606554604080516001600160a01b03928316815291831660208301527f6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6910160405180910390a1606580546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082019091526000808252602082015261364c6141de565b835181526020808501519082015260408082018490526000908360608460076107d05a03fa905080801561367f57613681565bfe5b50806136bf5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5b5d5b0b59985a5b1959609a1b6044820152606401610674565b505092915050565b60408051808201909152600080825260208201526136e36141fc565b835181526020808501518183015283516040808401919091529084015160608301526000908360808460066107d05a03fa905080801561367f5750806136bf5760405162461bcd60e51b815260206004820152600d60248201526c1958cb5859190b59985a5b1959609a1b6044820152606401610674565b61376361421a565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f275dc4a288d1afb3cbb1ac09187524c7db36395df7be3b99e673b13a075a65ec82527f1d9befcd05a5323e6da4d435f3b617cdb3af83285c2df711ef39c01571827f9d60208381019190915281019190915290565b60408051808201909152600080825260208201526000808061384b60008051602061581b83398151915286615019565b90505b61385781613f27565b909350915060008051602061581b833981519152828309831415613891576040805180820190915290815260208101919091529392505050565b60008051602061581b83398151915260018208905061384e565b6040805180820182528681526020808201869052825180840190935286835282018490526000918291906138dd61423f565b60005b6002811015613aa25760006138f68260066156e5565b905084826002811061390a5761390a615003565b6020020151518361391c83600061532c565b600c811061392c5761392c615003565b602002015284826002811061394357613943615003565b6020020151602001518382600161395a919061532c565b600c811061396a5761396a615003565b602002015283826002811061398157613981615003565b602002015151518361399483600261532c565b600c81106139a4576139a4615003565b60200201528382600281106139bb576139bb615003565b60200201515160016020020151836139d483600361532c565b600c81106139e4576139e4615003565b60200201528382600281106139fb576139fb615003565b602002015160200151600060028110613a1657613a16615003565b602002015183613a2783600461532c565b600c8110613a3757613a37615003565b6020020152838260028110613a4e57613a4e615003565b602002015160200151600160028110613a6957613a69615003565b602002015183613a7a83600561532c565b600c8110613a8a57613a8a615003565b60200201525080613a9a8161510a565b9150506138e0565b50613aab61425e565b60006020826101808560088cfa9151919c9115159b50909950505050505050505050565b600080613adb84613fa9565b90508015613b81578260ff168460018651613af691906152fc565b81518110613b0657613b06615003565b016020015160f81c10613b815760405162461bcd60e51b815260206004820152603f60248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206269746d61702065786365656473206d61782076616c7565006064820152608401610674565b90505b92915050565b6000805b8215613b8457613b9f6001846152fc565b9092169180613bad816157c9565b915050613b8e565b60408051808201909152600080825260208201526102008261ffff1610613c115760405162461bcd60e51b815260206004820152601060248201526f7363616c61722d746f6f2d6c6172676560801b6044820152606401610674565b8161ffff1660011415613c25575081613b84565b6040805180820190915260008082526020820181905284906001905b8161ffff168661ffff1610613c8e57600161ffff871660ff83161c81161415613c7157613c6e84846136c7565b93505b613c7b83846136c7565b92506201fffe600192831b169101613c41565b509195945050505050565b60408051808201909152600080825260208201528151158015613cbe57506020820151155b15613cdc575050604080518082019091526000808252602082015290565b60405180604001604052808360000151815260200160008051602061581b8339815191528460200151613d0f9190615019565b613d279060008051602061581b8339815191526152fc565b905292915050565b919050565b6033546001600160a01b031633146124485760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610674565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60606000805b610100811015613e36576001811b915083821615613e2657828160f81b604051602001613e149291906157eb565b60405160208183030381529060405292505b613e2f8161510a565b9050613de6565b5050919050565b6065546001600160a01b0316158015613e5e57506001600160a01b03821615155b613ee05760405162461bcd60e51b815260206004820152604760248201527f5061757361626c652e5f696e697469616c697a655061757365723a205f696e6960448201527f7469616c697a6550617573657228292063616e206f6e6c792062652063616c6c6064820152666564206f6e636560c81b608482015260a401610674565b606681905560405181815233907fab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d9060200160405180910390a2613f2382613539565b5050565b6000808060008051602061581b833981519152600360008051602061581b8339815191528660008051602061581b833981519152888909090890506000613f9d827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f5260008051602061581b833981519152614136565b91959194509092505050565b6000610100825111156140325760405162461bcd60e51b8152602060048201526044602482018190527f4269746d61705574696c732e6f72646572656442797465734172726179546f42908201527f69746d61703a206f7264657265644279746573417272617920697320746f6f206064820152636c6f6e6760e01b608482015260a401610674565b815161404057506000919050565b6000808360008151811061405657614056615003565b0160200151600160f89190911c81901b92505b845181101561412d5784818151811061408457614084615003565b0160200151600160f89190911c1b91508282116141195760405162461bcd60e51b815260206004820152604760248201527f4269746d61705574696c732e6f72646572656442797465734172726179546f4260448201527f69746d61703a206f72646572656442797465734172726179206973206e6f74206064820152661bdc99195c995960ca1b608482015260a401610674565b918117916141268161510a565b9050614069565b50909392505050565b60008061414161425e565b61414961427c565b602080825281810181905260408201819052606082018890526080820187905260a082018690528260c08360056107d05a03fa925082801561367f5750826141d35760405162461bcd60e51b815260206004820152601a60248201527f424e3235342e6578704d6f643a2063616c6c206661696c7572650000000000006044820152606401610674565b505195945050505050565b60405180606001604052806003906020820280368337509192915050565b60405180608001604052806004906020820280368337509192915050565b604051806040016040528061422d61429a565b815260200161423a61429a565b905290565b604051806101800160405280600c906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180604001604052806002906020820280368337509192915050565b6001600160a01b038116811461068657600080fd5b6000602082840312156142df57600080fd5b8135613b81816142b8565b6000602082840312156142fc57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561433b5761433b614303565b60405290565b60405161010081016001600160401b038111828210171561433b5761433b614303565b604051601f8201601f191681016001600160401b038111828210171561438c5761438c614303565b604052919050565b6000604082840312156143a657600080fd5b6143ae614319565b9050813581526020820135602082015292915050565b600082601f8301126143d557600080fd5b604051604081018181106001600160401b03821117156143f7576143f7614303565b806040525080604084018581111561440e57600080fd5b845b81811015613c8e578035835260209283019201614410565b60006080828403121561443a57600080fd5b614442614319565b905061444e83836143c4565b815261445d83604084016143c4565b602082015292915050565b600080600080610120858703121561447f57600080fd5b843593506144908660208701614394565b925061449f8660608701614428565b91506144ae8660e08701614394565b905092959194509250565b63ffffffff8116811461068657600080fd5b8035613d2f816144b9565b6000602082840312156144e857600080fd5b8135613b81816144b9565b60008060006060848603121561450857600080fd5b8335614513816142b8565b92506020848101356001600160401b038082111561453057600080fd5b818701915087601f83011261454457600080fd5b81358181111561455657614556614303565b614568601f8201601f19168501614364565b9150808252888482850101111561457e57600080fd5b80848401858401376000848284010152508094505050506145a1604085016144cb565b90509250925092565b600081518084526020808501808196508360051b810191508286016000805b8681101561462c578385038a52825180518087529087019087870190845b81811015614617578351805184528a01516001600160601b03168a840152928901926040909201916001016145e7565b50509a87019a955050918501916001016145c9565b509298975050505050505050565b60208152600061464d60208301846145aa565b9392505050565b801515811461068657600080fd5b60006020828403121561467457600080fd5b8135613b8181614654565b60008083601f84011261469157600080fd5b5081356001600160401b038111156146a857600080fd5b6020830191508360208285010111156146c057600080fd5b9250929050565b600080600080600080608087890312156146e057600080fd5b86356146eb816142b8565b955060208701356146fb816144b9565b945060408701356001600160401b038082111561471757600080fd5b6147238a838b0161467f565b9096509450606089013591508082111561473c57600080fd5b818901915089601f83011261475057600080fd5b81358181111561475f57600080fd5b8a60208260051b850101111561477457600080fd5b6020830194508093505050509295509295509295565b600081518084526020808501945080840160005b838110156147c057815163ffffffff168752958201959082019060010161479e565b509495945050505050565b6000602080835283516080828501526147e760a085018261478a565b905081850151601f1980868403016040870152614804838361478a565b92506040870151915080868403016060870152614821838361478a565b60608801518782038301608089015280518083529194508501925084840190600581901b8501860160005b82811015614878578487830301845261486682875161478a565b9588019593880193915060010161484c565b509998505050505050505050565b60ff8116811461068657600080fd5b6000602082840312156148a757600080fd5b8135613b8181614886565b60006001600160401b038211156148cb576148cb614303565b5060051b60200190565b600082601f8301126148e657600080fd5b813560206148fb6148f6836148b2565b614364565b82815260059290921b8401810191818101908684111561491a57600080fd5b8286015b84811015610cc0578035614931816144b9565b835291830191830161491e565b600082601f83011261494f57600080fd5b8135602061495f6148f6836148b2565b82815260069290921b8401810191818101908684111561497e57600080fd5b8286015b84811015610cc0576149948882614394565b835291830191604001614982565b600082601f8301126149b357600080fd5b813560206149c36148f6836148b2565b82815260059290921b840181019181810190868411156149e257600080fd5b8286015b84811015610cc05780356001600160401b03811115614a055760008081fd5b614a138986838b01016148d5565b8452509183019183016149e6565b60006101808284031215614a3457600080fd5b614a3c614341565b905081356001600160401b0380821115614a5557600080fd5b614a61858386016148d5565b83526020840135915080821115614a7757600080fd5b614a838583860161493e565b60208401526040840135915080821115614a9c57600080fd5b614aa88583860161493e565b6040840152614aba8560608601614428565b6060840152614acc8560e08601614394565b6080840152610120840135915080821115614ae657600080fd5b614af2858386016148d5565b60a0840152610140840135915080821115614b0c57600080fd5b614b18858386016148d5565b60c0840152610160840135915080821115614b3257600080fd5b50614b3f848285016149a2565b60e08301525092915050565b600080600080600060808688031215614b6357600080fd5b8535945060208601356001600160401b0380821115614b8157600080fd5b614b8d89838a0161467f565b909650945060408801359150614ba2826144b9565b90925060608701359080821115614bb857600080fd5b50614bc588828901614a21565b9150509295509295909350565b600081518084526020808501945080840160005b838110156147c05781516001600160601b031687529582019590820190600101614be6565b6040815260008351604080840152614c266080840182614bd2565b90506020850151603f19848303016060850152614c438282614bd2565b925050508260208301529392505050565b600080600080600080600080600060a08a8c031215614c7257600080fd5b89356001600160401b0380821115614c8957600080fd5b614c958d838e0161467f565b909b50995060208c0135915080821115614cae57600080fd5b614cba8d838e0161467f565b909950975060408c0135915080821115614cd357600080fd5b614cdf8d838e0161467f565b909750955060608c01359150614cf4826144b9565b90935060808b01359080821115614d0a57600080fd5b50614d178c828d0161467f565b915080935050809150509295985092959850929598565b600060c08284031215614d4057600080fd5b50919050565b600060408284031215614d4057600080fd5b600080600060608486031215614d6d57600080fd5b83356001600160401b0380821115614d8457600080fd5b614d9087838801614d2e565b94506020860135915080821115614da657600080fd5b614db287838801614d46565b93506040860135915080821115614dc857600080fd5b50614dd586828701614a21565b9150509250925092565b60008060008060a08587031215614df557600080fd5b84356001600160401b0380821115614e0c57600080fd5b614e1888838901614d2e565b95506020870135915080821115614e2e57600080fd5b614e3a88838901614d46565b9450614e498860408901614d46565b93506080870135915080821115614e5f57600080fd5b50614e6c8782880161493e565b91505092959194509250565b600080600060608486031215614e8d57600080fd5b8335614e98816142b8565b9250602084013591506040840135614eaf816144b9565b809150509250925092565b828152604060208201526000614ed360408301846145aa565b949350505050565b60008060008060808587031215614ef157600080fd5b8435614efc816142b8565b93506020850135614f0c816142b8565b92506040850135614f1c816142b8565b91506060850135614f2c816142b8565b939692955090935050565b600060208284031215614f4957600080fd5b8151613b81816142b8565b6020808252602a908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526939903ab73830bab9b2b960b11b606082015260800190565b600060208284031215614fb057600080fd5b8151613b8181614654565b60208082526028908201527f6d73672e73656e646572206973206e6f74207065726d697373696f6e6564206160408201526739903830bab9b2b960c11b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b60008261503657634e487b7160e01b600052601260045260246000fd5b500690565b6000602080838503121561504e57600080fd5b82516001600160401b0381111561506457600080fd5b8301601f8101851361507557600080fd5b80516150836148f6826148b2565b81815260059190911b820183019083810190878311156150a257600080fd5b928401925b828410156150c0578351825292840192908401906150a7565b979650505050505050565b6000602082840312156150dd57600080fd5b81516001600160601b0381168114613b8157600080fd5b634e487b7160e01b600052601160045260246000fd5b600060001982141561511e5761511e6150f4565b5060010190565b63ffffffff84168152604060208201819052810182905260006001600160fb1b0383111561515257600080fd5b8260051b8085606085013760009201606001918252509392505050565b6000602080838503121561518257600080fd5b82516001600160401b0381111561519857600080fd5b8301601f810185136151a957600080fd5b80516151b76148f6826148b2565b81815260059190911b820183019083810190878311156151d657600080fd5b928401925b828410156150c05783516151ee816144b9565b825292840192908401906151db565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b63ffffffff841681526040602082015260006152466040830184866151fd565b95945050505050565b60006020828403121561526157600080fd5b81516001600160c01b0381168114613b8157600080fd5b60006020828403121561528a57600080fd5b8151613b81816144b9565b600060ff821660ff8114156152ac576152ac6150f4565b60010192915050565b6040815260006152c96040830185876151fd565b905063ffffffff83166020830152949350505050565b6000602082840312156152f157600080fd5b8151613b8181614886565b60008282101561530e5761530e6150f4565b500390565b60006020828403121561532557600080fd5b5051919050565b6000821982111561533f5761533f6150f4565b500190565b60006020828403121561535657600080fd5b815167ffffffffffffffff1981168114613b8157600080fd5b60006001600160601b038381169083168181101561538f5761538f6150f4565b039392505050565b63ffffffff60e01b8360e01b1681526000600482018351602080860160005b838110156153d2578151855293820193908201906001016153b6565b5092979650505050505050565b60005b838110156153fa5781810151838201526020016153e2565b83811115615409576000848401525b50505050565b600081518084526154278160208601602086016153df565b601f01601f19169290920160200192915050565b60208152600063ffffffff808451166020840152602084015160c0604085015261546860e085018261540f565b90506040850151601f1980868403016060870152615486838361540f565b925060608701519150808684030160808701526154a3838361540f565b925060808701519150808684030160a0870152506154c1828261540f565b9150508160a08601511660c0850152809250505092915050565b600063ffffffff8083168185168083038211156154fa576154fa6150f4565b01949350505050565b6000808335601e1984360301811261551a57600080fd5b8301803591506001600160401b0382111561553457600080fd5b6020019150368190038213156146c057600080fd5b6000808335601e1984360301811261556057600080fd5b83016020810192503590506001600160401b0381111561557f57600080fd5b8036038313156146c057600080fd5b602081526000823561559f816144b9565b63ffffffff80821660208501526155b96020860186615549565b925060c060408601526155d060e0860184836151fd565b9250506155e06040860186615549565b601f19808786030160608801526155f88583856151fd565b94506156076060890189615549565b93509150808786030160808801526156208584846151fd565b945061562f6080890189615549565b93509150808786030160a0880152506156498483836151fd565b9350505060a085013561565b816144b9565b1660c0939093019290925250919050565b60008135615679816144b9565b63ffffffff16835261568e6020830183615549565b604060208601526152466040860182846151fd565b60208152600061464d602083018461566c565b60006001600160601b03808316818516818304811182151516156156dc576156dc6150f4565b02949350505050565b60008160001904831182151516156156ff576156ff6150f4565b500290565b606081526000615717606083018561566c565b905063ffffffff8351166020830152602083015160408301529392505050565b60608152600061574a606083018561566c565b90508235615757816144b9565b63ffffffff8116602084015250602083013560408301529392505050565b60006040820163ffffffff851683526020604081850152818551808452606086019150828701935060005b818110156157bc578451835293830193918301916001016157a0565b5090979650505050505050565b600061ffff808316818114156157e1576157e16150f4565b6001019392505050565b600083516157fd8184602088016153df565b6001600160f81b031993909316919092019081526001019291505056fe30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47424c535369676e6174757265436865636b65722e636865636b5369676e617475a264697066735822122073df32a6a5343e080a21454e295803cbb764f424dd6159359c299a609f3212ea64736f6c634300080c0033",
}

// ContractMfssiaTaskManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMfssiaTaskManagerMetaData.ABI instead.
var ContractMfssiaTaskManagerABI = ContractMfssiaTaskManagerMetaData.ABI

// ContractMfssiaTaskManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMfssiaTaskManagerMetaData.Bin instead.
var ContractMfssiaTaskManagerBin = ContractMfssiaTaskManagerMetaData.Bin

// DeployContractMfssiaTaskManager deploys a new Ethereum contract, binding an instance of ContractMfssiaTaskManager to it.
func DeployContractMfssiaTaskManager(auth *bind.TransactOpts, backend bind.ContractBackend, _registryCoordinator common.Address, _taskResponseWindowBlock uint32) (common.Address, *types.Transaction, *ContractMfssiaTaskManager, error) {
	parsed, err := ContractMfssiaTaskManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractMfssiaTaskManagerBin), backend, _registryCoordinator, _taskResponseWindowBlock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractMfssiaTaskManager{ContractMfssiaTaskManagerCaller: ContractMfssiaTaskManagerCaller{contract: contract}, ContractMfssiaTaskManagerTransactor: ContractMfssiaTaskManagerTransactor{contract: contract}, ContractMfssiaTaskManagerFilterer: ContractMfssiaTaskManagerFilterer{contract: contract}}, nil
}

// ContractMfssiaTaskManager is an auto generated Go binding around an Ethereum contract.
type ContractMfssiaTaskManager struct {
	ContractMfssiaTaskManagerCaller     // Read-only binding to the contract
	ContractMfssiaTaskManagerTransactor // Write-only binding to the contract
	ContractMfssiaTaskManagerFilterer   // Log filterer for contract events
}

// ContractMfssiaTaskManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaTaskManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaTaskManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractMfssiaTaskManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractMfssiaTaskManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractMfssiaTaskManagerSession struct {
	Contract     *ContractMfssiaTaskManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractMfssiaTaskManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractMfssiaTaskManagerCallerSession struct {
	Contract *ContractMfssiaTaskManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ContractMfssiaTaskManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractMfssiaTaskManagerTransactorSession struct {
	Contract     *ContractMfssiaTaskManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ContractMfssiaTaskManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerRaw struct {
	Contract *ContractMfssiaTaskManager // Generic contract binding to access the raw methods on
}

// ContractMfssiaTaskManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerCallerRaw struct {
	Contract *ContractMfssiaTaskManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractMfssiaTaskManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractMfssiaTaskManagerTransactorRaw struct {
	Contract *ContractMfssiaTaskManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractMfssiaTaskManager creates a new instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManager(address common.Address, backend bind.ContractBackend) (*ContractMfssiaTaskManager, error) {
	contract, err := bindContractMfssiaTaskManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManager{ContractMfssiaTaskManagerCaller: ContractMfssiaTaskManagerCaller{contract: contract}, ContractMfssiaTaskManagerTransactor: ContractMfssiaTaskManagerTransactor{contract: contract}, ContractMfssiaTaskManagerFilterer: ContractMfssiaTaskManagerFilterer{contract: contract}}, nil
}

// NewContractMfssiaTaskManagerCaller creates a new read-only instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManagerCaller(address common.Address, caller bind.ContractCaller) (*ContractMfssiaTaskManagerCaller, error) {
	contract, err := bindContractMfssiaTaskManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerCaller{contract: contract}, nil
}

// NewContractMfssiaTaskManagerTransactor creates a new write-only instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractMfssiaTaskManagerTransactor, error) {
	contract, err := bindContractMfssiaTaskManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTransactor{contract: contract}, nil
}

// NewContractMfssiaTaskManagerFilterer creates a new log filterer instance of ContractMfssiaTaskManager, bound to a specific deployed contract.
func NewContractMfssiaTaskManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractMfssiaTaskManagerFilterer, error) {
	contract, err := bindContractMfssiaTaskManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerFilterer{contract: contract}, nil
}

// bindContractMfssiaTaskManager binds a generic wrapper to an already deployed contract.
func bindContractMfssiaTaskManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMfssiaTaskManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMfssiaTaskManager.Contract.ContractMfssiaTaskManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.ContractMfssiaTaskManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.ContractMfssiaTaskManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractMfssiaTaskManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.contract.Transact(opts, method, params...)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TASKCHALLENGEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "TASK_CHALLENGE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// TASKCHALLENGEWINDOWBLOCK is a free data retrieval call binding the contract method 0xf63c5bab.
//
// Solidity: function TASK_CHALLENGE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TASKCHALLENGEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKCHALLENGEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TASKRESPONSEWINDOWBLOCK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "TASK_RESPONSE_WINDOW_BLOCK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// TASKRESPONSEWINDOWBLOCK is a free data retrieval call binding the contract method 0x1ad43189.
//
// Solidity: function TASK_RESPONSE_WINDOW_BLOCK() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TASKRESPONSEWINDOWBLOCK() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TASKRESPONSEWINDOWBLOCK(&_ContractMfssiaTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Aggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "aggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Aggregator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Aggregator(&_ContractMfssiaTaskManager.CallOpts)
}

// Aggregator is a free data retrieval call binding the contract method 0x245a7bfc.
//
// Solidity: function aggregator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Aggregator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Aggregator(&_ContractMfssiaTaskManager.CallOpts)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) AllTaskHashes(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "allTaskHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskHashes(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// AllTaskHashes is a free data retrieval call binding the contract method 0x2d89f6fc.
//
// Solidity: function allTaskHashes(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) AllTaskHashes(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskHashes(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) AllTaskResponses(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "allTaskResponses", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskResponses(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// AllTaskResponses is a free data retrieval call binding the contract method 0x2cb223d5.
//
// Solidity: function allTaskResponses(uint32 ) view returns(bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) AllTaskResponses(arg0 uint32) ([32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.AllTaskResponses(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.BlsApkRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.BlsApkRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) CheckSignatures(opts *bind.CallOpts, msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "checkSignatures", msgHash, quorumNumbers, referenceBlockNumber, params)

	if err != nil {
		return *new(IBLSSignatureCheckerQuorumStakeTotals), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(IBLSSignatureCheckerQuorumStakeTotals)).(*IBLSSignatureCheckerQuorumStakeTotals)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.CheckSignatures(&_ContractMfssiaTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// CheckSignatures is a free data retrieval call binding the contract method 0x6efb4636.
//
// Solidity: function checkSignatures(bytes32 msgHash, bytes quorumNumbers, uint32 referenceBlockNumber, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) params) view returns((uint96[],uint96[]), bytes32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) CheckSignatures(msgHash [32]byte, quorumNumbers []byte, referenceBlockNumber uint32, params IBLSSignatureCheckerNonSignerStakesAndSignature) (IBLSSignatureCheckerQuorumStakeTotals, [32]byte, error) {
	return _ContractMfssiaTaskManager.Contract.CheckSignatures(&_ContractMfssiaTaskManager.CallOpts, msgHash, quorumNumbers, referenceBlockNumber, params)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Delegation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "delegation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Delegation() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Delegation(&_ContractMfssiaTaskManager.CallOpts)
}

// Delegation is a free data retrieval call binding the contract method 0xdf5cf723.
//
// Solidity: function delegation() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Delegation() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Delegation(&_ContractMfssiaTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Generator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "generator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Generator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Generator(&_ContractMfssiaTaskManager.CallOpts)
}

// Generator is a free data retrieval call binding the contract method 0x7afa1eed.
//
// Solidity: function generator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Generator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Generator(&_ContractMfssiaTaskManager.CallOpts)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetCheckSignaturesIndices(opts *bind.CallOpts, registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getCheckSignaturesIndices", registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)

	if err != nil {
		return *new(OperatorStateRetrieverCheckSignaturesIndices), err
	}

	out0 := *abi.ConvertType(out[0], new(OperatorStateRetrieverCheckSignaturesIndices)).(*OperatorStateRetrieverCheckSignaturesIndices)

	return out0, err

}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractMfssiaTaskManager.Contract.GetCheckSignaturesIndices(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetCheckSignaturesIndices is a free data retrieval call binding the contract method 0x4f739f74.
//
// Solidity: function getCheckSignaturesIndices(address registryCoordinator, uint32 referenceBlockNumber, bytes quorumNumbers, bytes32[] nonSignerOperatorIds) view returns((uint32[],uint32[],uint32[],uint32[][]))
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetCheckSignaturesIndices(registryCoordinator common.Address, referenceBlockNumber uint32, quorumNumbers []byte, nonSignerOperatorIds [][32]byte) (OperatorStateRetrieverCheckSignaturesIndices, error) {
	return _ContractMfssiaTaskManager.Contract.GetCheckSignaturesIndices(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, referenceBlockNumber, quorumNumbers, nonSignerOperatorIds)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetOperatorState(opts *bind.CallOpts, registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getOperatorState", registryCoordinator, quorumNumbers, blockNumber)

	if err != nil {
		return *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, err

}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState is a free data retrieval call binding the contract method 0x3563b0d1.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes quorumNumbers, uint32 blockNumber) view returns((bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetOperatorState(registryCoordinator common.Address, quorumNumbers []byte, blockNumber uint32) ([][]OperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, quorumNumbers, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetOperatorState0(opts *bind.CallOpts, registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getOperatorState0", registryCoordinator, operatorId, blockNumber)

	if err != nil {
		return *new(*big.Int), *new([][]OperatorStateRetrieverOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][]OperatorStateRetrieverOperator)).(*[][]OperatorStateRetrieverOperator)

	return out0, out1, err

}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState0(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetOperatorState0 is a free data retrieval call binding the contract method 0xcefdc1d4.
//
// Solidity: function getOperatorState(address registryCoordinator, bytes32 operatorId, uint32 blockNumber) view returns(uint256, (bytes32,uint96)[][])
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetOperatorState0(registryCoordinator common.Address, operatorId [32]byte, blockNumber uint32) (*big.Int, [][]OperatorStateRetrieverOperator, error) {
	return _ContractMfssiaTaskManager.Contract.GetOperatorState0(&_ContractMfssiaTaskManager.CallOpts, registryCoordinator, operatorId, blockNumber)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) GetTaskResponseWindowBlock(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "getTaskResponseWindowBlock")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractMfssiaTaskManager.CallOpts)
}

// GetTaskResponseWindowBlock is a free data retrieval call binding the contract method 0xf5c9899d.
//
// Solidity: function getTaskResponseWindowBlock() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) GetTaskResponseWindowBlock() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.GetTaskResponseWindowBlock(&_ContractMfssiaTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) LatestTaskNum(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "latestTaskNum")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) LatestTaskNum() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.LatestTaskNum(&_ContractMfssiaTaskManager.CallOpts)
}

// LatestTaskNum is a free data retrieval call binding the contract method 0x8b00ce7c.
//
// Solidity: function latestTaskNum() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) LatestTaskNum() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.LatestTaskNum(&_ContractMfssiaTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Owner() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Owner(&_ContractMfssiaTaskManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Owner() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.Owner(&_ContractMfssiaTaskManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Paused(index uint8) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.Paused(&_ContractMfssiaTaskManager.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Paused(index uint8) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.Paused(&_ContractMfssiaTaskManager.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Paused0() (*big.Int, error) {
	return _ContractMfssiaTaskManager.Contract.Paused0(&_ContractMfssiaTaskManager.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) Paused0() (*big.Int, error) {
	return _ContractMfssiaTaskManager.Contract.Paused0(&_ContractMfssiaTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) PauserRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.PauserRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) PauserRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.PauserRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) RegistryCoordinator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "registryCoordinator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.RegistryCoordinator(&_ContractMfssiaTaskManager.CallOpts)
}

// RegistryCoordinator is a free data retrieval call binding the contract method 0x6d14a987.
//
// Solidity: function registryCoordinator() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) RegistryCoordinator() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.RegistryCoordinator(&_ContractMfssiaTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) StakeRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "stakeRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) StakeRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.StakeRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// StakeRegistry is a free data retrieval call binding the contract method 0x68304835.
//
// Solidity: function stakeRegistry() view returns(address)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) StakeRegistry() (common.Address, error) {
	return _ContractMfssiaTaskManager.Contract.StakeRegistry(&_ContractMfssiaTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) StaleStakesForbidden(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "staleStakesForbidden")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) StaleStakesForbidden() (bool, error) {
	return _ContractMfssiaTaskManager.Contract.StaleStakesForbidden(&_ContractMfssiaTaskManager.CallOpts)
}

// StaleStakesForbidden is a free data retrieval call binding the contract method 0xb98d0908.
//
// Solidity: function staleStakesForbidden() view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) StaleStakesForbidden() (bool, error) {
	return _ContractMfssiaTaskManager.Contract.StaleStakesForbidden(&_ContractMfssiaTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TaskNumber(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "taskNumber")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TaskNumber() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TaskNumber(&_ContractMfssiaTaskManager.CallOpts)
}

// TaskNumber is a free data retrieval call binding the contract method 0x72d18e8d.
//
// Solidity: function taskNumber() view returns(uint32)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TaskNumber() (uint32, error) {
	return _ContractMfssiaTaskManager.Contract.TaskNumber(&_ContractMfssiaTaskManager.CallOpts)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TaskSuccesfullyChallenged(opts *bind.CallOpts, arg0 uint32) (bool, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "taskSuccesfullyChallenged", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// TaskSuccesfullyChallenged is a free data retrieval call binding the contract method 0x5decc3f5.
//
// Solidity: function taskSuccesfullyChallenged(uint32 ) view returns(bool)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TaskSuccesfullyChallenged(arg0 uint32) (bool, error) {
	return _ContractMfssiaTaskManager.Contract.TaskSuccesfullyChallenged(&_ContractMfssiaTaskManager.CallOpts, arg0)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCaller) TrySignatureAndApkVerification(opts *bind.CallOpts, msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	var out []interface{}
	err := _ContractMfssiaTaskManager.contract.Call(opts, &out, "trySignatureAndApkVerification", msgHash, apk, apkG2, sigma)

	outstruct := new(struct {
		PairingSuccessful bool
		SiganatureIsValid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PairingSuccessful = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SiganatureIsValid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractMfssiaTaskManager.Contract.TrySignatureAndApkVerification(&_ContractMfssiaTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// TrySignatureAndApkVerification is a free data retrieval call binding the contract method 0x171f1d5b.
//
// Solidity: function trySignatureAndApkVerification(bytes32 msgHash, (uint256,uint256) apk, (uint256[2],uint256[2]) apkG2, (uint256,uint256) sigma) view returns(bool pairingSuccessful, bool siganatureIsValid)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerCallerSession) TrySignatureAndApkVerification(msgHash [32]byte, apk BN254G1Point, apkG2 BN254G2Point, sigma BN254G1Point) (struct {
	PairingSuccessful bool
	SiganatureIsValid bool
}, error) {
	return _ContractMfssiaTaskManager.Contract.TrySignatureAndApkVerification(&_ContractMfssiaTaskManager.CallOpts, msgHash, apk, apkG2, sigma)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa11c1a5d.
//
// Solidity: function createNewTask(string system1Value, string system2Value, string dkgValue, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) CreateNewTask(opts *bind.TransactOpts, system1Value string, system2Value string, dkgValue string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "createNewTask", system1Value, system2Value, dkgValue, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa11c1a5d.
//
// Solidity: function createNewTask(string system1Value, string system2Value, string dkgValue, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) CreateNewTask(system1Value string, system2Value string, dkgValue string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.CreateNewTask(&_ContractMfssiaTaskManager.TransactOpts, system1Value, system2Value, dkgValue, quorumThresholdPercentage, quorumNumbers)
}

// CreateNewTask is a paid mutator transaction binding the contract method 0xa11c1a5d.
//
// Solidity: function createNewTask(string system1Value, string system2Value, string dkgValue, uint32 quorumThresholdPercentage, bytes quorumNumbers) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) CreateNewTask(system1Value string, system2Value string, dkgValue string, quorumThresholdPercentage uint32, quorumNumbers []byte) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.CreateNewTask(&_ContractMfssiaTaskManager.TransactOpts, system1Value, system2Value, dkgValue, quorumThresholdPercentage, quorumNumbers)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) Initialize(opts *bind.TransactOpts, _pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "initialize", _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Initialize(&_ContractMfssiaTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _pauserRegistry, address initialOwner, address _aggregator, address _generator) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) Initialize(_pauserRegistry common.Address, initialOwner common.Address, _aggregator common.Address, _generator common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Initialize(&_ContractMfssiaTaskManager.TransactOpts, _pauserRegistry, initialOwner, _aggregator, _generator)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Pause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Pause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) PauseAll() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.PauseAll(&_ContractMfssiaTaskManager.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) PauseAll() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.PauseAll(&_ContractMfssiaTaskManager.TransactOpts)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xb31c72c6.
//
// Solidity: function raiseAndResolveChallenge((uint32,bytes,string,string,string,uint32) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RaiseAndResolveChallenge(opts *bind.TransactOpts, task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "raiseAndResolveChallenge", task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xb31c72c6.
//
// Solidity: function raiseAndResolveChallenge((uint32,bytes,string,string,string,uint32) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RaiseAndResolveChallenge(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RaiseAndResolveChallenge(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RaiseAndResolveChallenge is a paid mutator transaction binding the contract method 0xb31c72c6.
//
// Solidity: function raiseAndResolveChallenge((uint32,bytes,string,string,string,uint32) task, (uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata, (uint256,uint256)[] pubkeysOfNonSigningOperators) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) RaiseAndResolveChallenge(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, taskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata, pubkeysOfNonSigningOperators []BN254G1Point) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RaiseAndResolveChallenge(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, taskResponseMetadata, pubkeysOfNonSigningOperators)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RenounceOwnership(&_ContractMfssiaTaskManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RenounceOwnership(&_ContractMfssiaTaskManager.TransactOpts)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xaedde731.
//
// Solidity: function respondToTask((uint32,bytes,string,string,string,uint32) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) RespondToTask(opts *bind.TransactOpts, task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "respondToTask", task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xaedde731.
//
// Solidity: function respondToTask((uint32,bytes,string,string,string,uint32) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) RespondToTask(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RespondToTask(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// RespondToTask is a paid mutator transaction binding the contract method 0xaedde731.
//
// Solidity: function respondToTask((uint32,bytes,string,string,string,uint32) task, (uint32,string) taskResponse, (uint32[],(uint256,uint256)[],(uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint32[],uint32[],uint32[][]) nonSignerStakesAndSignature) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) RespondToTask(task IMfssiaTaskManagerTask, taskResponse IMfssiaTaskManagerTaskResponse, nonSignerStakesAndSignature IBLSSignatureCheckerNonSignerStakesAndSignature) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.RespondToTask(&_ContractMfssiaTaskManager.TransactOpts, task, taskResponse, nonSignerStakesAndSignature)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.SetPauserRegistry(&_ContractMfssiaTaskManager.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.SetPauserRegistry(&_ContractMfssiaTaskManager.TransactOpts, newPauserRegistry)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) SetStaleStakesForbidden(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "setStaleStakesForbidden", value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.SetStaleStakesForbidden(&_ContractMfssiaTaskManager.TransactOpts, value)
}

// SetStaleStakesForbidden is a paid mutator transaction binding the contract method 0x416c7e5e.
//
// Solidity: function setStaleStakesForbidden(bool value) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) SetStaleStakesForbidden(value bool) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.SetStaleStakesForbidden(&_ContractMfssiaTaskManager.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.TransferOwnership(&_ContractMfssiaTaskManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.TransferOwnership(&_ContractMfssiaTaskManager.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Unpause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _ContractMfssiaTaskManager.Contract.Unpause(&_ContractMfssiaTaskManager.TransactOpts, newPausedStatus)
}

// ContractMfssiaTaskManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerInitializedIterator struct {
	Event *ContractMfssiaTaskManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerInitialized)
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
		it.Event = new(ContractMfssiaTaskManagerInitialized)
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
func (it *ContractMfssiaTaskManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerInitialized represents a Initialized event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerInitializedIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerInitializedIterator{contract: _ContractMfssiaTaskManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerInitialized)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseInitialized(log types.Log) (*ContractMfssiaTaskManagerInitialized, error) {
	event := new(ContractMfssiaTaskManagerInitialized)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerNewTaskCreatedIterator is returned from FilterNewTaskCreated and is used to iterate over the raw logs and unpacked data for NewTaskCreated events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerNewTaskCreatedIterator struct {
	Event *ContractMfssiaTaskManagerNewTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerNewTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerNewTaskCreated)
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
		it.Event = new(ContractMfssiaTaskManagerNewTaskCreated)
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
func (it *ContractMfssiaTaskManagerNewTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerNewTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerNewTaskCreated represents a NewTaskCreated event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerNewTaskCreated struct {
	TaskIndex uint32
	Task      IMfssiaTaskManagerTask
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewTaskCreated is a free log retrieval operation binding the contract event 0x850fa919096ea51a0971805dcc782e21d54b81f21a8e3c76e9adfac51b678649.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32,bytes,string,string,string,uint32) task)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterNewTaskCreated(opts *bind.FilterOpts, taskIndex []uint32) (*ContractMfssiaTaskManagerNewTaskCreatedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerNewTaskCreatedIterator{contract: _ContractMfssiaTaskManager.contract, event: "NewTaskCreated", logs: logs, sub: sub}, nil
}

// WatchNewTaskCreated is a free log subscription operation binding the contract event 0x850fa919096ea51a0971805dcc782e21d54b81f21a8e3c76e9adfac51b678649.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32,bytes,string,string,string,uint32) task)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchNewTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerNewTaskCreated, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "NewTaskCreated", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerNewTaskCreated)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
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

// ParseNewTaskCreated is a log parse operation binding the contract event 0x850fa919096ea51a0971805dcc782e21d54b81f21a8e3c76e9adfac51b678649.
//
// Solidity: event NewTaskCreated(uint32 indexed taskIndex, (uint32,bytes,string,string,string,uint32) task)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseNewTaskCreated(log types.Log) (*ContractMfssiaTaskManagerNewTaskCreated, error) {
	event := new(ContractMfssiaTaskManagerNewTaskCreated)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "NewTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerOwnershipTransferredIterator struct {
	Event *ContractMfssiaTaskManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerOwnershipTransferred)
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
		it.Event = new(ContractMfssiaTaskManagerOwnershipTransferred)
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
func (it *ContractMfssiaTaskManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractMfssiaTaskManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerOwnershipTransferredIterator{contract: _ContractMfssiaTaskManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerOwnershipTransferred)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ContractMfssiaTaskManagerOwnershipTransferred, error) {
	event := new(ContractMfssiaTaskManagerOwnershipTransferred)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPausedIterator struct {
	Event *ContractMfssiaTaskManagerPaused // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerPaused)
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
		it.Event = new(ContractMfssiaTaskManagerPaused)
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
func (it *ContractMfssiaTaskManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerPaused represents a Paused event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*ContractMfssiaTaskManagerPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerPausedIterator{contract: _ContractMfssiaTaskManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerPaused)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParsePaused(log types.Log) (*ContractMfssiaTaskManagerPaused, error) {
	event := new(ContractMfssiaTaskManagerPaused)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPauserRegistrySetIterator struct {
	Event *ContractMfssiaTaskManagerPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerPauserRegistrySet)
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
		it.Event = new(ContractMfssiaTaskManagerPauserRegistrySet)
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
func (it *ContractMfssiaTaskManagerPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerPauserRegistrySet represents a PauserRegistrySet event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerPauserRegistrySetIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerPauserRegistrySetIterator{contract: _ContractMfssiaTaskManager.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerPauserRegistrySet)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParsePauserRegistrySet(log types.Log) (*ContractMfssiaTaskManagerPauserRegistrySet, error) {
	event := new(ContractMfssiaTaskManagerPauserRegistrySet)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerStaleStakesForbiddenUpdateIterator is returned from FilterStaleStakesForbiddenUpdate and is used to iterate over the raw logs and unpacked data for StaleStakesForbiddenUpdate events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerStaleStakesForbiddenUpdateIterator struct {
	Event *ContractMfssiaTaskManagerStaleStakesForbiddenUpdate // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerStaleStakesForbiddenUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerStaleStakesForbiddenUpdate)
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
		it.Event = new(ContractMfssiaTaskManagerStaleStakesForbiddenUpdate)
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
func (it *ContractMfssiaTaskManagerStaleStakesForbiddenUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerStaleStakesForbiddenUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerStaleStakesForbiddenUpdate represents a StaleStakesForbiddenUpdate event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerStaleStakesForbiddenUpdate struct {
	Value bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaleStakesForbiddenUpdate is a free log retrieval operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterStaleStakesForbiddenUpdate(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerStaleStakesForbiddenUpdateIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerStaleStakesForbiddenUpdateIterator{contract: _ContractMfssiaTaskManager.contract, event: "StaleStakesForbiddenUpdate", logs: logs, sub: sub}, nil
}

// WatchStaleStakesForbiddenUpdate is a free log subscription operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchStaleStakesForbiddenUpdate(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerStaleStakesForbiddenUpdate) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "StaleStakesForbiddenUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerStaleStakesForbiddenUpdate)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
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

// ParseStaleStakesForbiddenUpdate is a log parse operation binding the contract event 0x40e4ed880a29e0f6ddce307457fb75cddf4feef7d3ecb0301bfdf4976a0e2dfc.
//
// Solidity: event StaleStakesForbiddenUpdate(bool value)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseStaleStakesForbiddenUpdate(log types.Log) (*ContractMfssiaTaskManagerStaleStakesForbiddenUpdate, error) {
	event := new(ContractMfssiaTaskManagerStaleStakesForbiddenUpdate)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "StaleStakesForbiddenUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator is returned from FilterTaskChallengedSuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedSuccessfully events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator struct {
	Event *ContractMfssiaTaskManagerTaskChallengedSuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
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
		it.Event = new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
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
func (it *ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskChallengedSuccessfully represents a TaskChallengedSuccessfully event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedSuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedSuccessfully is a free log retrieval operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskChallengedSuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskChallengedSuccessfullyIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskChallengedSuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedSuccessfully is a free log subscription operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskChallengedSuccessfully(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskChallengedSuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskChallengedSuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
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

// ParseTaskChallengedSuccessfully is a log parse operation binding the contract event 0xc20d1bb0f1623680306b83d4ff4bb99a2beb9d86d97832f3ca40fd13a29df1ec.
//
// Solidity: event TaskChallengedSuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskChallengedSuccessfully(log types.Log) (*ContractMfssiaTaskManagerTaskChallengedSuccessfully, error) {
	event := new(ContractMfssiaTaskManagerTaskChallengedSuccessfully)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedSuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator is returned from FilterTaskChallengedUnsuccessfully and is used to iterate over the raw logs and unpacked data for TaskChallengedUnsuccessfully events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator struct {
	Event *ContractMfssiaTaskManagerTaskChallengedUnsuccessfully // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
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
		it.Event = new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
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
func (it *ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskChallengedUnsuccessfully represents a TaskChallengedUnsuccessfully event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskChallengedUnsuccessfully struct {
	TaskIndex  uint32
	Challenger common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTaskChallengedUnsuccessfully is a free log retrieval operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskChallengedUnsuccessfully(opts *bind.FilterOpts, taskIndex []uint32, challenger []common.Address) (*ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskChallengedUnsuccessfullyIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskChallengedUnsuccessfully", logs: logs, sub: sub}, nil
}

// WatchTaskChallengedUnsuccessfully is a free log subscription operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskChallengedUnsuccessfully(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskChallengedUnsuccessfully, taskIndex []uint32, challenger []common.Address) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskChallengedUnsuccessfully", taskIndexRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
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

// ParseTaskChallengedUnsuccessfully is a log parse operation binding the contract event 0xfd3e26beeb5967fc5a57a0446914eabc45b4aa474c67a51b4b5160cac60ddb05.
//
// Solidity: event TaskChallengedUnsuccessfully(uint32 indexed taskIndex, address indexed challenger)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskChallengedUnsuccessfully(log types.Log) (*ContractMfssiaTaskManagerTaskChallengedUnsuccessfully, error) {
	event := new(ContractMfssiaTaskManagerTaskChallengedUnsuccessfully)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskChallengedUnsuccessfully", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskCompletedIterator struct {
	Event *ContractMfssiaTaskManagerTaskCompleted // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskCompleted)
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
		it.Event = new(ContractMfssiaTaskManagerTaskCompleted)
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
func (it *ContractMfssiaTaskManagerTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskCompleted represents a TaskCompleted event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskCompleted struct {
	TaskIndex uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskCompleted(opts *bind.FilterOpts, taskIndex []uint32) (*ContractMfssiaTaskManagerTaskCompletedIterator, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskCompletedIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskCompleted, taskIndex []uint32) (event.Subscription, error) {

	var taskIndexRule []interface{}
	for _, taskIndexItem := range taskIndex {
		taskIndexRule = append(taskIndexRule, taskIndexItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskCompleted", taskIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskCompleted)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0x9a144f228a931b9d0d1696fbcdaf310b24b5d2d21e799db623fc986a0f547430.
//
// Solidity: event TaskCompleted(uint32 indexed taskIndex)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskCompleted(log types.Log) (*ContractMfssiaTaskManagerTaskCompleted, error) {
	event := new(ContractMfssiaTaskManagerTaskCompleted)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerTaskRespondedIterator is returned from FilterTaskResponded and is used to iterate over the raw logs and unpacked data for TaskResponded events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskRespondedIterator struct {
	Event *ContractMfssiaTaskManagerTaskResponded // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerTaskRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerTaskResponded)
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
		it.Event = new(ContractMfssiaTaskManagerTaskResponded)
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
func (it *ContractMfssiaTaskManagerTaskRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerTaskRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerTaskResponded represents a TaskResponded event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerTaskResponded struct {
	TaskResponse         IMfssiaTaskManagerTaskResponse
	TaskResponseMetadata IMfssiaTaskManagerTaskResponseMetadata
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterTaskResponded is a free log retrieval operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterTaskResponded(opts *bind.FilterOpts) (*ContractMfssiaTaskManagerTaskRespondedIterator, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerTaskRespondedIterator{contract: _ContractMfssiaTaskManager.contract, event: "TaskResponded", logs: logs, sub: sub}, nil
}

// WatchTaskResponded is a free log subscription operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchTaskResponded(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerTaskResponded) (event.Subscription, error) {

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "TaskResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerTaskResponded)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
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

// ParseTaskResponded is a log parse operation binding the contract event 0xd0e833257cf2514f7f416921130672e0a5099b141e19ad5285b881ec34e9eb54.
//
// Solidity: event TaskResponded((uint32,string) taskResponse, (uint32,bytes32) taskResponseMetadata)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseTaskResponded(log types.Log) (*ContractMfssiaTaskManagerTaskResponded, error) {
	event := new(ContractMfssiaTaskManagerTaskResponded)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "TaskResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractMfssiaTaskManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerUnpausedIterator struct {
	Event *ContractMfssiaTaskManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *ContractMfssiaTaskManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractMfssiaTaskManagerUnpaused)
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
		it.Event = new(ContractMfssiaTaskManagerUnpaused)
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
func (it *ContractMfssiaTaskManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractMfssiaTaskManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractMfssiaTaskManagerUnpaused represents a Unpaused event raised by the ContractMfssiaTaskManager contract.
type ContractMfssiaTaskManagerUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*ContractMfssiaTaskManagerUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractMfssiaTaskManagerUnpausedIterator{contract: _ContractMfssiaTaskManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ContractMfssiaTaskManagerUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ContractMfssiaTaskManager.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractMfssiaTaskManagerUnpaused)
				if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_ContractMfssiaTaskManager *ContractMfssiaTaskManagerFilterer) ParseUnpaused(log types.Log) (*ContractMfssiaTaskManagerUnpaused, error) {
	event := new(ContractMfssiaTaskManagerUnpaused)
	if err := _ContractMfssiaTaskManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
