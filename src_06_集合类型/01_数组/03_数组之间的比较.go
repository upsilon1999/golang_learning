// package main
package arrCompare

func main()  {
	/* 
		数组的比较原则：
		① count和type要相同，不相同认为是不同类型 
		② 在count和type一样时，对应的元素要相同才会认为相等
	*/
	course1 := [2]string{"a","b"}
	course2 := [3]string{"a","b","c"}
	course3 := [3]string{"a","c","b"}
	course4 := [...]string{"a","b","c"}

	/* 
		类型不相同无法比较 
		[2]string和[3]string是不同类型
	*/
	// if course1 == course2 {
	// 	println("相等")
	// }
	println(course1)

	/* 
		类型相同，对应位置值不同 
		不相等
	*/
	if course3 == course2 {
		println("相等")
	}else{
		println("不相等")
	}

	/* 
		[...]string{"a","b","c"} 最终解析为 [3]string{"a","b","c"}
		类型相同，都是 [3]string 
		对应位置元素相同,都是{"a","b","c"}
	*/
	if course2 == course4 {
		println("相等")
	}else{
		println("不相等")
	}
}