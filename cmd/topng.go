package cmd

import (
	"fmt"
	"julypdf/mupdf"

	"github.com/spf13/cobra"
)

// toPngCmd represents the toPng command
var toPngCmd = &cobra.Command{
	Use:   "topng",
	Short: "Converts PDF to PNG",
	Long:  `Converts a PDF file to PNG format with optional watermarking.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Converting PDF to PNG...")
		mupdf.ToPng(fileName, outPath, concurrent, waterMarker)
	},
}

func init() {
	rootCmd.AddCommand(toPngCmd)
	toPngCmd.Flags().StringVarP(&outPath, "output", "o", "", "Output path for PNG files")
	toPngCmd.Flags().BoolVarP(&concurrent, "concurrent", "c", false, "Enable concurrent processing")
}
