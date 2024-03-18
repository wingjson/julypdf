package main

/*
#cgo CFLAGS: -I./include
#cgo LDFLAGS: -L./libs -lmupdf -lmupdfthird -lm -lcomdlg32 -lgdi32 -lgcc -Wl,--allow-multiple-definition
#include <stdlib.h>
#include "mupdf/fitz.h"

const char* getMupdfVersion() {
    return FZ_VERSION;
}

int add_watermark_to_pdf()
{
    fz_context *ctx;
    fz_document *doc;
    fz_page *page;
    fz_device *dev;
    fz_matrix ctm = { 1, 0, 0, 1, 0, 0 }; // 无旋转、无缩放的变换矩阵
    fz_rect bounds;
    fz_stext_page *text_page;
    fz_stext_options options;
    fz_font *font;
    char *input_file = "input.pdf";
    char *output_file = "output.pdf";
    char *watermark_text = "Confidential";
    int page_number = 0; // 页码从0开始

    ctx = fz_new_context(NULL, NULL, FZ_STORE_UNLIMITED);
    if (!ctx) { printf("Cannot create mupdf context\n"); return 1; }

    fz_try(ctx)
    {
        doc = fz_open_document(ctx, input_file);
        page = fz_load_page(ctx, doc, page_number);
        bounds = fz_bound_page(ctx, page);
// fz_device *fz_new_draw_device(fz_context *ctx, fz_matrix transform, fz_pixmap *dest);
        dev = fz_new_draw_device(ctx, &ctm, page);
        font = fz_new_font_from_name(ctx, NULL, "Helvetica", 0, FZ_FLAGS_NONE);

        fz_begin_page(ctx, dev, bounds, &ctm);
        fz_fill_text(ctx, dev, fz_new_text_line(ctx, font, watermark_text, &ctm, 0), &ctm, fz_device_rgb(ctx), 0.5); // 0.5透明度
        fz_end_page(ctx);

        fz_close_document(ctx, doc);
    }
    fz_always(ctx)
    {
        fz_drop_context(ctx);
    }
    fz_catch(ctx)
    {
        fprintf(stderr, "MuPDF error: %s\n", fz_caught_message(ctx));
    }

    return 0;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	version := C.GoString(C.getMupdfVersion())
	fmt.Println("MuPDF version:", version)
	inputPdf := C.CString("1.pdf")
	outputPdf := C.CString("2.pdf")
	defer C.free(unsafe.Pointer(inputPdf))
	defer C.free(unsafe.Pointer(outputPdf))

	C.add_watermark_to_pdf()
}
