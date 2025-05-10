package image

import (
	"creator/internal/metrics"
	v1 "creator/internal/templates/v1"
	v2 "creator/internal/templates/v2"
	"fmt"
	"log"
	"time"
)

type ImageService struct {
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (i *ImageService) ProcessImages(template string, payload []byte) {
	start := time.Now()
	defer func() {
		metrics.ObserveCreateImage(time.Since(start), template)
	}()

	fmt.Println("process image: ", template)

	switch template {
	case "v1":
		v1.Process(payload)
	case "v2":
		v2.Process(payload)
	default:
		log.Println("template doesn't exist", template)
	}
}
