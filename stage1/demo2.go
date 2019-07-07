package stage1

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Println(s)
	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)

	//fmt.Println(strings.Join(os.Args[1:], " "))

	//fmt.Println(os.Args[0:])

	//s, sep := "", ""
	t1:=time.Now()
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx,arg)
	}
	t2:=time.Now()

	fmt.Println(t2.Unix()- t1.Unix())



}