package cmd

import (
	"julypdf/qpdf"

	"github.com/spf13/cobra"
)

var decrypt = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt pdf",
	Long:  `decrypt -f file.pdf -o output.pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		qpdf.CryptoOnPDF(qpdf.Encryption, fileName, outFile, 2)
	},
}

// qpdf.MeargeOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
func init() {
	rootCmd.AddCommand(decrypt)
	decrypt.Flags().StringVarP(&outFile, "output", "o", "", "Output file for decrypt file")
}
