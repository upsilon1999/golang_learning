// package main
package otherType

import "fmt"

func main(){

	// 主要适用于存放字符
	var c byte
	c = 'a'

	//byte本质上是个数字，所以打印的是个数字
	//byte用于存放ASCII码，是其他语言中的char类型
	fmt.Println(c) //97

	//想要打印字符,需要格式化为char格式
	// %c 会把ASCII码转为码值对应字符
	fmt.Printf("%c",c)



	//byte的大小太小 
	//就是ASCII码和万国码的区别
	var c2 rune
	c2 = '幕'
	fmt.Printf("%c",c2)


}