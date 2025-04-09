package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/FelipeSoft/filesync-cli/internal/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ctx, stop := signal.NotifyContext(ctx, os.Kill, os.Interrupt)
	defer stop()

	fmt.Println("Welcome to FileSync CLI!")
	fmt.Println("Type a command or 'exit' to quit.")

	reader := bufio.NewReader(os.Stdin)

	go func() {
		fileSyncJsonNotExists := !fileExists("./.filesync.json")
		if fileSyncJsonNotExists {
			os.Create("./.filesync.json")
		}

		fileSyncKeyNotExists := !fileExists("./.filesync.key")
		if fileSyncKeyNotExists {
			os.Create("./.filesync.key")
		}

		fileSyncNotExists := !fileExists("./.filesync")
		if fileSyncNotExists {
			os.Create("./.filesync")
		}

		for {
			fmt.Print("filesync>")
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("[FileSync Error] Error reading input:", err)
				continue
			}

			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			if line == "exit" {
				cancel()
				fmt.Println("Bye!")
				break
			}

			args := strings.Fields(line)
			cmd.RootCmd.SetArgs(args)
			err = cmd.RootCmd.Execute()
			if err != nil {
				fmt.Printf("[FileSync Error] Command error: %s\n", err.Error())
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("Exited")
	stop()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}
