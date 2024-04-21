package main

import (
	"julypdf/caj"
	"julypdf/cpu"
	"julypdf/mupdf"
	"julypdf/qpdf"
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

//export DecryptPdf
func DecryptPdf(fileName string, outFile string) {
	qpdf.CryptoOnPDF(qpdf.Encryption, fileName, outFile, 2)
}

//export EncryptPdf
func EncryptPdf(fileName string, outFile string) {
	qpdf.CryptoOnPDF(qpdf.Encryption, fileName, outFile, 1)
}

//export Img2Pdf
func Img2Pdf(fileName string, outPath string) {
	cpu.Img2Pdf(fileName, outPath)
}

//export MeargeOnPDF
func MeargeOnPDF(fileName string, outFile string) {
	qpdf.MeargeOnPDF(qpdf.Merge, fileName, outFile)
}

//export AddWatermark
func AddWatermark(fileName string, outFile string, waterMarker string) {
	cpu.AddWatermark(fileName, outFile, waterMarker)
}

//export SplitPdf
func SplitPdf(fileName string, splitPage int) {
	qpdf.SplitOnPDF(qpdf.Split, fileName, splitPage)
}

func main() {}
