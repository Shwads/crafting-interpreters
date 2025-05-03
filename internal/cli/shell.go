package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Shell() error {
	var builder strings.Builder
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(">> ")
		scanner.Scan()
		if scanner.Text() == "exit" {
			return nil
		}
		fmt.Fprintf(&builder, "%s", scanner.Text())
		if scanner.Text() == "" {
			fmt.Printf("%s\n", builder.String())
			builder.Reset()
		}
	}
}
