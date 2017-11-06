package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	// Version of current release
	Version = "v0.0.1"
)

var (
	version   bool
	increment int
	usageText = `  _   _   _   _  
 / \ / \ / \ / \ 
( l | i | p | s )
 \_/ \_/ \_/ \_/   %s

     List IPs
                           
Usage:
	lips [options] <CIDR IP>
	lips [options] <start IP> <end IP>

Examples:
	lips -i 256 10.8.0.0/16
	lips 192.168.0.1 192.168.0.47

Options:
	-v  print version and exit
	-i  change value of increment (default 1)
`
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(usageText, Version))
	}
	flag.IntVar(&increment, "i", 1, "value to increment by")
	flag.BoolVar(&version, "v", false, "print version and exit")
	flag.Parse()
}

func main() {
	if version {
		fmt.Println(Version)
		return
	}
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
	default:
		fmt.Fprintf(os.Stderr, fmt.Sprintf(usageText, Version))
	}
}
