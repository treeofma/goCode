package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

func main() {
    // 创建Redis客户端
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis服务器地址
        Password: "",               // 密码，如果没有密码则为空字符串
        DB:       0,                // 使用的数据库编号
    })

    // Ping测试连接
    pong, err := client.Ping(context.Background()).Result()
    fmt.Println(pong, err)

    // 设置键值对
    err = client.Set(context.Background(), "key", "value", 0).Err()
    if err != nil {
        fmt.Println(err)
    }

    // 获取键值对
    val, err := client.Get(context.Background(), "key").Result()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("key:", val)
}

