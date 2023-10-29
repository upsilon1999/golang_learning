// package main
package sliceCopy

import "fmt"

func main()  {
	// 复制slice


	/* 
		方案1:
		直接赋值
	*/
	var sliceA = []string{"A","B"}
	sliceB := sliceA

	fmt.Printf("%#v\n",sliceB)
	//[]string{"A", "B"}





	/* 
		方案2:
		截取全部 
	*/
	var sliceC =[]string{"C","D"}
	sliceD := sliceC[:]
	fmt.Printf("%#v\n",sliceD)


	println("------------------------------------------------")
	/* 
		方案3:
		copy(①，②)
		① 拷贝到的切片 
		② 被拷贝的切片 
		把切片②拷贝给切片①,覆盖对应位置的元素
	*/
	var sliceE = []string{"E","F"}

	var sliceF []string
	copy(sliceF,sliceE)
	fmt.Printf("%#v\n",sliceF)
	//[]string(nil)
	// 原因,拷贝并不会自动扩容，sliceF没有空间，自然无法接收拷贝的元素


	var sliceG = make([]string,5)
	copy(sliceG,sliceE)
	fmt.Printf("%#v\n",sliceG)
	//[]string{"E", "F", "", "", ""}


	var sliceH = []string{"GO","GOD","GOOD"}
	copy(sliceH,sliceE)
	fmt.Printf("%#v\n",sliceH)
	// []string{"E", "F", "GOOD"}
}