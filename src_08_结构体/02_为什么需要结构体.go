// package main
package whyNeedStruct

func main()  {
	/* 
		为什么需要结构体，
		例如我们要存储人的信息，包含姓名，年龄，身高 
		之前我们只能采取多维数组，类型单一，使用时要进行类型转换 
	*/
	var persons [][]string
	persons = append(persons, []string{"zs","18","180"})

	/* 
		类型麻烦，可以使用interface{}，存储的时候可以任意类型，但是使用的时候需要类型断言 
	*/
	var persons01 [][]interface{}
	persons01 = append(persons01, []interface{}{"lisi",18,1.8})
}