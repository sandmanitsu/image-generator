package main

import (
	"context"
	"fmt"
	"generator/config"
	"generator/kafka"
)

func main() {
	config := config.MustLoad()
	fmt.Println(config)

	producer := kafka.NewProducer(config)

	msgs := CreateMessage()

	ctx := context.Background()
	for _, msg := range msgs.Data {
		producer.WriteMesage(ctx, msg.Template, msg.Payload)
	}
}
