// package main
package sliceAdd

import "fmt"

func main()  {
	allcourse := []string{"python","java","PHP","RUST"}

	/* 
		
		append(切片A,元素B)
		将元素B添加到切片A的末尾,返回值是一个切片

		根据打印结果可以得知:
		这个操作不改变原始切片,返回的是一个全新的切片 

		我们可以采取吧全新切片赋值给原切片的方式实现对原切片的添加
	*/
	newCourse := append(allcourse,"Ruby")
	fmt.Printf("新切片%#v\n 原始切片%#v\n",newCourse,allcourse)
	// 新切片[]string{"python", "java", "PHP", "RUST", "Ruby"}
 	// 原始切片[]string{"python", "java", "PHP", "RUST"}

	/* 
		append(切片A,元素B,元素C,...)
		可以一次性添加多个元素
	*/
	var sliceA []string
	// 将全新切片赋值给初始切片,就实现了对原始切片的数据修改
	sliceA = append(sliceA,"vue","react")
	fmt.Printf("一次性添加多个值%#v\n",sliceA)
	// 一次性添加多个值[]string{"vue", "react"}

	/* 
		append(切片A,切片B...)
		将切片A,切片B进行合并
		切片B被展开,使用其内部元素

		在append方法中使用 切片... 相当于js中的拓展运算符,展开其所有子元素
	*/
	sliceB := []string{"go","grpc"}
	sliceC := []string{"gin","beego"}
	sliceD := []string{"mysql","redis","es"}

	sliceE := append(sliceB,sliceC...)
	// 只想把silceD的一部分放进去,可以使用 silceD[x:y]对silceD进行截取
	sliceF := append(sliceB,sliceD[0:1]...)

	// 只能扩展一组
	// sliceE := append(sliceB,sliceC...,sliceD...)

	// 扩展不能结合其他插入一起用
	// sliceE := append(sliceB,sliceC...,"abc")
	// sliceE := append(sliceB,"grom",sliceD...)

	fmt.Printf("%#v\n %#v\n",sliceE,sliceF)
	// []string{"go", "grpc", "gin", "beego"}
 	// []string{"go", "grpc", "mysql"}
	
}