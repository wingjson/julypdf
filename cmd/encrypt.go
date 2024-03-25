package cmd

import (
	"julypdf/qpdf"

	"github.com/spf13/cobra"
)

var encrypt = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypt pdf",
	Long:  `encrypt -f file.pdf -o output.pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		qpdf.CryptoOnPDF(qpdf.Encryption, fileName, outFile, 1)
	},
}

// qpdf.MeargeOnPDF(qpdf.Merge, "1.pdf 2.pdf 3.pdf", "555.pdf")
func init() {
	rootCmd.AddCommand(encrypt)
	encrypt.Flags().StringVarP(&outFile, "output", "o", "", "Output file for encrypt file")
}
