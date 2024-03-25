package cmd

import (
	"julypdf/cpu"

	"github.com/spf13/cobra"
)

var watermark = &cobra.Command{
	Use:   "watermarker",
	Short: "PDF add watermark",
	Long:  `PDF add watermark.`,
	Run: func(cmd *cobra.Command, args []string) {
		cpu.AddWatermark(fileName, outFile, waterMarker)
	},
}

// qpdf.MeargeOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
func init() {
	rootCmd.AddCommand(watermark)
	watermark.Flags().StringVarP(&outFile, "output", "o", "", "Output path for PNG files")
}
