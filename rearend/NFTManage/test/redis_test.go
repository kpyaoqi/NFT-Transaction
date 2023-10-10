package test

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"yqnft/NFTManage/common"
	"yqnft/NFTManage/config"
	"yqnft/NFTManage/service"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var people = []Person{
	{"John", "Doe", 30},
	{"Alice", "Johnson", 25},
	{"Bob", "Smith", 35},
}
var DS = config.RedisClient.Get()

func TestRedisString(t *testing.T) {
	var res []Person
	//json序列化
	datas, _ := json.Marshal(people)
	//缓存数据
	DS.Do("set", "struct3", datas)
	//读取数据
	rebytes, _ := redis.Bytes(DS.Do("GET", "struct3"))
	//json反序列化
	json.Unmarshal(rebytes, &res)

}

func TestRedis(t *testing.T) {
	//以hash类型保存
	DS.Do("hmset", redis.Args{"struct1"}.AddFlat(people)...)
	//获取缓存
	value, _ := redis.Values(DS.Do("hgetall", "struct1"))
	//将values转成结构体
	object := &[]Person{}
	redis.ScanStruct(value, object)
	fmt.Println(object)
}

func TestName(t *testing.T) {
	collectionID := "20230908155342317331776148"
	var colservice service.NftCollectionService
	nftCollection := colservice.FindById(collectionID)
	data, _ := json.Marshal(nftCollection)
	// 设置缓存
	config.SetRedis(common.CollectionByID+collectionID, data)
}
func TestName2(t *testing.T) {
	collectionID := "20230908155342317331776148"
	//获取缓存
	collectionRedis := config.GetNftCollectionRedis(common.CollectionByID + collectionID)
	fmt.Println(collectionRedis)
}
