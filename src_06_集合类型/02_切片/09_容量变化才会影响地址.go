// package main
package capEdit

import "fmt"

func main()  {
	// 直接赋值初始化 
	var sliceC = []string{"vue","react","Nuxt","Nest","flutter"}
	fmt.Printf("sliceC的初始容量为%v\n",cap(sliceC))
	// sliceC的初始容量为5

	//如果切片的操作没有触发扩容，那么地址不发生改变都是在操作这个数组
	sliceC1 := sliceC[:]
	fmt.Printf("sliceC1的初始容量为%v\n",cap(sliceC1))
	//sliceC1是从sliceC开头进行截取，即从底层的"vue"开始，容量是5
	sliceC2 := sliceC[1:3]
	fmt.Printf("sliceC2的初始容量为%v\n",cap(sliceC2))
	//sliceC2是从sliceC下标1进行截取，即从底层的"react"开始，容量是4

	//切片截取是指针在底层数组上移动，且只考虑开头，位置根据下标加1得到
	//sliceC[:]  指针在(0+1)位，所以容量为5
	//sliceC[1:3]指针在(1+1)位，所以容量为4
	// 然后计算长度，截取对应内容展示，但是指针在截取开始的位置 

	sliceC1[2] = "Umi"
	// 容量没有发生改变，所以底层地址每变,映射出来的切片全部被修改 
	fmt.Printf("sliceC的为%#v\n",sliceC)
	// sliceC的为[]string{"vue", "react", "Umi", "Nest", "flutter"}
	fmt.Printf("sliceC1的为%#v\n",sliceC1)
	// sliceC1的为[]string{"vue", "react", "Umi", "Nest", "flutter"}
	fmt.Printf("sliceC2的为%#v\n",sliceC2)
	// sliceC2的为[]string{"react", "Umi"}



	/* 
		sliceC2为[]string{"react", "Umi"},
		只有2个元素，但是容量为4
		增加一个元素,并没有超过容量，所以地址不变 
		sliceC3和sliceC2共用地址，即sliceC底层地址 
		对应的就修改了底层的值
	*/

	sliceC3:=append(sliceC2,"AJAX")
	fmt.Printf("sliceC2的为%#v\n",sliceC2)
	// sliceC2的为[]string{"react", "Umi"}
	fmt.Printf("sliceC3的为%#v\n",sliceC3)
	// sliceC3的为[]string{"react", "Umi", "AJAX"}
	fmt.Printf("sliceC的为%#v\n",sliceC)
	// sliceC的为[]string{"vue", "react", "Umi", "AJAX", "flutter"}


	/* 
		sliceC2为[]string{"react", "Umi"},
		只有2个元素，但是容量为4
		增加3个元素,超过了容量，触发扩容，地址改变
		不影响sliceC2的底层地址 

		这就是为什么我们给某个切片使用append增加元素，需要用它来接收 
		因为append可能会改变地址
	*/
	sliceC4 := append(sliceC2,"a","b","c")
	fmt.Printf("sliceC2的为%#v\n",sliceC2)
	// sliceC2的为[]string{"react", "Umi"}
	fmt.Printf("sliceC3的为%#v\n",sliceC4)
	// sliceC3的为[]string{"react", "Umi", "a", "b", "c"}
	fmt.Printf("sliceC的为%#v\n",sliceC)
	// sliceC的为[]string{"vue", "react", "Umi", "AJAX", "flutter"}

	}