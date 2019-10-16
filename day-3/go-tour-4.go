package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v *vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := vertex{10, 20}
	fmt.Println(v.abs())
}
