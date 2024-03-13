package main

import "julypdf/qpdf"

func main() {
	// qpdf.cryptoOnPDF(qpdf.Decryption)
	// qpdf.CryptoOnPDF(qpdf.Encryption)
	qpdf.OperateOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
}
