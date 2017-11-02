package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// https://groups.google.com/forum/#!topic/golang-nuts/zlcYA4qk-94

func main() {
	addr := os.Args[1]
	ip, ipnet, err := net.ParseCIDR(addr)
	if err != nil {
		log.Fatal(err)
	}
	for i := ip.Mask(ipnet.Mask); ipnet.Contains(i); inc(i) {
		fmt.Println(i)
	}
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
