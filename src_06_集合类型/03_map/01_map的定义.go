// package main
package defineMap

import "fmt"

func main()  {
	/* 
		map是一个key(索引)和value(值)的无序集合 
		主要查询方便，查询性能比切片高 
		时间复杂度为o(1)
	*/

	/* 
		map的定义 
		var ① map[②]③
		①map的变量名 
		②key的类型,
		key的类型不是任意的，例如切片就不行，因为不定，但是数组是可以的
		③value的类型
		可以是任何类型
	*/
	var course map[string]string

	fmt.Printf("%#v",course)
	// map[string]string(nil)


	// 报错，map的nil不能赋值(切片的nil是可以的)
	//即map类型想要设值，必须要先进行初始化
	// course["level"] = "hight"
}