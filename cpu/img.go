package cpu

import (
	"github.com/signintech/gopdf"
)

func Img2Pdf(fileName string, outFile string) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	pdf.Image(fileName, 200, 50, nil)
	pdf.SetXY(250, 200)
	pdf.WritePdf(outFile)
}
