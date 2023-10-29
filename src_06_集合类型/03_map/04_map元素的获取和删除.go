// package main
package getMap

import "fmt"

func main()  {
	var course = map[string]string{
		"name":"zs",
		"gender":"man",
	}

	/* 
		我们不能使用 变量名[key]来判断map是否有对应的键 
	*/
	fmt.Printf("%#v\n",course["java"])
	// ""
	//如果key不存在，会返回对应的零值，但是如果有一个key的值就是零值就无法判断，所以不能

	/* 
		实际上 变量名[key] 是有两个返回值的， 
		①,② := 变量名[key]
		① key对应的值 
		② 布尔值,true代表存在，false代表不存在
	*/
	d,ok := course["sex"]
	if ok {
		println("存在",d)
	}else{
		println("不存在",d)
	}

	/* 
		删除元素 
		delete(map变量,key)

		删除不存在的元素不会报错
	*/
	delete(course,"age")
	

	/* 
		很重要的提示,
		map不是线程安全的，当你有多个线程同时对一个map进行操作，会报错

		并发时使用的map必须使用sync包下的Map，即syncMap
	*/
}