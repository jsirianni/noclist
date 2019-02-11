package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var file string
var port string
var host string
var tls  bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "noclist",
	Short: "manage the noc server",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "hostname or ip address of the target server")
	rootCmd.PersistentFlags().StringVar(&port, "port", "port", "tcp port for the target server")
	rootCmd.PersistentFlags().BoolVar(&tls, "tls", false, "enable tls")
}
