// package main
package pointerDown

import "fmt"

func main()  {
	
	/* 
		①声明变量
			var 变量 类型

		②声明变量对应的指针
			var 变量 *类型
	*/


	var a int= 100
	//存储变量a 的地址
	//b 是个指针类型的变量，他存储的是一个地址，这个地址指向变量a的值所在的空间
	var b *int = &a

	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	println("----------------------")
	

	/* 
			指针类型的变量B存储的是变量A的地址     &变量A
			想要通过变量B获取变量A的值  					*变量B
	*/
	c := &map[string]string{
		"name":"张三",
	}
	(*c)["name"]= "李四"
	fmt.Printf("%#v\n",c)
	fmt.Printf("%#v\n",*c)

	/* 
		结构体实例的指针变量和结构体变量之间可以互换
	*/
	type Test struct{
		name string
	}

	//这是右边结构体的指针变量
	p1 := &Test{
		name:"王五",
	}

	//常规的值修改 
	(*p1).name = "李四"
	fmt.Printf("%#v\n",p1)

	//结构体的特例，结构体的指针变量可以直接使用和值的用法一样 
	//精简结构体的使用
	p1.name = "二虎"
	fmt.Printf("%#v\n",p1)

}