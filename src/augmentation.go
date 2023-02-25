package src

import (
	"fmt"
	"image"

	"github.com/disintegration/imaging"
)

func Rotate(path string, savePath string, name string, state int) {
	filePath := fmt.Sprintf("%s%s", path, name)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	var transformedImage *image.NRGBA
	if err != nil {
		panic(err)
	}

	if state == 0 {
		transformedImage = imaging.Rotate(src, float64(0), src.Bounds().RGBA64At(src.Bounds().Size().X, src.Bounds().Size().Y))
	}

	if state == 1 {
		transformedImage = imaging.Rotate90(src)
	}

	if state == 2 {
		transformedImage = imaging.Rotate180(src)
	}

	if state == 3 {
		transformedImage = imaging.Rotate270(src)
	}

	imaging.Save(transformedImage, savePathName)

	print := fmt.Sprintf("image %s saved after rotate", name)
	fmt.Println(print)
}

func Flip(path string, savePath string, name string, state int) {
	filePath := fmt.Sprintf("%s%s", path, name)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	var transformedImage *image.NRGBA
	if err != nil {
		panic(err)
	}

	if state == 5 {
		transformedImage = imaging.FlipH(src)
	}

	if state == 6 {
		transformedImage = imaging.FlipV(src)
	}

	imaging.Save(transformedImage, savePathName)
	print := fmt.Sprintf("image %s saved after Flip", name)
	fmt.Println(print)
}

func AdjustBringhtness(path string, savePath string, name string, state int) {
	filePath := fmt.Sprintf("%s%s", path, name)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	var transformedImage *image.NRGBA
	if err != nil {
		panic(err)
	}

	if state == 7 {
		transformedImage = imaging.AdjustBrightness(src, 30)
	}

	if state == 8 {
		transformedImage = imaging.AdjustBrightness(src, -30)
	}

	imaging.Save(transformedImage, savePathName)
	print := fmt.Sprintf("image %s saved after adjust bringhtsness", name)
	fmt.Println(print)
}

func Blur(path string, savePath string, name string) {
	filePath := fmt.Sprintf("%s%s", path, name)
	savePathName := fmt.Sprintf("%s%s", savePath, name)
	src, err := imaging.Open(filePath)
	if err != nil {
		panic(err)
	}

	transformedImage := imaging.Blur(src, 2.25)

	imaging.Save(transformedImage, savePathName)
	print := fmt.Sprintf("image %s saved after adjust blur", name)
	fmt.Println(print)
}
