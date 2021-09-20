package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sql2StructCmd)
}

var (
	rootCmd = cobra.Command{
		Use: "",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
