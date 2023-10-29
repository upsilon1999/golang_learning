// package main
package constant

func main()  {
	//常量，定义的时候就指定的值，不能修改

	/* 
		①显式定义指明类型
		const 常量名 类型 = 常量值
	*/
	const PI float32= 3.1415

	/* 
		②隐式定义(不指名类型由值的类型推断)
		const 常量名 =  常量值
	*/
	const MY_INFO = "小白"

	/* 
		③定义多个常量，相当于抽离const
	*/
	const (
		UNKNOWN = 1
		FEMALE = 2
		MALE = 3
	)


	/* 
		④组合定义
		组合定义的规则:
		1.如果没有给定值和类型将沿用上面最近的一个值和类型
		2.所以第一个必须有值
	*/
	const (
		// 第一个必须有值
		// P


		X int = 16
		//沿用 X，相当于 const Y int = 16
		Y
		S = "abc"
		//沿用S，相当于 const Z="abc"
		Z
		//沿用S，相当于 const M="abc"
		M
	)


	/* 
		⑤常量必须赋初始值
	*/
	// const XI string

	/* 
		规范：①常量全部大写   PI
				 ②多个单词之间采用下划线分割 MY_INFO
		
		规则:
		①常量类型只可以定义bool、数字(整数、浮点数、复数)和字符串
		②不曾使用的常量不报错,不强制使用
		③显示指定类型的时候，必须指定常量左右值类型一致
		④常量必须赋初始值
	*/
}