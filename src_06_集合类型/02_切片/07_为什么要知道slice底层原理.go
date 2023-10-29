// package main
package whyKnowSlice

import "fmt"

func printSlice(data []string)  {
	data[1] = "nodeJs"

	data = append(data,"ssh")
}


func main()  {
	/* 
		go的slice在函数传递的时候是值传递还是引用传递
		本质上： 值传递 
		但是又呈现出了不完全的引用传递的效果
	*/
	course := []string{"go","java","php"}
	printSlice(course)
	fmt.Printf("%#v",course)
	// []string{"go", "nodeJs", "php"}

	//?? 如果是值传递，那么在函数内部修改不应该对外部有影响
	//?? 如果是引用传递，为什么在函数内给切片增加元素没有对外部生效

}