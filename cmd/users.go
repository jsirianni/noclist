package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Get a list of users from the noc api",
	Run: func(cmd *cobra.Command, args []string) {
		getUsers()
	},
}

func init() {
	getCmd.AddCommand(usersCmd)
}

func getUsers() {
	users, err := nocClient.GetUsers()
	if err != nil {
		// log the error to standard error
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	j, err := users.ToJson()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(string(j))
}
