package util

import (
	"image/jpeg"
	"log"
	"os"
)

func Compress(imagePath string, quality int) {
	imgfile, err := os.Open(imagePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer imgfile.Close()

	jpgimg, err := jpeg.Decode(imgfile)
	if err != nil {
		log.Println(err.Error())
		return
	}
	newfile, err := os.Create(imagePath)
	if err != nil {
		log.Println("creating:" + err.Error())
		return
	}
	defer newfile.Close()

	err = jpeg.Encode(newfile, jpgimg, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Println("encoding:" + err.Error())
		return
	}

	return

}
