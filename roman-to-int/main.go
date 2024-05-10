package main

import (
	"fmt"
	"strings"
)

func main() {
	res := romanToInt("MCMXCIV")
	fmt.Println("res=", res)
}

/*  3ms
romVal := map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var total int = 0

for i := 0; i < len(s)-1; i++ {
	if romVal[s[i]] < romVal[s[i+1]] {
		total -= romVal[s[i]]
		continue
	}
	total += romVal[s[i]]
}
total += romVal[s[len(s)-1]]
return total
*/

// leetcode 9ms
func romanToInt(s string) int {
	romans := make(map[string]int)
	romans["I"] = 1
	romans["V"] = 5
	romans["X"] = 10
	romans["L"] = 50
	romans["C"] = 100
	romans["D"] = 500
	romans["M"] = 1000

	result := 0
	isStop := false
	isSkipInd := -1

	for k, v := range s {
		if k == isSkipInd {
			isSkipInd = -1
			continue
		}

		if k == len(s)-1 {
			isStop = true
		}

		switch letter := string(v); letter {
		case "I":
			if !isStop {
				nextLetter := string(s[k+1])
				if strings.Contains("VX", nextLetter) {
					isSkipInd = k + 1
					if nextLetter == "V" {
						result += 4
					} else {
						result += 9
					}
				} else {
					result += romans[string(v)]
				}
			} else {
				result += romans[string(v)]
			}
		case "X":
			if !isStop {
				nextLetter := string(s[k+1])
				if strings.Contains("LC", nextLetter) {
					isSkipInd = k + 1
					if nextLetter == "L" {
						result += 40
					} else {
						result += 90
					}
				} else {
					result += romans[string(v)]
				}
			} else {
				result += romans[string(v)]
			}
		case "C":
			if !isStop {
				nextLetter := string(s[k+1])
				if strings.Contains("DM", nextLetter) {
					isSkipInd = k + 1
					if nextLetter == "D" {
						result += 400

					} else {
						result += 900
					}
				} else {
					result += romans[string(v)]
				}
			} else {
				result += romans[string(v)]
			}
		default:
			result += romans[string(v)]
		}
	}
	return result
}
