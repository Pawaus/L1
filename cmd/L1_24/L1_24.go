package L1_24

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

// Method to calculate length from one to another point
func (p *Point) Length(p2 *Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(p.x)-float64(p2.x)), 2) + math.Pow(math.Abs(float64(p.y)-float64(p2.y)), 2))
}

func TwentyFour() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(2, 2)
	fmt.Println(p1.Length(p2))
}
