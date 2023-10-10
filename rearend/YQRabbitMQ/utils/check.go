package MQUtils

import (
	"strings"
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/config"
)

var MqMessage = map[string]string{
	"ChangeNftWorkByIdMessage":       common.WorksByID,
	"ChangeNftCollectionByIdMessage": common.CollectionByID,
}

func ChangeRedis(message []byte) string {
	parts := strings.Split(string(message), "_")
	prefix := MqMessage[parts[0]]
	if len(parts) == 2 {
		config.DelRedis(prefix + parts[1])
		return prefix + parts[1]
	} else {
		config.DelRedis(prefix)
		return prefix
	}
}
