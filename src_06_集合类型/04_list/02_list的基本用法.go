// package main
package listBase

import (
	"container/list"
	"fmt"
)


func main()  {
	// list不是个golang全局内置关键字，需要引入包 
	// container/list引入后就可以使用list了 

	// 定义
	/* 
		方式①：var 变量名 list.List
		方式②: 变量名 := list.New()
	*/
	var myList list.List
	// myList := list.New()

	
	// list方法非常多 
	/* 
		PushBack 尾部放数据 
		PushFront 头部放数据 
	*/
	myList.PushBack("golang") //在最后放一个 go 
	myList.PushBack("mySql")

	fmt.Println(myList)//直接打印不到值，是地址 
	// {{0xc0000a0180 0xc0000a01b0 <nil> <nil>} 2}

	/* 
		遍历打印值
		用法需要记
	*/
	// 正向遍历
	for i:=myList.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value)
	}
	// golang
	// mySql


	//倒序遍历 
	for i:=myList.Back();i!=nil;i=i.Prev(){
		fmt.Println(i.Value)
	}
	// mySql
	// golang

	/* 
		list的很多操作需要用到element元素，一种他自己返回的元素
	*/
	// 得到一个 list的element元素 
	ele:=myList.Front()
	// 通过循环遍历拿到golang对应的element元素
	for ;ele!=nil;ele=ele.Next(){
		if ele.Next().Value.(string)=="golang"{
			break
		}
	}
	// 在golang前面插入 gin
	myList.InsertBefore("gin",ele)

	//更多用法需要的时候看源码
}