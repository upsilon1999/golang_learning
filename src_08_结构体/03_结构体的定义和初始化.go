// package main
package defineStruct

import "fmt"

/*
	结构体语法
	核心：结构体本质上也是一种自定义类型

	type 类型名 struct {
		字段A 类型
		字段B 类型
		...
	}

	注意：每行对应一个字段，不需要逗号结尾
*/
type Persons struct{
	name string 
	age  int 
	height float32
}

func main()  {
	// 如何初始化结构体
	/* 
	方式①
		结构体名{值A,值B,...}
	注意： 
		顺序和个数要与定义时一一对应
	*/
	p1 := Persons{"zs",18,1.8}
	fmt.Printf("%#v\n",p1)
	println("-----------------")
	/* 
		方式②：
			结构体名{
				字段A:值A,
				字段B:值B,
			}
		注意： 
		最后一个必须逗号结尾，字段需要在结构体中 
		没有顺序，没赋值的会采用类型零值
	*/
	p2 := Persons{
		name:"lisi",
	}
	fmt.Printf("%#v\n",p2)
	println("-----------------")


	/* 
		方式③：实例的切片，即多个实例 
	*/
	var p3 []Persons
	// 因为分离了，所以需要指明元素是Persons的实例
	p3 = append(p3,Persons{"zs",18,1.8})
	p3 = append(p3,p2)
	fmt.Printf("%#v\n",p3)
	println("-----------------")

	/* 
		方式④：实例切片的优化

		注意：
		1.这是省略了Persons
		2.最后一项一定要以逗号结尾
	*/
	var p4 = []Persons{
		{name:"lisi"},
		{name:"wangwu",age:18},
		// 可以嵌套着混合写，他会自动识别
		{"zs",18,1.8},
	}
	fmt.Printf("%#v\n",p4)
}