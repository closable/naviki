package main

import (
	"fmt"
)

func main() {
	res := isPalindrome(1222)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	// 8 ms
	tmp := x
	rev, rem := 0, 0
	for tmp > 0 {
		rem = tmp % 10
		rev = rev*10 + rem
		tmp /= 10
	}
	return rev == x
	// leetcode 28 ms
	// v := strconv.Itoa(x)
	// length := len(v) / 2
	// next := 0
	// if len(v)%2 != 0 {
	// 	next += 1
	// }
	// for i := 0; i < length; i++ {
	// 	//fmt.Println(i, length, length-1-i, string(v[i]), string(v[len(v)/2:][length-1-i]))
	// 	if v[i] != v[length+next:][length-1-i] {
	// 		return false
	// 	}
	// }

	//return true
}
