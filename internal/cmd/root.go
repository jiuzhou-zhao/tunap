package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tools",
	Short: "tools for vpn",
	Long:  "tools for vpn",
	Args:  cobra.MinimumNArgs(1),
}

func Execute() error {
	return rootCmd.Execute()
}
