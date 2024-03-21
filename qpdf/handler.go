package qpdf

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -lgcc_s -lstdc++ -ljpeg -lz  -lssl -lcrypto  -lqpdf
#include <stdlib.h>
#include "qpdf/qpdf-c.h"
*/
import "C"
import (
	"fmt"
	"julypdf/ignoreerror"
	"unsafe"
)

type Crypto func(qpdf C.qpdf_data) error

type MergeType func(qpdf C.qpdf_data, originFile string)

type SplitType func(originFile string, splitNum int)

func freeCString(cstr *C.char) {
	C.free(unsafe.Pointer(cstr))
}

// CryptoOnPDF encrypts a PDF file using the specified Crypto operation.
//
// op: the Crypto operation to apply to the PDF.
// originFileName: the name of the PDF file to encrypt.
// cryptType: the type of cryption to use. 1:encryption, 2:decryption
func CryptoOnPDF(op Crypto, originFileName string, targetFileName string, cryptType int) (string, error) {
	qpdf := C.qpdf_init()
	defer C.qpdf_cleanup(&qpdf)

	inputPDF := C.CString(originFileName)
	defer freeCString(inputPDF)
	password := C.CString("")
	defer freeCString(password)

	if C.qpdf_read(qpdf, inputPDF, password) != C.QPDF_SUCCESS {
		return "", fmt.Errorf("failed to read PDF: %s", originFileName)
	}

	outputPDF := C.CString(targetFileName)
	defer freeCString(outputPDF)

	C.qpdf_init_write(qpdf, outputPDF)

	op(qpdf)

	//some pdf not standard
	if C.qpdf_write(qpdf) != C.QPDF_SUCCESS {
		return targetFileName, &ignoreerror.IgnorableError{Msg: "ignore"}
	}
	return targetFileName, nil
}

func MeargeOnPDF(op MergeType, originFile string, outputFile string) {
	qpdf := C.qpdf_init()
	defer C.qpdf_cleanup(&qpdf)

	C.qpdf_empty_pdf(qpdf)

	outputFilename := C.CString(outputFile)
	defer C.free(unsafe.Pointer(outputFilename))

	op(qpdf, originFile)

	C.qpdf_init_write(qpdf, outputFilename)
	C.qpdf_write(qpdf)
}

func SplitOnPDF(op SplitType, originFile string, splitNum int) {
	op(originFile, splitNum)
}
