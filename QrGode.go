package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	w, b := QrGodeInit()

	QrCodeStart(w, b, 500)
}

const (
	eachBlockLong int = 10
)

// QrBlock 颜色快
type QrBlock struct {
	long    int
	isBlack bool
}

// QrGodeInit start
func QrGodeInit() (QrBlock, QrBlock) {

	fmt.Println("start QrGode.go")

	var whiteBlock, blackBlock QrBlock
	whiteBlock.isBlack = false
	whiteBlock.long = eachBlockLong
	blackBlock.isBlack = true
	blackBlock.long = eachBlockLong

	return whiteBlock, blackBlock
}

// QrCodeStart ss
func QrCodeStart(whiteBlock, blackBlock QrBlock, width int) {
	fmt.Println(whiteBlock.long, whiteBlock.isBlack)

	t := image.Rect(0, 0, whiteBlock.long, whiteBlock.long)

	fmt.Println(t)

	tagQrImage, err := os.Create("tagQrImage.png")

	if err != nil {
		fmt.Println(err)
	}
	defer tagQrImage.Close()

	tagImg := image.NewGray(image.Rect(0, 0, width, width))

	fmt.Println(tagImg.PixOffset(width, 10))

	for x := 10; x < width-10; x += eachBlockLong {
		for y := 10; y < width-10; y += eachBlockLong {
			// if ((x/eachBlockLong)+(y/eachBlockLong))%2 == 0 {
			// 	var thisGray color.Gray
			// 	thisGray.Y = 255
			// 	tagImg.SetGray(x, y, thisGray)
			// }
			if (x >= 20 && x <= 60) && (y >= 20 && y <= 60) {
				colorBlock(whiteBlock, x, y, tagImg)
			}
		}
	}

	printErr := png.Encode(tagQrImage, tagImg)
	if printErr != nil {
		log.Fatal(printErr)
	}
}

func colorBlock(thisBlock QrBlock, x, y int, tagImg *image.Gray) {
	// fmt.Println("coloring thisBlock is Black?", thisBlock.isBlack)
	var thisBlockColor color.Gray
	if thisBlock.isBlack {
		thisBlockColor.Y = 0
	} else {
		thisBlockColor.Y = 255
	}
	for i := x; i < x+thisBlock.long; i++ {
		for j := y; j < y+thisBlock.long; j++ {
			tagImg.SetGray(i, j, thisBlockColor)
		}
	}
}
