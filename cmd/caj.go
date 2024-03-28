package cmd

import (
	"julypdf/caj"

	"github.com/spf13/cobra"
)

var cajpdf = &cobra.Command{
	Use:   "caj",
	Short: "convert caj to pdf",
	Long:  `caj -f file.caj -o output.pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		caj.Caj2pdf(fileName, outFile)
	},
}

// qpdf.MeargeOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
func init() {
	rootCmd.AddCommand(cajpdf)
	cajpdf.Flags().StringVarP(&outFile, "output", "o", "", "Output file for caj file")
}
