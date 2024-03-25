package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fileName string
var outPath string
var outFile string
var waterMarker string
var concurrent bool

var rootCmd = &cobra.Command{
	Use:   "julypdf",
	Short: "PDF processing application",
	Long:  `A Fast and Flexible PDF processing tool built with love by Go and Cobra.`,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "Source PDF file")
	rootCmd.PersistentFlags().StringVarP(&waterMarker, "watermark", "w", "", "Watermark text")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
