//go:build !lib

package qpdf

/*
#cgo CFLAGS: -I../include
#cgo linux,amd64,!musl LDFLAGS: -lgcc_s -lstdc++ -ljpeg -lz  -lssl -lcrypto  -lqpdf
#cgo windows,amd64 LDFLAGS: -lgcc_s -lstdc++ -lmingwex -ljpeg -lz -lgnutls -lssl -lcrypto -lws2_32 -lqpdf
*/
import "C"
