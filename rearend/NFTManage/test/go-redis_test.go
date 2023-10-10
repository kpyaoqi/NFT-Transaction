package test

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:         "127.0.0.1:6379",
	Password:     "", // no password set
	DB:           0,  // use default DB
	MinIdleConns: 1,
	PoolSize:     1000,
})

func init() {
	ctx := context.Background()
	res, err := rdb.Ping(ctx).Result()
	fmt.Printf("ping res=%v\n", res)
	if err != nil {
		fmt.Printf("ping err=%v\n", err.Error())
	}
}

var _ encoding.BinaryMarshaler = new(User)
var _ encoding.BinaryUnmarshaler = new(User)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

var users = []User{
	{"abc", 10},
	{"efg", 20},
}

var user = &User{
	"abc", 10,
}
var ctx = context.Background()

func Struct1() {
	rdb.Set(ctx, "key1", user, 0)
	result := &User{}
	rdb.Get(ctx, "key1").Scan(result)
	fmt.Println(result)
}

func Struct2() {
	b, _ := json.Marshal(user)
	rdb.Set(ctx, "key1", string(b), 0)
	res, _ := rdb.Get(ctx, "key1").Result()
	var user User
	// 使用json.Unmarshal解码JSON数组到结构体切片
	json.Unmarshal([]byte(res), &user)
	//fmt.Println(user)
}

func Struct3() {
	b, _ := json.Marshal(users)
	rdb.Set(ctx, "key1", string(b), 0)
	res, _ := rdb.Get(ctx, "key1").Result()
	var user []User
	// 使用json.Unmarshal解码JSON数组到结构体切片
	json.Unmarshal([]byte(res), &user)
	//fmt.Println(user)
}

func Struct4() {
	for i := 0; i < len(users); i++ {
		rdb.SAdd(ctx, "key2", &users[i])
	}
	members, _ := rdb.SMembers(ctx, "key2").Result()
	var users []User
	// 循环遍历JSON字符串切片并解码每个JSON字符串
	for _, jsonStr := range members {
		var user User
		err := json.Unmarshal([]byte(jsonStr), &user)
		if err != nil {
			fmt.Println("解码JSON失败：", err)
			return
		}
		users = append(users, user)
	}
	//fmt.Println(users)
}
func BenchmarkSplit1(b *testing.B) {
	Struct1()
}
func BenchmarkSplit2(b *testing.B) {
	Struct2()
}
func BenchmarkSplit3(b *testing.B) {
	Struct3()
}
func BenchmarkSplit4(b *testing.B) {
	Struct4()
}

func TestString(t *testing.T) {
	res, err := rdb.Set(ctx, "name", "测试001", 0).Result()
	fmt.Printf("set res=%v\n", res)
	if err != nil {
		fmt.Printf("set err=%v\n", err.Error())
	}
	res, _ = rdb.Get(ctx, "name").Result()
	fmt.Println(res)

}
