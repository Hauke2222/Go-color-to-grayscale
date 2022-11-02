package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"os"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

var img image.Image

func main() {
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(8))

	start := time.Now()

	loadImage("The_Sun_in_high_resolution.jpg")

	// wg.Add(1)
	// var bounds = img.Bounds()
	// var gray = image.NewGray(bounds)
	// go rgbaToGray(img, gray)
	// wg.Wait()

	// wg.Add(1)
	// go createGrayImage(gray)
	// wg.Wait()

	var bounds = img.Bounds()
	var gray = image.NewGray(bounds)
	rgbaToGray(img, gray)

	createGrayImage(gray)

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Printf("Processed image in: %s", elapsed)
}

func loadImage(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic("Error opening image.")
	}
	defer file.Close()
	img, _, err = image.Decode(file)
	if err != nil {
		panic("Error decoding image.")
	}
}

func rgbaToGray(img image.Image, gray *image.Gray) {
	// defer wg.Done()

	bounds := img.Bounds()
	nbLines := 1 //bounds.Max.Y / 1

	// process each line in image individually
	for y := 0; y < bounds.Max.Y; y += nbLines {
		// fmt.Println(y)
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
	// defer wg.Done()
	f, _ := os.Create("gray.png")
	defer f.Close()
	png.Encode(f, gray)
}
