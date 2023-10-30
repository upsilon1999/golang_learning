// package main
package error01

import (
	"errors"
	"fmt"
)

func fn01() (int,error) {
	// 我们可以自定意任何错误为error
	//但是在golang中实际上有个专门的包生成error对象 
	/* 
		errors.New(字符串)
		这个包还有好几种用法，来返回不同的error
	*/
	return 0,errors.New("这是一个错误")
}

func main()  {
/* 
	一个函数可能出错，其他语言是用try...catch包住函数，防止整个代码崩溃
	golang则抛弃了这种思路，而是认为，如果出错了，就抛出一个错误对象，通过判断error是否为nil，来知道函数是否出错了

	开发函数的人需要有一个返回值去告诉调用者是否成功，返回值的处理更灵活，因为你可以决定是否使用或者接收他，也可以决定是否对外继续暴露
	不过按照开发规范，golang设计者要求我们必须要处理这个error,所以代码中会大量出现 if err !=nil，这种就叫做防御性编程，程序健壮性好
*/

	if _,err := fn01();err != nil{
		fmt.Printf("出错原因是:%#v", err)
	}
}