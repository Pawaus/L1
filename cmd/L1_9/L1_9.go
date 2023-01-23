package L1_9

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Worker1(ch *chan int64, n int, wg *sync.WaitGroup) {
	//write to the channel numbers from array
	for i := 0; i < n; i++ {
		*ch <- int64(rand.Intn(1000))
		time.Sleep(100 * time.Millisecond)
	}
	//say wait group that finish the gorutine
	wg.Done()

}

func Worker2(ch *chan int64, ch2 *chan int64, n int, wg *sync.WaitGroup) {
	//read from first channel and put square into second
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		res := <-*ch
		*ch2 <- int64(res * res)
	}
	//say wait group that gorutine is finish
	wg.Done()
}

func Nine() {
	//declare 2 channels
	var ch, ch2 chan int64
	//n is len of array
	n := 10
	//first channel
	ch = make(chan int64, n)
	//second channel
	ch2 = make(chan int64, n)
	//wait group to wait gorutines
	var wg sync.WaitGroup

	wg.Add(2)
	//start first worker
	go Worker1(&ch, n, &wg)
	//start second worker
	go Worker2(&ch, &ch2, n, &wg)
	//wait to finish all workers
	wg.Wait()
	//read from second channel and print to stdout
	for i := 0; i < n; i++ {
		fmt.Println(<-ch2)
	}

}
