// package main
package switchLearn

import "fmt"

func main()  {
	// 由于switch的功能if基本都能做，所以if用的多 

	/* 
		基本语法:
		switch var1 {
				case val1:
						...
				case val2:
						...
				default:
						...
		}

		①用var1依次去匹配case后面的值 
		②变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
		③测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
	*/
	var num = 100
	var grade string
	
	switch num {
		// 相当于判断 num == 80?
		case 80:
			grade = "B"	
		//多个条件相当于或，满足一个即可
		case 90,100:
			grade = "A"
		// default写在哪都可以，如果没有满足条件的就会使用它
		default:
			grade = "D"
		case 70:
			grade="C"
	}
	println(grade)

	/* 
		纯条件匹配： 
		switch {
			case 表达式1:
				逻辑 
			case 表达式2:
				逻辑 
			default:
				逻辑 
			...
		}

		实际上基本语法是这种形式的变种，基本语法的完整形式 
		是拿值进行比较，也是表达式，算简写形式
	*/
	var score int
	switch {
		case num == 80:
			score = 80
		case num ==90,num==100:
			score = 95
		default:
			score = 50
		case num==70:
			score = 100
	}
	println(score)

	// 常见用法，interface的类型判断
	/* 
	前提 x必须是interface类型的

		switch x.(type){
				case type1:
					statement(s);      
				case type2:
					statement(s); 
				default:
					statement(s);
		}
	
		x.(type)  获取x的真实类型

		同理可改写为
			switch {
				case x.(type) == type1:
					statement(s);      
				case x.(type) == type2:
					statement(s); 
				default:
					statement(s);
			}
	*/

	var userName interface{}
	
	switch userName.(type){
		case int:
			println("int类型")     
		case bool,float32:
			println("布尔或浮点类型") 
		case string:
			println("字符串类型")
		default:
			println("其他类型")
	}

	/* 
		switch默认情况下case最后自带break语句,匹配成功后就不会执行其他case,如果我们需要执行后面的case,可以使用fallthrough。

		使用fallthrough会强制执行后面的case语句,fallthrough不会判断下一条case的表达式结果是否为true。

		switch从第一个判断表达式为true的case开始执行，如果case带有fallthrough,程序会继续执行下一条case,且它不会去判断下一个 case的表达式是否为true
	*/

	switch {
		case false:
						fmt.Println("1、case 条件语句为 false")
						fallthrough
		case true:
						fmt.Println("2、case 条件语句为 true")
						fallthrough
		case false:
						fmt.Println("3、case 条件语句为 false")
						fallthrough
		case true:
						fmt.Println("4、case 条件语句为 true")
		case false:
						fmt.Println("5、case 条件语句为 false")
						fallthrough
		default:
						fmt.Println("6、默认 case")
	}
}