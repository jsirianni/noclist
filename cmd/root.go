package cmd

import (
	"fmt"
	"os"

	"noclist/noc"

	"github.com/spf13/cobra"
)

// client object
var nocClient noc.Noc

// command line flags
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
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "hostname or ip address of the target server")
	rootCmd.PersistentFlags().StringVar(&port, "port", "8888", "tcp port for the target server")
	rootCmd.PersistentFlags().BoolVar(&tls, "tls", false, "enable tls")
}

func initConfig() {
	var err error

	// set the host address
	err = nocClient.InitNoc(host, port, tls)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// aquire a token on startup
	err = nocClient.SetAuth()
	if err != nil {
		// just exit, standard error already has the errors printed
		os.Exit(1)
	}
}
