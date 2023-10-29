// package main

package multipleVariables

import "fmt"

func main()  {
	/* 
		①定义多个同类型的变量

		要求同类型，因为显式指定了类型
	*/
	var user1,user2,user3 string

	fmt.Printf("%v %v %v\n",user1,user2,user3)

	/* 
		② 定义多个变量的时候赋值

		由于没有显式指定类型，所以会采用类型推断的形式
		这样就允许不同类型的存在，变量与值一一对应
	*/
	var age1,age2,age3 = 18,19,"dddd"
	fmt.Printf("%v %v %v\n",age1,age2,age3)

	/* 
		③限定类型的同时赋值
	*/
	var he,she,it string = "他","她","它"
	fmt.Println(he,she,it)


	/* 
		④使用简洁定义多个变量，
		由于简洁变量没有类型定义，所以只能采用值推导的形式
	*/
	sex1,sex2 := "男","女"
	fmt.Println(sex1,sex2)

	/* 
		⑤ 相当于抽离var，是一种非常常用的多变量定义方式
	*/
	var (
		ok bool
		firstName string = "zs"
		lastName = "李四"
	)
	fmt.Println(ok,firstName,lastName)

	/* 
		注意：
			①变量必须先定义才能使用
			②go语言是静态语言，要求变量的类型和赋值类型一致

			③变量不能重复的定义，但是全局变量和局部变量可以重名，使用时优先使用本地的局部变量
			④ 简洁变量定义，即 (:=) 的形式不能用于全局变量定义
			⑤ 变量是有零值的，也就是默认值 
			⑥ 局部变量定义了一定要使用

			⑦多个变量定义与赋值要注意变量与值的位置一一对应
	*/
}