package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

/**
方法/变量定义
数据基本类型的认识
 */

//  定义方法
//  func  [funcName] ([p] pType) [return type]

func add(x int, y int) int {
	return x + y;
}

//  参数类型 相同时，可以统一在最后一个参数后标明参数类型
func add2(x, y int) int {
	return x + y;
}

//可以返回多个值   //交换值
func swap(x, y string) (string, string) {
	return y, x;
}

//返回值被命名，则视为当前函数中的变量，可以不用在 return 后返回，可直接返回
func split(sum int) (x, y int) {
	x = sum * 4 / 9; //强类型 直接转为int类型
	y = sum - x;
	fmt.Println(sum)
	fmt.Println(sum*4/9, x, y)
	return
}

//变量声明
//声明变量时，如需带上类型，如果相同类型的可以在最后一个变量后带上类型，
// 如果没有对其初始化/赋值 则为当前类型的默认值
var c, python, java, nodejs bool;

//变量初始化 如果初始化值已存在，则不用带对应的类型，变量会自动从初始中获取对应的数据类型
var x, y int = 1, 2;
//函数体中声明变量可以用简便方式，:=的方式赋值，不可在函数体外使用
//aa:=123;

//类型转换
func transType() {
	//i := 42
	//f := float64(i)
	//u := uint(f)
	//fmt.Println(i,f,u)

	i := 42
	f := i
	u := f
	fmt.Println(i, f, u)

}

//类型推导
func checkType() {
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("i %T %v \n", i, i)
	fmt.Printf("f %T %v \n", f, f)
	fmt.Printf("g %T %v \n", g, g)
}

//常量  常量不能用 :=的方式声明
const Pi = 3.14;

func constData() {
	const WORLD = "世界";
	const HELLO = "你好";
	fmt.Println(HELLO, WORLD)
	fmt.Println("HAPPY ", Pi, "DAY");
	const TRUTH = true;
	fmt.Println("Tell me the truth", TRUTH)
}

//数值常量
func needInt(x int) int           { return x*10 + 1; }
func needFloat(x float64) float64 { return x * 0.1 }
func digitConst() {
	const (
		Big   = 1 << 100;
		Small = Big >> 99;
		Big64 = 1 << 62;
	)
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Small))
	fmt.Println(Big64)

}

func main() {

	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())   //当前时间
	fmt.Println("rand num ", rand.Intn(100)) //随机数 rand.Intn(n) n 为范围
	fmt.Println("math.sqt ", math.Sqrt(7))   //返回7的平方根

	fmt.Println(math.Pi);
	fmt.Println(add(10, 11));
	a, b := swap("hello", "world"); //值互换位置
	fmt.Println(a, b);
	fmt.Println(split(10))

	var i int;
	fmt.Println(i, c, python, java, nodejs)

	var cc, jaja, pypy = true, 123, "hello python"; //变量初始化 如果初始化值已存在，则不用带对应的类型，变量会自动从初始中获取对应的数据类型
	fmt.Println(x, y, cc, jaja, pypy)

	//函数体中声明变量可以用简便方式，:=的方式赋值，不可在函数体外使用
	jack := "Jack";
	fmt.Println(jack)

	//数据基本类型
	/**
	bool
	string
	int  int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64

	byte //uint8的别名
	rune //int32的别名  表示一个Unicode码点
	float32  float64
	complex64 complex128  复数类型

	*/

	//分组命名
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	//fmt.Println("Type :%T Value:%v\n",ToBe,ToBe)
	//fmt.Println("Type :%T Value:%v\n",MaxInt,MaxInt)
	//fmt.Println("Type :%T Value:%v\n",z,z)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//零值 未明确赋值则取值为0值
	var i1 int     //默认0
	var f1 float64 // 0
	var b1 bool    //false
	var s1 string  // ""
	fmt.Printf("%v %v %v %v \n ", i1, f1, b1, s1)

	//类型转换
	transType();

	//类型推导 根据右边值推导类型
	checkType();

	//常量
	constData();
	//数值常量
	digitConst()
}
