// package main
package arrayBase

import (
	"fmt"
	"strconv"
)

func main()  {
	
	/* 
		数组的定义:

		var 数组名 [count]type
		1.count 数组的元素个数 
		2.type  数组存放的值的类型

		含义：只有count个type类型的元素的数组类型 
		理解： 
		① 不同的[count]type类型是不一样的
		② 由于count和type强关联成一个新类型，所以count和type是不能修改的 
		③ 即数组定义了之后不能扩容

	*/

	/* 
		在golang中，[count]type会被看做一个独立的类型 
		所以不同的count和type不能转换
	*/
	var courses [3]string
	var courses1 [4]string 
	fmt.Printf("%T\n",courses)//[3]string
	fmt.Printf("%T\n",courses1)//[4]string



	/* 
		数组的赋值和值的修改 
		数组名[下标] = 值 

		① 初始赋值就是覆盖零值 
		② 第二次赋值就是替换上一次的值 
		③下标不能超过(数组元素长度-1),否则会报越界错误 
		④ 值的类型必须要满足数组允许的值类型
	*/
	courses[0] = "go"
	courses[1] = "java"
	courses[2] = "python"

	/* 
		数组的遍历，使用for...range或for循环 
		数组的值可以通过 数组名[下标] 来访问 
		所以可以通过遍历索引等手段来遍历数组
	*/ 
	for key,value := range courses{
		println(value)
		println("我是数组的第"+strconv.Itoa(key)+"个元素，值为"+courses[key])
	}


}