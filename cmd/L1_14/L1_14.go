package L1_14

import "fmt"

func Responce(l1 interface{}) {
	//%T print type of value
	fmt.Printf("%T", l1)
}

func Fourteen() {
	//any variable is interface, transfer it to function
	Responce(123)
}
