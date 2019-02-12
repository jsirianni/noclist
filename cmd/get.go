package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Performs GET request against the noc server",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
