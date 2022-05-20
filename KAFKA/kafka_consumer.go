package main

import (
        "github.com/segmentio/kafka-go"
        "fmt"
        "context"
)

const (
        topic          = "test"
        broker1Address = "localhost:9092"
)

func consume(ctx context.Context) {
        r := kafka.NewReader(kafka.ReaderConfig{
                Brokers: []string{broker1Address},
                Topic:   topic,
                GroupID: "mygroup",
        })
        for {
                msg, err := r.ReadMessage(ctx)
                if err != nil {
                        panic("could not read message " + err.Error())
                }
                fmt.Println("received: ", string(msg.Value))
        }
}

func main() {
        ctx := context.Background()
        consume(ctx)
}

