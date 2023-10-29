// package main
package rangeLearn

import "fmt"

func main()  {
	
	/* 
		又名 for-each range语法
		这种格式的循环可以对字符串、数组、切片、map、channel等进行迭代输出元素。
	*/
	/* 
		for key,value:= range 数组/字符串/切片/map/channel{

		}
	
		相当于其他语言中的forEach,
		对于 数组/字符串/切片 key就是下标
	*/
	var name = "abcdefg"
	for key,val := range name {
		//println(key)//下标
		// println(val)//直接打印是ASCII码
		fmt.Printf("%v %v\n",key,val)
	}

	//可以省略不想要的值，用匿名变量接收 
	for _,val := range name {
		fmt.Printf("字符值为%c\n",val)
	}


	var str = "好好学习"
	str1:=[]rune(str)
	println(str1[0])//拿到对应的码值 
	fmt.Printf("%c\n",str1[0])//好


	// 处理含中文字符串
	var mystr = "aa起来了"
	//①直接获取值是正常的
	
	for key,value := range mystr{
		println(key)
		fmt.Printf("%c\n",value)
	}

	//②通过索引获取值 
	opStr := []rune(mystr)
	// for value := range opStr {
	// 	fmt.Printf("%c\n",value)
	// }
	for i:=0;i<len(opStr);i++{
		fmt.Printf("%c\n",opStr[i])
	}

}