// package main
package capLearn

import (
	"fmt"
)

// go的slice本质上是个结构体
// type slice struct{
// 	array unsafe.Printer //用来存储实际数据的数组指针，指向一块连续的内存
// 	len   int            //切片中元素的数量
// 	cap   int            //array数组的长度
// }

func main()  {
	// 未初始化
	var ms []string 
	fmt.Printf("ms的初始容量为%v\n",cap(ms)) 
	//ms的初始容量为0

	//使用数组初始化
	var firstArr = [5]string{"go","java","php","C","C++"}
	sliceA := firstArr[:3]
	sliceB := firstArr[2:3]
	fmt.Printf("sliceA的初始容量为%v\n",cap(sliceA))
	fmt.Printf("sliceB的初始容量为%v\n",cap(sliceB))
	// sliceA的初始容量为5
	//sliceB的初始容量为3

	// 直接赋值初始化 
	var sliceC = []string{"vue","react","Nuxt","Nest"}
	fmt.Printf("sliceC的初始容量为%v\n",cap(sliceC))
	// sliceC的初始容量为4


	//make指定切片的容量
	var sliceD = make([]string, 0,5)
	fmt.Printf("sliceD的初始容量为%v\n",cap(sliceD))
	// sliceD的初始容量为5
}