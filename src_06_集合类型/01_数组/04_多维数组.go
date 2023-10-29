// package main
package multidimensional

import "fmt"

func main()  {
	
	/* 
		多维数组的定义 

		var 数组名 [①][②][...]类型

		从左到右，维度从高到低,左边代表右边的个数，层层嵌套
		[2][4]string => 有2个[4]string,
								 =>[4]string,每一项有4个string类型的元素
		数学书写格式:
		[
			[],
			[]
		]
		代码赋值格式 
		{
			{},
			{}
		}

		[2][3][4]string => 有2个[3][4]string, 
										=>[3][4]string 有3个[4]string
										=>[4]string,每一项有4个string类型的元素
		数学书写格式:
		[
			[
				[],
				[],
				[]
			],
			[
				[],
				[],
				[]
			]
		]
		代码书写格式：
		{
			{
				{},
				{},
				{}
			},
			{
				{},
				{},
				{}
			}
		}
	*/

	var arr [2][4]string 
	
	// 具体元素赋值
	arr[0][1] = "abc"
	fmt.Printf("%#v\n",arr)
	/* 
		{
			[4]string{"", "abc", "", ""}, 
			[4]string{"", "", "", ""}
		}
	*/

	// 按维度赋值 
	arr[1] = [4]string{2:"我",3:"好"}
	fmt.Printf("%#v\n",arr)
	/* 
		{
			[4]string{"", "abc", "", ""},
			[4]string{"", "", "我", "好"}
		}
		
	*/
}