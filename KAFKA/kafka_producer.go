package main

import (
        "github.com/segmentio/kafka-go"
        "fmt"
        "time"
        "context"
        "strconv"
)

const (
        topic          = "test"
        broker1Address = "localhost:9092"
)

func produce(ctx context.Context) {
        i := 0
        var name string

        w := kafka.NewWriter(kafka.WriterConfig{
                Brokers: []string{broker1Address},
                Topic:   topic,
        })

        for {
                fmt.Scan(&name)
                err := w.WriteMessages(ctx, kafka.Message{
                        Key: []byte(strconv.Itoa(i)),
                        Value: []byte("this is message  " + strconv.Itoa(i) +  "  name :" + name),
                })
                if err != nil {
                        panic("could not write message " + err.Error())
                }

                fmt.Println("writes: ", i , name)
                i++
                time.Sleep(time.Second)
        }
}

func main() {
        ctx := context.Background()
        produce(ctx)
}
