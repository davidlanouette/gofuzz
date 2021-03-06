package main

import "fmt"

func Reverse2(s string) string {
	fmt.Printf("input: '%q'\n", s)
	r := []rune(s)
	fmt.Printf("runes: '%q'\n", r)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
