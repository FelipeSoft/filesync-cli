package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var setKeyCmd = &cobra.Command{
	Use:   "set-key <key>",
	Short: "Set the installation key for FileSync CLI starts the backup source environment configurations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Service started.")
	},
}

func init() {
	RootCmd.AddCommand(setKeyCmd)
}
