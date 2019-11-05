package cobra

import (
	"fmt"

	"github.com/shubhanshu7/Gophercises/Secret_CLI"
	"github.com/spf13/cobra"
)

func get(cmd *cobra.Command, args []string) {
	v := Secret_CLI.File(encodingKey, secretsPath())
	key := args[0]
	value, err := v.Get(key)
	if err != nil {
		fmt.Println("no value set")
		// os.Exit(1)
	}
	fmt.Printf("%s = %s\n", key, value)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret in your secret storage",
	Run:   get,
}

func init() {
	RootCmd.AddCommand(getCmd)
}
