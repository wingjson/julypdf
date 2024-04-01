package cmd

import (
	"fmt"
	"julypdf/cpu"

	"github.com/spf13/cobra"
)

// toPngCmd represents the toPng command
var toPdfCmd = &cobra.Command{
	Use:   "topdf",
	Short: "Converts IMG to PDF",
	Long:  `Converts a IMG file to PDF format.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Converting IMG to PDF...")
		cpu.Img2Pdf(fileName, outPath)
	},
}

func init() {
	rootCmd.AddCommand(toPdfCmd)
	toPdfCmd.Flags().StringVarP(&outPath, "output", "o", "", "Output file for PDF files")
}
