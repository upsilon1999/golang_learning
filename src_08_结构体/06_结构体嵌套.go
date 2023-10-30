// package main
package NestedStruct

import "fmt"

type Persons struct{
	name string
	age  int
	height float32
}

type Persons01 struct{
	name string
	age  int
	height float32
}
/* 
	第一种嵌套方式，给定一个字段
	让该字段是要嵌套的结构体类型

	取值赋值都是嵌套的形式,通过该字段去实现
	这个字段就相当于桥梁 

	这种叫具名嵌套
	优势： 
	避免了结构体的字段冲突 
	劣势： 
	书写繁琐
*/
type Student struct{
	p1 Persons01
	score float32
	p Persons
}


/* 
	匿名嵌套 直接给要嵌套的结构体名

	优势： 
	书写方便，访问直接 
	不需要再通过中间桥梁 
	
*/
type Teacher struct{
	Persons
	salary int
} 

func main()  {
	s := Student{
		p: Persons{
			name:"张三",
		},
		score: 95.6,
		p1: Persons01{},
	}

	fmt.Printf("学生姓名为%#v\r\n", s.p.name)
	println("-------------------------")

	t := Teacher{}
	t.name = "王二"
	t.salary = 10000
	fmt.Printf("老师信息为%#v\r\n", t)
}