package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var setPathCmd = &cobra.Command{
	Use:   "set-path <path>",
	Short: "Set the path for FileSync CLI starts the backup of files and folders",
	Run:   setPath,
}

func setPath(cmd *cobra.Command, args []string) {
	for _, path := range args {
		if len(path) > 255 {
			log.Printf("[FileSync Error] The provided key exceeded size limit (permitted per path 255 bytes):  %s", path[0:255]+"...")
		}
		err := os.WriteFile(".filesync", []byte(path + "\n"), 0600)
		if err != nil {
			log.Printf("[FileSync Error] Could not save the path: %v", err)
		}
	}
}

func init() {
	RootCmd.AddCommand(setPathCmd)
}
