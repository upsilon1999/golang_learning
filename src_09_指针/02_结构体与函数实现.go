// package main
package structAndFun

import "fmt"

type Student struct{
	name string
	age int
}

// 使用函数
func changeName(c Student){
	c.name = "张三"
}

func changeAge(c *Student){
	// 结构体的指针和值再使用点语法时不需要转换
	c.age = 20
}

func main()  {
	s:=&Student{
		name:"王五",
		age:15,
	}

	// 函数传参时，参数要求是什么类型就必须什么类型
	changeName(*s)
	fmt.Printf("改名%#v\n",s)

	changeAge(s)
	fmt.Printf("改岁%#v\n",s)

}