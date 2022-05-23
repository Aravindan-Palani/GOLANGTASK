package main

import (
        "fmt"
        "github.com/go-redis/redis"
)

func main() {
        fmt.Println("Go Redis Tutorial")

        client := redis.NewClient(&redis.Options{
                Addr: "localhost:6379",
                Password: "",
                DB: 0,
        })

        pong, err := client.Ping().Result()
        fmt.Println(pong, err)

        err1 := client.Set("name", "Elliot", 0).Err()
        if err1 != nil {
                fmt.Println(err1)
        }

        val, err := client.Get("name").Result()
        if err != nil {
                fmt.Println(err)
        }

        fmt.Println(val)
}

