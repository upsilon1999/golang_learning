// package main
package sliceDelete

import (
	"fmt"
)

func main()  {
	
	/* 
		切片中元素的删除
		思路:
		对切片需要保留的部分进行截取,拼接成新切片 

		这样就把不需要的元素删除了
	*/
	allcourse := []string{"python","java","PHP","RUST","mysql"}
	// 例如我们要删除元素 PHP 
	sliceA := allcourse[:2] //[]string{"python", "java"}

	// 注意我们获取到的是元素还是切片

	sliceB := allcourse[3:] //[]string{"RUST","mysql"}

	// sliceB如果是切片就要扩展开
	sliceNew := append(sliceA,sliceB...)
	fmt.Printf("删除PHP后%#v\n",sliceNew)
	// ?? 末尾的mysql来源
	// 删除PHP后[]string{"python", "java", "RUST", "mysql"}

	// 原始切片
	fmt.Printf("原始切片%#v\n",allcourse)
	// 原始切片[]string{"python", "java", "RUST", "mysql", "mysql"}
	
	//验证是否会影响到初始切片 
	sliceNew = append(sliceNew, "typeScript")
	fmt.Printf("sliceNew被改变后:\n %#v\n 是否影响到allcourse\n %#v\n",sliceNew,allcourse)

	// sliceNew被改变后:
	// []string{"python", "java", "RUST", "mysql", "typeScript"}
	// 是否影响到allcourse
	//	?? 为什么受到影响
	// []string{"python", "java", "RUST", "mysql", "typeScript"}


	/* 
		思路2:
		实际上和上面的思路是同源的,采用切片访问的方式获得切片,
		赋值给原切片形成覆盖,
		或者赋值给新切片等价于删除
	*/
	sliceC := []string{"python","java","PHP","RUST","mysql"}
	sliceD := sliceC[2:3]
	sliceD = append(sliceD,"redis")
	fmt.Printf("sliceD被改变后:\n %#v\n 是否影响到sliceC\n %#v\n",sliceD,sliceC)
	// sliceD被改变后:
	// []string{"PHP", "redis"}
	// 是否影响到sliceC
	// []string{"python", "java", "PHP", "redis", "mysql"}
}