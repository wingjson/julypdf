package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "julypdf",
	Short: "PDF processing application",
	Long:  `A Fast and Flexible PDF processing tool built with love by Go and Cobra.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
