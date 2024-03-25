package cmd

import (
	"julypdf/qpdf"

	"github.com/spf13/cobra"
)

var splitPage int
var split = &cobra.Command{
	Use:   "split",
	Short: "split pdf",
	Long:  `split -f file1.pdf -p 2(page you want to split)`,
	Run: func(cmd *cobra.Command, args []string) {
		qpdf.SplitOnPDF(qpdf.Split, fileName, splitPage)
	},
}

// qpdf.SplitOnPDF(qpdf.Split, "555.pdf", 1)
func init() {
	rootCmd.AddCommand(split)
	split.Flags().IntVarP(&splitPage, "split", "s", 0, "Number of pages to split")
}
