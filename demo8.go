package main

import (
	"fmt"
	"math"
	"strings"
)

/**
函数 的使用
函数值
闭包
*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcValue() {
	//匿名函数
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}

/**
把函数作为返回值
闭包=函数+外层变量的引用 //判断标准
*/
func adder() func(int) int {
	sum := 0;
	return func(x int) int {
		sum += x;
		return sum;
	}

}

/**
定义文件名后缀
*/
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix;
		}
		return name;
	}
}

//计算
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i;
		return base;
	}
	sub := func(i int) int {
		base -= i;
		return base;
	}
	return add, sub;
}

func funcClosure() {
	pos, neg := adder(), adder();
	for i := 0; i < 10; i++ {
		pos(i);
		//p := pos(i);
		//fmt.Println("i:", i, "p:", p)
		neg(-2 * i)
		//n := neg(-2 * i)
		//fmt.Println("i:", i, "n:", n)
	}

	r := makeSuffixFunc(".txt");
	ret := r("test");
	fmt.Println("ret", ret);

	//闭包计算
	calcR1,calR2 := calc(100);
	cr1 := calcR1(200)
	fmt.Println("cr1:", cr1) //cr1: 300
	cr2 := calR2(100)
	fmt.Println("cr2:", cr2)//cr2: 200
}

func main() {
	funcValue();
	funcClosure()
}
