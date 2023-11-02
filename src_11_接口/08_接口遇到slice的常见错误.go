package main

import "fmt"

/*
	...interface{}允许传递0~多个interface{}的值
	他形式上是[]interface{}，可是实际上就是多个interface{}
	所以传入任意值都可以

		不会出现上述的切片问题，例如我传入一个 []string
	这里也只会把它作为 interface{}
	因为他是把所有参数组合为 []interface{},而不是指datas的类型是 []interface{}，datas的实际形成在逻辑处理时出现，并不是在某个参数传递时确定
*/
func mPrint(datas ...interface{}){
	for _,val := range datas{
		fmt.Println(val)
	}
}


func main()  {
	/* 
		[]interface{}本质上是一个切片
		和interface{}不一样 

		所以 []type ≠ []interface{}
	*/

	var data = []interface{}{
		"body",18,1.80,
	}
	// var a = "你好"
	mPrint(data)

	println("---------------------")
	var name = []string{
		"zs","lisi",
	}
	mPrint(name)

}