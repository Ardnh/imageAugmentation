package main

import (
	"fmt"
	_ "image/jpeg"
	"io/ioutil"

	"github.com/Ardnh/imageAugmentation/src"
)

func main() {
	folder := "jambu"
	files, _ := ioutil.ReadDir(fmt.Sprintf("./data/%s/", folder))
	path := fmt.Sprintf("./data/%s/", folder)
	savePath := fmt.Sprintf("./saved/%s/", folder)

	state := 0
	count := 0

	for _, f := range files {
		name := f.Name()
		if state == 0 {
			src.Rotate(path, savePath, name, state)
			state = 1
		} else if state == 1 {
			src.Rotate(path, savePath, name, state)
			state = 2
		} else if state == 2 {
			src.Rotate(path, savePath, name, state)
			state = 3
		} else if state == 3 {
			src.Rotate(path, savePath, name, state)
			state = 4
		} else if state == 4 {
			src.Blur(path, savePath, name)
			state = 5
		} else if state == 5 {
			src.Flip(path, savePath, name, state)
			state = 6
		} else if state == 6 {
			src.Flip(path, savePath, name, state)
			state = 7
		} else if state == 7 {
			src.AdjustBringhtness(path, savePath, name, state)
			state = 8
		} else if state == 8 {
			src.AdjustBringhtness(path, savePath, name, state)
			state = 0
		}
		count++
	}

	fmt.Println(count)
}
