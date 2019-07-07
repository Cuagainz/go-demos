package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64;
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.x*v.x + v.y*v.y)
//}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2);
	//v := Vertex{3, 4};
	a = f; //a MyFloat 实现了 Abser
	//a = &v //a *Vertex 实现了Abser
	//
	//a = v;
	fmt.Println(a.Abs())

}
