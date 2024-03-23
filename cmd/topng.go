package cmd

import (
	"fmt"
	"julypdf/mupdf"

	"github.com/spf13/cobra"
)

var fileName string
var outPath string
var waterMarker string

// toPngCmd represents the toPng command
var toPngCmd = &cobra.Command{
	Use:   "toPng",
	Short: "Converts PDF to PNG",
	Long:  `Converts a PDF file to PNG format with optional watermarking.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Converting PDF to PNG...")
		mupdf.ToPng(fileName, outPath, waterMarker)
	},
}

func init() {
	rootCmd.AddCommand(toPngCmd)
	toPngCmd.Flags().StringVarP(&fileName, "file", "f", "", "Source PDF file")
	toPngCmd.Flags().StringVarP(&outPath, "output", "o", "", "Output path for PNG files")
	toPngCmd.Flags().StringVarP(&waterMarker, "watermark", "w", "", "Watermark text (optional)")

}
