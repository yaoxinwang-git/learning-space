package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.37.160:6379",
		Password: "root123", // no password set
		DB:       0,         // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func main() {
	initClient()

	rdb.Set("yxw", "test", time.Hour)

}
