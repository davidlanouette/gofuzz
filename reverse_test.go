package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		doubleRev := Reverse(rev)
		t.Logf("orig='%s', rev='%s', doublerev='%v'", tc.in, rev, doubleRev)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
		if doubleRev != tc.in {
			t.Errorf("Before: %q, after: %q", tc.in, doubleRev)
		}
	}
}
