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
		i = inc(i, increment)
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
		a = inc(a, increment)
	}
	return nil
}

func int32toIP(a int32) net.IP {
	return net.IPv4(byte(a>>24), byte(a>>16), byte(a>>8), byte(a))
}

func IPtoInt32(ip net.IP) int32 {
	ip = ip.To4()
	return int32(ip[0])<<24 | int32(ip[1])<<16 | int32(ip[2])<<8 | int32(ip[3])
}

func inc(ip net.IP, increment int) net.IP {
	a := int(IPtoInt32(ip[len(ip)-4:]))
	return int32toIP(int32(a + increment))
}
