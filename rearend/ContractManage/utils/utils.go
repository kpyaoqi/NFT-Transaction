package utils

import (
	"chainmaker.org/chainmaker/common/v2/crypto"
	bcx509 "chainmaker.org/chainmaker/common/v2/crypto/x509"
	"chainmaker.org/chainmaker/common/v2/evmutils"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	sdkutils "chainmaker.org/chainmaker/sdk-go/v2/utils"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"io/ioutil"
)

const (
	OrgId1 = "TestCMorg1"
	OrgId2 = "TestCMorg2"
	OrgId3 = "TestCMorg3"
	OrgId4 = "TestCMorg4"

	UserNameOrg1Client1 = "cmtestnode1"
	UserNameOrg2Client1 = "cmtestnode2"
	UserNameOrg3Client1 = "cmtestnode3"
	UserNameOrg4Client1 = "cmtestnode4"

	UserNameOrg1Admin1 = "cmtestuser1"
	UserNameOrg2Admin1 = "cmtestuser2"
	UserNameOrg3Admin1 = "cmtestuser3"
	UserNameOrg4Admin1 = "cmtestuser4"

	Version        = "1.0.0"
	UpgradeVersion = "2.0.0"
)

var users = map[string]*User{
	"cmtestnode1": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/client/cmtestnode1.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/client/cmtestnode1.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/client/cmtestnode1.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/client/cmtestnode1.sign.crt",
	},
	"cmtestnode2": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/client/cmtestnode2.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/client/cmtestnode2.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/client/cmtestnode2.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/client/cmtestnode2.sign.crt",
	},
	"cmtestnode3": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/client/cmtestnode3.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/client/cmtestnode3.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/client/cmtestnode3.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/client/cmtestnode3.sign.crt",
	},
	"cmtestnode4": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/client/cmtestnode4.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/client/cmtestnode4.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/client/cmtestnode4.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/client/cmtestnode4.sign.crt",
	},
	"cmtestuser1": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/admin/cmtestuser1.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/admin/cmtestuser1.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/admin/cmtestuser1.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg1/admin/cmtestuser1.sign.crt",
	},
	"cmtestuser2": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/admin/cmtestuser2.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/admin/cmtestuser2.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/admin/cmtestuser2.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg2/admin/cmtestuser2.sign.crt",
	},

	"cmtestuser3": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/admin/cmtestuser3.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/admin/cmtestuser3.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/admin/cmtestuser3.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg3/admin/cmtestuser3.sign.crt",
	},
	"cmtestuser4": {
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/admin/cmtestuser4.tls.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/admin/cmtestuser4.tls.crt",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/admin/cmtestuser4.sign.key",
		"F:/MyGopath/src/yqnft/ContractManage/config/crypto-config/TestCMorg4/admin/cmtestuser4.sign.crt",
	},
}

var permissionedPkUsers = map[string]*PermissionedPkUsers{
	"org1client1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/permissioned-with-key/wx-org1/user/client1/client1.key",
		OrgId1,
	},
	"org2client1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/permissioned-with-key/wx-org2/user/client1/client1.key",
		OrgId2,
	},
	"org1admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/permissioned-with-key/wx-org1/user/admin1/admin1.key",
		OrgId1,
	},
	"org2admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/permissioned-with-key/wx-org2/user/admin1/admin1.key",
		OrgId2,
	},
	"org3admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/permissioned-with-key/wx-org3/user/admin1/admin1.key",
		OrgId3,
	},
	"org4admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/permissioned-with-key/wx-org4/user/admin1/admin1.key",
		OrgId4,
	},
}

var pkUsers = map[string]*PkUsers{
	"org1client1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/public/user/user1/user1.key",
	},
	"org2client1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/public/user/user2/user2.key",
	},
	"org1admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/public/admin/admin1/admin1.key",
	},
	"org2admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/public/admin/admin2/admin2.key",
	},
	"org3admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/public/admin/admin3/admin3.key",
	},
	"org4admin1": {
		"C:/Users/zhuba/Desktop/sdk-go/testdata/crypto-config-pk/public/admin/admin4/admin4.key",
	},
}

type PkUsers struct {
	SignKeyPath string
}

type PermissionedPkUsers struct {
	SignKeyPath string
	OrgId       string
}

type User struct {
	TlsKeyPath, TlsCrtPath   string
	SignKeyPath, SignCrtPath string
}

func GetUser(username string) (*User, error) {
	u, ok := users[username]
	if !ok {
		return nil, errors.New("user not found")
	}

	return u, nil
}

func CheckProposalRequestResp(resp *common.TxResponse, needContractResult bool) error {
	if resp.Code != common.TxStatusCode_SUCCESS {
		if resp.Message == "" {
			resp.Message = resp.Code.String()
		}
		return errors.New(resp.Message)
	}

	if needContractResult && resp.ContractResult == nil {
		return fmt.Errorf("contract result is nil")
	}

	if resp.ContractResult != nil && resp.ContractResult.Code != 0 {
		return errors.New(resp.ContractResult.Message)
	}

	return nil
}

// CreateChainClientWithSDKConf create a chain client with sdk config file path
func CreateChainClientWithSDKConf(sdkConfPath string) (*sdk.ChainClient, error) {
	return sdk.NewChainClient(
		sdk.WithConfPath(sdkConfPath),
	)
}

func CalcContractName(contractName string) string {
	return hex.EncodeToString(evmutils.Keccak256([]byte(contractName)))[24:]
}

func GetEndorsers(payload *common.Payload, usernames ...string) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	for _, name := range usernames {
		u, ok := users[name]
		if !ok {
			return nil, errors.New("user not found")
		}

		var err error
		var entry *common.EndorsementEntry
		p11Handle := sdk.GetP11Handle()
		if p11Handle != nil {
			entry, err = sdkutils.MakeEndorserWithPathAndP11Handle(u.SignKeyPath, u.SignCrtPath, p11Handle, payload)
			if err != nil {
				return nil, err
			}
		} else {
			entry, err = sdkutils.MakeEndorserWithPath(u.SignKeyPath, u.SignCrtPath, payload)
			if err != nil {
				return nil, err
			}
		}

		endorsers = append(endorsers, entry)
	}

	return endorsers, nil
}

func GetEndorsersWithAuthType(hashType crypto.HashType, authType sdk.AuthType, payload *common.Payload, usernames ...string) ([]*common.EndorsementEntry, error) {
	var endorsers []*common.EndorsementEntry

	for _, name := range usernames {
		var entry *common.EndorsementEntry
		var err error
		switch authType {
		case sdk.PermissionedWithCert:
			u, ok := users[name]
			if !ok {
				return nil, errors.New("user not found")
			}
			entry, err = sdkutils.MakeEndorserWithPath(u.SignKeyPath, u.SignCrtPath, payload)
			if err != nil {
				return nil, err
			}

		case sdk.PermissionedWithKey:
			u, ok := permissionedPkUsers[name]
			if !ok {
				return nil, errors.New("user not found")
			}
			entry, err = sdkutils.MakePkEndorserWithPath(u.SignKeyPath, hashType, u.OrgId, payload)
			if err != nil {
				return nil, err
			}

		case sdk.Public:
			u, ok := pkUsers[name]
			if !ok {
				return nil, errors.New("user not found")
			}
			entry, err = sdkutils.MakePkEndorserWithPath(u.SignKeyPath, hashType, "", payload)
			if err != nil {
				return nil, err
			}

		default:
			return nil, errors.New("invalid authType")
		}
		endorsers = append(endorsers, entry)
	}

	return endorsers, nil
}

func MakeAddrAndSkiFromCrtFilePath(crtFilePath string) (string, string, string, error) {
	crtBytes, err := ioutil.ReadFile(crtFilePath)
	if err != nil {
		return "", "", "", err
	}

	blockCrt, _ := pem.Decode(crtBytes)
	crt, err := bcx509.ParseCertificate(blockCrt.Bytes)
	if err != nil {
		return "", "", "", err
	}

	ski := hex.EncodeToString(crt.SubjectKeyId)
	addrInt, err := evmutils.MakeAddressFromHex(ski)
	if err != nil {
		return "", "", "", err
	}

	fmt.Sprintf("0x%s", addrInt.AsStringKey())

	return addrInt.String(), fmt.Sprintf("0x%x", addrInt.AsStringKey()), ski, nil
}

func PrintPrettyJson(data interface{}) {
	output, err := prettyjson.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}
