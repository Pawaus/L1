package L1_8

import (
	"fmt"
	"math"
)

func Eight(state bool) {
	//read number
	var n int64
	fmt.Scan(&n)
	//read bit number
	var i int
	fmt.Scan(&i)
	//if state is true, set the bit in true, else set in false
	st := int64(math.Pow(2, float64(i)))
	if state {
		n += st
	} else {
		n -= st
	}
	fmt.Print(n)
}
