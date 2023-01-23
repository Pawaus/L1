package L1_2

import (
	"fmt"
	"sync"
)

//calculation of squares in concurrent place
func Square() {
	//Wait Group for waiting end of gorutines
	var wg sync.WaitGroup
	//array of int's
	var Mas []int = []int{2, 4, 6, 8, 10}
	//Will wait for n times of calling Done from gorutines
	wg.Add(len(Mas))
	//run n gorutines(n - count of numbers in array)
	for _, v := range Mas {
		//start gorutine
		go PrintSquare(v, &wg)
	}
	//wait for ending all gorutines
	wg.Wait()
}

//func to print square number
func PrintSquare(i int, wg *sync.WaitGroup) {
	//Print
	fmt.Println(i * i)
	//Say to WaitGroup of end of func and printing
	wg.Done()
}
