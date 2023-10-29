// package main
package globalVariablesScope

func a() int{

	return abc
}

// 全局变量
/* 
	①任意函数和文件任意位置都可以使用，
	注意全局变量和函数的书写位置没有关系,例如a 函数在全局变量之前定义，但是他仍然能正常使用全局变量
	因为函数的实际执行位置在全局变量定义之后
*/

/* 
	②但全局变量之间的书写位置有关联,如果val全局变量要用到abc全局变量，那么abc应该先于val定义
		函数外要使用全局变量，必须保证全局变量已定义
*/


var abc = 5
var val = abc -3

func main(){
	println(bb)
}

//虽然main函数写在前面，但是真正执行时，全局变量先定义
var bb = 9