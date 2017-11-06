package main

import (
	"net"
)

// ListCIDR ...
func ListCIDR(cidr string) (ips []net.IP, err error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return ips, err
	}
	for i := ip.Mask(ipnet.Mask); ipnet.Contains(i); inc(i) {
		ips = append(ips, dupIP(i))
	}
	return ips, nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func dupIP(ip net.IP) net.IP {
	// To save space, try and only use 4 bytes
	if x := ip.To4(); x != nil {
		ip = x
	}
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}
