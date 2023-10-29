// package main
package compare

import "fmt"

func main()  {
	
	//字符串的比较
	a := "hello"
	b := "hello"
	c := "hboy"
	fmt.Println(a==b)//true
	fmt.Println(a!=b)//false

	//字符串的大小比较
	fmt.Println(a>c)//true

	/* 
		字符串的比较原则：
		从左到右挨个字符比较，一样大就比较下一个字符，否则大的哪个整个字符串就大 
		例如：
		①name :="hello"
		②age :="hboy"

		h与h比较，一样大，比较下一位 
		e与b比较，e比b大，所以①比②大
		比的是字符对应的码值
	*/

}