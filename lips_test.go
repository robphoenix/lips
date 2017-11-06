package main

import (
	"fmt"
	"net"
	"testing"
)

func TestListCIDR(t *testing.T) {
	cidr := "192.168.0.0/30"
	expected := []net.IP{net.ParseIP("192.168.0.0"), net.ParseIP("192.168.0.1"), net.ParseIP("192.168.0.2"), net.ParseIP("192.168.0.3")}
	actual, err := ListCIDR(cidr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fmt.Sprintf("%s", actual) != fmt.Sprintf("%s", expected) {
		t.Fatalf("expected:\n\t%q\nactual:\n\t%q", expected, actual)
	}
}
