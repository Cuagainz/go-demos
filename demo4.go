package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/**
for循环
条件判断 if-else;switch-case
推迟函数 defer 以及defer 栈的工作原理
 */

//条件判断，for循环
func forFunc() {
	//sum := 0;
	////与其他语言相比 没有小括号
	//for i := 0; i < 10; i++ {
	//	sum += i;
	//}
	//fmt.Println(sum)

	//初始化语句和后置语句可选
	sum := 1
	for ; sum < 1000; {
		sum += sum;
	}
	fmt.Println(sum)
	whileFunc();
	unLimitLoop(); //无限循环
}

// for 是go中的while 可去掉分号
func whileFunc() {
	sum := 1;
	for sum < 1000 {
		sum += sum;
	}
	fmt.Println(sum)

}

//无限循环
func unLimitLoop() {
	//for {
	//	fmt.Println("1")
	//}
}

//if条件语句
func ifFunc() {
	s := sqrt(1.111);
	fmt.Println(s)

	//可在if中执行简短的语句
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func pow(x, n, lim float64) float64 {
	//v的作用域仅在if开头的语句中以及else中有效
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
	//return strconv.FormatFloat(math.Sqrt(x), 'E', -1, 64)
}

//switch 条件判断
func switchFunc() {
	fmt.Print("Go runs on ss");
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func swFallThrough() {
	switch {
	case false:
		fmt.Println("The integer was <= 4")
		fallthrough
	case true:
		fmt.Println("The integer was <= 5")
		fallthrough //强制执行后面的case语句，不判断是否true
	case false:
		fmt.Println("The integer was <= 6")
		fallthrough
	case true:
		fmt.Println("The integer was <= 7")
	case false:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

//求值顺序
func swReturnSort() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday();
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today");
	case today + 1:
		fmt.Println("Tomorrow");
	case today + 2:
		fmt.Println("In two days");
	default:
		fmt.Println("Too far away");

	}
}

//没有条件的switch 同 switch true相同 可以替换if-then-else 更清晰
func swNoCond() {
	t := time.Now();
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning！")
	case t.Hour() < 17:
		fmt.Println("Good afternoon");
	default:
		fmt.Println("Good evening");

	}
}

func deferFunc()  {
	//defer 将函数推迟到当前函数执行完后再执行

	defer fmt.Println("hello")
	fmt.Println("world")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	//推迟函数调用会被加入一个栈中，当前函数返回时，被推迟的函数会按照后进先出的顺序调用
	fmt.Println("done")

}

func main() {
	//defer 将函数推迟到当前函数执行完后再执行
	defer fmt.Println("hello")
	forFunc();
	ifFunc();
	switchFunc();
	swFallThrough(); //强制执行
	swReturnSort()   //switch 求值执行顺序
	swNoCond();//没有条件的switch



	fmt.Println("world")

	deferFunc();
}
