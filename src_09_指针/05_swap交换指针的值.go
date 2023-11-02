// package main
package swap

import "fmt"

/*
	需求：定义一个方法，交换两个变量的值

	传参本质上都是值传递，将变量复制一份进行传递
	这里传入的是指针类型的值
*/
func swap(a,b *int) {
	*a,*b = *b,*a
}
func main()  {
	a := 5
	b := 7
	
	swap(&a,&b)
	fmt.Printf("%#v\n%#v",a,b)

}