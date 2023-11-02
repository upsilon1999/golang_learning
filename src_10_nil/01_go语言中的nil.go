package main

func main()  {
	/* 
		不同类型的数据零值不一样 
		bool      false
		numbers  0 
		string     ""
		pointer    nil
		slice      nil
		map        nil
		channel    nil
		interface  nil
		function   nil

		struct 默认值不是nil，默认值是具体字段的默认值
	*/

	/* 
		判断一个值是否为 nil

		变量名 == nil

		等于就为true，反之false

		注意： 
		除了基本数据类型外都可以使用此方法判断
		基本数据类型(bool\数字\字符串)不能和nil做判断
	*/
	var a map[string]string
	if a==nil {
		println("yes")
	}else{
		println("No")
	}

	/* 
		make初始化后不是nil，而是空的 

		所以make和new函数都是定义加初始化 

		例如我们会被切片分为 nil slice 和empty slice
		map分为 nil map 和empty map
		其他同理 

		nil的无和empty的空的区别 
		取值的时候都一样，nil的不能赋值
	*/
	var b = make([]string,0,0)
	if b == nil {
		println("yes")
	}else{
		println("No")
	}


	/* 
		nil与空的异同:
		nil的无和empty的空的区别 
		取值的时候都一样，nil的不能赋值
	*/
	c := make(map[string]string,0)
	var d map[string]string
	// 取值都一样 
	println(c["name"])
	println(d["name"])

	//赋值有差别 
	c["name"] = "李四" //make的进行了初始化，所以可以赋值
	d["name"] = "王五" //nil不可以赋值，会panic

}