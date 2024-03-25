package cmd

import (
	"julypdf/qpdf"

	"github.com/spf13/cobra"
)

var merge = &cobra.Command{
	Use:   "merge",
	Short: "Merge pdf",
	Long:  `merge -f file1.pdf file2.pdf file3.pdf -o output.pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		qpdf.MeargeOnPDF(qpdf.Merge, fileName, outFile)
	},
}

// qpdf.MeargeOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
func init() {
	rootCmd.AddCommand(merge)
	merge.Flags().StringVarP(&outFile, "output", "o", "", "Output file for merge file")
}
