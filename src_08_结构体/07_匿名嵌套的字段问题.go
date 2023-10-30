// package main
package NestedWord

import "fmt"

type Persons struct{
	name string
	age  int
	height float32
}

type Persons01 struct{
	name string
	age  int
	height float32
}

// ①本地有字段，本地优先级最高
type Teacher struct{
	Persons
	name string 
	Persons01
}

// 字段存在冲突，在使用时会报错 
type Student struct {
	Persons01
	score float32
	Persons
}
func main()  {
	var t = Teacher{}
	t.name = "大老师"
	fmt.Printf("%#v\r\n", t)
	/* 
		main.Teacher{
			Persons:main.Persons{name:"", age:0, height:0},
			name:"大老师", 
			Persons01:main.Persons01{name:"", age:0, height:0},
		}
	*/

	var s = Student{}
	// 无法确定取的是哪个name,故报错
	// s.name = "李四"
	fmt.Printf("%#v\r\n",s)
	/* 
		main.Student{
			Persons01:main.Persons01{name:"", age:0, height:0},
			score:0, 
			Persons:main.Persons{name:"", age:0, height:0},
		}
	*/



}