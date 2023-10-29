package main

func main()  {
	/* 
		算术运算符 
		+ - * / %(取余) ++ --
	*/
	var a,b = 1,2
	println(a+b)
	println(a-b)
	println(a*b)
	//运算并不会改变类型，所以还是int，所以这里相当于 
	// var c = int(a/b)
	println(a/b) //0
	println(a%b)
}