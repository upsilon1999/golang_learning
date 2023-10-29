// package main
package learnFmt

import "fmt"

func main()  {
	
	/* 
		简单输出，不换行，可以配合转义字符换行
		fmt.Print(①,②,③,...)
		最后输出为一个字符串
	*/
	fmt.Print("你好")

	/* 
		输出，自动换行 
		fmt.Println(①,②,③,...)
		结果中间以空格隔开，放不下自动换行
	*/
	fmt.Println("大家好","你好")

	/* 
		格式化输出
		fmt.Printf(①,②,③,...)
		①格式化模板，里面有很多格式化占位符号 
		②及之后按顺序顶替占位符 

		占位符： 
		%s 字符串 
		%d 数字
		
	*/
	// 可读性强，但是性能没有前面两种好
	fmt.Printf("用户名:%s,年龄:%d","张三",18)

	/* 
		格式化转换 
		fmt.Stringf(①,②,③,...)
		①格式化模板 
		②及之后按顺序顶替占位符 

		这个方法返回一个字符串
		同理还有 fmt.String fmt.Stringln
		都是用于格式化生成字符串
	*/
	myStr := fmt.Sprintf("用户名:%s,年龄:%d","张三",18)
	print(myStr)
}