package util

import (
	"image"
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

func CropPic(imagePath string, x, y, x1, y1 int) error {
	log.Println(imagePath)
	imgfile, err := os.Open(imagePath)
	if err != nil {
		log.Println("opening:" + err.Error())
		return err
	}
	defer imgfile.Close()

	jpgimg, err := jpeg.Decode(imgfile)
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
