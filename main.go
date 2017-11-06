package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		cidr := args[1]
		err := ListCIDR(cidr)
		if err != nil {
			fmt.Fprintln(errOut, err)
		}
	}
}
