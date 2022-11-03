package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"time"
)

func main() {
	start := time.Now()

	var img, _ = loadImage("The_Sun_in_high_resolution.jpg")
	var bounds = img.Bounds()
	var gray = image.NewGray(bounds)

	rgbaToGray(img, gray)
	createGrayImage(gray)

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("Processed image in: %s", elapsed)
}

func loadImage(filepath string) (image.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func rgbaToGray(img image.Image, gray *image.Gray) {
	bounds := img.Bounds()
	nbLines := 1
	for y := 0; y < bounds.Max.Y; y += nbLines {
		go lineToGray(img, gray, bounds.Max.X, y, nbLines)
	}
}

func lineToGray(img image.Image, gray *image.Gray, width int, y int, nbLines int) {
	for x := 0; x < width; x++ {
		for n := 0; n < nbLines; n++ {
			var rgba = img.At(x, y+n)
			gray.Set(x, y+n, rgba)
		}
	}
}

func createGrayImage(gray *image.Gray) {
	f, _ := os.Create("gray.png")
	defer f.Close()
	png.Encode(f, gray)
}
