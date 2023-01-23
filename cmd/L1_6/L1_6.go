package L1_6

import (
	"fmt"
	"time"
)

//stop gorutine with global variable
var StopGoRutine bool = false

func Worker1() {
	//while not stop, do something
	for !StopGoRutine {
		time.Sleep(100 * time.Microsecond)
		fmt.Println("Hello")
	}
}

func Worker2(stop *bool) {
	//stop gorutine by variable transferred by pointer
	for !*stop {
		time.Sleep(100 * time.Microsecond)
		fmt.Println("Hello2")
	}
}

func Worker3(done *chan bool) {
	//gorutine stop by recieving data from channel, witch transferred by pointer
	//(actually it can be locally if gorutine declarated in the same function or globally)
	for {
		select {
		case <-*done:
			fmt.Println("Done!")
			time.Sleep(100 * time.Microsecond)
			return
		}
	}
}

func Six() {
	//channel for stopping gorutine
	var done chan bool
	var n int
	//variable for stopping gorutine
	stopG := false
	//how much seconds we will wait to stop the gorutine(it can be the situation, where push to channel data)
	fmt.Scan(&n)
	//run first gorutine to stop
	go Worker1()
	//run second gorutine to stop
	go Worker2(&stopG)
	//run third gorutine to stop
	go Worker3(&done)
	time.Sleep(time.Duration(n) * time.Second)
	//stop third gorutine
	done <- true
	//stop second gorutine
	stopG = true
	//stop first gorutine
	StopGoRutine = true

}
