// package main
package multipleVariablesQuestion


func main()  {
	/* 
		①var 形式的多变量定义不允许重复定义已存在的变量
	*/
	var name = "zs"
	
	// var name,age = "wanwu",19
	println(name)

	/* 
		② 简洁形式
		多变量定义，至少有一个变量未定义才能通过
		如果变量已存在就相当于赋值，否则相当于定义
	*/
	var firstName = "zhang"

	//变量有未定义的，未定义进行定义，定义了的进行赋值
	firstName,lastName := "Li","shijie"
	println(firstName,lastName)


	/* 
		③简洁形式
		多变量定义，如果变量都已定义就不能重复定义
	*/
	var age1,age2 = 18,20

	//变量都已定义，不能重复定义
	// age1,age2 := 21,22
	println(age1,age2)

}