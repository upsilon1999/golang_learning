// package main
package breakLearn

import (
	"fmt"
	"strconv"
)

func main()  {
	/* 
		跳出switch
		在其他语言中需要显式使用来防止穿透，在go中默认携带

		也可以显式写出防止继续执行
	*/
	var num = 40
	switch num {
		case 20:
			println("我很好")
			//可以省略
			break
		case 40:
			println("第一步")
			// 显式写出，后续的不再执行
			break
			println("被打断吟唱") 
		default:
			println("空空如也")
	}

	var str = "aaaaaaaaa"
	for key,_:= range str{
		if(key>3){
			break
		}
		println("我很"+strconv.Itoa(key))
	}

	/* 
		跳出标识的循环，标识不需要提前用变量声明
	*/
	mytab:
	for i := 1; i <= 3; i++ {
			fmt.Printf("i: %d\n", i)
			for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break mytab
		}
	}
}