package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2

	res := searchInsert(nums, target)
	fmt.Println(res)
}

// leetcode 3ms
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	return l

}
