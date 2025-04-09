package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var setKeyCmd = &cobra.Command{
	Use:   "set-key <key>",
	Short: "Set the installation key for FileSync CLI starts the backup source environment configurations",
	Run:   setKey,
}

func setKey(cmd *cobra.Command, args []string) {
	fileSyncKeyFile := ".filesync.key"
	err := os.WriteFile(fileSyncKeyFile, []byte(""), 0600)
	if err != nil {
		log.Print(err)
	}
	key := args[0]
	if len(key) > 255 {
		log.Print("[FileSync Error] The provided key exceeded size limit (permitted per key 255 bytes).")
	}
	err = os.WriteFile(fileSyncKeyFile, []byte(key), 0600)
	if err != nil {
		log.Print(err)
	}
}

func init() {
	RootCmd.AddCommand(setKeyCmd)
}
