// package main
package stringsBuilder

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	/* 
		计算字符串长度
	*/
	name := "abcdefg"
	println(len(name))

	/* 
		高性能字符串拼接
		通过strings的builder进行字符串拼接

		对性能要求高使用strings的builder，要求不高使用fmt的方法
	*/
	myname := "张无忌"
	age := 18
	var msg = fmt.Sprintf("我的名字是%s,今年%d岁",myname,age)
	println(msg)


	// 使用strings的builder实现上述同样的字符串
	// ① 定义一个builder的实例 
	var myBaby strings.Builder
	// ②往builder实例里面进行写入 
	/* 
		builder.Write(byte切片) 返回字符串长度和错误对象
		builder.WriteString(字符串) 返回字符串长度和错误对象
		builder.WriteByte(一个byte类型的参数) 返回错误对象，和其他三个方法有点不一样
		builder.WriteRune(一个rune类型的参数) 返回字符串长度和错误对象
	*/
	myBaby.WriteString("我的名字是")
	myBaby.WriteString(myname)
	myBaby.WriteString(",今年")
	myBaby.WriteString(strconv.Itoa(age))
	myBaby.WriteString("岁")

	//③拿到结果，使用builder对象上的String()
	re := myBaby.String()
	println(re)



}