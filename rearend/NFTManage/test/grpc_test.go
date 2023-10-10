package test

import (
	"context"
	"fmt"
	"testing"
	"yqnft/NFTManage/config"
	pb "yqnft/proto"
)

var client = config.GrpcConn()

func TestCreateContract(t *testing.T) {
	req := new(pb.CreateContractRequest)
	req.ContractName = "test3"
	req.Name = "YNFT"
	req.Symbol = "YAOQI"
	req.Version = "1.0.0"
	response, err := client.CreateContract(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("响应结果： %v\n", response)
}
func TestMintNFT(t *testing.T) {
	req := new(pb.MintNftRequest)
	req.ContractName = "test3"
	req.AddressTo = "0x3a61dd3fb25a7ee3e088b6390b9bc66b299b15c1"
	req.TokenUrl = "YAOQI"
	result, tokenId := client.MintNft(context.Background(), req)
	fmt.Println(result, tokenId)
}
func TestNFTUrl(t *testing.T) {
	req := new(pb.DetailsNftRequest)
	req.ContractName = "test4"
	req.TokenId = 1
	result, _ := client.DetailsNft(context.Background(), req)
	fmt.Println(result)
}
