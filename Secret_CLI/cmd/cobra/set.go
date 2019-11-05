package cobra

import (
	"fmt"

	"github.com/shubhanshu7/Gophercises/Secret_CLI"
	"github.com/spf13/cobra"
)

func set(cmd *cobra.Command, args []string) {
	v := Secret_CLI.File(encodingKey, secretsPath())
	key, value := args[0], args[1] /*"check", "it"*/

	// fmt.Println("Key", key, "Value:", value, "Enc Key", encodingKey, "SecretPath", secretsPath())

	err := v.Set(key, value)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("Value set successfully!")
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run:   set,
}

func init() {
	RootCmd.AddCommand(setCmd)
}
