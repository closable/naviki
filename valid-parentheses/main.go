package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "(){}}{"
	res := isValid(s)
	fmt.Println(s, res)
}

// leetcode 1ms
func isValid(s string) bool {
	types := make(map[string]int)
	types["("] = 1
	types[")"] = -1
	types["{"] = 2
	types["}"] = -2
	types["["] = 3
	types["]"] = -3
	res := make([]int, 0)
	par := strings.Split(s, "")
	for i := 0; i < len(par); i++ {
		if i == 0 && types[par[i]] < 0 {
			return false
		}

		if types[par[i]] > 0 {
			res = append(res, types[par[i]])
			//fmt.Println("0000", i, res)
		} else {
			if len(res) > 0 && res[len(res)-1]+types[par[i]] == 0 {
				res = res[:len(res)-1]
				//fmt.Println("1111", i, res, types[par[i-1]], types[par[i]])
			} else {
				//fmt.Println("2222", i, res, types[par[i]])
				return false
			}
		}
	}

	return len(res) == 0
}
