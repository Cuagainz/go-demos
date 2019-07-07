package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func scale(v *Vertex, f float64) {
	v.x = v.x * f;
	v.y = v.y * f;
}

func main() {
	v := Vertex{3, 4};
	scale(&v, 10)
	fmt.Println(abs(v))
}
