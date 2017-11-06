package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

var (
	out    io.Writer = os.Stdout
	errOut io.Writer = os.Stderr
)

// ListCIDR ...
func ListCIDR(cidr string) error {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	for i := ip.Mask(ipnet.Mask); ipnet.Contains(i); inc(i) {
		fmt.Fprintln(out, i)
	}
	return nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
