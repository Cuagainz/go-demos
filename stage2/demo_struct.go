package stage2

import (
	"fmt"
	"math"
)

/**
go没有类，但可以在结构体类型上面定义方法
*/

type Test struct {
	x, y float64;
}

//结构体Test的方法 abs
func (t Test) abs() float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y);
}

//普通方法
func Abs(t Test) float64 {
	return math.Sqrt(t.x*t.x + t.y*t.y)
}

func (t Test) sayHello() {
	fmt.Println("hello ,Li Lei")
}

//以type方法定义不止可以定义结构体，其他类型也可以
type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	t := Test{3, 4};
	fmt.Println(t.abs())
	t.sayHello();
	fmt.Println(Abs(t))

	f := MyFloat(-math.Sqrt2);
	fmt.Println(f.abs());
}
