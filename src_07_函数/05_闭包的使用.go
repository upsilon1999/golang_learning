// package main
package useClosure

import "fmt"

/*
	将函数需要的变量在内部声明
	返回值是一个函数的调用结果，
	用这个被调用的函数来操作变量
*/
func autoIncrement() int{
	local := 0
	
	//由于内部的匿名函数访问到了其他函数中的局部变量，所以称为闭包
	// 这个函数调用，获得对应类型的返回值
	return func() int{
		local += 1
		return local
	}()
}

/* 
	和autoIncrement的区别，前者每次调用都是整体调用，会执行autoIncrement 
	就会使得local重置为0
	而autoIncrement01是将内部闭包函数抛出,只执行闭包函数，所以不会重置local

	想要重置local，只需要再调用一次autoIncrement01即可
*/
func autoIncrement01() func() int{
	local := 0
	
	//由于内部的匿名函数访问到了其他函数中的局部变量，所以称为闭包
	// 这个函数调用，获得对应类型的返回值
	return func() int{
		local += 1
		return local
	}
}
func main()  {
	/* 
		autoIncrement的使用，返回值是个int
		每次调用执行，local重置，计算完后抛出值，所以一直为1
	*/
	for i := 0; i < 3; i++ {
		fmt.Printf("我是autoIncrement的值%#v\n", autoIncrement())
	}

	println("------------------------------------")
	/* 
		autoIncrement01的返回值是个函数
	*/
	//赋值的时候，调用autoIncrement01，loacl初始化为0
	next := autoIncrement01()
	for i := 0; i < 3; i++ {
		fmt.Printf("我是autoIncrement01的值%#v\n", next())
	}

	println("------------------------------------")

	//想要重置loacl，只需要再次调用autoIncrement01
	next = autoIncrement01()
	fmt.Printf("我是local重置后再开始的值%#v\n", next())
}