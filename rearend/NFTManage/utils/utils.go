package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func EncoderByMd5With32Bit(instr string) string {
	hexDigits := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	if instr != "" {
		strTemp := []byte(instr)
		mdTemp := md5.New()
		mdTemp.Write(strTemp)
		md := mdTemp.Sum(nil)
		str := make([]byte, len(md)*2)
		k := 0
		for i := 0; i < len(md); i++ {
			byte0 := md[i]
			str[k] = hexDigits[byte0>>4&0xf]
			str[k+1] = hexDigits[byte0&0xf]
			k += 2
		}
		return string(str)
	} else {
		return ""
	}
}

func CalculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	md5Sum := hash.Sum(nil)
	md5Hex := hex.EncodeToString(md5Sum)
	return md5Hex, nil
}

func GetTimeNO() string {
	currentTime := time.Now()
	timeNo := currentTime.Format("20060102150405") // Format: YYYYMMDDHHmmss
	randomInt := rand.Intn(1000000000000)          // Generate a random integer
	timeNoWithRandom := fmt.Sprintf("%s%012d", timeNo, randomInt)
	return timeNoWithRandom
}

func GetUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func PhoneToEVMAddress(phoneNumber string) string {
	// 使用SHA-256哈希手机号码
	hasher := sha256.New()
	hasher.Write([]byte(phoneNumber))
	hashedPhoneNumberBytes := hasher.Sum(nil)
	hashedPhoneNumber := hex.EncodeToString(hashedPhoneNumberBytes)
	// 添加一个前缀以标识这是一个手机号码生成的地址（可选步骤）
	addressWithPrefix := "0x" + hashedPhoneNumber
	// 取前 40 个字符，这将成为最终的EVM地址
	evmAddress := addressWithPrefix[:42]
	return evmAddress
}
