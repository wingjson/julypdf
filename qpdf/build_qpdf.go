//go:build !lib

package qpdf

/*
#cgo CFLAGS: -I../include
#cgo linux,amd64,!musl LDFLAGS: -lgcc_s -lstdc++ -ljpeg -lz  -lssl -lcrypto -lgnutls -lqpdf
#cgo windows,amd64 LDFLAGS: -lgcc_s -lstdc++ -lqpdf  -lmingwex -ljpeg -lz -lgnutls -lssl -lcrypto -lws2_32 -LD:/Code/mysy2/ucrt64/lib
*/
import "C"
