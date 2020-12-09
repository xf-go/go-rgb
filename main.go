package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Point struct {
	X, Y int
}

type RGBA struct {
	R, G, B, A uint8
}

func main() {
	points := []Point{
		{
			X: 2519,
			Y: 2306,
		},
		{
			X: 2524,
			Y: 2306,
		},
		{
			X: 2529,
			Y: 2306,
		},
	}
	dir := "./100EOS5D"
	filelist, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("err: ", err)
	}
	for i := range filelist {
		filePath := dir + "\\" + filelist[i].Name()
		avgRGB(filePath, points)
	}

	// newImage("./2.jpg", nil)
	// watermark("./00B46C.JPG")

}

func avgRGB(imagePath string, points []Point) {
	f, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer f.Close()

	// buf := bufio.NewReader(f)

	m, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	bounds := m.Bounds()
	dx := bounds.Dx()
	// dy := bounds.Dy()

	var sumR, sunG, sumB int
	pointLen := len(points)
	for _, point := range points {
		c := m.At(point.X, point.Y)
		r, g, b, _ := c.RGBA()
		newR, newG, newB := r>>8, g>>8, b>>8
		sumR += int(newR)
		sunG += int(newG)
		sumB += int(newB)
	}

	newR, newG, newB := uint8(sumR/pointLen), uint8(sunG/pointLen), uint8(sumB/pointLen)

	fmt.Printf("%v- %v- %v \n", newR, newG, newB)

	// filenameWithoutSuffix(imagePath)
	data := fmt.Sprintf("%s - %02x%02x%02x", filenameWithoutSuffix(imagePath), newR, newG, newB)
	writeFile("1.txt", data)

	im, err := newImage(dx, 500, &RGBA{newR, newG, newB, 255})
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	// createImg("./3.jpg", dx, 500, &RGBA{newR, newG, newB, 255})

	watermark(imagePath, im)
}

// const (
// 	dx = 256
// 	dy = 256
// )

func newImage(width, height int, rgba *RGBA) (image.Image, error) {
	newRGBA := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			if rgba == nil {
				newRGBA.SetRGBA(x, y, color.RGBA{uint8(x % 256), uint8(y % 256), 0, 255})
			} else {
				newRGBA.SetRGBA(x, y, color.RGBA{uint8(rgba.R), uint8(rgba.G), uint8(rgba.B), uint8(rgba.A)})
			}
		}
	}
	return newRGBA, nil
}

func createImg(filePath string, width, height int, rgba *RGBA) {
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	im, err := newImage(width, height, rgba)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	encode(filePath, f, im)
}

func encode(inputName string, file *os.File, rgba image.Image) {
	if strings.HasSuffix(inputName, "jpg") || strings.HasSuffix(inputName, "jpeg") {
		if err := jpeg.Encode(file, rgba, nil); err != nil {
			fmt.Println("err: ", err)
			return
		}
	} else if strings.HasSuffix(inputName, "png") {
		png.Encode(file, rgba)
	} else if strings.HasSuffix(inputName, "gif") {
		gif.Encode(file, rgba, nil)
	} else {
		fmt.Println("不支持的图片格式")
	}
}

func watermark(src string, waterImg image.Image) {
	//原始图片是sam.jpg
	srcF, _ := os.Open(src)
	srcImg, _ := jpeg.Decode(srcF)
	defer srcF.Close()

	// water, _ := os.Open("./3.jpg")
	// waterImg, _ := jpeg.Decode(water)
	// defer water.Close()

	offset := image.Pt(0, 0)
	bounds := srcImg.Bounds()
	m := image.NewNRGBA(bounds)

	draw.Draw(m, bounds, srcImg, image.ZP, draw.Src)
	draw.Draw(m, waterImg.Bounds().Add(offset), waterImg, image.ZP, draw.Over)

	//生成新图片new.jpg，并设置图片质量..
	imgw, _ := os.Create(filename(src))
	jpeg.Encode(imgw, m, &jpeg.Options{100})

	defer imgw.Close()

	fmt.Println("水印添加结束...")
}

func filename(filePath string) string {
	n := strings.LastIndex(filePath, ".")
	return filePath[:n] + "-new" + filePath[n:]
}

func filenameWithoutSuffix(filePath string) string {
	_, file := filepath.Split(filePath)
	n := strings.LastIndex(file, ".")
	return file[:n]
}
