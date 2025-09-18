package main

import (
    "context"
    "fmt"
    "log"

    "github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/redis/go-redis/v9"
)

func main() {
    rdb := redis.NewClient(&redis.Options{Addr: "redis:6379"})
    reader, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
        "group.id":          "chronoslm-group",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer reader.Close()
    reader.Subscribe("multimodal-stream", nil)

    ctx := context.Background()
    for {
        msg, err := reader.ReadMessage(-1)
        if err == nil {
            fmt.Printf("Message: %s\n", string(msg.Value))
            rdb.LPush(ctx, "hotmemory", msg.Value)
        } else {
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }
    }
}
