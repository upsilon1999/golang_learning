// package main
package findVariables

import "fmt"

var name = "zs"

func main()  {
	/* 
		变量的核心查找逻辑
		①
		查找顺序，同一代码块中往上查找,也就是按照顺序结构，从起点往上查找
		②
		本代码块查找 => 本代码块的祖先代码块查找 =>全局变量
		按照顺序查找，最先查到的定义位置就使用，查不到就报错

		③变量的生命周期是从定义开始，他的作用域也是从定义开始
		
		④变量的查找指的是查找变量最近定义的位置,然后计算变量的值
	*/

	fmt.Println(name)// zs

	name := "王五"

	if 5>3 {
		name:="lisi"
		val := name
		println(val)
	}

	fmt.Println(name) //王五

	/* 
		第一个打印name，按照顺序结构，从起点往上找，本代码块没有，继续祖先照，最终找到定义位置全局变量name

		第二个打印name,按照顺序结构，从起点往上照，跳过子代码块，最后先找到本代码块定义位置的局部变量name
	*/
}