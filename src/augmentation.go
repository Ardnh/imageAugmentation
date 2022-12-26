package src

import (
	"fmt"
	"image"

	"github.com/disintegration/imaging"
)

func Rotate(path string, savePath string, name string, deg int) {
	filePath := fmt.Sprintf("%s%s", path, name)
	fmt.Println(filePath)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	var transformedImage *image.NRGBA
	if err != nil {
		panic(err)
	}

	if deg == 0 {
		transformedImage = imaging.Rotate(src, float64(deg), src.Bounds().RGBA64At(src.Bounds().Dx(), src.Bounds().Dy()))
	}

	if deg == 90 {
		transformedImage = imaging.Rotate90(src)
	}

	if deg == 180 {
		transformedImage = imaging.Rotate180(src)
	}

	if deg == 270 {
		transformedImage = imaging.Rotate270(src)
	}

	imaging.Save(transformedImage, savePathName)

	print := fmt.Sprintf("image %s saved after rotate %d", name, deg)
	fmt.Println(print)
}

func Flip(path string, savePath string, name string, deg int) {
	filePath := fmt.Sprintf("%s%s", path, name)
	fmt.Println(filePath)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	var transformedImage *image.NRGBA
	if err != nil {
		panic(err)
	}

	if deg == 1 {
		transformedImage = imaging.FlipH(src)
	}

	if deg == 0 {
		transformedImage = imaging.FlipV(src)
	}

	imaging.Save(transformedImage, savePathName)
	print := fmt.Sprintf("image %s saved after Flip %d", name, deg)
	fmt.Println(print)
}

func Bringhtness(path string, savePath string, name string) {
	filePath := fmt.Sprintf("%s%s", path, name)
	fmt.Println(filePath)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	if err != nil {
		panic(err)
	}

	transformedImage := imaging.AdjustBrightness(src, 20)

	imaging.Save(transformedImage, savePathName)
	print := fmt.Sprintf("image %s saved after adjust bringhtsness %d", name, 10)
	fmt.Println(print)
}

func Blur(path string, savePath string, name string) {
	filePath := fmt.Sprintf("%s%s", path, name)
	fmt.Println(filePath)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	if err != nil {
		panic(err)
	}

	transformedImage := imaging.Blur(src, 2)

	imaging.Save(transformedImage, savePathName)
	print := fmt.Sprintf("image %s saved after adjust blur %d", name, 1)
	fmt.Println(print)
}
