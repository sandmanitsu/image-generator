package v2

import (
	"creator/internal/templates"
	"encoding/json"
	"log"
	"path/filepath"

	"github.com/google/uuid"
	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	templateName = "v2"

	weight = 1000
	height = 1000

	imgW = 600
	imgH = 600

	backgroundColor = "#ffd663"
)

type Param struct {
	Size  string `json:"size"`
	Price string `json:"price"`
	Type  string `json:"type"`
	File  string `json:"file"`
}

func Process(payload []byte) {
	var param Param
	if err := json.Unmarshal(payload, &param); err != nil {
		log.Panicln("error unmarshal payload", err)
		return
	}

	// fmt.Println("v1", param)

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	mw.SetSize(weight, height)

	bc := imagick.NewPixelWand()
	defer bc.Destroy()
	bc.SetColor(backgroundColor)
	mw.NewImage(weight, height, bc)

	// main image
	imgPath := filepath.Join("../", "files", param.File)
	img := imagick.NewMagickWand()
	err := img.ReadImage(imgPath)
	if err != nil {
		log.Printf("%s : error read image %s : %v", templateName, imgPath, err)
	}

	err = img.ResizeImage(imgW, imgH, imagick.FILTER_LANCZOS)
	if err != nil {
		log.Printf("%s : error resize image %s : %v", templateName, imgPath, err)
	}

	err = mw.CompositeImage(img, imagick.COMPOSITE_OP_OVER, true, (weight-imgW)/2, (height-imgH)/2)
	if err != nil {
		log.Println(templateName, err)
	}
	img.Destroy()

	// text
	dw := imagick.NewDrawingWand()
	pw := imagick.NewPixelWand()
	defer dw.Destroy()
	defer pw.Destroy()

	pw.SetColor("black")
	dw.SetFillColor(pw)
	dw.SetFontSize(110)

	metrics := mw.QueryFontMetrics(dw, param.Type)
	mw.AnnotateImage(dw, weight/2-(float64(metrics.TextWidth)/2), 170, 0, param.Type)

	dw.SetFontSize(40)
	priceText := "Цена: " + param.Price
	metrics = mw.QueryFontMetrics(dw, priceText)

	mw.AnnotateImage(dw, 300-float64(metrics.TextWidth)/2, 800, 0, priceText)

	sizeText := "Размер: " + param.Size
	metrics = mw.QueryFontMetrics(dw, sizeText)
	mw.AnnotateImage(dw, 600-float64(metrics.TextWidth)/2, 800, 0, sizeText)

	templates.SaveImage(mw, templateName, uuid.NewString())
}
