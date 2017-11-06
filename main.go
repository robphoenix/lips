package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	switch len(args) {
	case 2:
		if err := PrintCIDR(args[1]); err != nil {
			fmt.Fprintln(errOut, err)
		}
	case 3:
		if err := PrintRange(args[1], args[2]); err != nil {
			fmt.Fprintln(errOut, err)
		}
	}
}
