// package main
package defineAndCheck

import "fmt"

func main()  {
	/* 
		指针的定义 

		var 变量 *类型 
	*/
	var a *string 

	/* 
		指针的初始化 
		①我们不能在基本类型上面取地址

		a = &"你好"

		这样是不被允许的

		②但是可以在变量上面取地址 
		var d = "王五"
		a = &d 
	*/
	var d = "你好"
	a = &d
	fmt.Printf("%#v\n", a)
	println("------------------------------")

	/* 
		指针需要初始化，不初始化是nil
		即空指针
	*/
	var b *int
	fmt.Printf("%#v\n",b)
	//空指针的错误常见于结构体 
	type Persons struct{
		name string
	}

	var p *Persons
	// 报错，nil上面没有这个
	//因为指针不初始化是nil，和map类似
	// p.name = "张三"
	println(p)

	var p1 Persons
	// 会使用零值
	fmt.Printf("%#v\n",p1.name)
	println("-------------------------------")


	// 复杂数据类型指针初始化
	var f1 *map[string]string
	var f2 *Persons
	var f3 *[]string
	f1 = &map[string]string{}

	f2 = &Persons{}

	f3 = &[]string{}
	fmt.Printf("%#v\r\n %#v\r\n %#v\r\n",f1,f2,f3)

	println("-----------------------")

	/* 
		使用new函数初始化 

		new(Type)
		new函数，参数传入一个类型，会开辟一块空间，然后返回这个类型对应的指针，
		这样就同时实现了指针类型的定义和初始化 
	*/
	var kk = new(int)
	fmt.Printf("%#v\r\n",kk)

	/* 
		tips：
			初始化关键字有两个 
			map、channel、slice定义兼初始化推荐使用make函数
			指针定义兼初始化推荐使用 new函数

			指针需要初始化，否则会出现nil空指针的问题
			map也需要初始化，否则会报nil问题
	*/
}