// package main
package useIota

import "fmt"

func main()  {
	/* 
		①iota 特殊常量
		可以认为是一个能被编译器修改的常量
		在编译的时候根据iota自动推断值
	*/

	const (
		ERROR = iota
		SUCCESS = iota
		WIN = iota
		FAIL = iota
	)
	fmt.Println(ERROR,SUCCESS,WIN,FAIL) //0 1 2 3

	fmt.Println(WIN,FAIL,ERROR,SUCCESS) //2 3 0 1

	//可以发现在编译时他自动推断了值，与执行顺序无关


	/* 
		②iota内部有一个计数器，默认从0开始
		按照组合式定义进行优化

		由于沿用之前的值和类型，那么AGE2与AGE3的值都会是iota
		根据iota计数器自增，最终得到
		 AGE2  = 1
		 AGE3 = 2

		 注意：继承的值是iota，而不是iota计算后的值
		 例如值是iota+5 ，那么继承得到的值就是 iota+5
	*/

	const (
		AGE1 = iota
		AGE2
		AGE3
	)

	/* 
		③由于iota初始是0，那么我们可以修改初始值来改变起点
	*/
	const (
		ERR6 = iota+2 //2
		ERR9 //3
		ERR4 //4
	)

	/* 
		④ iota计数器不会中断
		1.每一次值都会被记录，C为字符串，但是iota计数器仍然执行了

		2.如果中断了iota，后续还要用他必须显式的恢复
		理解:如果你不恢复他，后续的就继承不了他，虽然他计数器没断，但是没有值用他
	*/
	const (
		A = iota //0
		B  //1
		C = "haha" //haha
		D	//haha
		E	//haha
		F = iota //5
		G //6
		H = iota+5 // 7+5 = 12
		I //iota+5 = 8+5 =13
	)
	fmt.Println(A,B,C,D,E,F,G,H,I) //0 1 haha haha haha 5 6 12 13

	/* 
		⑤每次出现const关键字的时候 iota归零
	*/
	const ABC = iota //0
	const (
		BCD = iota //0
		GDP  //1
	)
	const TNT = iota //0
}