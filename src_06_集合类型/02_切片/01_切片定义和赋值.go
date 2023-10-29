// package main
package sliceLearn01

import "fmt"

func main()  {
	/* 
		切片slice可以理解成js中的array，python中的list
		因为数组不可改变长度，所以使用很不方便

		go的设计上弱化了本来的数组，从数组上抽离出切片来让大家操作
	*/

	/* 
		定义切片：
		var 切片名 []type 
		① 切片的底层是数组 
		② 由于切片长度可变，所以没有count
	*/
	var course []string
	fmt.Printf("%T\n", course) //[]string

	/* 
		①切片里面放切片多维切片，和多维数组一个性质
		 var 切片名 [][]int 
		②任意复杂类型都可以 
	*/


	/* 
		切片的赋值，使用append追加 

		① = append(②,元素)
		在②的基础上追加元素，返回一个新切片 

		如果①与②名字一样，就实现了赋值
	*/
	course = append(course, "go")
	course = append(course, "gin")
	fmt.Printf("%#v\n",course)

	/* 
		切片的取值 
		切片名[下标]

		如果下标超过了切片范围会报错
	*/
	fmt.Printf("%#v\n",course[0])
	fmt.Printf("%#v\n",course[2])
}