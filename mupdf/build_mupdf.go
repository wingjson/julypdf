//go:build !lib

package mupdf

/*
#cgo CFLAGS: -I../include
#cgo linux,amd64,!musl LDFLAGS: -L../libs/linux -lmupdf -lmupdfthird -lm -lcomdlg32 -lgdi32 -lgcc -Wl,--allow-multiple-definition
#cgo windows,amd64 LDFLAGS: -L../libs/win64 -lmupdf -lmupdfthird -lm -lcomdlg32 -lgdi32 -lcrypto -lgcc -Wl,--allow-multiple-definition
*/
import "C"
