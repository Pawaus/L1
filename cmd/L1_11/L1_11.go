package L1_11

import "fmt"

func Eleven() {

	var set1 []int = []int{10, 21, 51, 33, 49, 54}
	var set2 []int = []int{32, 33, 33, 56, 49}
	//map store values from set1 like keys
	m := make(map[int][]int)

	for _, v := range set1 {
		m[v] = nil
	}
	//if value key(number from set1) have number from set2? append it to map
	for _, v := range set2 {
		if _, ok := m[v]; ok {
			m[v] = append(m[v], v)
		}
	}
	//print intersections of sets
	for _, arr := range m {
		if len(arr) > 0 {
			for _, i := range arr {
				fmt.Printf("%d ", i)
			}
		}
	}
}
