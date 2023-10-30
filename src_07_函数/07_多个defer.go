// package main
package manyDefer

import "fmt"

func main()  {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")

	/* 
		defer都是在函数return之前执行 
		defer是一个栈的概念，先写的先进栈，所以后写的在上面 
		出栈时由于从上到下出栈，所以后写的先执行
	*/
	return
}