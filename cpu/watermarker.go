package cpu

import (
	"bytes"
	"fmt"
	"julypdf/ignoreerror"
	"julypdf/qpdf"
	"julypdf/utilts"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/font"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func AddWatermark(originfileName string, targetFileName string, waterMark string) {
	utilts.ExtractFont("SourceHanSansCN-Regular.ttf")
	// utilts.ExtractFont("SourceHanSansCN-Regular.ttf")
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v", err)
	}
	fontsDir := filepath.Join(userConfigDir, "julypdf", "fonts")
	Configuration := model.NewDefaultConfiguration()
	Configuration.Path = fontsDir
	// println(fontsDir)
	fontDir := fontsDir
	fontPath := filepath.Join(fontsDir, "SourceHanSansCN-Regular.ttf")
	// fontPath := filepath.Join(fontsDir, "SourceHanSansCN-Regular.ttf")

	font.UserFontDir = fontDir

	err = font.InstallTrueTypeFont(fontDir, fontPath)
	if err != nil {
		panic(err)
	}
	err = font.LoadUserFonts()
	if err != nil {
		panic(err)
	}
	start := time.Now()
	fmt.Printf("start: %s\n", start)

	//first decrypt
	tempFile := decrypt(originfileName, targetFileName)
	if tempFile == "" {
		return
	}

	content, err := os.ReadFile(tempFile)
	if err != nil {
		fmt.Printf("cannot read %s: %v\n", tempFile, err)
		return
	}

	if err := os.Remove(tempFile); err != nil {
		fmt.Printf("cannot remove input file: %v\n", err)
		return
	}

	inputFileReader := bytes.NewReader(content)

	outputFile, err := os.Create(targetFileName)
	if err != nil {
		fmt.Printf("cannot create %s: %v\n", targetFileName, err)
		return
	}
	defer outputFile.Close()

	onTop := false
	update := false
	wm, err := api.TextWatermark(waterMark, "font:SourceHanSansCN-Regular", onTop, update, types.POINTS)
	if err != nil {
		fmt.Printf("cannot create watermark: %v\n", err)
		return
	}

	//add watermark
	if err := api.AddWatermarks(inputFileReader, outputFile, nil, wm, nil); err != nil {
		fmt.Printf("cannot add watermark: %v\n", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("mission complete: %s\n", elapsed)
}

// because sometime pdfcpu maybe decrypt failed,so use qpdf to decrypt
func decrypt(originfileName string, targetFileName string) string {
	tempFile := dealFileName(originfileName, targetFileName)
	result, err := qpdf.CryptoOnPDF(qpdf.Decryption, originfileName, tempFile, 2)
	if err != nil {
		if _, ok := err.(*ignoreerror.IgnorableError); ok {
			return result
		} else {
			return ""
		}
	}
	return result
}

// generate temp file
func dealFileName(originFileName string, targetFileNameExt string) string {
	baseName := strings.TrimSuffix(path.Base(originFileName), path.Ext(originFileName))
	outputFilename := fmt.Sprintf("%s%s_tempencrypted.pdf", baseName, targetFileNameExt)
	return outputFilename
}
