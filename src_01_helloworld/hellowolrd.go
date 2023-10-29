// 包名任意，与文件名没有任何关系
// 但是main函数必须在main包中才能编译
// package main
package hello

//导包
import "fmt"

//main函数是go文件的入口
func main()  {
	
	/* 
		打印的方式
		① println() 不需要引入包
		② fmt包，功能丰富
	*/
	println("我是print -- Hello World")

	// 打印一行
	fmt.Println("fmt打印 -- HelloWorld")
}

