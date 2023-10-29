// package main
package ifLearn

import "fmt"

func main()  {
	var age = 18
	/* 
	情况一：
		if 布尔表达式 {
			布尔表达式为true时执行的逻辑
		}

		①布尔表达式：一个返回值为true或false的表达式
	*/
	if age>3{
		println("我超过三岁了")
	}

	/* 
		情况二： 
		if (布尔表达式) {
			布尔表达式为true时执行的逻辑
		}

		兼容了其他语言的代码习惯，将布尔表达式用括号包裹也是一样的
		但是为了规范起见，一般不去写
		括号多用于分组，例如 
		(age>3)&&(age<20)
	*/

	/* 
		情况三： 
		if 条件① {
			满足条件①
		} else if 条件② {
			满足条件②
		}
	*/
	if age < 18 {
		println("你还太年轻了")
	}else if age == 18{
		println("恭喜长大")
	} 

	/* 
		情况四： 
		if 条件① {
			满足条件①
		} else if 条件② {
			满足条件②
		}else{
			其他if条件都不满足时执行
		}
	*/
	if age < 16 {
		println("你还太年轻了")
	}else if age>18{
		println("你是个大人了")
	}else{
		println("至死是少年")
	}

	/* 
		if 变量初始化;布尔表达式 {
			逻辑
		}

		①在if之后，条件语句之前，可以添加变量初始化语句，使用;进行分隔
		②这里初始化的变量是一个局部变量，可以在之后的布尔表达式和逻辑块中使用
	*/
	if num:=10;num>5{
		println(num)
	}

	if num:=20;num<5{
		println("我比5小")
	}else if toy:= 15;num>toy{
		fmt.Printf("%v比%v大",num,toy)
	}
	/*
		一般都是把其他函数的执行结果用来进行初始化变量的添加，例如
	*/

}