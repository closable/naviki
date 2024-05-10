package main

import "fmt"

func main() {
	//nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 3}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}

// leetcode 27ms
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	// solution < 10ms
	// var m = make(map[int]int, len(nums))
	// for i := 0 ; i < len(nums); i++ {
	//     val, ok := m[nums[i]]
	//     if ok {
	//         return []int{val, i}
	//     }
	//     m[(nums[i] - target) * -1] = i
	// }

	return nil
}
