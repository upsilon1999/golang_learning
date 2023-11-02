// package main
package switchWithInterface

import (
	"fmt"
	"strings"
)

/*
	现在需要支持 任意类型的传入和返回
	switch进行类型判断，语法

	switch 空接口变量.(type) {
		case 类型①:
			逻辑
		case 类型②:
			逻辑
		default:
			逻辑
	}

	或者
	switch 变量:=空接口变量.(type) {
		case 判断①:
			逻辑
		case 判断①:
			逻辑
		default:
			逻辑
	}
*/
func addAll(a,b interface{}) interface{} {
	defer func ()  {
		if r:=recover();r!=nil {
			fmt.Println("这是不支持的运算")
		}
		
	}()


	// 由于我们需求中 a,b 是一样的类型，这里的写法会被简化 
	//正常情况处理这种，使用泛型最好 
	switch a.(type) {
		case int:
			ai,_ := a.(int)
			bi,_ := b.(int)
			return ai+bi
		case float64:
			ai,_ := a.(float64)
			bi,_ := b.(float64)
			return ai+bi
		case string:
			ai,_ := a.(string)
			bi,_ := b.(string)
			return ai+bi
		default:
			panic("not supported")
	}
}

func main()  {
	 a := "hello"
	 b := "world"
	 fmt.Printf("运行结果为%#v\n",addAll(a,b))

	//  虽然打印结果是个字符串，但是实际上addAll()的返回值是个interface{}
	// 是因为打印时做了隐式的断言
	// 所以我们要使用的时候需要先进行类型断言
	re := addAll(a,b)
	res,_ :=re.(string)
	// 只有断言后是字符串类型才可以用字符串的方法
	k := strings.Split(res,"l")
	fmt.Printf("%#v\n",k)
}