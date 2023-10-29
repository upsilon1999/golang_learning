// package main
package visitSlice

import "fmt"


func main()  {

	allCourses := []string{"javascript","html","css","python","golang"}

	/* 
			切片初始化完成后使用下标进行访问，
			访问超出长度的元素会产生越界错误

			直接访问某个下标，返回值是具体元素
	*/
	fmt.Printf("%v\n", allCourses[3])

	/* 
			取多个元素,用到的语法和用数组创建切片时的语法基本相似 

			切片[start:end]

			① start开始位置，end结束位置，左闭右开
			② start省略，相当于开始到end
			③	end省略，相当于从start到结束
			④	都省略，相当于取全部
	*/
	courses1 := allCourses[1:2]
	courses2 := allCourses[:3]
	courses3 := allCourses[1:]
	courses4 := allCourses[:]
	fmt.Printf("%#v\n %#v\n %#v\n %#v\n",courses1,courses2,courses3,courses4)
	println("-------------------------")

	/* 
			courses1容量为4，修改值不触发容量变化
			底层被修改，故courses1和allCourses都被修改
	*/
	courses1[0] = "Java"
	fmt.Printf("%#v\n %#v\n",courses1,allCourses)
	println("-------------------------")
	/* 
			courses1容量为5，元素长度为3,增加一个不触发容量变化
			返回的地址不变
			底层被修改，故courses2和allCourses都被修改
	*/
	courses2 = append(courses2, "PHP")
	fmt.Printf("%#v\n %#v\n",courses2,allCourses)
}