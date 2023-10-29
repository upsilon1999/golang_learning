// package main
package GlobalAndLocalVariables

import "fmt"

/*
	go文件采用函数式编程，
	所以定义在go文件内，函数体外的就是全部变量
	函数体内的就是局部变量

	局部变量定义了就必须要使用
	全局变量定义了可以不使用
*/

// 全局变量
var name = "wanwu"
var age int


func main()  {
	//局部变量
	isYes := true

	fmt.Printf("%v %v",isYes,name)
}
