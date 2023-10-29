package main

import "fmt"

func main()  {
	// 数组的初始化
	/* 
		方式1：
		
		var 数组名 [count]type = [count]type{元素1,元素2,...}

		由于赋值的时候已经指明类型，故左边类型可以省略 
		var 数组名 = [count]type{元素1,元素2,...}
	*/
	var courses = [3]string{"go","java","python"}
	fmt.Printf("%#v\n",courses)//[3]string{"go", "java", "python"}


	/* 
		方式2:
		数组名 := [count]type{元素1,元素2,...}
	*/
	courses1 := [4]string{"box1","box2","box3","box4"}
	fmt.Printf("%#v\n",courses1) //[4]string{"box1", "box2", "box3", "box4"}

	/* 
		方式3
		指定下标初始化，其他地方采用零值 
		其实之前的初始化可以理解为所有下标都给了值，所以省略了下标

		var 数组名 = [count]type{下标:值,下标:值}

		数组名 := [count]type{下标:值,下标:值}
	*/
	courses2 := [3]string{2:"gin"}
	var courses3 = [2]string{0:"beego"}
	fmt.Printf("%#v\n",courses2) //[3]string{"", "", "gin"}
	fmt.Printf("%#v\n",courses3) //[2]string{"beego", ""}


	/* 
		方式4 
		用省略号代替count，根据赋值自动推算count值 
		① 只有{}赋值的形式有效，
		②下标赋值法无效,因为没给定count不知道下标
		③会根据{}元素的个数推算出类型 

		例:
		[...]int{1，2} =>根据个数推算 =>[2]int{1,2}
		[...]int{1,2,3} =>[3]int{1,2,3}

		优势：方便代码的书写
	*/
	courses4 := [...]string{"abc","更好"}
	var courses5 = [...]int{1,2,3,4}
	
	// courses5 := [...]string{5:"天天"} //报错，下标语法无效

	fmt.Printf("%#v\n",courses4) //[2]string{"abc", "更好"}
	fmt.Printf("%#v\n",courses5)//[4]int{1, 2, 3, 4}

}