// package main
package numberType

func main(){
	/*
		最常见的类型转换是int类型的转换，由于go的int类型太多了 
		很多我们内部的库或使用别人的库，经常需要对int类型进行转换 
	*/

	//float32或float64转int
	var a float32 =  0.5

	b := int(a) //float32转int，最后丢失小数变为0

	println(b)
	//int8转int64
	var c int8 = 15
	var d int64= int64(c)
	println(d)



	/* 
		①go允许在底层结构相同的两个类型之间互换
	*/
	// IT类型的底层是 int
	type IT int

	//e的类型是IT，底层是int,所以可以接收int类型的值
	var e IT = 5
	println(e)

	// 但是直接将int类型的变量赋值给IT会报错
	//虽然底层一样，但是在go里面他们仍然属于不一样的类型，需要显式转换
	var f int= 8
	// var g IT = f 
	//显式转换
	var h IT = IT(f)
	println(h)

}