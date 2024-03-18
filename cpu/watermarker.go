package cpu

import (
	"fmt"
	"julypdf/ignoreerror"
	"julypdf/qpdf"
	"os"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func AddWatermark(fileName string) {
	start := time.Now()
	fmt.Printf("start: %s\n", start)

	waterFile := decrypt(fileName)
	if waterFile == "" {
		return
	}

	outputPDF := "output.pdf"
	inputFile, err := os.Open(waterFile)
	if err != nil {
		fmt.Printf("cannot open %s: %v\n", fileName, err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputPDF)
	fmt.Printf("cannot create %s", outputPDF)
	if err != nil {
		fmt.Printf("cannot create %s: %v\n", outputPDF, err)
		return
	}
	defer outputFile.Close()

	onTop := false
	update := false
	wm, err := api.TextWatermark("Demo", "", onTop, update, types.POINTS)
	if err != nil {
		fmt.Printf("cannot create watermark: %v\n", err)
		return
	}

	//add watermark
	if err := api.AddWatermarks(inputFile, outputFile, nil, wm, nil); err != nil {
		fmt.Printf("cannot add watermark: %v\n", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("mission complete: %s\n", elapsed)
}

// because sometime pdfcpu maybe decrypt failed,so use qpdf to decrypt
func decrypt(fileName string) string {
	result, err := qpdf.CryptoOnPDF(qpdf.Decryption, fileName, 2)
	if err != nil {
		if _, ok := err.(*ignoreerror.IgnorableError); ok {
			return result
		} else {
			return ""
		}
	}
	return result
}
