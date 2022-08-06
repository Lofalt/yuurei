package util

import (
	"bytes"
	"fmt"
	"github.com/ser163/png2j"
	"image"
	"image/color"
	"image/draw"
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
func CompressImageResource(imagePath string, quality int) error {
	imgFile, err := os.ReadFile(imagePath)
	log.Println(imagePath)
	if err != nil {
		log.Println("reading" + err.Error())
		return err
	}

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	imgSrc, _, err := image.Decode(bytes.NewReader(imgFile))
	if err != nil {
		log.Println("decoding:" + err.Error())

		return err
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	newFile, err := os.Create(imagePath + ".jpg")
	if err != nil {
		log.Println("creating:" + err.Error())

		return err
	}

	err = jpeg.Encode(newFile, newImg, &jpeg.Options{Quality: quality})
	if err != nil {
		log.Println("encoding" + err.Error())
		return err
	}
	return nil
}

func CompressPng(imagePath string, quality int) error {
	err := png2j.ConBase(imagePath, imagePath+".jpg", color.White, quality)
	if err != nil {
		return err
	}
	return nil
}
