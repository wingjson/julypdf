package qpdf

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L../libs -lpdf
#include <stdlib.h>
#include "qpdf/qpdf-c.h"
*/
import "C"
import (
	"fmt"
	"path"
	"strings"
	"unsafe"
)

type Crypto func(qpdf C.qpdf_data)

type MergeType func(qpdf C.qpdf_data, originFile string)

type SplitType func(originFile string, splitNum int)

func freeCString(cstr *C.char) {
	C.free(unsafe.Pointer(cstr))
}

func CryptoOnPDF(op Crypto, fileName string) {
	qpdf := C.qpdf_init()
	defer C.qpdf_cleanup(&qpdf)

	inputPDF := C.CString(fileName)
	defer freeCString(inputPDF)
	password := C.CString("")
	defer freeCString(password)

	if C.qpdf_read(qpdf, inputPDF, password) != C.QPDF_SUCCESS {
		fmt.Println("read PDF error")
		return
	}
	baseName := strings.TrimSuffix(fileName, path.Ext(fileName))
	outputFilename := fmt.Sprintf("%s_encrtpt.pdf", baseName)
	outputPDF := C.CString(outputFilename)
	defer freeCString(outputPDF)

	C.qpdf_init_write(qpdf, outputPDF)

	op(qpdf)

	if C.qpdf_write(qpdf) != C.QPDF_SUCCESS {
		fmt.Println("write error")
		return
	}
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
