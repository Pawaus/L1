package L1_4

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var is_start bool = false

func Worker(data chan int, wg *sync.WaitGroup) {
	//wait for start writing
	wg.Wait()
	//recieve first value
	val, is_opened := <-data
	//while channel in opened read from channel
	for is_opened {
		//print readed data
		fmt.Println(val)
		//recieve a new value
		val, is_opened = <-data
	}

}

func Four() {
	//make channel with big buffer
	channelData := make(chan int, 100000)
	//make channel for stopping program by the signals from system
	sigChan := make(chan os.Signal, 1)
	//define signals iwch want to recieve and channel where signals will write
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
	//flag stopping the program
	var finishProgram bool = false
	//gorutine wait for signal from system
	go func() {
		<-sigChan
		finishProgram = true
		fmt.Println("finish Program")
		time.Sleep(10 * time.Millisecond)
		close(channelData)

	}()
	//number of workers
	var n int
	fmt.Scan(&n)
	//wait group for starting workers
	var wg sync.WaitGroup
	wg.Add(1)
	//start workers
	for i := 0; i < n; i++ {
		fmt.Println("start worker")
		go Worker(channelData, &wg)
	}
	//run workers
	wg.Done()
	//while program is not finished, push data into channel
	for !finishProgram {
		//fmt.Println("Push")
		//push data in channel
		channelData <- rand.Intn(100000)
		//sleep(for educational purpose)
		time.Sleep(10 * time.Millisecond)
	}

}
