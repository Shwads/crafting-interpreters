package main

import (
	"fmt"
	"os"
	"github.com/Shwads/crafting-interpreters/internal/cli"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Printf("Usage glox [script]")
		os.Exit(1)
	} else if len(args) == 1 {
		err := cli.HandleFile(args[0])
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	} else {
		err := cli.Shell()
		if err != nil {
			fmt.Printf("%s", err.Error())
		}
	}
}
