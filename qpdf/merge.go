package qpdf

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -lgcc_s -lstdc++ -ljpeg -lz  -lssl -lcrypto  -lqpdf
#include <stdlib.h>
#include "qpdf/qpdf-c.h"
*/
import "C"
import (
	"strings"
	"unsafe"
)

// func Merge() {
func Merge(qpdf C.qpdf_data, originFile string) {
	filenames := strings.Split(originFile, " ")
	for _, filename := range filenames {
		inputFilename := C.CString(filename)
		defer C.free(unsafe.Pointer(inputFilename))
		addPagesFromPDF(qpdf, inputFilename)
	}

}

func addPagesFromPDF(qpdf C.qpdf_data, filename *C.char) {
	srcQPDF := C.qpdf_init()
	C.qpdf_read(srcQPDF, filename, nil)
	numPages := C.qpdf_get_num_pages(srcQPDF)
	for i := C.int(0); i < numPages; i++ {
		page := C.qpdf_get_page_n(srcQPDF, C.size_t(i))
		C.qpdf_add_page(qpdf, srcQPDF, page, C.QPDF_TRUE)
	}
	C.qpdf_cleanup(&srcQPDF)
}
