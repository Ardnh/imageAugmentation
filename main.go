package main

import (
	"fmt"
	_ "image/jpeg"
	"io/ioutil"

	"github.com/Ardnh/imageAugmentation/src"
)

func main() {
	folder := "Psidium Guajava (Guava)/"
	files, _ := ioutil.ReadDir(fmt.Sprintf("/home/ardanhilal/Documents/coding/golang/imageAugmentation/data/%s", folder))
	path := fmt.Sprintf("/home/ardanhilal/Documents/coding/golang/imageAugmentation/data/%s", folder)
	savePath := fmt.Sprintf("/home/ardanhilal/Documents/coding/golang/imageAugmentation/saved/%s", folder)

	state := 0
	count := 0
	// for i, f := range files {
	// 	if (i-0)%9 == 0 {
	// 		deg = 90
	// 		src.Rotate(path, savePath, f.Name(), deg)
	// 	} else if (i-1)%9 == 0 {
	// 		deg = 180
	// 		src.Rotate(path, savePath, f.Name(), deg)
	// 	} else if (i-2)%9 == 0 {
	// 		deg = 270
	// 		src.Rotate(path, savePath, f.Name(), deg)
	// 	} else if (i-3)%9 == 0 {
	// 		src.Blur(path, savePath, f.Name())
	// 	} else if (i-4)%9 == 0 {
	// 		src.Flip(path, savePath, f.Name(), 1)
	// 	} else if (i-5)%9 == 0 {
	// 		src.Flip(path, savePath, f.Name(), 0)
	// 	} else if (i-6)%9 == 0 {
	// 		src.AdjustBringhtness(path, savePath, f.Name(), 0)
	// 	} else if (i-7)%9 == 0 {
	// 		src.AdjustBringhtness(path, savePath, f.Name(), 1)
	// 	} else {
	// 		deg = 0
	// 		src.Rotate(path, savePath, f.Name(), deg)
	// 	}
	// 	count++
	// }

	for _, f := range files {
		if state == 0 {
			src.Rotate(path, savePath, f.Name(), 90)
			state = 1
		} else if state == 1 {
			src.Rotate(path, savePath, f.Name(), 180)
			state = 2
		} else if state == 2 {
			src.Rotate(path, savePath, f.Name(), 270)
			state = 3
		} else if state == 3 {
			src.Blur(path, savePath, f.Name())
			state = 4
		} else if state == 4 {
			src.Flip(path, savePath, f.Name(), 1)
			state = 5
		} else if state == 5 {
			src.Flip(path, savePath, f.Name(), 0)
			state = 6
		} else if state == 6 {
			src.AdjustBringhtness(path, savePath, f.Name(), 0)
			state = 7
		} else if state == 7 {
			src.AdjustBringhtness(path, savePath, f.Name(), 1)
			state = 8
		} else if state == 8 {
			src.Rotate(path, savePath, f.Name(), 0)
			state = 0
		}
		count++
	}

	fmt.Println(count)
}
