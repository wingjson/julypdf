package main

import (
	"julypdf/cmd"
)

func main() {
	// qpdf.cryptoOnPDF(qpdf.Decryption)
	// qpdf.CryptoOnPDF(qpdf.Encryption, "555.pdf", 1)
	// qpdf.MeargeOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
	// qpdf.SplitOnPDF(qpdf.Split, "555.pdf", 1)
	// res := utilts.CheckIfEncrypted("555.pdf")
	// fmt.Println(res)
	// cpu.AddWatermark("555.pdf", "666.pdf", "july")
	// mupdf.ToPng("555.pdf", "test", "yesyes")

	cmd.Execute()
}
