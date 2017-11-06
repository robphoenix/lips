package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	out    io.Writer = os.Stdout
	errOut io.Writer = os.Stderr
)

// PrintCIDR ...
func PrintCIDR(cidr string, increment int) error {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	for i := ip.Mask(ipnet.Mask); ipnet.Contains(i); inc(i, increment) {
		fmt.Fprintln(out, i)
	}
	return nil
}

// PrintRange ...
func PrintRange(start, finish string, increment int) error {
	a := net.ParseIP(start)
	b := net.ParseIP(finish)
	if a == nil || b == nil {
		var e bytes.Buffer
		if a == nil {
			e.WriteString(fmt.Sprintf("invalid IP address: %s\n", start))
		}
		if b == nil {
			e.WriteString(fmt.Sprintf("invalid IP address: %s\n", finish))
		}
		return fmt.Errorf(e.String())
	}

	for {
		if c := bytes.Compare(a, b); c > 0 {
			break
		}
		fmt.Fprintln(out, a)
		inc(a, increment)
	}
	return nil
}

func inc(ip net.IP, increment int) {
	for j := len(ip) - 1; j >= 0; j-- {
		// ip[j]++
		ip[j] += byte(increment)
		if ip[j] > 0 {
			break
		}
	}
}
