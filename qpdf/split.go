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
	"path"
	"strings"
	"unsafe"
)

// #cgo LDFLAGS: -lgcc_s -lstdc++ -lmingwex -ljpeg -lz -lgnutls -lssl -lcrypto -lws2_32 -lqpdf

// func Split(originFile string, splitNum int) {
// 	OriginFile := C.CString(originFile)
// 	defer C.free(unsafe.Pointer(OriginFile))
// 	baseName := strings.TrimSuffix(originFile, path.Ext(originFile))
// 	srcQPDF := C.qpdf_init()
// 	defer C.qpdf_cleanup(&srcQPDF)

// 	C.qpdf_read(srcQPDF, OriginFile, nil)
// 	i := 0
// 	outQPDF := C.qpdf_init()
// 	C.qpdf_empty_pdf(outQPDF)
// 	outputFilename := fmt.Sprintf("%spart%d.pdf", baseName, i/splitNum+1)
// 	cOutputFilename := C.CString(outputFilename)
// 	defer C.free(unsafe.Pointer(cOutputFilename))
// 	fmt.Println("QPDF版本:", outputFilename)
// 	page := C.qpdf_get_page_n(srcQPDF, C.size_t(i))
// 	C.qpdf_add_page(outQPDF, srcQPDF, page, C.QPDF_TRUE)
// 	C.qpdf_init_write(outQPDF, cOutputFilename)
// 	C.qpdf_write(outQPDF)
// 	C.qpdf_cleanup(&outQPDF)

// }

func Split(originFile string, splitNum int) {
	OriginFile := C.CString(originFile)
	defer C.free(unsafe.Pointer(OriginFile))
	baseName := strings.TrimSuffix(originFile, path.Ext(originFile))
	srcQPDF := C.qpdf_init()
	defer C.qpdf_cleanup(&srcQPDF)

	C.qpdf_read(srcQPDF, OriginFile, nil)
	//get page number
	numPages := int(C.qpdf_get_num_pages(srcQPDF))
	for i := 0; i < numPages; i += splitNum {
		outQPDF := C.qpdf_init()
		C.qpdf_empty_pdf(outQPDF)

		outputFilename := fmt.Sprintf("%s_part_%d.pdf", baseName, i/splitNum+1)
		cOutputFilename := C.CString(outputFilename)
		defer C.free(unsafe.Pointer(cOutputFilename))

		endPage := i + splitNum
		if endPage > numPages {
			endPage = numPages
		}

		for j := i; j < endPage; j++ {
			page := C.qpdf_get_page_n(srcQPDF, C.size_t(j))
			C.qpdf_add_page(outQPDF, srcQPDF, page, C.QPDF_TRUE)
		}

		C.qpdf_init_write(outQPDF, cOutputFilename)
		C.qpdf_write(outQPDF)
		C.qpdf_cleanup(&outQPDF)
	}

}
