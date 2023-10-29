// package main
package localVariablesScope

import "fmt"

func main()  {

/* 	
	① 
		局部变量只能在定义他的函数中使用，内部代码块也属于函数自身的一部分
		外部函数不能直接使用这个变量 -- 可以传参实现间接访问
		
		实际上传参给外部函数，也相当于把外部函数暂时变成了自身的代码块
*/
	err := "错误"
	println(err)

	
/* 
	② 
	变量的作用域和代码块绑定，实际在Go中局部变量是参考代码块的 
	即最近的一组大括号
*/
	{
		localName := "local"
		//localName是这个代码块中的局部变量
		//在{}外无法访问
		fmt.Println(localName)

		//err是外面父函数的变量，相当于闭包的概念，所以可以访问
		fmt.Println(err)
	}

	// 无法访问
	// fmt.Println(localName)

	/* 
		③ if 条件 { 代码块 }
	*/
	var maname string 
	if err == "错误" {
		// 这里重新赋值了，则只能在代码块中识别
		maname := "it"
		fmt.Println(maname)
	}
	//这里的值是最初在if代码块外面定义的 maname
	fmt.Println(maname)

	/* 
		④深度理解，局部变量是相对于代码块的
	*/
	if err == "错误" {
		abc := 2
		fmt.Println(abc)
	}else{
		// abc在这个代码块中没有
		// fmt.Println(abc)
	}

	/* 
		变量的查找逻辑 :本代码块查找 => 本代码块的祖先代码块查找 =>全局变量
		按照顺序查找，最先查到的就使用，查不到就报错
	*/

}