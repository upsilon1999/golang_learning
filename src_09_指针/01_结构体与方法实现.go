// package main
package structAndPointer

import "fmt"

type Persons struct{
	name string
	age int
}

/*
	接收的是值类型 也就是值传递
	所谓的值传递，就是变量将自己复制一份，然后将复制的传入
*/
func (c Persons) AddAge() int{
	//修改的是变量复制的内容
	//如果复制后地址发生了变化，那么就无法修改到变量
c.age ++
return c.age
}

/*
	接收的是指针类型，也就是说传参要传地址
	接收者方法对结构体做了优化，结构体的值和地址在传入时会根据需要自动转化
*/

func (c *Persons) changeName() string{
	//引用传递，传递的是地址
	 //对地址的修改，即修改了变量底层
c.name ="李四"
return c.name
}

func main()  {
	p := Persons{}

	p.age = 20

	fmt.Printf("调用方法得到的年龄%v\n",p.AddAge())
	fmt.Printf("age变量对应的年龄%v\n",p.age)
	println("--------------------------------")

	p.name = "张三"
	fmt.Printf("调用方法得到的名字%v\n",p.changeName())
	fmt.Printf("name变量对应的名字%v\n",p.name)
}