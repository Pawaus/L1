package L1_18

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
}

// add method to the Counter struct
func (counter *Counter) increment(wg *sync.WaitGroup) {
	// devide WG counter using defer
	defer wg.Done()
	atomic.AddInt64(&counter.value, 1)
}

func Eigthteen() {
	var wg sync.WaitGroup

	value := 10
	n := int64(5)

	val := Counter{n}

	for i := 0; i < value; i++ {
		// add counter to WG
		wg.Add(1)
		// increment in goroutines
		go val.increment(&wg)
	}

	// wait for all goroutines
	wg.Wait()

	fmt.Println("Result of increment:", val.value)
}
