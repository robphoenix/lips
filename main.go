package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		err := ListCIDR(args[1])
		if err != nil {
			fmt.Fprintln(errOut, err)
		}
	}

	if len(args) == 3 {
		err := ListRange(args[1], args[2])
		if err != nil {
			fmt.Fprintln(errOut, err)
		}
	}
}
