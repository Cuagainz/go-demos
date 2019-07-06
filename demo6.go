package main

import "fmt"

/**
range 用于遍历切片
*/

//
func rangeExample() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128};
	// i 元素下标，v 元素数据
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
	//若只要索引，忽略第二个变量
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	//将下表或值 用 _ 作为变量来忽略当前_所在的元素
	for _, value := range pow {
		fmt.Println(value)
	}
	//todo 练习切片

}
func main() {
	rangeExample();
}
