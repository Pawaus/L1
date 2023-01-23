package L1_3

import (
	"fmt"
	"sync"
)

//func to print sum of squares
func SumSquare() {
	//Use mutex to lock a shared for gorutines variable
	var mtx sync.Mutex
	//wait for all gorutines
	var wg sync.WaitGroup
	//array of numbers
	var Mas []int = []int{2, 4, 6, 8, 10}
	//Will wait n threads(n - count of numbers in array)
	wg.Add(len(Mas))
	//shared variable
	var Sum int = 0
	for _, v := range Mas {
		//start gorutines and give a pointer of shared variable
		go doSquare(&Sum, v, &wg, &mtx)
	}
	//wait for all gorutines
	wg.Wait()
	//print result of sum squares
	fmt.Println(Sum)
}

//func to calculate square and put it to shared variable
func doSquare(sum *int, val int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	//wait mutex to lock it for add calculated result in shared variable
	mtx.Lock()
	//add calculated value into shared variable when mutex is locked
	*sum += val * val
	//unlock mutex
	mtx.Unlock()
	//say to wait group
	wg.Done()
}
