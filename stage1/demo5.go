package stage1

import (
	"fmt"
	"reflect"
	"strings"
)

/**
指针
结构体
数组
切片
*/

//指针
func pointer() {
	i, j := 42, 2701
	//“间接引用”或“重定向”
	//& 操作符会生成一个指向其操作数的指针。
	p := &i // 指向 i
	fmt.Println("p0", p)
	//* 操作符表示指针指向的底层值。
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	fmt.Println("p1", p)
	p = &j // 指向 j
	fmt.Println("p2", p)
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

//结构体

func structFunc() {
	//一个结构体（struct）就是一组字段（field）。
	type Test struct {
		X int
		Y int
	}
	fmt.Println(Test{1, 2})

	//访问并赋值结构体属性值
	v := Test{2, 3};
	v.X = 4;
	fmt.Println(v.X);
}

//结构体指针
func pointStruct() {
	type Test struct {
		X int
		Y int
	}
	v := Test{3, 4}
	p := &v;
	p.X = 1e9; //==等同于 *p.X 此处可隐式间接引用
	fmt.Println(v)
}

//结构体语法/文法
func structSyntax() {
	type Test struct {
		x, y int
	}
	var (
		v1 = Test{1, 2}  //创建一个Test类型的结构体
		v2 = Test{x: 1}  // y:0 被隐式的赋予
		v3 = Test{}      //x,y默认0
		p  = &Test{1, 2} // 创建一个*Test类型的结构体（指针）
	)
	fmt.Println(v1, v2, v3, p)

}

//数组
func arrayFunc() {
	var a [10] int; //创建一个长度为10，元素默认为int 0值的数组
	fmt.Println(a);

	//元素访问与赋值
	var a1 [2] string
	a1[0] = "hello";
	a1[1] = "world";
	fmt.Println(a1[0], a1[1])
	fmt.Println(a1);

	//初始长度并初始具体的元素
	list := [6]int{2, 3, 4, 5, 7, 9};
	fmt.Println(list);
	list1 := [6]int{2, 3, 4}; //未声明的默认值为0
	fmt.Println(list1);
	list2 := [...]int{2, 3, 4}; //不指定长度，用...代替，自动推导数组长度
	fmt.Println(list2);

	//a := [3]int{5, 78, 8} 不允许将两个不同长度的数组进行相互赋值
	//var b [5]int
	//b = a //not possible since [3]int and [5]int are distinct types
}

//切片
func slicesFn() {
	//切片提供动态大小,灵活视角
	primes := [6]int{2, 3, 4, 5, 6, 9}
	var s []int = primes[1:4];
	fmt.Println(s)

	//类型 []T 表示一个元素类型为 T 的切片。
	//a[low : high] 含头不含尾

	//切片不存储任何数据，只是描述底层数组中的一小段数据
	//更改切片元素会修改底层数组中对应元素
	//与它共享的底层数组切片会检测到对应的修改
	names := [4]string{"John", "Paul", "George", "Ringo"};
	fmt.Println(names); //[John Paul George Ringo]

	a := names[0:2];
	b := names[1:3];
	fmt.Println(a, b); //[John Paul] [Paul George]
	b[0] = "xxx"
	fmt.Println(a, b)   //[John xxx] [xxx George]
	fmt.Println(names); // [John xxx George Ringo]
}

//切片语法/文法
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slicesSyntax() {
	//类似没有长度的数组语法
	q := []int{2, 3, 4, 5, 6, 7};
	fmt.Println(q);
	fmt.Println(reflect.TypeOf(q)) //判断类型

	r := []bool{true, false, false, true, false};
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, false},
	}
	fmt.Println(s)
	fmt.Println(s[0].i, s[0].b) //取值

	//切片默认行为

	s1 := []int{2, 3, 4, 5, 6};
	s1 = s1[1:4];
	fmt.Println(s1)

	s1 = s1[:2];
	fmt.Println(s1)

	s1 = s1[1:];
	fmt.Println(s1);
	//切片下边界默认值0，上边界为该切片的长度
	var arr [10]int        //以下等价
	fmt.Println(arr[0:10]) //等价
	fmt.Println(arr[:10])
	fmt.Println(arr[0:])
	fmt.Println(arr[:])

	//切片长度和容量

	sl1 := []int{2, 3, 5, 7, 11, 13}
	printSlice(sl1)
	//截取切片使其长度为0
	sl1 = sl1[:0];
	printSlice(sl1);
	//拓展其长度
	sl1 = sl1[:5];
	printSlice(sl1)
	//舍弃前两个值
	sl1 = sl1[2:];
	printSlice(sl1)

}

//nil切片
func nilSlice() {
	var s []int; //切片的默认值为 nil  长度和容量为0 没有底层数组
	fmt.Println(s, len(s), cap(s));
	if s == nil {
		fmt.Println("nil")
	}
}

// 用make来创建切片 也是创建动态数组的方式
func makeSlice() {
	fmt.Println("make slice");
	//    make([type],len,cap) //不传cap默认与len相同
	//a := make([]int, 5); // len(a)=5  cap(a)=5
	a := make([]int, 5); // len(a)=5  cap(a)=5
	fmt.Println(a, len(a), cap(a));
	b := make([]int, 0, 5) //len(b)=0  cap(b)=5
	fmt.Println("b1", b, len(b), cap(b));
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	fmt.Println("b2", b, len(b), cap(b));
	b = b[1:] // len(b)=4, cap(b)=4
	fmt.Println("b3", b, len(b), cap(b));
}

/**
切片的切片，切片的切片包含任何类型，甚至包含其他切片
*/
func sliceOfSlice() {
	//创建一个＃字版
	board := [][]string{
		[] string{"_", "_", "_"},
		[] string{"_", "_", "_"},
		[] string{"_", "_", "_"},
	}
	fmt.Println("board before", board);
	//两玩家轮流打X/O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], ""));
	}
}

/**
向切片中追加元素
*/
func appendSlice() {
	var s []int;
	s = append(s, 0);
	printSlice(s);
	s = append(s, 1);
	printSlice(s);
	//可一次性添加多个元素，按需增长
	s = append(s, 2, 3, 4);
	printSlice(s);
	/**
	append（s []T,vs ...T） 第一个参数s是类型为T的切片
	其余类型为T的值将追加到该切片的末尾。
	当s的底层数组太小不足容纳给定值时，会分配更大的数组，返回的切片会指向新数组
	 */
}

func main() {
	pointer();
	structFunc();
	pointStruct();
	structSyntax();
	arrayFunc();
	slicesFn();
	slicesSyntax();
	nilSlice();
	makeSlice();
	sliceOfSlice();
	appendSlice();
}
