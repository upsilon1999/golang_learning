// package main
package useType

import "fmt"

/*
	1.定义结构体
	2.定义接口
	3.共享底层，类型定义
	4.定义类型别名
	5.类型判断
*/


func main()  {
		/* 
			①定义类型别名 
			type 类型 = 类型别名

			别名实际上是为了更好的理解代码

			底层类型是一样的，只是取了一个别名，
			通过打印可知，根本类型没变，赋值也可以直接使用

			原理：
			在编译的时候，类型别名会直接替换为底层类型 
			即编译时，MyInt会直接变成int

			用途：
			增加代码可读性
		*/
		type MyInt = int
		var i MyInt
		fmt.Printf("%T\r\n", i)//int
		var a  int = 8
		i = a
		fmt.Printf("%v\r\n",i)//8

		println("---------------------------")
		/* 
			② 共享底层，自定义类型
			type 类型A 类型B

			少了等号，这个实际上是定义了一个新类型，
			基于已有的类型B定义一个新的类型A
			和底层类型B可以接收完全一样的值，但是二者属于不同的类型 
			变量之间需要类型转换

			用途： 
			既希望使用原类型的特性，又需要在其上扩展方法特性
			内置类型没办法扩展，但是自定义类型可以使用接收者函数扩展，结构体可以理解为一种特殊的自定义类型
		*/
		type newInt int 
		var b newInt
		fmt.Printf("%T\r\n", b)
		//main.newInt  包名.类型
		var c int = 9
		// 类型转换
		b = newInt(c)
		fmt.Printf("%v\r\n",b)
		println("---------------------------")
		/* 
			③类型判断
			对于interface{}类型的变量，使用
			变量.(type) 可以拿到他的实际类型,配合switch使用

			变量.(类型) 进行类型断言，直接强转
		*/
		var f interface{} = "abc"
		switch f.(type) {
			case string:
				println("是字符串类型")
			default:
				println("不是字符串")
		}

		// 类型断言，interface{}不能使用其他类型的方法 
		//经过类型断言为string，只要成立就可以使用string的方法了 
		m := f.(string)
		fmt.Println("m的长度为",len(m))
	}