package L1_10

import "fmt"

func Ten() {
	//array of temperatures
	var mas []float64 = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//map of temperatures with step
	res := make(map[int][]float64)
	for _, v := range mas {
		//temperature div 10 gives key to put it in map
		m := int(v) / 10
		res[m*10] = append(res[m*10], v)
	}
	fmt.Printf("%+v", res)
}
