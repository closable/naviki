package main

import (
	"fmt"
	"sort"
)

func main() {
	list2 := []int{1, 7, 7}
	list1 := []int{5, 6, 8, 9, 10}
	res := mergeTwoLists(list1, list2)
	fmt.Println(res)
}

func mergeTwoLists(list1 []int, list2 []int) []int {
	res := make([]int, 0)

	for i := 0; i < len(list1); i++ {
		if i < len(list2) {
			a, b := list1[i], list2[i]

			if i == 0 {
				if a > b {
					res = append(res, b, a)
				} else {
					res = append(res, a, b)
				}
			} else {
				a := setOrderSlice(res[len(res)-1], a, b)
				res = append(res[:len(res)-1], a...)
			}

		} else {
			res = append(res, list1[i])
		}
	}
	if len(list2) > len(list1) {
		res = append(res, list2[len(list1):]...)
	}

	return res
}

func setOrderSlice(a, b, c int) []int {
	x := make([]int, 3)

	x[0] = a
	x[1] = b
	x[2] = c

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	// fmt.Println("!!!", a, b, c, x)
	return x
}
