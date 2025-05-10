package templates

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	PNG = ".png"
)

func SaveImage(mw *imagick.MagickWand, template string, filename string) {
	dirPath := filepath.Join("images", template)

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		fmt.Printf("dir %s doesn't exist, creating dir", template)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			log.Println("error create dir", template, err)

			return
		}
	} else if err != nil {
		log.Printf("error check if dir %s exist %v", template, err)

		return
	}

	mw.WriteImage(filepath.Join(dirPath, filename+PNG))
}
