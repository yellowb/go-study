package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v1 *Vertex) Scale(f float64, v2 *Vertex) {
	v1.X = v1.X * f
	v1.Y = v1.Y * f
	fmt.Println(v2)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(20, &v)
	fmt.Println(v.Abs())
}
