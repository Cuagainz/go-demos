package main

import (
	"fmt"
	"math"
)

/**
指针接收者
*/

type Test struct {
	x, y float64
}

func (t Test) abs() float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y)
}

func (t *Test) scale(f float64) {
	t.x = t.x * f;
	t.y = t.y * f;
}

func main() {
	v := Test{3, 4};
	v.scale(10); //接收引用类型，修改指向的原始数据
	fmt.Println(v)
	fmt.Println(v.abs())
}
