package main

import (
	"fmt"
	"strings"
)

func main() {
	a, b := "111", "1010"
	res := addBinary(a, b)
	fmt.Println(res)
}

// leetcode 1ms
func addBinary(a string, b string) string {
	res := ""
	overflow := 0
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	} else {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}

	for i := len(b) - 1; i >= 0; i-- {
		dig := ""
		if len(a)-1 >= i {
			if string(a[i])+string(b[i]) == "11" {
				if overflow == 1 {
					dig = "1"
				} else {
					dig = "0"
				}
				overflow = 1
			} else if string(a[i])+string(b[i]) == "00" {
				if overflow == 1 {
					dig = "1"
				} else {
					dig = "0"
				}
				overflow = 0

			} else {
				if overflow == 1 {
					dig = "0"
					overflow = 1
				} else {
					dig = "1"
					overflow = 0
				}
			}
		}
		res = dig + res
	}

	if overflow == 1 {
		res = "1" + res
	}

	return res
}
