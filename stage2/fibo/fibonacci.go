package main

import "fmt"

/**
闭包实现斐波那契数
*/

func fibonacci() func() int {
	back1, back2 := 0, 1;
	return func() int {
		back1, back2 = back2, (back1 + back2);
		return back1
	}
}

func main() {
	f := fibonacci();
	const LEN = 40;
	var array [LEN]int
	for i := 0; i < LEN; i++ {
		array[i] = f();
	}
	fmt.Println(array)
}
