package main

import "fmt"

func main() {
	strs := []string{"flower", "flower", "flower", "flower"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)
}

// leetcode 2ms
func longestCommonPrefix(strs []string) string {
	//pref := make(map[string]int)
	stop := false
	lastKey, pr := "", ""
	if len(strs) == 1 {
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ {
		pr = strs[0][:i+1]
		for k := range strs {
			if len(pr) > len(strs[k]) {
				lastKey = pr[:len(pr)-1]
				stop = true
				break
			}

			if pr != strs[k][:i+1] {
				lastKey = pr[:len(pr)-1]
				stop = true
				break
			}
		}
		if stop {
			break
		}
	}
	if !stop {
		lastKey = pr
	}
	return lastKey
}
