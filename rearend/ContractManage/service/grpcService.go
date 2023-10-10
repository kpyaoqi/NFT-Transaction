package service

import (
	"context"
	"yqnft/ContractManage/serviceImpl"
	pb "yqnft/proto"
)

type ContractService struct{}

var service serviceImpl.ContractService

func (s *ContractService) CreateContract(ctx context.Context, req *pb.CreateContractRequest) (resp *pb.TransactionResponse, err error) {
	resp = new(pb.TransactionResponse)
	contractName := req.ContractName
	name := req.Name
	symbol := req.Symbol
	version := req.Version
	resp.Status = service.CreateContract(contractName, name, symbol, version)
	return resp, nil
}
func (s *ContractService) MintNft(ctx context.Context, req *pb.MintNftRequest) (resp *pb.MintNftResponse, err error) {
	resp = new(pb.MintNftResponse)
	contractName := req.ContractName
	addressTo := req.AddressTo
	tokenUrl := req.TokenUrl
	resp.Status, resp.TokenId = service.MintNft(contractName, addressTo, tokenUrl)
	return resp, nil
}
func (s *ContractService) TransferNft(ctx context.Context, req *pb.TransferNftRequest) (resp *pb.TransactionResponse, err error) {
	resp = new(pb.TransactionResponse)
	contractName := req.ContractName
	addressTo := req.AddressFrom
	addressFrom := req.AddressTo
	tokenId := req.TokenId
	resp.Status = service.TransferNft(contractName, addressTo, addressFrom, tokenId)
	return resp, nil
}
func (s *ContractService) ApproveNft(ctx context.Context, req *pb.ApproveNftRequest) (resp *pb.TransactionResponse, err error) {
	resp = new(pb.TransactionResponse)
	contractName := req.ContractName
	addressTo := req.AddressTo
	tokenId := req.TokenId
	resp.Status = service.ApproveNft(contractName, addressTo, tokenId)
	return resp, nil
}
func (s *ContractService) DetailsNft(ctx context.Context, req *pb.DetailsNftRequest) (resp *pb.DetailsNftResponse, err error) {
	resp = new(pb.DetailsNftResponse)
	contractName := req.ContractName
	tokenId := req.TokenId
	resp.TokenUrl = service.DetailsNft(contractName, tokenId)
	return resp, nil
}
func (s *ContractService) DetailsNftContract(ctx context.Context, req *pb.DetailsNftContractRequest) (resp *pb.DetailsNftContractResponse, err error) {
	resp = new(pb.DetailsNftContractResponse)
	contractName := req.ContractName
	resp.Details = service.DetailsNftContract(contractName)
	return resp, nil
}
