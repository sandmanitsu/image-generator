package kafka

import (
	"context"
	"creator/internal/config"
	"creator/internal/semaphore"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type ImageService interface {
	ProcessImages(template string, payload []byte)
}

const (
	maxConc = 5
)

func StartConsumer(cfg *config.Config, imageSrv ImageService) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.KafkaBroker},
		Topic:    cfg.KafkaTopic,
		GroupID:  "image-create",
		MaxBytes: 10e6,
		// Logger:      kafka.LoggerFunc(logf),
		ErrorLogger: kafka.LoggerFunc(logf),
	})

	ctx := context.Background()

	log.Println("consumer started!")
	sem := semaphore.NewSemaphore(maxConc)
	go func() {
		for {
			m, err := r.FetchMessage(ctx)
			if err != nil {
				log.Println("error fetch message", err)
				continue
			}

			sem.Acquire()
			go func(msg kafka.Message) {
				defer sem.Release()

				imageSrv.ProcessImages(string(msg.Key), msg.Value)
			}(m)

			// if err = r.CommitMessages(ctx, m); err != nil {
			// 	log.Println("error commit message", err)
			// }
		}
	}()
}

func logf(msg string, a ...any) {
	fmt.Printf(msg, a...)
	fmt.Println()
}
