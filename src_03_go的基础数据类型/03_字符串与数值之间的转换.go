// package main
package stringAndNumber

import "strconv"

func main(){
	/* 
		需要用到strconv包，
		①字符串转数值，会得到两个返回值，第一个是成功的结果，第二个是错误对象err
		②数值转字符串，只有一个返回值，即成功的结果，因为理论上不可能失败
	*/
	var str = "32"

	myint,err := strconv.Atoi(str)
	if err != nil {
		println(err)
	}else{
		println(myint)
	}


	var myAge int = 32
	myStr := strconv.Itoa(myAge)
	println(myStr)
}