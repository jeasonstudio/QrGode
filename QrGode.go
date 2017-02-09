package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

const (
	eachBlockLong int = 10
)

// QrBlock 颜色快
type QrBlock struct {
	long    int
	isBlack bool
}

func main() {

	// input
	versionNum := 5
	tagPathName := "tagQrImage.png"
	// input end

	whiteBlock, blackBlock := QrGodeInit()
	thisPNGWidth := (qrVersion(versionNum) + 2) * 10
	QrCodeStart(whiteBlock, blackBlock, thisPNGWidth, tagPathName)

}

func qrVersion(vNum int) int {
	if vNum < 1 || vNum > 40 {
		fmt.Println("Error Version!!!")
		return -1
	}
	return 21 + 4*(vNum-1)
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

// QrCodeStart 开始主程序
func QrCodeStart(whiteBlock, blackBlock QrBlock, width int, tagPathName string) {

	tagQrImage, err := os.Create(tagPathName)

	if err != nil {
		fmt.Println(err)
	}
	defer tagQrImage.Close()

	tagImg := image.NewGray(image.Rect(0, 0, width, width))

	finalArr := initResultArr(width / 10)

	changeResultArr(finalArr, width/10)

	printArr(finalArr, width/10)

	// for x := 0; x < width; x += eachBlockLong {
	// 	for y := 0; y < width; y += eachBlockLong {
	// 		colorBlock(whiteBlock, x, y, tagImg)
	// 	}
	// }

	for x := 0; x < width; x += eachBlockLong {
		for y := 0; y < width; y += eachBlockLong {

			if finalArr[x/10][y/10] == 0 {
				colorBlock(whiteBlock, x, y, tagImg)
			} else {
				colorBlock(blackBlock, x, y, tagImg)
			}
			// // 基准线
			// if x == 7*eachBlockLong || y == 7*eachBlockLong {
			// 	colorBlock(blackBlock, x, y, tagImg)
			// }
		}
	}

	printErr := png.Encode(tagQrImage, tagImg)
	if printErr != nil {
		log.Fatal(printErr)
	}
}

// 初始化结果数组
func initResultArr(width int) [][]int {
	arr := make([][]int, width, width)
	for i := 0; i < width; i++ {
		arr2 := make([]int, width, width)
		for j := 0; j < width; j++ {
			arr2[j] = 0
		}
		arr[i] = arr2
	}
	return arr
}

func changeResultArr(arr [][]int, width int) {
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			// 三个黑圈
			if (y == width-2*1 || y == width-8*1) && (x >= 1 && x <= 7*1) || (x == width-2*1 || x == width-8*1) && (y >= 1 && y <= 7*1) || (x == 1 || x == 7*1) && ((y >= 1 && y <= 7*1) || (y >= width-8*1 && y <= width-2)) || (y == 1 || y == 7*1) && ((x >= 1 && x <= 7*1) || (x >= width-8*1 && x <= width-2)) {
				arr[x][y] = 1
			}
			// 三个黑点
			if (x >= 3 && x <= 5 && y >= 3 && y <= 5) || (x >= width-6 && x <= width-4 && y >= 3 && y <= 5) || (x >= 3 && x <= 5 && y >= width-6 && y <= width-4) {
				arr[x][y] = 1
			}
			if (x == width/2 && y == width/2) || (x == width/2 && (y == 7 || y == width-8)) || (y == width/2 && (x == 7 || x == width-8)) {
				drawLittleBlock(arr, x, y)
			}
		}
	}
}

func drawLittleBlock(arr [][]int, x, y int) {
	arr[x][y] = 1
	arr[x-2][y-2] = 1
	arr[x-1][y-2] = 1
	arr[x][y-2] = 1
	arr[x+1][y-2] = 1
	arr[x+2][y-2] = 1
	arr[x-2][y-1] = 1
	arr[x+2][y-1] = 1
	arr[x-2][y] = 1
	arr[x+2][y] = 1
	arr[x-2][y+1] = 1
	arr[x+2][y+1] = 1
	arr[x-2][y+2] = 1
	arr[x+2][y+2] = 1
	arr[x+1][y+2] = 1
	arr[x-1][y+2] = 1
	arr[x][y+2] = 1
}

// colorBlock 方块着色
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

//////

func printArr(arr [][]int, width int) {
	for i := 0; i < width; i++ {
		fmt.Println(arr[i])
	}
}
