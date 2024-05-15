package main

import (
	"fmt"
)

func main() {
	digits := []int{1, 9, 9}
	res := plusOne(digits)
	fmt.Println(res)
}

// leetcode 1ms
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	lastDigit := digits[len(digits)-1]
	if lastDigit < 9 {
		digits[len(digits)-1] = digits[len(digits)-1] + 1
		return digits
	} else {
		isUpdated := false
		for i := len(digits) - 1; i >= 0; i-- {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				isUpdated = true
				digits[i] = digits[i] + 1
				break
			}
		}
		if !isUpdated {
			digits = append([]int{1}, digits...)
		}

	}

	return digits
}
