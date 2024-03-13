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
	"unsafe"
)

type Crypto func(qpdf C.qpdf_data)

type Operater func(qpdf C.qpdf_data, originFile string)

func freeCString(cstr *C.char) {
	C.free(unsafe.Pointer(cstr))
}

func CryptoOnPDF(op Crypto) {
	qpdf := C.qpdf_init()
	defer C.qpdf_cleanup(&qpdf)

	version := C.GoString(C.qpdf_get_qpdf_version())
	fmt.Println("QPDF版本:", version)

	inputPDF := C.CString("./2.pdf")
	defer freeCString(inputPDF)
	password := C.CString("")
	defer freeCString(password)

	if C.qpdf_read(qpdf, inputPDF, password) != C.QPDF_SUCCESS {
		fmt.Println("读取 PDF 文件失败")
		return
	}
	outputPDF := C.CString("./output.pdf")
	defer freeCString(outputPDF)

	C.qpdf_init_write(qpdf, outputPDF)

	op(qpdf)

	if C.qpdf_write(qpdf) != C.QPDF_SUCCESS {
		fmt.Println("写入 PDF 文件失败")
		return
	}
	fmt.Println("成功移除 PDF 权限并保存为新文件")
}

func OperateOnPDF(op Operater, originFile string, outputFile string) {
	qpdf := C.qpdf_init()
	defer C.qpdf_cleanup(&qpdf)

	version := C.GoString(C.qpdf_get_qpdf_version())
	fmt.Println("QPDF版本:", version)

	C.qpdf_empty_pdf(qpdf)

	outputFilename := C.CString(outputFile)
	defer C.free(unsafe.Pointer(outputFilename))

	op(qpdf, originFile)

	C.qpdf_init_write(qpdf, outputFilename)
	C.qpdf_write(qpdf)
}
