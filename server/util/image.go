package util

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func Compress(imagePath string, quality int) error {
	imgfile, err := os.ReadFile(imagePath)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	jpgimg, err := jpeg.Decode(bytes.NewReader(imgfile))
	if err != nil {
		fmt.Println(imgfile[:5])
		log.Println("decoding:" + err.Error())
		return err
	}
	newfile, err := os.Create(imagePath)
	if err != nil {
		log.Println("creating:" + err.Error())
		return err
	}
	defer newfile.Close()

	err = jpeg.Encode(newfile, jpgimg, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Println("encoding:" + err.Error())
		return err
	}

	return nil

}

func CropPic(imagePath string, x, y, x1, y1 int) error {
	log.Println(imagePath)
	imgfile, err := os.ReadFile(imagePath)
	if err != nil {
		log.Println("opening:" + err.Error())
		return err
	}

	jpgimg, err := jpeg.Decode(bytes.NewReader(imgfile))
	if err != nil {
		log.Println("decoding:" + err.Error())
		return err
	}

	newfile, err := os.Create(imagePath)
	if err != nil {
		log.Println("creating:" + err.Error())
		return err
	}
	defer newfile.Close()

	rgbImg, ok := jpgimg.(*image.YCbCr)
	if !ok {
		log.Println("转码" + err.Error())
		return err
	}
	subImg := rgbImg.SubImage(image.Rect(x, y, x1, y1)).(*image.YCbCr)
	if err = jpeg.Encode(newfile, subImg, nil); err != nil {
		log.Println("encoding:" + err.Error())
		return err

	}
	return nil
}
