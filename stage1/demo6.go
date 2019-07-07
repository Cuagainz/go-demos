package stage1

import "fmt"
import "golang.org/x/tour/pic"

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
}
func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dx);
	for x := range a {
		b := make([]uint8, dy);
		for y := range b {
			b[y] = uint8(x ^ y - 1)
		}
		a[x] = b;
	}
	return a
}

func practiceSli() {
	//返回一个长度为dyg的切片，每个元素长度为dx,元素类型为unit8的切片
	pic.Show(Pic)
}

func main() {
	rangeExample();
	practiceSli();
}
