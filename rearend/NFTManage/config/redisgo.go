package config

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
	"yqnft/NFTManage/models/model"
	apiResult "yqnft/NFTManage/models/request_result/api"
)

var RedisClient = &redis.Pool{
	MaxIdle:     100,             //最大连接数，即最多的tcp连接数
	MaxActive:   10,              // MaxIdle 最大空闲连接数
	IdleTimeout: 5 * time.Second, // IdleTimeout 空闲连接超时时间，但应该设置比redis服务器超时时间短。否则服务端超时了，客户端保持着连接也没用。
	Wait:        true,
	Dial: func() (redis.Conn, error) {
		con, err := redis.Dial("tcp", "127.0.0.1:6379",
			//redis.DialPassword("123456"),
			redis.DialDatabase(0),
			redis.DialConnectTimeout(3*time.Second),
			redis.DialReadTimeout(5*time.Second),
			redis.DialWriteTimeout(5*time.Second))
		if err != nil {
			return nil, err
		}
		return con, nil
	},
}

func SetRedis(key string, data []byte) error {
	DS := RedisClient.Get()
	defer DS.Close()
	_, err := DS.Do("SET", key, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func SetRedisWithTime(key string, data []byte, expiration_time int) {
	DS := RedisClient.Get()
	defer DS.Close()
	DS.Do("SET", key, data)
	DS.Do("expire", key, expiration_time)
}

func DelRedis(key string) {
	DS := RedisClient.Get()
	defer DS.Close()
	DS.Do("DEL", key)
}

func GetNFTMarketDTOListRedis(key string) []apiResult.NFTMarketDTO {
	DS := RedisClient.Get()
	defer DS.Close()
	var results []apiResult.NFTMarketDTO
	t_data, _ := redis.Bytes(DS.Do("GET", key))
	json.Unmarshal(t_data, &results)
	return results
}

func GetNftWorksRedis(key string) model.NftWork {
	DS := RedisClient.Get()
	defer DS.Close()
	var results model.NftWork
	t_data, _ := redis.Bytes(DS.Do("GET", key))
	json.Unmarshal(t_data, &results)
	return results
}

func GetNftCollectionRedis(key string) model.NftCollection {
	DS := RedisClient.Get()
	defer DS.Close()
	var results model.NftCollection
	t_data, _ := redis.Bytes(DS.Do("GET", key))
	json.Unmarshal(t_data, &results)
	return results
}
