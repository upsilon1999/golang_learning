// package main
package mapstart

import "fmt"

func main(){
	/* 
		map的初始化方式1:
		var 变量名 map[keyType]valueType = map[keyType]valueType{ key1:value1,...}
		
		简化 
		var 变量名 = map[keyType]valueType{ key1:value1,...}
	*/
	var course = map[string]string{
		// 最后一对值一定要以逗号结尾
		"name":"zs",
		"gender":"man",
	}
	fmt.Printf("%#v\n",course)
	// map[string]string{"gender":"man", "name":"zs"}

	/* 
	  取值:
			变量名[key]
		赋值和修改： 
			变量名[key] = value

		给不存在的key赋值就相当于增加对应的key，扩展方便
	*/
	course["level"] = "hight" 
	fmt.Printf("%#v\n",course)
	// map[string]string{"gender":"man", "level":"hight", "name":"zs"}

	/* 
		初始化也可以不放值
	*/
	var animal = map[string]string{}
	fmt.Printf("%#v\n",animal)
	//map[string]string{}


	/* 
		make初始化 
		make(类型,长度)
		map也有长度,底层的结构需要,不过我们用不到，会自动扩展
	*/
	books := make(map[string]string, 3)
	books["study"] = "php"

	/* 
		map必须初始化后才能赋值，赋值方式： 
		①写在大括号里 {key:value}
		②变量名[key] = value 
	*/
}