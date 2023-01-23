package L1_19

import "fmt"

func reverseString(s string) (result string) {
	for _, val := range s {
		result = string(val) + result
	}
	return
}

func Nineteen() {
	str := "главрыба"
	fmt.Println("Reversed string:", reverseString(str))
}
