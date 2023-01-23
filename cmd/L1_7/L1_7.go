package L1_7

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//map with mutex
type MyMap struct {
	ShareMap map[int]string
	mtx      sync.RWMutex
	end      bool
}

//create new map with mutex
func (m *MyMap) NewMyMap() *MyMap {
	//make map
	m.ShareMap = make(map[int]string)
	//program is not finished
	m.end = false
	return m
}

//Concurrent write into map
func Worker(done *chan bool, mp *MyMap) {
	//while not stop the program
	for !mp.end {
		//try to lock the map
		mp.mtx.Lock()
		//add value into the amp
		mp.ShareMap[rand.Intn(100)] += "Hello,"
		//free the mutex
		mp.mtx.Unlock()
	}
}

func Seven() {
	//number of gorutines(read from stdin)
	var n int
	fmt.Scan(&n)
	done := make(chan bool)
	//create map with mutex
	var myMap MyMap
	myMap.NewMyMap()
	//start gorutines
	for i := 0; i < n; i++ {
		go Worker(&done, &myMap)
	}
	//sleep and stop cocurrent work
	time.Sleep(1000 * time.Millisecond)
	myMap.end = true
	//print map
	for k, v := range myMap.ShareMap {
		fmt.Println(k, " ", len(v))
	}
}
