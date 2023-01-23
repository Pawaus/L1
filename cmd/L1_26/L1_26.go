package L1_26

import (
	"fmt"
	"strings"
)

func TwentySix(inp string) {
	// Store chars
	m := make(map[rune]int)
	// make lower all chars in string
	res := strings.ToLower(inp)
	//if char was in string, increment number of chars
	for _, v := range res {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 0
		}
	}
	//if more then one time char in string, return false, else return true
	for _, v := range m {
		if v > 0 {
			fmt.Print("false")
			return
		}
	}
	fmt.Print("true")

}
