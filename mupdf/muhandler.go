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
	"sync"
	"time"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
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

func ToPng(fileName string, outPath string, concurrent bool, waterMarker ...string) {
	start := time.Now()
	fmt.Printf("start: %s\n", start)

	doc, err := New(fileName)
	if err != nil {
		panic(err)
	}
	defer doc.Close()
	var font *truetype.Font
	if len(waterMarker) > 0 && waterMarker[0] != "" {
		fontBytes, err := os.ReadFile("font.ttf")
		if err != nil {
			panic(err)
		}
		font, err = freetype.ParseFont(fontBytes)
		if err != nil {
			panic(err)
		}
	}
	processPage := func(page int) {
		imgBytes, err := doc.Png(page)
		if err != nil {
			fmt.Println("Error processing page", page, ":", err)
			return
		}

		img, _, err := image.Decode(bytes.NewReader(imgBytes))
		if err != nil {
			fmt.Println("Error decoding image:", err)
			return
		}
		var outputImage image.Image = img
		if len(waterMarker) > 0 && waterMarker[0] != "" {
			bounds := img.Bounds()
			watermarkedImg := image.NewRGBA(bounds)
			draw.Draw(watermarkedImg, bounds, img, bounds.Min, draw.Src)

			drawString(watermarkedImg, 10, 20, waterMarker[0], font)
			outputImage = watermarkedImg
		}

		filePath := filepath.Join(outPath, fmt.Sprintf("test%03d.png", page))
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		if err := png.Encode(file, outputImage); err != nil {
			fmt.Println("Error writing watermarked image:", err)
			return
		}
		file.Close()
		fmt.Printf("page:%d done\n", page)
	}
	// check if use sync or not
	if concurrent {
		var wg sync.WaitGroup
		for n := 0; n < doc.NumPage(); n++ {
			wg.Add(1)

			go func(page int) {
				defer wg.Done()
				fmt.Println("Processing page:", page)
				processPage(page)
			}(n)

		}
		wg.Wait()
	} else {
		for n := 0; n < doc.NumPage(); n++ {
			processPage(n)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Mission complete: %s\n", elapsed)
}

// drawString draws a string on the image at the specified coordinates.
//
// img *image.RGBA, x, y int, label string
func drawString(img *image.RGBA, x, y int, label string, font *truetype.Font) {

	context := freetype.NewContext()
	context.SetDPI(72)
	context.SetFont(font)
	context.SetFontSize(24)
	context.SetClip(img.Bounds())
	context.SetDst(img)
	context.SetSrc(image.NewUniform(color.RGBA{255, 0, 0, 255}))
	pt := freetype.Pt(x, y+int(context.PointToFixed(24)>>6))
	context.DrawString(label, pt)

}
