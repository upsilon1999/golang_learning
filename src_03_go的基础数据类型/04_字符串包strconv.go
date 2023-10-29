// package main
package learnStrconv

import "strconv"

func main()  {
	// 字符串转浮点数 
	/* 
		strconv.ParseFloat(①,②)
		①要转换的浮点数 
		②要转换多少位的浮点数,32或64

		返回值(①,②)
		①返回的是个float64的值，也就是无论我们第二个是64还是32，他返回的都是float64,
		如果我们填32，无非在用的时候会按32位去转，所以推荐填64
		②error错误对象
	*/
	myFloat,err := strconv.ParseFloat("3.14",64)

	if err!=nil {
		println(err)
	}else{
		println(myFloat)
	}
	
	//字符串转整数
	/* 
		strconv.ParseInt(①，②，③)
		①要转换的字符串 
		②要转换的进制，2 8 10 16常用这几个
		③字节数，64=>int64 8=>int8 ...

		返回值(①，②)
		①返回的是一个int64的值,如果错了，就是int64的零值
		②错误对象 
	*/
	/*
		这里的err不是重新定义而是赋值，逻辑上不会有冲突 
		因为 如果后面执行错误了，这个err就被赋予了错误的对象 
		如果后面执行成功了，err就会被赋值为nil，即使之前err有值也会被覆盖 
		因为之前err被赋值的类型也是错误对象error类型 

		牢记：就近原则
	*/
	myint,err := strconv.ParseInt("12",8,64)

	if err!=nil {
		println(err)
	}else{
		println(myint)//10 (12的八进制是10)
	}
	// ParseInt主要用于进制转换


		/* 
		strconv.ParseBool(①)
		①要转换的字符串 

		返回值(①，②)
		①返回的是一个bool类型的值,如果错了，就是bool的零值false
		②错误对象 
	*/
	// "true"、"1"" =>true "false"、"0"=>false ，其他都是错的
	myBool,err :=strconv.ParseBool("false")
	if err!=nil{
		println(err)
	}else{
		println(myBool)
	}


	//基本类型转字符串 
	/* 
		反过来转使用strconv包对应的FormatXxx方法
		strconv.FormatBool(①)
		①为bool类型
	*/
	println(strconv.FormatBool(true))
	
	/* 
		浮点数转字符串
		strconv.FormatFloat(①,②,③,④)
		①浮点数值 
		②类型，看源码 
		③看源码 
		④位数，32或64
	*/
	println(strconv.FormatFloat(3.1245,'E',-1,64))

	/* 
		数字转字符串
		strconv.FormatInt(①,②)
		①数字 
		②对应的进制
	*/
	println(strconv.FormatInt(42,16))
}