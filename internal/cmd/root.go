package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "filesync",
	Short: "filesync is a CLI tool to prepare backup source environment inside machines, VMs, containers and others to upload files and disks for FileSync Cloud, our backup application.",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
	fmt.Println("Root command runned.")
}
