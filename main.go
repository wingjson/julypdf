package main

import "julypdf/qpdf"

func main() {
	// qpdf.cryptoOnPDF(qpdf.Decryption)
	qpdf.CryptoOnPDF(qpdf.Encryption, "555.pdf")
	// qpdf.MeargeOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
	// qpdf.SplitOnPDF(qpdf.Split, "555.pdf", 1)
}
