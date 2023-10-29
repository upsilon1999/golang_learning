// package main
package sliceStart

import "fmt"

func main()  {
		//切片的初始化 
		/* 
			三种方法:
				①从数组直接创建
				②赋值语法 
				③make
		*/

		/* 
			①从数组中创建
			切片名 := 数组名[下标A:下标B]

			左闭右开,将从 数组[下标A]截取到数组[下标B],但是不取数组[下标B]

			tips:
			1）取数组的所有元素 
			数组[0:len(数组名)]

			2）从某个位置开始到结束 
			数组[下标A:]

			3）从0开始到某个位置 
			数组[:下标B]

			4）所有元素
			数组[:]
		*/
		allCourses := [...]string{"javascript","html","css","python","golang"}
		// 取前两个
		slice1 := allCourses[0:2] 
		// 取数组所有的元素 
		slice2 := allCourses[0:len(allCourses)]
		//从某个位置开始到结束 
		slice3:=allCourses[2:]
		//从0开始到某个位置 
		slice4:=allCourses[:2]
		// 所有元素
		slice5:=allCourses[:]

		fmt.Printf("%#v\n",slice1)
		fmt.Printf("%#v\n",slice2)
		fmt.Printf("%#v\n",slice3)
		fmt.Printf("%#v\n",slice4)
		fmt.Printf("%#v\n",slice5)

		/* 
			②直接赋值
			注意值的类型要与设定的类型一致

			初始化之后，就有了初始长度，就能用下标访问了
		*/
		var myCourse = []string{"Japan","USA"}
		fmt.Printf("%#v\n",myCourse)

		/* 
			③make函数 
			make(type,len,cap)

			type: 切片类型 
			len :切片初始大小，相当于准备好了一个空间，这样在len范围内就可以用下标赋值了 
			cap 容量，可以不写，有扩容策略

			优势：
			1.准备好了空间，初始化的时候不会反复扩容，性能好
			2.空间中没有赋值的部分采用零值

			3.超出len用下标赋值会报越界错误，需要使用append添加元素
		*/
		myHobby := make([]string,3)

		myHobby[0] = "read"
		myHobby[1] = "study"
		myHobby[2] = "game"

		fmt.Printf("%#v\n",myHobby)
}