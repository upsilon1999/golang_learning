// package main
package NamelessStruct

import "fmt"

func main()  {
	/* 
		有些时候我们定义结构体只是为了临时使用，不需要全局使用他 
		例如一次性的接收前端参数，我们只想做一个类型校验
		这时候就需要使用匿名结构体
	*/
	/* 
		局部结构体定义的语法:

		变量名 := struct{
			字段A 类型A 
			字段B 类型B
		}{
			值A,
			值B
		}

		var 变量名 = struct{
			字段A 类型A 
			字段B 类型B
		}{
			值A,
			值B
		}

		这种实际上是定义的时候初始化，所以很多人也不认为他是匿名结构体 
		不过他为完全的匿名结构体提供了思路
	*/
	// 第一个大括号是定义，第二个大括号是赋值 
	address := struct {
		province string
		city string
		age int
	}{
		age:200,
	}
	fmt.Printf("%#v\r\n",address)

	/* 
		上述也叫做匿名结构体，不过是可以多次使用的，是相对于全局结构体而言的 
		然后可以通过 结构体名.字段 来赋值或修改 
	*/
	println("-----------------------------")
	/* 
		一次性的结构体，真正的匿名结构体 
	*/
	fmt.Printf("%#v\r\n",struct {
		province string
		city string
		age int
	}{
		city:"Rome",
	})

}