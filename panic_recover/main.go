package main

import "fmt"

/**
用于错误的处理
defer recover panic搭配使用处理异常 defer 必须定义在可能引发panic的语句之前

*/
func a()  {
	fmt.Println("func a")
}

func b()  {
	defer func() {
		err:=recover();
		if err!=nil{
			fmt.Println("fun b error")
		}
	}();
	//fmt.Println("func b")
	panic("panic in b")
}

func c()  {
	fmt.Println("func c")
}




func main() {
	a();
	b();
	c();
}
