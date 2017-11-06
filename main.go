package main

import (
	"flag"
	"fmt"
)

func init() {
	flag.IntVar(&increment, "i", 1, "value to increment by")
	flag.Parse()
}

func main() {
	args := flag.Args()

	switch len(args) {
	case 1:
		if err := PrintCIDR(args[0], increment); err != nil {
			fmt.Fprintln(errOut, err)
		}
	case 2:
		if err := PrintRange(args[0], args[1], increment); err != nil {
			fmt.Fprintln(errOut, err)
		}
	}
}
