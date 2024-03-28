package main

import (
	"julypdf/caj"
	"julypdf/mupdf"
)
import "C"

//export CajToPdf
func CajToPdf(fileName string, outFile string) {
	caj.Caj2pdf(fileName, outFile)
}

//export PdfToPng
func PdfToPng(fileName string, outPath string, concurrent bool, waterMarker string) {
	mupdf.ToPng(fileName, outPath, concurrent, waterMarker)
}
func main() {}
