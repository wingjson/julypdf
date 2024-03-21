package mupdf

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
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

func ToPng(fileName string, outPath string) {
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
		img, err := doc.Png(n)
		if err != nil {
			panic(err)
		}

		file, err := os.Create(filepath.Join(outPath, fmt.Sprintf("test%03d.png", n)))
		if err != nil {
			panic(err)
		}

		defer file.Close()

		_, err = file.Write(img)
		if err != nil {
			panic(err)
		}
	}

	// for n := 0; n < doc.NumPage(); n++ {
	// 	fmt.Println("Processing page:", n)
	// 	imgBytes, err := doc.Png(n)
	// 	if err != nil {
	// 		fmt.Println("Error processing page", n, ":", err)
	// 		continue
	// 	}
	// 	// waterImg := utilts.MarkerWithImg(imgBytes)

	// 	filePath := filepath.Join(outPath, fmt.Sprintf("test%03d_watermarked.png", n))
	// 	file, err := os.Create(filePath)
	// 	if err != nil {
	// 		fmt.Println("Error creating file:", err)
	// 		continue
	// 	}
	// 	defer file.Close()

	// 	if err := png.Encode(file, waterImg); err != nil {
	// 		fmt.Println("Error writing watermarked image:", err)
	// 		continue
	// 	}
	// }
	elapsed := time.Since(start)
	fmt.Printf("mission complete: %s\n", elapsed)
}

func drawString(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 0, 0, 255} // 红色水印
	point := fixed.Point26_6{X: fixed.Int26_6(x * 64), Y: fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
