package main

import "fmt"

type Student struct {
	name string
}

// 接收者函数不能作为嵌套函数
func (c Student) getName() string {
	return c.name
}

func main()  {
	/* 
		接收者函数是为类型绑定方法的手段 
		接口、结构体、自定义类型都属于类型
	*/
	/* 
		Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
	*/
	// 语法 
	/* 
		func (variable_name variable_data_type) function_name() [return_type]{
    	函数体
		}

		作用：在某个类型身上挂载新方法 
		func (变量名① 变量数据类型②) 函数③ {
			函数体
		}

		和函数的区别，在函数前加了一组括号，表名接收这个函数的对象 

		(变量名① 变量数据类型②) 叫做接收器,默认使用值传递，就把值复制一份传进来 
		如果我们想对调用者进行修改，或者值太大了，可以考虑引用传递，就是把目标的指针传进来

		变量名① 代表调用者，方法的调用者会将自身传入 
		变量的数据类型② 将函数体作为函数③的逻辑体挂载到②上，成为类型②的一个内置方法 
		函数③ 和普通函数是一样的，不过这里叫做方法 
	*/
	
	s:= Student{
		name:"张三",
	}

	fmt.Printf("%#v\r\n",s.getName())
}