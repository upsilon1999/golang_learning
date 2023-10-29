package main

import "fmt"

func addMore(a, b int) int {
	return a + b
}

/*
函数既然可以作为一个变量，那么他也能作为返回值
*/
func funA(op string, items ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("这是加法")
		}
	case "-":
		return func() {
			fmt.Println("这是减法")
		}
	default:
		return func() {
			fmt.Println("什么都不是")
		}
	}

}

/*
函数是一个变量，那么就也可以作为参数传递
myfunc func(item ...int) int
这是一个整体，
myfunc 函数名
item ...int 函数的参数
int 函数的返回值类型

所以这种可以嵌套写的很复杂，需要把函数整体拿出来看
他和函数单独写在外面的区别就是，函数名被抽离了
*/
func funB(myfunc func(item ...int) int) int {
	return myfunc()
}
func main() {

	// 函数可以作为一个变量传递
	// 注意传递的是函数名，加括号传递的是函数调用结果
	fun01 := addMore
	fun01(5, 6)

	//	返回是一个函数，函数需要调用
	funA("+")()

}
