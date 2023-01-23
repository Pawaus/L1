package L1_5

import (
	"fmt"
	"math/rand"
	"time"
)

func Tick() {
	//ticker for check spending time
	ticker := time.NewTicker(time.Second)
	//in the end we will stop the ticker
	defer ticker.Stop()
	//channel for stoping the program
	done := make(chan bool)
	//channel for data
	chanData := make(chan int)
	//variable for number of seconds
	var n int
	fmt.Println("Input count seconds")
	fmt.Scan(&n)
	// gorutine witch wait fon n seconds
	go func() {
		time.Sleep(time.Duration(n) * time.Second)
		//after n seconds send signal for end the program
		done <- true
	}()
	//gorutine read data from data channel
	go func() {
		for v := range chanData {
			fmt.Println(v)
		}
	}()
	//if we haeven't signal from done channel, send random number into channel
	for {
		select {
		case <-done:
			//printing that program is finished
			fmt.Println("Done!")
			//close data channel
			close(chanData)
			return
		case <-ticker.C:
			//put random number into channel
			chanData <- rand.Intn(100)
		}
	}
}
