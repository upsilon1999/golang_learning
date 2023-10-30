// package main
package whyNeedClosure

import "fmt"

// 这样虽然实现了需求，但是污染了外部变量
// 且由于变量在外部，所以会被其他修改，不安全
var local = 0

func autoIncrement() int {
	local += 1
	return local
}


func main()  {
	/*
		需求：有个函数每次调用一次返回的结果值都是增加一次之后的值
	*/
	for i := 0; i < 5; i++ {
		fmt.Println(autoIncrement())
	}
	
}