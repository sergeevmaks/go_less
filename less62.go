package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	red := color.RGBA{200, 30, 30, 255}
	//draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{red}, image.ZP, draw.Src)

	for x := 0; x < 200; x++ {
		for i := 1; i < 10; i++ {
			y := 20 * i
			rectImg.Set(x, y, red)
		}
	}
	for y := 0; y < 200; y++ {
		for i := 1; i < 10; i++ {
			x := 20 * i
			rectImg.Set(x, y, red)
		}
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
