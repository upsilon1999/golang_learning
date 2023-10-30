// package main
package panicLearning

import "fmt"

func testA(){
	/* 
		panic 这个函数会导致程序退出
		相当于javascript中的throw
		golang中不推荐随便使用panic

	
	*/
	a := 10

	// 从这抛出异常，结束程序，后面的都不执行
	panic("这是一个错误")

	fmt.Println(a)
}

/* 
		一般用到panic的场景： 
		服务启动的过程中，
		要启动服务必须某些依赖服务要存在，例如日志文件、mysql等等
		如果任一依赖服务不满足，就主动调用panic

		但是你的服务一旦启动，如果某行代码中不小心使用panic，可能会导致该核心程序挂掉，整个服务崩溃 --重大事故
		所以不推荐
*/

/* 
	但是有些时候，我们的代码也会被动触发panic
	例如 我们使用map之前没有初始化
*/
func testB(){
	var name map[string]string
	// map没进行初始化就被赋值了 
	//会触发panic
	name["go"] = "go工程师"
}

/* 
	recover 用于捕捉panic，我们可以在捕捉到后做任何操作 
	常见的是捕捉到后进行日志打印 
	语法： r:= recover()
	如果有panic被捕获，recover()就会返回一个错误对象 
	处理时机 一般配合defer，因为defer的执行时机在return前
*/
func testC()  {
	defer func ()  {
		if r:=recover();r!=nil{
			fmt.Printf("捕获到的错误为%#v",r)
		}
	}()

	var name map[string]string
	name["go"] = "go工程师"
}
func main()  {
	testC()
	// 捕获到的错误为"assignment to entry in nil map
}

/* 
注意事项：
	1.defer需要放在panic之前定义，放在panic后就不会执行了，所以写在最开始
	2.recover只有在defer调用的函数中才会生效
	3.recover处理异常后，逻辑并不会恢复到panic的那个点继续执行
	4.多个defer会形成栈，后定义的defer会先执行
*/