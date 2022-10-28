package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"runtime"
	"time"
)

// var wg sync.WaitGroup
var gray *image.Gray
var img image.Image

func main() {
	// runtime.GOMAXPROCS(1)

	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(0))

	start := time.Now()

	img, _ = loadImage("The_Sun_in_high_resolution.jpg")

	gray = rgbaToGray(img)

	f, _ := os.Create("gray.png")
	defer f.Close()
	png.Encode(f, gray)

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

func rgbaToGray(img image.Image) *image.Gray {
	// defer wg.Done()
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}
