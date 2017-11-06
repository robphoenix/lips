package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		cidr := args[1]
		if err := ListCIDR(cidr); err != nil {
			fmt.Fprintf(os.Stderr, "could not list IPs from %s: %v", cidr, err)
		}
	}
}
