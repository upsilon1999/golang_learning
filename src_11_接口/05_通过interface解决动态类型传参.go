// package main
package interfaceOk

import "fmt"

/*
	我们只要一个简单的加法功能，但由于数字类型繁多
	需要定义非常多的函数(因为函数不能重名)

	在其他语言中会用泛型解决这个问题，但go的泛型是后期加入的

	所以通常都用空接口来解决这种问题，因为空接口可以是任意类型
*/
func add(a int,b int) int {
	return a+b 
}

func addInt32(a int32,b int32) int32 {
	return a+b 
}

func addInt64(a int64,b int64) int64 {
	return a+b 
}

//空接口 
func addAll(a,b interface{}) int {
	/* 
		为了解决传参时的类型限制，我们使用了空接口类型 
		当我们接收到参数需要使用时，需要变回具体类型 
		这时候就需要用到空接口类型独有的类型转换 -- 类型断言 

		语法： 
			空接口类型变量.(具体类型)
			返回值有两个：
			①	断言的结果 
			② 布尔值，是否成功 
	*/
	ai,_ := a.(int)

	bi,_ := b.(int)
	return ai+bi
}

func main()  {
	a := 1
	b := 2
	fmt.Printf("值为%#v\n", addAll(a,b))
}