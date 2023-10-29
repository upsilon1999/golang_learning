// package main
package main

import "fmt"

/*
基本定义:

	func 函数名(参数 参数类型) 返回值类型{
		函数体
	}

①如果多个参数类型一样，可以简写，例如
(a int,b int) ==简写==> (a,b int)
②有多个返回值，需要用括号包裹
(int error)
*/
func add(a int, b int) (int, error) {
	return a + b, nil
}

func sub(a, b int, c float32) int {
	a = a + b
	return a - b
}

/*
方式二:

	var 函数名 = func(){}
	//局部可以采用
	函数名 := func(){}

	因为函数可以看作一个变量，
	变量名就是函数名，类型就是他本身，包括参数和返回值
*/
var localFun = func(a int) int {
	return a
}

func main() {
	//	go函数支持普通函数、匿名函数、闭包函数
	/*
		go中函数是一等公民：
			① 函数本身可以当作变量
			② 函数可以满足接口
	*/

	//函数调用
	sum, _ := add(2, 5)
	fmt.Println(sum)

	/*
		函数使用的是值传递，调用函数时传入的变量不是变量本身
		而是将变量拷贝一份之后进行的传递


	*/
	a := 5
	b := 3
	//这里传入函数的是变量的拷贝
	subVal := sub(a, b, 3.14)
	fmt.Println(subVal)
}
