package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	rdb *redis.Client
)

func main() {
	V8Example()
}

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	_, err = rdb.Ping().Result()
	return err
}

func V8Example() {
	//ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}

	err := rdb.Set("key", "value", time.Second*60).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get("key2").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
