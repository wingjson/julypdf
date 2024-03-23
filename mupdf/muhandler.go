package mupdf

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func ToJpeg(fileName string, outPath string) {
	start := time.Now()
	fmt.Printf("start: %s\n", start)
	doc, err := New(fileName)
	if err != nil {
		panic(err)
	}
	defer doc.Close()
	println(doc.NumPage())

	for n := 0; n < doc.NumPage(); n++ {
		println(n)
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(outPath, fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
	}
	elapsed := time.Since(start)
	fmt.Printf("mission complete: %s\n", elapsed)
}

func ToPng(fileName string, outPath string, waterMarker ...string) {
	start := time.Now()
	fmt.Printf("start: %s\n", start)

	doc, err := New(fileName)
	if err != nil {
		panic(err)
	}
	defer doc.Close()

	for n := 0; n < doc.NumPage(); n++ {
		fmt.Println("Processing page:", n)
		imgBytes, err := doc.Png(n)
		if err != nil {
			fmt.Println("Error processing page", n, ":", err)
			continue
		}

		img, _, err := image.Decode(bytes.NewReader(imgBytes))
		if err != nil {
			fmt.Println("Error decoding image:", err)
			continue
		}
		var outputImage image.Image = img
		if len(waterMarker) > 0 && waterMarker[0] != "" {
			bounds := img.Bounds()
			watermarkedImg := image.NewRGBA(bounds)
			draw.Draw(watermarkedImg, bounds, img, bounds.Min, draw.Src)

			drawString(watermarkedImg, 10, 20, waterMarker[0])
			outputImage = watermarkedImg
		}

		filePath := filepath.Join(outPath, fmt.Sprintf("test%03d.png", n))
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer file.Close()

		if err := png.Encode(file, outputImage); err != nil {
			fmt.Println("Error writing watermarked image:", err)
			continue
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Mission complete: %s\n", elapsed)
}

// drawString draws a string on the image at the specified coordinates.
//
// img *image.RGBA, x, y int, label string
func drawString(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 0, 0, 255}
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
