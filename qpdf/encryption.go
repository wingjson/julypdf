package qpdf

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L../libs -lpdf
#include <stdlib.h>
#include "qpdf/qpdf-c.h"

// decryption
void decrypt(qpdf_data qpdf) {
    qpdf_set_r3_encryption_parameters_insecure(
        qpdf,
        "",
        "",
        QPDF_TRUE,
        QPDF_TRUE,
        QPDF_TRUE,
        QPDF_TRUE,
        QPDF_TRUE,
        QPDF_TRUE,
        qpdf_r3p_full
    );
}


// encryption
void encrypt(qpdf_data qpdf) {
    qpdf_set_r6_encryption_parameters2(
        qpdf,
        "",
        "",
        QPDF_FALSE,
        QPDF_FALSE,
        QPDF_FALSE,
        QPDF_FALSE,
        QPDF_FALSE,
        QPDF_FALSE,
        qpdf_r3p_none,
        QPDF_TRUE
	);
}
*/
import "C"

func Decryption(qpdf C.qpdf_data) {
	C.decrypt(qpdf)
}

func Encryption(qpdf C.qpdf_data) {
	C.encrypt(qpdf)
}