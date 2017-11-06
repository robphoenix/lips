package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		cidr := args[1]
		ips, err := ListCIDR(cidr)
		if err != nil {
			fmt.Printf("could not list IPs from %s: %v", cidr, err)
		}
		for _, ip := range ips {
			fmt.Println(ip)
		}
	}
}
