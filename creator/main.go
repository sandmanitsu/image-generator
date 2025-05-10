package main

import (
	"creator/internal/config"
	"creator/internal/kafka"
	"creator/internal/metrics"
	image "creator/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := config.MustLoad()

	imageSrv := image.NewImageService()

	go func() {
		err := metrics.Listen(config.MetricAddress)
		if err != nil {
			panic(err)
		}
	}()

	kafka.StartConsumer(config, imageSrv)

	// shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown
	log.Println("gracefully shutdown")
}
