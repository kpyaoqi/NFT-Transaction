package config

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
	pb "yqnft/proto"
)

var GrpcClient pb.ContractServiceClient

type GrpcMethod struct {
}

func init() {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:7080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
	}
	//defer conn.Close()
	// 2. 实例化gRPC客户端
	GrpcClient = pb.NewContractServiceClient(conn)
}

func GrpcConn() pb.ContractServiceClient {
	// 1.连接
	conn, err := grpc.Dial("127.0.0.1:7080", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
	}
	//defer conn.Close()
	// 2. 实例化gRPC客户端
	client := pb.NewContractServiceClient(conn)
	return client
}

// 部署合约
func (grpc GrpcMethod) CreateContract(userId string, nftCollectionVO apiResult.NftCollectionVO) string {
	req := new(pb.CreateContractRequest)
	req.ContractName = userId + "_" + nftCollectionVO.Name
	req.Name = nftCollectionVO.Name
	req.Symbol = nftCollectionVO.Symbol
	req.Version = "1.0.0"
	res, err := GrpcClient.CreateContract(context.Background(), req)
	if err != nil {
		return err.Error()
	}
	return res.Status
}

func (grpc GrpcMethod) TransferNft(works model.NftWork, collection model.NftCollection, fromUser string, toUser string) string {
	req := new(pb.TransferNftRequest)
	req.ContractName = collection.UserID + "_" + collection.Name
	req.AddressFrom = fromUser
	req.AddressFrom = "0x91664c03c49ad7f38c8d20835b5ebf535dc42fb0"
	req.AddressTo = toUser
	req.TokenId = works.TokenID
	res, err := GrpcClient.TransferNft(context.Background(), req)
	if err != nil {
		return err.Error()
	}
	return res.Status
}

func (grpc GrpcMethod) MintNft(path string, address string, collection model.NftCollection) (int64, string) {
	req := new(pb.MintNftRequest)
	req.ContractName = collection.UserID + "_" + collection.Name
	req.AddressTo = address
	req.AddressTo = "0x91664c03c49ad7f38c8d20835b5ebf535dc42fb0"
	req.TokenUrl = path
	nft, err := GrpcClient.MintNft(context.Background(), req)
	if err != nil {
		return 0, err.Error()
	}
	return nft.TokenId, nft.Status
}
