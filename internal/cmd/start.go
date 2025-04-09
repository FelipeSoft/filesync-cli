package cmd

import (
	"bufio"
	"log"
	"os"

	chunk "github.com/FelipeSoft/filesync-cli/internal/domain"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Before start backup your files, you need to set your application key and the path to backup",
	Run:   start,
}

func start(cmd *cobra.Command, args []string) {
	log.Print("Service started.")
	pathsFile, err := os.Open(".filesync")
	if err != nil {
		log.Printf("[FileSync Error] Could not open the file with paths to backup of FileSyn: %v", err)
	}
	defer pathsFile.Close()

	s := bufio.NewScanner(pathsFile)
	for s.Scan() {
		line := s.Text()
		log.Print(line)
		err := chunk.ProcessFileInChunks(line, 10_000_000)
		if err != nil {
			log.Printf("[FileSync Error] Could not read the file chunk: %v", err)
		}
	}
}

func init() {
	RootCmd.AddCommand(startCmd)
}
