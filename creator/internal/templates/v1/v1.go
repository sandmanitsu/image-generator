package v1

import (
	"creator/internal/templates"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/google/uuid"
	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	templateName = "v1"

	weight = 1000
	height = 1000
)

type Param struct {
	Size  string `json:"size"`
	Price string `json:"price"`
	Type  string `json:"type"`
	File  string `json:"file"`
}

func ProcessTemplate(payload []byte) {
	var Param Param
	if err := json.Unmarshal(payload, &Param); err != nil {
		log.Panicln("error unmarshal payload", err)
		return
	}

	fmt.Println("v1", Param)
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.SetSize(weight, height)

	bc := imagick.NewPixelWand()
	bc.SetColor("white")
	mw.NewImage(weight, height, bc)

	// main image
	imgPath := filepath.Join("../", "files", Param.File)
	img := imagick.NewMagickWand()
	err := img.ReadImage(imgPath)
	if err != nil {
		log.Printf("%s : error read image %s : %v", templateName, imgPath, err)
	}

	err = img.ResizeImage(600, 600, imagick.FILTER_LANCZOS)
	if err != nil {
		log.Printf("%s : error resize image %s : %v", templateName, imgPath, err)
	}

	err = mw.CompositeImage(img, imagick.COMPOSITE_OP_OVER, true, 500, 500)
	if err != nil {
		log.Println(templateName, err)
	}

	templates.SaveImage(mw, templateName, uuid.NewString())
}
