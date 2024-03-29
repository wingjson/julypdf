package utilts

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//go:embed SourceHanSansCN-Regular.ttf
var embeddedTTF embed.FS

func checkAndExtractFont(targetDir, fontName string) error {
	targetPath := filepath.Join(targetDir, fontName)

	if _, err := os.Stat(targetPath); err == nil {
		log.Printf("Font file already exists: %s", targetPath)
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}
	var fontFile fs.File
	var err error
	fontFile, err = embeddedTTF.Open(fontName)
	if err != nil {
		return err
	}

	defer fontFile.Close()

	outFile, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, fontFile)
	return err
}

func ExtractFont(fontName string) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting user home directory: %v", err)
	}

	fontsDir := filepath.Join(userConfigDir, "julypdf", "fonts")
	if err := os.MkdirAll(fontsDir, 0755); err != nil {
		log.Fatalf("Error creating fonts directory: %v", err)
	}
	if err := checkAndExtractFont(fontsDir, fontName); err != nil {
		log.Fatalf("Error checking and extracting font: %v", err)
	}

	// log.Printf("Font successfully checked and extracted to %s", fontsDir)
}
