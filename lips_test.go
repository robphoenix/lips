package main

import (
	"bytes"
	"testing"
)

var cidrTestCases = []struct {
	cidr        string
	increment   int
	expected    string
	description string
}{
	{
		cidr:        "192.168.0.0/30",
		expected:    "192.168.0.0\n192.168.0.1\n192.168.0.2\n192.168.0.3\n",
		increment:   1,
		description: "normal increment",
	},
	{
		cidr:        "192.168.0.0/28",
		expected:    "192.168.0.0\n192.168.0.2\n192.168.0.4\n192.168.0.6\n192.168.0.8\n192.168.0.10\n192.168.0.12\n192.168.0.14\n",
		increment:   2,
		description: "even increment",
	},
	{
		cidr:        "192.168.0.0/28",
		expected:    "192.168.0.0\n192.168.0.3\n192.168.0.6\n192.168.0.9\n192.168.0.12\n192.168.0.15\n",
		increment:   3,
		description: "odd increment",
	},
}

func TestPrintCIDR(t *testing.T) {
	var actual bytes.Buffer
	originalOut := out
	out = &actual
	defer func() { out = originalOut }()

	for _, test := range cidrTestCases {
		actual.Reset()
		err := PrintCIDR(test.cidr, test.increment)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if actual.String() != test.expected {
			t.Errorf("PrintCIDR(%s, %d) [%s]\nexpected:\n\t%q\nactual:\n\t%q", test.cidr, test.increment, test.description, test.expected, actual.String())
		}
	}
}

var rangeTestCases = []struct {
	start       string
	finish      string
	increment   int
	expected    string
	description string
}{
	{
		start:       "192.168.0.0",
		finish:      "192.168.0.3",
		increment:   1,
		expected:    "192.168.0.0\n192.168.0.1\n192.168.0.2\n192.168.0.3\n",
		description: "normal increment",
	},
	{
		start:       "192.168.0.0",
		finish:      "192.168.0.6",
		increment:   2,
		expected:    "192.168.0.0\n192.168.0.2\n192.168.0.4\n192.168.0.6\n",
		description: "even increment, even finish",
	},
	{
		start:       "192.168.0.0",
		finish:      "192.168.0.7",
		increment:   2,
		expected:    "192.168.0.0\n192.168.0.2\n192.168.0.4\n192.168.0.6\n",
		description: "even increment, odd finish",
	},
	{
		start:       "192.168.0.0",
		finish:      "192.168.0.9",
		increment:   3,
		expected:    "192.168.0.0\n192.168.0.3\n192.168.0.6\n192.168.0.9\n",
		description: "odd increment, odd finish",
	},
	{
		start:       "192.168.0.0",
		finish:      "192.168.0.10",
		increment:   3,
		expected:    "192.168.0.0\n192.168.0.3\n192.168.0.6\n192.168.0.9\n",
		description: "odd increment, even finish",
	},
}

func TestPrintRange(t *testing.T) {
	var actual bytes.Buffer
	originalOut := out
	out = &actual
	defer func() { out = originalOut }()

	for _, test := range rangeTestCases {
		actual.Reset()
		err := PrintRange(test.start, test.finish, test.increment)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if actual.String() != test.expected {
			t.Errorf("PrintRange(%s, %s, %d) [%s]\nexpected:\n\t%q\nactual:\n\t%q", test.start, test.finish, test.increment, test.description, test.expected, actual.String())
		}
	}
}
