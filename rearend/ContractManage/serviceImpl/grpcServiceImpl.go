package serviceImpl

import (
	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/common/v2/evmutils/abi"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	"chainmaker.org/chainmaker/sdk-go/v2/examples"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
	"yqnft/ContractManage/utils"
)

const (
	createContractTimeout    = 5
	withSyncResult           = true
	sdkConfigOrg1Client1Path = "D:/OpenNFT/NFT交易平台/rearend/ContractManage/config/sdk_config1.yml"
	nftByteCodePath          = "D:/OpenNFT/NFT交易平台/rearend/ContractManage/contract_file/nft/nftmanger.bin"
	nftABIPath               = "D:/OpenNFT/NFT交易平台/rearend/ContractManage/contract_file/nft/nftmanger.abi"
)

var client, _ = examples.CreateChainClientWithSDKConf(sdkConfigOrg1Client1Path)
var usernames = [...]string{utils.UserNameOrg1Admin1, utils.UserNameOrg2Admin1, utils.UserNameOrg3Admin1}
var abiJson, _ = ioutil.ReadFile(nftABIPath)
var myAbi, _ = abi.JSON(strings.NewReader(string(abiJson)))

type ContractService struct {
}

func (con ContractService) CreateContract(contractName string, name string, symbol string, version string) string {
	dataByte, err := myAbi.Pack("", name, symbol)
	if err != nil {
		return "交易失败"
	}
	data := hex.EncodeToString(dataByte)
	pairs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(data),
		},
	}
	payload, err := client.CreateContractCreatePayload(contractName, version, nftByteCodePath, common.RuntimeType_EVM, pairs)
	if err != nil {
		return "交易失败"
	}
	endorsers, err := utils.GetEndorsersWithAuthType(client.GetHashType(),
		client.GetAuthType(), payload, usernames[:]...)
	if err != nil {
		return "交易失败"
	}
	// 发送创建合约请求
	resp, err := client.SendContractManageRequest(payload, endorsers, createContractTimeout, withSyncResult)
	if err != nil {
		return "交易失败"
	}
	err = utils.CheckProposalRequestResp(resp, true)
	if err != nil {
		return "交易失败"
	}
	fmt.Println("创建" + contractName + "合约成功")
	return "交易成功"
}

func (con ContractService) MintNft(contractName string, addressTo string, tokenURL string) (string, int64) {
	//address := evmutils.BigToAddress(evmutils.FromDecimalString(addressTo))
	address := addressTo
	method := "mint"
	dataByte, err := myAbi.Pack(method, address, tokenURL)
	if err != nil {
		return "交易失败", 0
	}
	dataString := hex.EncodeToString(dataByte)
	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	result, err := utils.InvokeUserContractWithResult(client, contractName, method, "", kvs, withSyncResult)
	if err != nil {
		return "交易失败", 0
	}
	tokenId, err := myAbi.Unpack(method, result)
	parseInt, _ := strconv.ParseInt(tokenId[0].(string), 10, 64)
	fmt.Println(contractName + "合约铸造NFT成功")
	return "交易成功", parseInt
}

func (con ContractService) TransferNft(contractName string, addressFrom string, addressTo string, tokenId int64) string {
	//addressfrom := evmutils.BigToAddress(evmutils.FromDecimalString(addressFrom))
	//addressto := evmutils.BigToAddress(evmutils.FromDecimalString(addressTo))
	method := "transferFrom"
	dataByte, err := myAbi.Pack(method, addressFrom, addressTo, big.NewInt(tokenId))
	if err != nil {
		return "交易失败"
	}
	dataString := hex.EncodeToString(dataByte)
	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	err = utils.InvokeUserContract(client, contractName, method, "", kvs, withSyncResult)
	if err != nil {
		return "交易失败"
	}
	fmt.Println(contractName + "合约交易NFT成功：" + addressFrom + "---->" + addressTo)
	return "交易成功"
}

func (con ContractService) DetailsNft(contractName string, tokenId int64) string {
	method := "tokenURI"
	dataByte, err := myAbi.Pack(method, tokenId)
	if err != nil {
		return "交易失败"
	}
	dataString := hex.EncodeToString(dataByte)
	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	//result, err := utils.InvokeUserContractWithResult(client, contractName, method, "", kvs, withSyncResult)
	result, err := client.QueryContract(contractName, method, kvs, -1)

	if err != nil {
		return "交易失败"
	}
	res, err := myAbi.Unpack(method, result.ContractResult.Result)
	if err != nil {
		return "交易失败"
	}
	return res[0].(string)
}

func (con ContractService) ApproveNft(contractName string, addressTo string, tokenId int64) string {
	addressto := evmutils.BigToAddress(evmutils.FromDecimalString(addressTo))
	method := "approve"
	dataByte, err := myAbi.Pack(method, addressto, tokenId)
	if err != nil {
		return "交易失败"
	}
	dataString := hex.EncodeToString(dataByte)
	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	err = utils.InvokeUserContract(client, contractName, method, "", kvs, withSyncResult)
	if err != nil {
		return "交易失败"
	}
	return ""
}

func (con ContractService) DetailsNftContract(contractName string) []string {
	var rseults []string
	method := "name"
	dataByte, _ := myAbi.Pack(method)
	dataString := hex.EncodeToString(dataByte)
	kvs := []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	result, _ := client.QueryContract(contractName, method, kvs, -1)
	nameRes, _ := myAbi.Unpack(method, result.ContractResult.Result)
	rseults = append(rseults, nameRes[0].(string))

	method = "symbol"
	dataByte, _ = myAbi.Pack(method)
	dataString = hex.EncodeToString(dataByte)
	kvs = []*common.KeyValuePair{
		{
			Key:   "data",
			Value: []byte(dataString),
		},
	}
	result, _ = client.QueryContract(contractName, method, kvs, -1)
	symbolRes, _ := myAbi.Unpack(method, result.ContractResult.Result)
	rseults = append(rseults, symbolRes[0].(string))
	return rseults
}
