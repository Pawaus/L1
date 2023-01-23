package L1_16

import (
	"fmt"
	"math/rand"
)

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Do it for two parts of current arr
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}

func Sixteen() {
	//array of int
	var set2 []int
	n := 20
	// fill array by values
	for i := 0; i < n; i++ {
		set2 = append(set2, rand.Intn(100))
	}
	//sorting the array
	qsort(set2)
	//print the array
	fmt.Println(set2)
}
