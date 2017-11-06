package main

import (
	"bytes"
	"testing"
)

func TestPrintCIDR(t *testing.T) {
	cidr := "192.168.0.0/30"
	expected := "192.168.0.0\n192.168.0.1\n192.168.0.2\n192.168.0.3\n"
	var actual bytes.Buffer
	originalOut := out
	out = &actual
	defer func() { out = originalOut }()

	err := PrintCIDR(cidr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual.String() != expected {
		t.Fatalf("expected:\n\t%q\nactual:\n\t%q", expected, actual.String())
	}
}

func TestPrintRange(t *testing.T) {
	start := "192.168.0.0"
	finish := "192.168.0.3"
	expected := "192.168.0.0\n192.168.0.1\n192.168.0.2\n192.168.0.3\n"
	var actual bytes.Buffer
	originalOut := out
	out = &actual
	defer func() { out = originalOut }()

	err := PrintRange(start, finish)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if actual.String() != expected {
		t.Fatalf("expected:\n\t%q\nactual:\n\t%q", expected, actual.String())
	}
}
