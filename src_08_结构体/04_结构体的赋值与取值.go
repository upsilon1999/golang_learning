package main

import "fmt"

func main()  {
	type Persons struct{
		name string
		age  int
		height float32
	}

	var p Persons

	// 赋值 
	/* 
		结构体实例.字段 = 值

		这也是一种初始化的方式,其他字段会采用零值 
	*/
	p.age = 18
	fmt.Printf("%#v\n", p)
	println("----------------------")
	// 取值 
	/* 
		结构体实例.字段 
	*/
	fmt.Println(p.height)
	// 无法取不存在的字段
	// fmt.Println(p.abc)
}