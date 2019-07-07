package main

import (
	"fmt"
	"golang.org/x/tour/wc"
)

/**
map 的使用
*/

func mapBase() {
	// map 将键映射到值
	// map 0值/默认值为nil 没有键值，也不能添加键值
	//make函数返回给定的类型map，并初始化

	type Test struct {
		Lat, Long float64
	}
	//初始map
	m := make(map[string]Test)
	//赋值
	m["chengdu"] = Test{100.2333, 33.34}
	fmt.Println(m);
	//取值
	fmt.Println(m["chengdu"]);

	//语法
	//（latitude纬度，longitude 经度）
	m1 := map[string]Test{
		"chengdu":   Test{30.67, 104.07},
		"chongqing": Test{30.02, 106.15},
	}
	fmt.Println(m1)
	//若底层只有一个类型名 如Test，可以在赋值的时候省略类型
	m2 := map[string]Test{
		"chengdu":   {30.67, 104.07},
		"chongqing": {30.02, 106.15},
	}
	fmt.Println(m2);
	//修改map值
	m2["deyang"] = Test{30.67, 105.07}
	fmt.Println(m2);
	//删除元素值
	delete(m2, "deyang");
	fmt.Println(m2);
	//检查某个键是否存在  element,ok:=m[key] 如果key存在,ok为true,否则为false
	v, ok := m2["chongqing"];
	fmt.Println("Value:", v, "existed:", ok);
}


func wordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

//练习
func practiceMap() {
	wc.Test(wordCount)
}

func main() {
	mapBase();
	practiceMap()
}
