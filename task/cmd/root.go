package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd will attach all commands
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a CLI manager",
}
