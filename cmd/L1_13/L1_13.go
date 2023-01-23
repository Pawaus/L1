package L1_13

import "fmt"

func Thirteen() {
	a := 10
	b := 20
	//change values of variables
	a, b = b, a
	fmt.Print(a, " ", b, "\n")
}
