package kafka

import (
	"context"
	"fmt"
	"generator/config"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	w *kafka.Writer
}

func NewProducer(cfg *config.Config) *Producer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(cfg.KafkaBroker),
		Topic:    cfg.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
		Logger:   kafka.LoggerFunc(logf),
	}

	return &Producer{
		w: writer,
	}
}

func (p *Producer) WriteMesage(ctx context.Context, tamplete string, payload []byte) error {
	err := p.w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(tamplete),
		Value: payload,
	})
	if err != nil {
		return err
	}

	return nil
}

func logf(msg string, a ...any) {
	fmt.Printf(msg, a...)
	fmt.Println()
}
