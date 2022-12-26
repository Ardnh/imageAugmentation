package main

import (
	"fmt"
	_ "image/jpeg"
	"io/ioutil"

	"github.com/Ardnh/imageAugmentation/src"
)

func main() {
	folder := "Santalum Album (Sandalwood)/"
	files, _ := ioutil.ReadDir(fmt.Sprintf("/home/ardanhilal/Documents/coding/golang/imageAugmentation/augmentasi/%s", folder))
	path := fmt.Sprintf("/home/ardanhilal/Documents/coding/golang/imageAugmentation/augmentasi/%s", folder)
	savePath := fmt.Sprintf("/home/ardanhilal/Documents/coding/golang/imageAugmentation/saved/%s", folder)
	fmt.Println(path)
	fmt.Println(savePath)
	deg := 0
	count := 0
	for i, f := range files {
		if (i+0)%8 == 0 {
			deg = 90
			src.Rotate(path, savePath, f.Name(), deg)
		} else if (i-1)%8 == 0 {
			deg = 180
			src.Rotate(path, savePath, f.Name(), deg)
		} else if (i-2)%8 == 0 {
			deg = 270
			src.Rotate(path, savePath, f.Name(), deg)
		} else if (i-3)%8 == 0 {
			src.Blur(path, savePath, f.Name())
		} else if (i-4)%8 == 0 {
			src.Flip(path, savePath, f.Name(), 1)
		} else if (i-5)%8 == 0 {
			src.Flip(path, savePath, f.Name(), 0)
		} else if (i-6)%8 == 0 {
			src.Bringhtness(path, savePath, f.Name())
		} else {
			deg = 0
			src.Rotate(path, savePath, f.Name(), deg)
		}
		count++
	}
	fmt.Println(count)
}
