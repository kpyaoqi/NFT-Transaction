package test

import (
	"fmt"
	"testing"
	"yqnft/ContractManage/serviceImpl"
)

var ser serviceImpl.ContractService

func TestCreateContract(t *testing.T) {
	result := ser.CreateContract("NFT1", "NFT", "YAOQI", "1.0.0")
	fmt.Println(result)
}
func TestMintNFT(t *testing.T) {
	result, tokenId := ser.MintNft("NFT1", "0x3a61dd3fb25a7ee3e088b6390b9bc66b299b15c1", "YAOQI")
	fmt.Println(result, tokenId)
}
func TestNFTUrl(t *testing.T) {
	result := ser.DetailsNft("NFT1", 4)
	fmt.Println(result)
}
