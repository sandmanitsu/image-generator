package image

import (
	v1 "creator/internal/templates/v1"
	"fmt"
)

type ImageService struct {
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func (i *ImageService) ProcessImages(template string, payload []byte) {
	fmt.Println("process image: ", template)

	switch template {
	case "v1":
		v1.ProcessTemplate(payload)
	}
}
