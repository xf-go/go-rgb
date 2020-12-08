package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func main() {
	rgb()

}

func rgb() {
	source := "E:/gowwwroot/xu/go-rgb/0000B4.JPG"

	f, err := os.Open(source)
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
	// bounds := m.Bounds()
	// dx := bounds.Dx()
	// dy := bounds.Dy()

	c := m.At(0, 0)
	fmt.Println("c: ", c)
	r, g, b, a := c.RGBA()
	fmt.Printf("%v-%v-%v-%v: ", r>>8, g>>8, b>>8, a>>8)
}
