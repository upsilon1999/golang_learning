// 运行时用这个
// package main

package defineVariable

import "fmt"

func main()  {
	/* 
		静态语言的特性
			①变量必须先定义后使用
			②变量必须有类型
			③类型定下来后不能改变
	*/

	/* 
		定义变量的方式① 
		var 变量名 变量类型
	*/
	var name string
	name = "zsf"

	/* 
		定义变量的方式②
		var 变量名 变量类型 = 变量值
	*/
	var firstName string = "wanwu"

	/* 
		定义变量的方式③ 
		定义时如果赋了初始值，那么类型可以省略，因为他会自动推断
		var 变量名 = 变量值
	*/
	var lastName = "shijie"

	/* 
		定义变量的方式④ 
		变量名 := 变量值

		优势：好用，相当于方式③的缩写
		缺点：无法显式指明类型

		只能用来声明局部变量，不能用来声明全局变量
	*/
	age := 18

	/* 
		go语言中局部变量定义了必须要使用
		fmt是格式化字符串，%v是占位符
	*/
	fmt.Printf("%v %v %v %v",name,firstName,lastName,age)
}