package test

import (
	"fmt"
	"testing"
	"yqnft/NFTManage/utils"
)

func TestPhoneToEVMAddress(t *testing.T) {
	evmAddress := utils.PhoneToEVMAddress("18312386077")
	fmt.Println(evmAddress)
}
