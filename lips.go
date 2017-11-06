package main

import (
	"fmt"
	"net"
)

// ListCIDR ...
func ListCIDR(addr string) error {
	ip, ipnet, err := net.ParseCIDR(addr)
	if err != nil {
		return err
	}
	for i := ip.Mask(ipnet.Mask); ipnet.Contains(i); inc(i) {
		fmt.Println(i)
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
