# Golang基础

## helloworld

1.新建任意文件夹

2.安装go的模块管理工具

```sh
go mod init 项目名
# 使用模块管理工具进行项目管理
```

3.创建`helloworld.go`

```go
// 包名任意，与文件名没有任何关系
//但是main函数必须在main包中才能编译
// package main
package hello

//导包
import "fmt"

//main函数是go文件的入口
func main()  {
	
	/* 
		打印的方式
		① println() 不需要引入包
		② fmt包，功能丰富
	*/
	println("我是print -- Hello World")

	// 打印一行
	fmt.Println("fmt打印 -- HelloWorld")
}
```

4.运行

```sh
# 包名与文件名没有关联，但是main函数的包必须是main
# 如果我们main函数不在main包就会报如下错误
package command-line-arguments is not a main package
```

在学习时，我们都以文件为中心，所以书写时不用main包，编译测试时改为main包

```sh
go run helloworld.go
```

>输出

```sh
我是print -- Hello World
fmt打印 -- HelloWorld
```

5.编译产生可执行文件

```sh
go build ./src/helloworld.go
# 会生成可执行文件，windows下就是exe
```

运行可执行文件

```sh
./src/helloworld.exe
#直接输入生成的可执行文件就可以运行
```

### 注意事项

同一个目录下只能有一个main函数，新版本goland有一个设置，会自动按照文件来执行，就没有这个问题

## 变量与常量

### 变量的定义

>变量定义的方式总结

```go
// 运行时用这个
// package main

package defineVariable

import "fmt"

func main()  {
	/* 
		静态语言的特性
			①变量必须先定义后使用
			②变量必须有类型
			③类型定下来后不能改变
	*/

	/* 
		定义变量的方式① 
		var 变量名 变量类型
	*/
	var name string
	name = "zsf"

	/* 
		定义变量的方式②
		var 变量名 变量类型 = 变量值
	*/
	var firstName string = "wanwu"

	/* 
		定义变量的方式③ 
		定义时如果赋了初始值，那么类型可以省略，因为他会自动推断
		var 变量名 = 变量值
	*/
	var lastName = "shijie"

	/* 
		定义变量的方式④ 
		变量名 := 变量值

		优势：好用，相当于方式③的缩写
		缺点：无法显式指明类型
		
		只能用来声明局部变量，不能用来声明全局变量
	*/
	age := 18

	/* 
		go语言中局部变量定义了必须要使用
		fmt是格式化字符串，%v是占位符
	*/
	fmt.Printf("%v %v %v %v",name,firstName,lastName,age)
}
```

>运行

```sh
go run ./src/01_定义变量.go
```

>结果

```sh
zsf wanwu shijie 18
```

### 全局变量和局部变量

>全局变量和局部变量

```go
// package main
package GlobalAndLocalVariables

import "fmt"

/*
	go文件采用函数式编程，
	所以定义在go文件内，函数体外的就是全部变量
	函数体内的就是局部变量

	局部变量定义了就必须要使用
	全局变量定义了可以不使用
*/

// 全局变量
var name = "wanwu"
var age int

//简写形式 把公共的var抽离
var (
	ok bool
	firstName string = "zs"
	lastName = "李四"
)

func main()  {
	//局部变量
	isYes := true

	fmt.Printf("%v %v",isYes,name)
}
```

>运行

```sh
go run ./02_全局变量和局部变量.go 
```

>结果

```sh
true wanwu
```

#### 知识补充

```sh
局部变量：在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

全局变量：在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

备注：

1.局部变量不会一直存在，在函数被调用时存在，函数调用结束后变量就会被销毁，即生命周期。

2.Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
```



### 多变量的定义

>语法

```go
// package main
package multipleVariables

import "fmt"

func main()  {
	/* 
		①定义多个同类型的变量

		要求同类型，因为显式指定了类型
	*/
	var user1,user2,user3 string

	fmt.Printf("%v %v %v\n",user1,user2,user3)

	/* 
		② 定义多个变量的时候赋值

		由于没有显式指定类型，所以会采用类型推断的形式
		这样就允许不同类型的存在，变量与值一一对应
	*/
	var age1,age2,age3 = 18,19,"dddd"
	fmt.Printf("%v %v %v\n",age1,age2,age3)

	/* 
		③限定类型的同时赋值
	*/
	var he,she,it string = "他","她","它"
	fmt.Println(he,she,it)


	/* 
		④使用简洁定义多个变量，
		由于简洁变量没有类型定义，所以只能采用值推导的形式
	*/
	sex1,sex2 := "男","女"
	fmt.Println(sex1,sex2)


    /* 
		⑤ 相当于抽离var，是一种非常常用的多变量定义方式
	*/
	var (
		ok bool
		firstName string = "zs"
		lastName = "李四"
	)
	fmt.Println(ok,firstName,lastName)
    
    
	/* 
		注意：
			①变量必须先定义才能使用
			②go语言是静态语言，要求变量的类型和赋值类型一致

			③变量不能重复的定义，但是全局变量和局部变量可以重名，使用时优先使用本地的局部变量
			④ 简洁变量定义，即 (:=) 的形式不能用于全局变量定义
			⑤ 变量是有零值的，也就是默认值 
			⑥ 局部变量定义了一定要使用

			⑦多个变量定义与赋值要注意变量与值的位置一一对应
	*/
}
```

>运行

```sh
go run ./03_多变量的定义.go 
```

>结果

```sh
  
18 19 dddd
他 她 它
男 女
false zs 李四
```

#### 多变量定义的注意事项

>小结

```go
// package main
package multipleVariablesQuestion


func main()  {
	/* 
		①var 形式的多变量定义不允许重复定义已存在的变量
	*/
	var name = "zs"
	
	// var name,age = "wanwu",19
	println(name)

	/* 
		② 简洁形式
		多变量定义，至少有一个变量未定义才能通过
		如果变量已存在就相当于赋值，否则相当于定义
	*/
	var firstName = "zhang"

	//变量有未定义的，未定义进行定义，定义了的进行赋值
	firstName,lastName := "Li","shijie"
	println(firstName,lastName)


	/* 
		③简洁形式
		多变量定义，如果变量都已定义就不能重复定义
	*/
	var age1,age2 = 18,20

	//变量都已定义，不能重复定义
	// age1,age2 := 21,22
	println(age1,age2)

}
```

>运行

```sh
 go run ./04_多变量定义的注意事项.go
```

>结果

```sh
zs
Li shijie
18 20
```

### 常量的定义

>语法

```go
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
```

#### 组合定义

这个尤为注意

```go
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
```

### iota--特殊常量

```go
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
```

#### 注意事项

>1.iota是个特殊常量，继承时继承的是iota而不是iota计算后的值

```go
const (
		ERR6 = iota+2 //0+2 = 2
		ERR9 //iota+2 = 1+2 = 3
		ERR4 //iota+2 = 2+2 =4
	)
```

>2.计算器不会被打断,iota从0开始

```go
	const (
		A = iota+1 // 0+1 = 1   iota为0
		B  // iota+1 = 1+1 =2   iota为1
		C = "haha" //haha       iota为2
		D	//haha              iota为3
		E	//haha	            iota为4
		F = iota //5            iota为5
		G //6                   iota为6
		H = iota+5 // 7+5 = 12  iota为7
		I //iota+5 = 8+5 =13    iota为8
	) 
```

### 匿名变量

```go
// package main
package anonymity


func main()  {
	//匿名变量
	/* 
		对于局部变量如果我们定义了不使用就会报错
		需求：我们希望定义一个变量,不使用他也不会报错

		这就是匿名变量 
	*/

	/* 
		① 匿名变量就是一个短下划线 
	*/
	var _ int

	/* 
		②简洁变量定义不能单独定义短下划线 
		需要和其他变量一起使用，用于占位 
		因为正常情况变量和赋值要一一对应，这样就可以用来占位

		例如：右边有三个值，这样就可以实现name和age接收对应的值， 
		匿名变量用来忽略那个不想接收的
	*/
	name,_,age := "zs",true,10
	println(name,age)


	/* 
		③简洁变量定义左侧必须有一个未定义的变量
		匿名变量不属于未定义的变量，
		所以下述用法属于重复定义
	*/
	xray := "xray"

	println(xray)
	// xray,_ := "yray",20

	/* 
		常见用途：用于占位，例如函数返回多个返回值，我们只想要第三个，其他就可以用匿名变量占位
		这样就不会创建新的变量污染内存
	*/
}
```

### 变量的作用域

#### 全局变量的作用域

```go
// package main
package globalVariablesScope

func a() int{

	return abc
}

// 全局变量
/* 
	①任意函数和文件任意位置都可以使用，
	注意全局变量和函数的书写位置没有关系,例如a 函数在全局变量之前定义，但是他仍然能正常使用全局变量
	因为函数的实际执行位置在全局变量定义之后
*/

/* 
	②但全局变量之间的书写位置有关联,如果val全局变量要用到abc全局变量，那么abc应该先于val定义
		函数外要使用全局变量，必须保证全局变量已定义
*/


var abc = 5
var val = abc -3

func main(){
	println(bb)
}

//虽然main函数写在前面，但是真正执行时，全局变量先定义
var bb = 9
```

#### 局部变量的作用域

>语法

```go
// package main
package localVariablesScope


import "fmt"

func main()  {

/* 	
	① 
		局部变量只能在定义他的函数中使用，内部代码块也属于函数自身的一部分
		外部函数不能直接使用这个变量 -- 可以传参实现间接访问
		
		实际上传参给外部函数，也相当于把外部函数暂时变成了自身的代码块
*/
	err := "错误"
	println(err)

	
/* 
	② 
	变量的作用域和代码块绑定，实际在Go中局部变量是参考代码块的 
	即最近的一组大括号
*/
	{
		localName := "local"
		//localName是这个代码块中的局部变量
		//在{}外无法访问
		fmt.Println(localName)

		//err是外面父函数的变量，相当于闭包的概念，所以可以访问
		fmt.Println(err)
	}

	// 无法访问
	// fmt.Println(localName)

	/* 
		③ if 条件 { 代码块 }
	*/
	var maname string 
	if err == "错误" {
		// 这里重新赋值了，则只能在代码块中识别
		maname := "it"
		fmt.Println(maname)
	}
	//这里的值是最初在if代码块外面定义的 maname
	fmt.Println(maname)

	/* 
		④深度理解，局部变量是相对于代码块的
	*/
	if err == "错误" {
		abc := 2
		fmt.Println(abc)
	}else{
		// abc在这个代码块中没有
		// fmt.Println(abc)
	}

	/* 
		变量的查找逻辑 :本代码块查找 => 本代码块的祖先代码块查找 =>全局变量
		按照顺序查找，最先查到的就使用，查不到就报错
	*/

}
```

#### 变量查找逻辑

```go
// package main
package findVariables

import "fmt"

var name = "zs"

func main()  {
	/* 
		变量的核心查找逻辑
		①
		查找顺序，同一代码块中往上查找,也就是按照顺序结构，从起点往上查找
		②
		本代码块查找 => 本代码块的祖先代码块查找 =>全局变量
		按照顺序查找，最先查到的定义位置就使用，查不到就报错

		③变量的生命周期是从定义开始，他的作用域也是从定义开始
		
		④变量的查找指的是查找变量最近定义的位置,然后计算变量的值
	*/

	fmt.Println(name)// zs

	name := "王五"

	if 5>3 {
		name:="lisi"
		val := name
		println(val)
	}

	fmt.Println(name) //王五

	/* 
		第一个打印name，按照顺序结构，从起点往上找，本代码块没有，继续祖先照，最终找到定义位置全局变量name

		第二个打印name,按照顺序结构，从起点往上照，跳过子代码块，最后先找到本代码块定义位置的局部变量name
	*/
}
```

>作用域

```sh
① 顺序结构，从起点往上找
② 本代码块 => 祖先代码块 => 全局变量
③ 满足①②的最近的同名变量定义位置，从这开始作用域
```

## 基础数据类型

### 布尔类型

```sh
布尔型的值只可以是常量 true 或者 false。⼀个简单的例⼦：var b bool = true
```

>类型名

```sh
bool
```

>类型值

```sh
true
false
```

### 数值类型

#### 整数类型

```sh
# 第一位表示符号 正负
int8 有符号 8 位整型 (-128 到 127) ⻓度：8bit

int16 有符号 16 位整型 (-32768 到 32767)

int32 有符号 32 位整型 (-2147483648 到 2147483647)

int64 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
```

```sh
uint8 ⽆符号 8 位整型 (0 到 255) 8位都⽤于表示数值

uint16 ⽆符号 16 位整型 (0 到 65535)

uint32 ⽆符号 32 位整型 (0 到 4294967295)

uint64 ⽆符号 64 位整型 (0 到 18446744073709551615)
```

>使用举例

```go
var a int8
var b int64
```

>动态类型 int

```go
var e int
//int是一个动态类型，根据操作系统计算得到
操作系统 32位 为： int <=> int32
操作系统 64位 为： int <=> int64
```

>类型是不同的

```go
虽然都是整数类型，但他们实际上也是不同类型，是需要类型装换的
var a int8
var b int32

//报错
//a = b

//需要类型转换
a = int8(b)
```

#### 浮点类型

```sh
float32 32位浮点型数 #大约是3.4e38

float64 64位浮点型数 #大约是1.8e308
```

>注意

```sh
go没有float类型
```

>赋值

```go
var f1 float32 = 3.14
var f2 float32 = 3.14

var f3 float32 = 3
//go会自己做转换变成浮点数
```

#### 字符类型

```sh
在go中，字符使用 '',单引号包裹
字符串使用""，双引号包裹
```

##### byte类型

>源码

```go
type byte unit8
//byte就是unit8的类型别名
```

>作用

```sh
主要用来存放字符
```

>举例

```go
// 主要适用于存放字符
var c byte
c = 'a'

//byte本质上是个数字，所以打印的是个数字
//byte用于存放ASCII码，是其他语言中的char类型
fmt.Println(c) //97

//想要打印字符,需要格式化为char格式
// %c 会把ASCII码转为码值对应字符
fmt.Printf("%c",c)
```

>注意

```sh
由于本质上是个数字，所以是可以计算的
```

```go
var k byte
k = 'a'+1 //98
// 98对应的ASCII码为字符 b
```

##### rune类型

>源码

```go
type rune int32
//由于byte类型的大小限制，他表示不了很多文字
//实际上就是ASCII码和万国码的区别
```

>作用

```sh
用于存放字符
```

>举例

```go
//byte的大小太小 
//就是ASCII码和万国码的区别
var c2 rune
c2 = '幕'
fmt.Printf("%c",c2)
```

##### string类型

>区别

```sh
字符类型用'',字符是指单个的
字符串类型用"",顾名思义，存放一个或多个字符
```

>举例

```go
var abc string
abc = "sjsjsj"
```

## 基本数据类型的装换

### 数值类型之间的转换

>语法

```go
var valueOfTypeB = typeB (valueOfTypeA)

valueOfTypeB := typeB (valueOfTypeA)
```

>背景

```sh
最常见的类型转换是int类型的转换，由于go的int类型太多了,很多我们内部的库或使用别人的库，经常需要对int类型进行转换 

所有的数值类型都可以这样进行转换，但是注意转换时可能会丢失精度，例如大的数往小的转，或者小数转整数
```

>举例

```go
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
```

```sh
go允许在底层结构相同的两个类型之间互换
理解：
①底层一样，所以数据可以互换，即赋值的时候接收的数据范围一样
②但是仍然属于两个类型，他们的变量不可以互换，仍然需要显式转换
```

```go
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
```

### 整数与字符串之间的转换

```sh
需要用到strconv包，
①字符串转数值，会得到两个返回值，第一个是成功的结果，第二个是错误对象err
②数值转字符串，只有一个返回值，即成功的结果，因为理论上不可能失败
```

>举例

```go
// package main
package stringAndNumber

import "strconv"

func main(){
	/* 
		需要用到strconv包，
		①字符串转数值，会得到两个返回值，第一个是成功的结果，第二个是错误对象err
		②数值转字符串，只有一个返回值，即成功的结果，因为理论上不可能失败
	*/
	var str = "32"

	myint,err := strconv.Atoi(str)
	if err != nil {
		println(err)
	}else{
		println(myint)
	}


	var myAge int = 32
	myStr := strconv.Itoa(myAge)
	println(myStr)
}
```

### strconv包

常用的方法

```sh
ParseXxx,将字符串转其他类型
FormatXxx,将其他转字符串
Itoa int转字符串
Atoi 字符串转int
```

>举例

```go
package main

import "strconv"

func main()  {
	// 字符串转浮点数 
	/* 
		strconv.ParseFloat(①,②)
		①要转换的浮点数 
		②要转换多少位的浮点数,32或64

		返回值(①,②)
		①返回的是个float64的值，也就是无论我们第二个是64还是32，他返回的都是float64,
		如果我们填32，无非在用的时候会按32位去转，所以推荐填64
		②error错误对象
	*/
	myFloat,err := strconv.ParseFloat("3.14",64)

	if err!=nil {
		println(err)
	}else{
		println(myFloat)
	}
	
	//字符串转整数
	/* 
		strconv.ParseInt(①，②，③)
		①要转换的字符串 
		②要转换的进制，2 8 10 16常用这几个
		③字节数，64=>int64 8=>int8 ...

		返回值(①，②)
		①返回的是一个int64的值,如果错了，就是int64的零值
		②错误对象 
	*/
	/*
		这里的err不是重新定义而是赋值，逻辑上不会有冲突 
		因为 如果后面执行错误了，这个err就被赋予了错误的对象 
		如果后面执行成功了，err就会被赋值为nil，即使之前err有值也会被覆盖 
		因为之前err被赋值的类型也是错误对象error类型 

		牢记：就近原则
	*/
	myint,err := strconv.ParseInt("12",8,64)

	if err!=nil {
		println(err)
	}else{
		println(myint)//10 (12的八进制是10)
	}
	// ParseInt主要用于进制转换


		/* 
		strconv.ParseBool(①)
		①要转换的字符串 

		返回值(①，②)
		①返回的是一个bool类型的值,如果错了，就是bool的零值false
		②错误对象 
	*/
	// "true"、"1"" =>true "false"、"0"=>false ，其他都是错的
	myBool,err :=strconv.ParseBool("true1")
	if err!=nil{
		println(err)
	}else{
		println(myBool)
	}


	//基本类型转字符串 
	/* 
		反过来转使用strconv包对应的FormatXxx方法
		strconv.FormatBool(①)
		①为bool类型
	*/
	println(strconv.FormatBool(true))
	
	/* 
		浮点数转字符串
		strconv.FormatFloat(①,②,③,④)
		①浮点数值 
		②类型，看源码 
		③看源码 
		④位数，32或64
	*/
	println(strconv.FormatFloat(3.1245,'E',-1,64))

	/* 
		数字转字符串
		strconv.FormatInt(①,②)
		①数字 
		②对应的进制
	*/
	println(strconv.FormatInt(42,16))
}
```

>结果

```sh
+3.140000e+000
10
(0x488b88,0xc000106f40)
true
3.1245E+00
2a
```

### 推荐类型转换包

```sh
"gconv"  
类型转换的包
```

## 运算符和表达式

### 算术运算符

下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例               |
| :----- | :--- | :----------------- |
| +      | 相加 | A + B 输出结果 30  |
| -      | 相减 | A - B 输出结果 -10 |
| *      | 相乘 | A * B 输出结果 200 |
| /      | 相除 | B / A 输出结果 2   |
| %      | 求余 | B % A 输出结果 0   |
| ++     | 自增 | A++ 输出结果 11    |
| --     | 自减 | A-- 输出结果 9     |

>举例

```go
package main

import "fmt"

func main() {

   var a int = 21
   var b int = 10
   var c int

   c = a + b
   fmt.Printf("第一行 - c 的值为 %d\n", c )
   c = a - b
   fmt.Printf("第二行 - c 的值为 %d\n", c )
   c = a * b
   fmt.Printf("第三行 - c 的值为 %d\n", c )
   c = a / b
   fmt.Printf("第四行 - c 的值为 %d\n", c )
   c = a % b
   fmt.Printf("第五行 - c 的值为 %d\n", c )
   a++
   fmt.Printf("第六行 - a 的值为 %d\n", a )
   a=21   // 为了方便测试，a 这里重新赋值为 21
   a--
   fmt.Printf("第七行 - a 的值为 %d\n", a )
}
```

>结果

```go
第一行 - c 的值为 31
第二行 - c 的值为 11
第三行 - c 的值为 210
第四行 - c 的值为 2
第五行 - c 的值为 1
第六行 - a 的值为 22
第七行 - a 的值为 20
```

#### 注意事项

Go 的自增，自减只能作为表达式使用，而不能用于赋值语句。

```go
a++ // 这是允许的，类似 a = a + 1,结果与 a++ 相同
a-- //与 a++ 相似
a = a++ // 这是不允许的，会出现编译错误 syntax error: unexpected ++ at end of statement
```

两个字符串使用加法运算是拼接

```go
var a string ="abc"
var b string = "def"
c := a+b 
Println(c)//abcdef
```

### 关系运算符

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述                                                         | 实例              |
| :----- | :----------------------------------------------------------- | :---------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       | (A == B) 为 False |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   | (A != B) 为 True  |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   | (A > B) 为 False  |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   | (A < B) 为 True   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True  |

>举例

```go
package main

import "fmt"

func main() {
   var a int = 21
   var b int = 10

   if( a == b ) {
      fmt.Printf("第一行 - a 等于 b\n" )
   } else {
      fmt.Printf("第一行 - a 不等于 b\n" )
   }
   if ( a < b ) {
      fmt.Printf("第二行 - a 小于 b\n" )
   } else {
      fmt.Printf("第二行 - a 不小于 b\n" )
   } 
   
   if ( a > b ) {
      fmt.Printf("第三行 - a 大于 b\n" )
   } else {
      fmt.Printf("第三行 - a 不大于 b\n" )
   }
   /* Lets change value of a and b */
   a = 5
   b = 20
   if ( a <= b ) {
      fmt.Printf("第四行 - a 小于等于 b\n" )
   }
   if ( b >= a ) {
      fmt.Printf("第五行 - b 大于等于 a\n" )
   }
}
```

>结果

```sh
第一行 - a 不等于 b
第二行 - a 不小于 b
第三行 - a 大于 b
第四行 - a 小于等于 b
第五行 - b 大于等于 a
```

### 逻辑运算符

下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。

| 运算符 | 描述                                                         | 实例               |
| :----- | :----------------------------------------------------------- | :----------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False  |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 | (A \|\| B) 为 True |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 | !(A && B) 为 True  |

>举例

```go
package main

import "fmt"

func main() {
   var a bool = true
   var b bool = false
   if ( a && b ) {
      fmt.Printf("第一行 - 条件为 true\n" )
   }
   if ( a || b ) {
      fmt.Printf("第二行 - 条件为 true\n" )
   }
   /* 修改 a 和 b 的值 */
   a = false
   b = true
   if ( a && b ) {
      fmt.Printf("第三行 - 条件为 true\n" )
   } else {
      fmt.Printf("第三行 - 条件为 false\n" )
   }
   if ( !(a && b) ) {
      fmt.Printf("第四行 - 条件为 true\n" )
   }
}
```

>结果

```sh
第二行 - 条件为 true
第三行 - 条件为 false
第四行 - 条件为 true
```

### 位运算符

很多底层库会用到，也叫做性能运算符

位运算符对整数在内存中的二进制位进行操作。

下表列出了位运算符 &, |, 和 ^ 的计算：

| p    | q    | p & q | p \| q | p ^ q |
| :--- | :--- | :---- | :----- | :---- |
| 0    | 0    | 0     | 0      | 0     |
| 0    | 1    | 0     | 1      | 1     |
| 1    | 1    | 1     | 1      | 0     |
| 1    | 0    | 0     | 1      | 1     |

假定 A = 60; B = 13; 其二进制数转换为：

```sh
A = 0011 1100

B = 0000 1101

-----------------

A&B = 0000 1100

A|B = 0011 1101

A^B = 0011 0001
```

Go 语言支持的位运算符如下表所示。假定 A 为60，B 为13：

| 运算符 | 描述                                                         | 实例                                   |
| :----- | :----------------------------------------------------------- | :------------------------------------- |
| &      | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
| \|     | 按位或运算符"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 | (A \| B) 结果为 61, 二进制为 0011 1101 |
| ^      | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 | (A ^ B) 结果为 49, 二进制为 0011 0001  |
| <<     | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >>     | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 | A >> 2 结果为 15 ，二进制为 0000 1111  |

>举例

```go
package main

import "fmt"

func main() {

   var a uint = 60      /* 60 = 0011 1100 */  
   var b uint = 13      /* 13 = 0000 1101 */
   var c uint = 0          

   c = a & b       /* 12 = 0000 1100 */ 
   fmt.Printf("第一行 - c 的值为 %d\n", c )

   c = a | b       /* 61 = 0011 1101 */
   fmt.Printf("第二行 - c 的值为 %d\n", c )

   c = a ^ b       /* 49 = 0011 0001 */
   fmt.Printf("第三行 - c 的值为 %d\n", c )

   c = a << 2     /* 240 = 1111 0000 */
   fmt.Printf("第四行 - c 的值为 %d\n", c )

   c = a >> 2     /* 15 = 0000 1111 */
   fmt.Printf("第五行 - c 的值为 %d\n", c )
}
```

>结果

```sh
第一行 - c 的值为 12
第二行 - c 的值为 61
第三行 - c 的值为 49
第四行 - c 的值为 240
第五行 - c 的值为 15
```

### 赋值运算符

下表列出了所有Go语言的赋值运算符。

| 运算符 | 描述                                           | 实例                                  |
| :----- | :--------------------------------------------- | :------------------------------------ |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C |
| +=     | 相加后再赋值                                   | C += A 等于 C = C + A                 |
| -=     | 相减后再赋值                                   | C -= A 等于 C = C - A                 |
| *=     | 相乘后再赋值                                   | C *= A 等于 C = C * A                 |
| /=     | 相除后再赋值                                   | C /= A 等于 C = C / A                 |
| %=     | 求余后再赋值                                   | C %= A 等于 C = C % A                 |
| <<=    | 左移后赋值                                     | C <<= 2 等于 C = C << 2               |
| >>=    | 右移后赋值                                     | C >>= 2 等于 C = C >> 2               |
| &=     | 按位与后赋值                                   | C &= 2 等于 C = C & 2                 |
| ^=     | 按位异或后赋值                                 | C ^= 2 等于 C = C ^ 2                 |
| \|=    | 按位或后赋值                                   | C \|= 2 等于 C = C \| 2               |

>举例

```go
package main

import "fmt"

func main() {
   var a int = 21
   var c int

   c =  a
   fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )

   c +=  a
   fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )

   c -=  a
   fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )

   c *=  a
   fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )

   c /=  a
   fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )

   c  = 200; 

   c <<=  2
   fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )

   c >>=  2
   fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )

   c &=  2
   fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )

   c ^=  2
   fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )

   c |=  2
   fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )

}
```

>结果

```sh
第 1 行 - =  运算符实例，c 值为 = 21
第 2 行 - += 运算符实例，c 值为 = 42
第 3 行 - -= 运算符实例，c 值为 = 21
第 4 行 - *= 运算符实例，c 值为 = 441
第 5 行 - /= 运算符实例，c 值为 = 21
第 6行  - <<= 运算符实例，c 值为 = 800
第 7 行 - >>= 运算符实例，c 值为 = 200
第 8 行 - &= 运算符实例，c 值为 = 0
第 9 行 - ^= 运算符实例，c 值为 = 2
第 10 行 - |= 运算符实例，c 值为 = 2
```

### 其他运算符号

下表列出了Go语言的其他运算符。

| 运算符 | 描述             | 实例                       |
| :----- | :--------------- | :------------------------- |
| &      | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
| *      | 指针变量。       | *a; 是一个指针变量         |

>举例

```go
package main

import "fmt"

func main() {
   var a int = 4
   var b int32
   var c float32
   var ptr *int

   /* 运算符实例 */
   fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

   /*  & 和 * 运算符实例 */
   ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
   fmt.Printf("a 的值为  %d\n", a);
   fmt.Printf("*ptr 为 %d\n", *ptr);
}
```

>结果

```sh
第 1 行 - a 变量类型为 = int
第 2 行 - b 变量类型为 = int32
第 3 行 - c 变量类型为 = float32
a 的值为  4
*ptr 为 4
```

#### 注意事项

指针变量 ***** 和地址值 **&** 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。当变量前面有 ***** 标识时，才等同于 **&** 的用法，否则会直接输出一个整型数字。

>举例

```go
func main() {
   var a int = 4
   var ptr *int
   ptr = &a
   println("a的值为", a);    // 4
   println("*ptr为", *ptr);  // 4
   println("ptr为", ptr);    // 824633794744
}
```

### 运算符优先级

有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| 优先级 | 运算符           |
| :----- | :--------------- |
| 5      | * / % << >> & &^ |
| 4      | + - \| ^         |
| 3      | == != < <= > >=  |
| 2      | &&               |
| 1      | \|\|             |

>举例

```go
package main

import "fmt"

func main() {
   var a int = 20
   var b int = 10
   var c int = 15
   var d int = 5
   var e int;

   e = (a + b) * c / d;      // ( 30 * 15 ) / 5
   fmt.Printf("(a + b) * c / d 的值为 : %d\n",  e );

   e = ((a + b) * c) / d;    // (30 * 15 ) / 5
   fmt.Printf("((a + b) * c) / d 的值为  : %d\n" ,  e );

   e = (a + b) * (c / d);   // (30) * (15/5)
   fmt.Printf("(a + b) * (c / d) 的值为  : %d\n",  e );

   e = a + (b * c) / d;     //  20 + (150/5)
   fmt.Printf("a + (b * c) / d 的值为  : %d\n" ,  e );  
}
```

>结果

```sh
(a + b) * c / d 的值为 : 90
((a + b) * c) / d 的值为  : 90
(a + b) * (c / d) 的值为  : 90
a + (b * c) / d 的值为  : 50
```

## 字符串的操作补充

### 下标取值

```go
// 字符串[下标] 能拿到对应字符
var str = "abc"
println(str[0])//对应ASCII码
fmt.Printf("%c",str[0])//a

//但是对应中文会有问题
var str1 = "好好学习"
println(str1[0])//乱码

//处理方式
str2 := []rune(str1)
println(str2[0])//对应码值
fmt.Printf("%c",str2[0])//好
```



### 转义字符

| 转义字符 |                 意义                  | ASCII 码值（⼗进制） |
| :------: | :-----------------------------------: | :------------------: |
|    \n    | 换⾏ (LF) ，将当前位置移到下⼀⾏开头  |         010          |
|    \r    |  回⻋ (CR) ，将当前位置移到本⾏开头   |         013          |
|    \t    | ⽔平制表 (HT) （跳到下⼀个 TAB 位置） |         009          |
|   `\\`   |         代表⼀个反斜线字符`\`         |         092          |
|   `\'`   |      代表⼀个单引号（撇号）字符       |         039          |
|   `\"`   |          代表⼀个双引号字符           |         034          |
|    \?    |             代表⼀个问号              |         063          |

### 格式化输出

>包

```sh
主要用到fmt包，用于打印或者得到格式化字符串
# 专门学习fmt包
```

>举例

```go
// package main
package learnFmt

import "fmt"

func main()  {
	
	/* 
		简单输出，不换行，可以配合转义字符换行
		fmt.Print(①,②,③,...)
		最后输出为一个字符串
	*/
	fmt.Print("你好")

	/* 
		输出，自动换行 
		fmt.Println(①,②,③,...)
		结果中间以空格隔开，放不下自动换行
	*/
	fmt.Println("大家好","你好")

	/* 
		格式化输出
		fmt.Printf(①,②,③,...)
		①格式化模板，里面有很多格式化占位符号 
		②及之后按顺序顶替占位符 

		占位符： 
		%s 字符串 
		%d 数字
		
	*/
	// 可读性强，但是性能没有前面两种好
	fmt.Printf("用户名:%s,年龄:%d","张三",18)

	/* 
		格式化转换 
		fmt.Stringf(①,②,③,...)
		①格式化模板 
		②及之后按顺序顶替占位符 

		这个方法返回一个字符串
		同理还有 fmt.String fmt.Stringln
		都是用于格式化生成字符串
	*/
	myStr := fmt.Sprintf("用户名:%s,年龄:%d","张三",18)
	print(myStr)
}
```

>常用打印或转换占位符

```sh
"%v"
任意值
"%T"
输出类型
"%#v"
输出go的格式，用于打印结构体数组切片等

fmt包对性能消耗比较大，追求性能少用它
```

### 字符串常用函数

#### 字符串长度`len`

```go
/* 
    计算字符串长度
*/
name := "abcdefg"
println(len(name))//7
```

#### 字符串高性能拼接

fmt包的字符串拼接对性能消耗太大，所以一般不采用

高性能使用strings上的Builder方法

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	/* 
		高性能字符串拼接
		通过strings的Builder进行字符串拼接

		对性能要求高使用strings的Builder，要求不高使用fmt的方法
	*/
	myname := "张无忌"
	age := 18
	var msg = fmt.Sprintf("我的名字是%s,今年%d岁",myname,age)
	println(msg)


	// 使用strings的Builder实现上述同样的字符串
	// ① 定义一个Builder的实例 
	var myBaby strings.Builder
	// ②往Builder实例里面进行写入 
	/* 
		builder.Write(byte切片) 返回字符串长度和错误对象
		builder.WriteString(字符串) 返回字符串长度和错误对象
		builder.WriteByte(一个byte类型的参数) 返回错误对象，和其他三个方法有点不一样
		builder.WriteRune(一个rune类型的参数) 返回字符串长度和错误对象
	*/
	myBaby.WriteString("我的名字是")
	myBaby.WriteString(myname)
	myBaby.WriteString(",今年")
	myBaby.WriteString(strconv.Itoa(age))
	myBaby.WriteString("岁")

	//③拿到结果，使用Builder对象上的String()
	re := myBaby.String()
	println(re)

}
```

### 字符串的比较

```sh
从左到右比较每一个字符的码值，如果一样大就比较下一位，如果不一样，那么大的那个字符对应的那个字符串整体大
```

>举例

```go
// package main
package compare

import "fmt"

func main()  {
	
	//字符串的比较
	a := "hello"
	b := "hello"
	c := "hboy"
	fmt.Println(a==b)//true
	fmt.Println(a!=b)//false

	//字符串的大小比较
	fmt.Println(a>c)//true

	/* 
		字符串的比较原则：
		从左到右挨个字符比较，一样大就比较下一个字符，否则大的哪个整个字符串就大 
		例如：
		①name :="hello"
		②age :="hboy"

		h与h比较，一样大，比较下一位 
		e与b比较，e比b大，所以①比②大
		比的是字符对应的码值
	*/

}
```

### 字符串的常用方法

```sh
字符串的常用方法基本都在strings的包中
包里面很多方法记不住很正常，用的时候直接看源码

tips:Ctrl+O 可以知道goland文件中有哪些方法
```

>举例

```go
package main

import (
	"fmt"
	"strings"
)

func main()  {
	
	/* 
		字符串的常用方法基本都在strings的包中
		包里面很多方法记不住很正常，用的时候直接看源码
	*/

	// 是否包含某个子串
	/* 
		strings.Contains(①,②)
		①是目标字符串 
		②是子串 
		判断①中是否包含了②，返回true或false
	*/
	res1:=strings.Contains("hello","lo")
	println(res1) //true

	// 查询子串出现的次数 
	/* 
			strings.Contains(①,②)
			①是目标字符串 
			②是子串 
			判断②在①中出现了多少次，返回int
	*/
	res2 := strings.Count("my boy","y")
	println(res2)//2

	// 字符串分割
  /* 
		strings.Split(①,②)
		①是目标字符串 
		②是子串 
		①按照②进行分割，得到一个字符串切片
	*/
	res3 := strings.Split("my-a-low","-")
	fmt.Printf("%#v\n",res3)

	// 字符串是否包含前缀，是否包含后缀 
	/* 
		strings.HasPrefix(①,②)
		①目标字符串 
		②前缀字符串 
		①是否以②开头
		返回值布尔值
	*/
	println(strings.HasPrefix("IMfree","IM"))//true
	
	/* 
		strings.HasSuffix(①,②)
		①目标字符串 
		②后缀字符串 
		①是否以②结尾
	*/
	println(strings.HasSuffix("youAre","re"))//true

	// 查询子串出现的位置
	/* 
		strings.Index(①,②)
		①目标字符串 
		②子串 
		查询②首次出现在①中的下标索引,返回值是一个int类型
		如果存在就返回下标索引，不存在就返回-1

		注意：他找的是byte字节的位置，所以有中文就不准确,因为每个中文算4个字节
	*/
	res4:=strings.Index("my god!go~","go")//3
	res5:=strings.Index("hello","go")//-1
	res6:=strings.Index("时代先锋go","go")
	println(res4,res5)
	println(res6)//12

	// 子串替换 
	/* 
		strings.Replace(①,②,③,④)
		①目标字符串 
		②要被替换的部分 
		③替换成什么 
		④替换几个，如果小于0，则默认替换全部
	*/
	println(strings.Replace("I love Go","Go","Python",-1))//I love Python


	// 大小写转换 
	/* 
		strings.ToLower(①) 
		①目标字符串
		将目标字符串小写

		strings.ToUpper(①) 
		①目标字符串
		将目标字符串大写
	*/
	println(strings.ToLower("HEllo"))//hello

	println(strings.ToUpper("HEllo"))//HELLO

	// 去掉首尾指定的字符
	/* 
		strings.Trim(①,②)
		①目标字符串 
		②减除字符串 
		从①的首尾移除②包含的字符，常用来去除空格，前缀或后缀

		他是把②拆成多个字符，只要①的开头和末尾有这些字符就删除，删到有不满足条件的为止

		strings.TrimLeft()
		只解决左边的 
		strings.TrimRight()
		只解决右边的
	*/
	res7:=strings.Trim(" hello go"," ")
	res8:=strings.Trim("##hello go","##")
	res9:=strings.Trim("$#hel$lo$","#$")
	println(res7)//hello go
	println(res8)//hello go
	println(res9)//hel$lo
}
```

## 条件判断

### if条件判断

>举例

```go
package main

func main()  {
	var age = 18
	/* 
	情况一：
		if 布尔表达式 {
			布尔表达式为true时执行的逻辑
		}

		①布尔表达式：一个返回值为true或false的表达式
	*/
	if age>3{
		println("我超过三岁了")
	}

	/* 
		情况二： 
		if (布尔表达式) {
			布尔表达式为true时执行的逻辑
		}

		兼容了其他语言的代码习惯，将布尔表达式用括号包裹也是一样的
		但是为了规范起见，一般不去写
		括号多用于分组，例如 
		(age>3)&&(age<20)
	*/

	/* 
		情况三： 
		if 条件① {
			满足条件①
		} else if 条件② {
			满足条件②
		}
	*/
	if age < 18 {
		println("你还太年轻了")
	}else if age == 18{
		println("恭喜长大")
	} 

	/* 
		情况四： 
		if 条件① {
			满足条件①
		} else if 条件② {
			满足条件②
		}else{
			其他if条件都不满足时执行
		}
	*/
	if age < 16 {
		println("你还太年轻了")
	}else if age>18{
		println("你是个大人了")
	}else{
		println("至死是少年")
	}
    
    /* 
		if 变量初始化;布尔表达式 {
			逻辑
		}

		①在if之后，条件语句之前，可以添加变量初始化语句，使用;进行分隔
		②这里初始化的变量是一个局部变量，可以在之后的布尔表达式和逻辑块中使用
	*/
	if num:=10;num>5{
		println(num)
	}

	if num:=20;num<5{
		println("我比5小")
	}else if toy:= 15;num>toy{
		fmt.Printf("%v比%v大",num,toy)
	}
}
```

>注意

```sh
用括号包含条件在go里面是兼容的，只是一般不用，go里面括号一般用来分组
```

>常用

```go
/* 
    if 变量初始化;布尔表达式 {
        逻辑
    }

    ①在if之后，条件语句之前，可以添加变量初始化语句，使用;进行分隔
    ②这里初始化的变量是一个局部变量，可以在之后的布尔表达式和逻辑块中使用
*/
if num:=10;num>5{
    println(num)
}

if num:=20;num<5{
    println("我比5小")
}else if toy:= 15;num>toy{
    fmt.Printf("%v比%v大",num,toy)
}

/*
	一般都是把其他函数或操作的执行结果用来进行初始化变量的添加
*/
```

### switch条件判断

switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。

switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 **fallthrough** 。

>基本匹配

```go
/* 
    基本语法:
    switch var1 {
            case val1:
                    ...
            case val2:
                    ...
            default:
                    ...
    }

    ①用var1依次去匹配case后面的值 
    ②变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
    ③测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
*/
var num = 100
var grade string
switch num {
    // 相当于判断 num == 80?
    case 80:
        grade = "B"	
    //多个条件相当于或，满足一个即可
    case 90,100:
        grade = "A"
    // default写在哪都可以，如果没有满足条件的就会使用它
    default:
        grade = "D"
    case 70:
        grade="C"
}
println(grade)
```

>纯条件匹配

```go
/* 
    纯条件匹配： 
    switch {
        case 表达式1:
            逻辑 
        case 表达式2:
            逻辑 
        default:
            逻辑 
        ...
    }

    实际上基本语法是这种形式的变种，基本语法的完整形式 
    是拿值进行比较，也是表达式，算简写形式
*/
var score int
switch {
    case num == 80:
        score = 80
    case num ==90,num==100:
        score = 95
    default:
        score = 50
    case num==70:
        score = 100
}
println(score)
```

>interface类型匹配

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

实际上就是参与判断的是类型

```go
/* 
前提 x必须是interface类型的,目的获取他存储的实际内容的类型

    switch x.(type){
            case type1:
                statement(s);      
            case type2:
                statement(s); 
            default:
                statement(s);
    }

    x.(type)  获取x的真实类型

    同理可改写为
        switch {
            case x.(type) == type1:
                statement(s);      
            case x.(type) == type2:
                statement(s); 
            default:
                statement(s);
        }
*/

var userName interface{}

switch userName.(type){
    case int:
        println("int类型")     
    case bool,float32:
        println("布尔或浮点类型") 
    case string:
        println("字符串类型")
    default:
        println("其他类型")
}

//特殊用法
//可以声明interface的局部变量
switch i := userName.(type){
    
}
//其他的都是不行的
```

>穿透

 switch默认情况下case最后自带break语句,匹配成功后就不会执行其他case,如果我们需要执行后面的case,可以使用`fallthrough`。

使用fallthrough会强制执行后面的case语句,fallthrough不会判断下一条case的表达式结果是否为true。

switch从第一个判断表达式为true的case开始执行，如果case带有fallthrough,程序会继续执行下一条case,且它不会去判断下一个 case的表达式是否为true

```go
switch {
    case false:
            fmt.Println("1、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("2、case 条件语句为 true")
            fallthrough
    case false:
            fmt.Println("3、case 条件语句为 false")
            fallthrough
    case true:
            fmt.Println("4、case 条件语句为 true")
    case false:
            fmt.Println("5、case 条件语句为 false")
            fallthrough
    default:
            fmt.Println("6、默认 case")
}
```

## for循环

#### 基本语法

go里面只有一个for循环，因为他足够做所有的事

```go
package main

func main()  {
	/* 
		for ①;②;③{
			④
		}

		①初始化 临时变量，非必填
		②条件判断  非必填
		③循环步进逻辑 
		④循环的内容

		② --> ④ --> ③-->② --> ④-->...
	*/

	for i:=0;i<10;i++{
		println(i)
	}

	// 省略局部初始化,但要保证判断逻辑的变量存在 
	//for也是个代码块，他也会遵循向外寻找变量的逻辑 
	var b = 100
	for ;b<105;b++{
		println(b)
	}

	// 步进条件可以写在循环体内 
	for ;b<104; {
		println("我来巡航")
		b++
	}

	/* 
		for会根据我们写的内容补全到对应位置，所以如果那一项为空，则分号可以去掉 
		原因是每个位置的内容不一样
		
		for ;b<105;b++{} => for b<105;b++{}
		for ;b<104;{} => for b<104{}
		for ;;{} =>for{} 相当于while true死循环
		
		死循环也是可以跳出的，只要设定了条件，死循环多用于做轮询
	*/

}
```

#### For-each range循环

这种格式的循环可以对字符串、数组、切片等进行迭代输出元素。

>语法

```go
/* 
    又名 for-each range语法
    这种格式的循环可以对字符串、数组、切片、map、channel等进行迭代输出元素。
*/
/* 
    for key,value:= range 数组/字符串/切片/map/channel{

    }

    相当于其他语言中的forEach,
    对于 数组/字符串/切片 key就是下标
*/
var name = "abcdefg"
for key,val := range name {
    //println(key)//下标
    // println(val)//直接打印是ASCII码
    fmt.Printf("%v %v\n",key,val)
}

//可以省略不想要的值，用匿名变量接收 
for _,val := range name {
    fmt.Printf("字符值为%c\n",val)
}



```

>结果

```sh
0 97
1 98
2 99
3 100
4 101
5 102
6 103
字符值为a
字符值为b
字符值为c
字符值为d
字符值为e
字符值为f
字符值为g
```

>注意点

```sh
如果遍历的是字符串、数组、切片
key 索引
value 对应的索引的值的"拷贝",所以我们在for循环里面修改遍历的值改不到原始值
如果只返回一个参数，返回的是索引
```

```sh
如果遍历的是map
key map的键
value 键对应的值的"拷贝"
`注意`：如果只返回一个参数，返回的是值
```

```sh
如果遍历的是channel
只有一个返回值，channel接收的数据
```

>对含中文字符的遍历

```go
var str = "好好学习"
str1:=[]rune(str)
println(str1[0])//拿到对应的码值 
fmt.Printf("%c\n",str1[0])//好

// 处理含中文字符串
var mystr = "aa起来了"
//①直接获取值是正常的

for key,value := range mystr{
    println(key)
    fmt.Printf("%c\n",value)
}

//②通过索引获取值 
opStr := []rune(mystr)

//无法获取
// for value := range opStr {
// 	fmt.Printf("%c\n",value)
// }

//正常获取
for i:=0;i<len(opStr);i++{
    fmt.Printf("%c\n",opStr[i])
}
```

### break

在 Go 语言中，break 语句用于终止当前循环或者 switch 语句的执行，并跳出该循环或者 switch 语句的代码块。

break 语句可以用于以下几个方面：。

- 用于循环语句中跳出循环，并开始执行循环之后的语句。
- break 在 switch 语句中在执行一条 case 后跳出语句的作用。
- break 可应用在 select 语句中。
- 在多重循环中，可以用标号 label 标出想 break 的循环。

select未来结合实际新知识研究

#### 跳出switch判断

```go
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
```

>结果

```sh
第一步
```

#### 跳出最近的循环体

```go
var str = "aaaaaaaaa"
for key,_:= range str{
   
    println("我很"+strconv.Itoa(key))
    //因为判断在执行之后，所以会有 key为4的输出
    if(key>3){
        break
    }
}
```

>结果

```sh
我很0
我很1
我很2
我很3
我很4
```

#### 多重循环跳出到标签循环体

给循环体加一个标签，

```sh
break 标签
# 如果break在标签标识的循环体中，就代表跳出这个循环体
"要在这个循环体中"
```

>语法

```sh
标签名:
	for循环
```

>举例

```go
mytab:
for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
        for i2 := 11; i2 <= 13; i2++ {
        fmt.Printf("i2: %d\n", i2)
        break mytab
    }
}
```

>输出

```sh
i: 1
i2: 11
```

>注意

```sh
标签是临时的，不需要提前用变量声明
```

### continue

Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。

for 循环中，执行 continue 语句会触发 for 增量语句的执行。

在多重循环中，可以用标号 label 标出想 continue 的循环。

#### 跳出单次循环

```go
/* 定义局部变量 */
var a int = 10

/* for 循环 */
for a < 20 {
  if a == 15 {
     /* 跳过此次循环 */
     a = a + 1;
     continue;
  }
  fmt.Printf("a 的值为 : %d\n", a);
  a++;    
}  
```

> 结果

```sh
a 的值为 : 10
a 的值为 : 11
a 的值为 : 12
a 的值为 : 13
a 的值为 : 14
a 的值为 : 16
a 的值为 : 17
a 的值为 : 18
a 的值为 : 19
```

#### 跳出标识的循环

```go
re:
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                continue re
            }
    }
```

>结果

```sh
i: 1
i2: 11
i: 2
i2: 11
i: 3
i2: 11
```

### goto

Go 语言的 goto 语句可以无条件地转移到过程中指定的行。

goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。

但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。

>语法

```go
goto label;
..
.
label: statement;
```

>举例

在变量 a 等于 15 的时候跳过本次循环并回到循环的开始语句 LOOP 处：

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* 循环 */
   LOOP: for a < 20 {
      if a == 15 {
         /* 跳过迭代 */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++     
   }  
}
```

>结果

```sh
a的值为 : 10
a的值为 : 11
a的值为 : 12
a的值为 : 13
a的值为 : 14
a的值为 : 16
a的值为 : 17
a的值为 : 18
a的值为 : 19
```

## 集合类型

之前都是基本的数据类型，实际开发中都把基本数据类型放到不同的容器中来满足我们的需求，这些容器就是一个个的集合，go语言中的集合类型都有哪些？

```sh
①数组
②切片slice
③map
④list
```

### 数组

#### 定义

```go
var 数组名 [count]type
1.count 数组的元素个数 
2.type  数组存放的值的类型

含义：只有count个type类型的元素的数组类型 
理解： 
① 不同的[count]type类型是不一样的
② 由于count和type强关联成一个新类型，所以count和type是不能修改的 
③ 即数组定义了之后不能扩容
```

>举例

```go
/* 
    在golang中，[count]type会被看做一个独立的类型 
    所以不同的count和type不能转换
*/
var courses [3]string
var courses1 [4]string 
fmt.Printf("%T\n",courses)//[3]string
fmt.Printf("%T",courses1)//[4]string
```

>结果

```sh
[3]string
[4]string
```

#### 数组的访问

>语法

```sh
数组的赋值和值的修改 
数组名[下标] = 值 

① 初始赋值就是覆盖零值 
② 第二次赋值就是替换上一次的值 
③下标不能超过(数组元素长度-1),否则会报越界错误 
④ 值的类型必须要满足数组允许的值类型
```

>举例

```go
courses[0] = "go"
courses[1] = "java"
courses[2] = "python"
```

#### 数组的遍历

>原理

```sh
数组的遍历，使用for...range或for循环 
数组的值可以通过 数组名[下标] 来访问 
所以可以通过遍历索引等手段来遍历数组
```

>举例

```go
for key,value := range courses{
    println(value)
    println("我是数组的第"+strconv.Itoa(key)+"个元素，值为"+courses[key])
}
```

>结果

```sh
go
我是数组的第0个元素，值为go
java
我是数组的第1个元素，值为java
python
我是数组的第2个元素，值为python
```

#### 数组的初始化

##### 方案1

>语法

```go
var 数组名 [count]type = [count]type{元素1,元素2,...}

由于赋值的时候已经指明类型，故左边类型可以省略 
var 数组名 = [count]type{元素1,元素2,...}
```

>举例

```go
var courses = [3]string{"go","java","python"}
fmt.Printf("%#v\n",courses)
```

>结果

```sh
[3]string{"go", "java", "python"}
```

##### 方案2

>语法

```go
数组名 := [count]type{元素1,元素2,...}
```

>举例

```go
courses1 := [4]string{"box1","box2","box3","box4"}
fmt.Printf("%#v\n",courses1)
```

>结果

```sh
[4]string{"box1", "box2", "box3", "box4"}
```

##### 方案3

>语法

```go
指定下标初始化，其他地方采用零值 
其实之前的初始化可以理解为所有下标都给了值，所以省略了下标

var 数组名 = [count]type{下标:值,下标:值}

数组名 := [count]type{下标:值,下标:值}
```

>举例

```go
courses2 := [3]string{2:"gin"}
var courses3 = [2]string{0:"beego"}
fmt.Printf("%#v\n",courses2)
fmt.Printf("%#v\n",courses3)
```

>结果

```sh
[3]string{"", "", "gin"}
[2]string{"beego", ""}
```

##### 方案4

>语法

```go
用省略号代替count，根据赋值自动推算count值 
① 只有{}赋值的形式有效，
②下标赋值法无效,因为没给定count不知道下标
③会根据{}元素的个数推算出类型 

例:
[...]int{1，2} =>根据个数推算 =>[2]int{1,2}
[...]int{1,2,3} =>[3]int{1,2,3}

优势：方便代码的书写
```

>举例

```go
courses4 := [...]string{"abc","更好"}
var courses5 = [...]int{1,2,3,4}

//报错，下标语法无效
// courses5 := [...]string{5:"天天"} 

fmt.Printf("%#v\n",courses4)
fmt.Printf("%#v\n",courses5)
```

>结果

```sh
[2]string{"abc", "更好"}
[4]int{1, 2, 3, 4}
```

#### 数组的比较

>原则

```sh
① count和type要相同，不相同认为是不同类型 
② 在count和type一样时，对应位置的元素要相同才会认为相等
```

>举例

```go
course1 := [2]string{"a","b"}
course2 := [3]string{"a","b","c"}
course3 := [3]string{"a","c","b"}
course4 := [...]string{"a","b","c"}

/* 
    类型不相同无法比较 
    [2]string和[3]string是不同类型
*/
// if course1 == course2 {
// 	println("相等")
// }
println(course1)

/* 
    类型相同，对应位置值不同 
    不相等
*/
if course3 == course2 {
    println("相等")
}else{
    println("不相等")
}

/* 
    [...]string{"a","b","c"} 最终解析为 [3]string{"a","b","c"}
    类型相同，都是 [3]string 
    对应位置元素相同,都是{"a","b","c"}
*/
if course2 == course4 {
    println("相等")
}else{
    println("不相等")
}
```

#### 多维数组

##### 维度理念

多维数组的定义

```go
var 数组名 [①][②][...]类型
//从左到右，维度从高到低,左边代表右边的个数，层层嵌套
```

> 举例1

```sh
[2][4]string 
=> 有2个[4]string,
=>[4]string,每一项有4个string类型的元素
```

数学书写格式:

```sh
[
    [],
    []
]
```

代码赋值格式

```sh
{
    {},
    {}
}
```

>举例2

```sh
[2][3][4]string 
=> 有2个[3][4]string, 
=>[3][4]string 有3个[4]string
=>[4]string,每一项有4个string类型的元素
```

数学书写格式

```sh
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
```

代码赋值格式

```sh
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
```

##### 多维数组的赋值

和一维数组一样，本质上多维数组就是用低维数组嵌套得到，换个角度，多维数组就是个巨大的一维数组，所以访问方法和一维数组一样

注意下标即可

```go
var arr [2][4]string 
	
// 具体元素赋值
arr[0][1] = "abc"
fmt.Printf("%#v\n",arr)


// 按维度赋值 
arr[1] = [4]string{2:"我",3:"好"}
fmt.Printf("%#v\n",arr)
```

>打印结果

```sh
# 具体元素赋值对应
{
    [4]string{"", "abc", "", ""}, 
    [4]string{"", "", "", ""}
}

# 按维度赋值
{
    [4]string{"", "abc", "", ""},
    [4]string{"", "", "我", "好"}
}
```

同理，多维数组的遍历和一维数组一样

### 切片

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

#### 为什么要了解切片的底层原理

```go
package main

import "fmt"

func printSlice(data []string)  {
	data[1] = "nodeJs"

	data = append(data,"ssh")
}


func main()  {
	/* 
		go的slice在函数传递的时候是值传递还是引用传递
		本质上： 值传递 
		但是又呈现出了不完全的引用传递的效果
	*/
	course := []string{"go","java","php"}
	printSlice(course)
	fmt.Printf("%#v",course)
	// []string{"go", "nodeJs", "php"}

	//?? 如果是值传递，那么在函数内部修改不应该对外部有影响
	//?? 如果是引用传递，为什么在函数内给切片增加元素没有对外部生效

}
```

#### slice底层原理

slice本质上是一个结构体

```go
// go的slice本质上是个结构体
type slice struct{
	array unsafe.Printer //用来存储实际数据的数组指针，指向一块连续的内存 
	len   int            //切片中元素的数量 
	cap   int            //array数组的长度
}

①所有的切片都可以看做是底层数组在外面的投影
②切片的操作，会作用到投影他的底层数组上
③切片触发扩容操作，底层数组改变，即投影他的数组不再是之前的数组(长大了，抄一份离开家，别毁了这个家)
④切片的容量和他的长度无关，是从他在底层指针位置到指针结束
⑤切片增加元素超出自己长度，但是没超过容量，地址就不会改变，就会对原底层数组进行改变。
```

#### 容量

容量是切片中一个至关重要的东西，只要容量发生变化，切片底层地址就会发生改变。

>语法

```sh
# 获取切片的容量
cap(切片)
```

##### 容量的大小

```go
// 未初始化
var ms []string 
fmt.Printf("ms的初始容量为%v\n",cap(ms)) 
//ms的初始容量为0

//使用数组初始化
var firstArr = [5]string{"go","java","php","C","C++"}
sliceA := firstArr[:3]
sliceB := firstArr[2:3]
fmt.Printf("sliceA的初始容量为%v\n",cap(sliceA))
fmt.Printf("sliceB的初始容量为%v\n",cap(sliceB))
// sliceA的初始容量为5
//sliceB的初始容量为3

// 直接赋值初始化 
var sliceC = []string{"vue","react","Nuxt","Nest"}
fmt.Printf("sliceC的初始容量为%v\n",cap(sliceC))
// sliceC的初始容量为4


//make指定切片的容量
var sliceD = make([]string, 0,5)
fmt.Printf("sliceD的初始容量为%v\n",cap(sliceD))
// sliceD的初始容量为5
```

##### 容量与地址

```go
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
```

##### 小结

```sh
①容量不改变，地址不发生变化，所有的修改和各种操作都会影响到底层数组
②容量发生变化，地址发生变化，底层分离
```

#### 切片的定义

```sh
var 切片名 []type 
① 切片的底层是数组 
② 由于切片长度可变，所以没有count
```

>举例

```go
var course []string
fmt.Printf("%T\n", course) 
```

>结果

```sh
[]string
```

>注意

```sh
①切片里面放切片多维切片，和多维数组一个性质
 var 切片名 [][]int 
②任意复杂类型都可以 
```

#### 切片的初始化

切片没有初始化之前没有长度，我们只能使用`append`进行元素添加，不能使用下标赋值

```sh
初始化的三种方法:
①从数组直接创建
②赋值语法 
③make
```

##### 从数组初始化

```sh
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
```

>举例

```go
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
```

>结果

```sh
[]string{"javascript", "html"}
[]string{"javascript", "html", "css", "python", "golang"}
[]string{"css", "python", "golang"}
[]string{"javascript", "html"}
[]string{"javascript", "html", "css", "python", "golang"}
```

##### 直接赋值

```sh
②直接赋值
注意值的类型要与设定的类型一致

初始化之后，就有了初始长度，就能用下标访问了
```

>举例

```go
var myCourse = []string{"Japan","USA"}
fmt.Printf("%#v\n",myCourse)
```

>结果

```sh
[]string{"Japan", "USA"}
```

##### make语法

性能最好

```sh
③make函数 
make(type,len,cap)

type: 切片类型 
len :切片初始大小，相当于准备好了一个空间，这样在len范围内就可以用下标赋值了 
cap 容量，可以不写，有扩容策略

优势：
1.准备好了空间，初始化的时候不会反复扩容，性能好
2.空间中没有赋值的部分采用零值

3.超出len用下标赋值会报越界错误，需要使用append添加元素
```

>举例

```go
myHobby := make([]string,3)

myHobby[0] = "read"
myHobby[1] = "study"
myHobby[2] = "game"

fmt.Printf("%#v\n",myHobby)
```

>结果

```sh
[]string{"read", "study", "game"}
```

#### 切片的访问

```go
allCourses := []string{"javascript","html","css","python","golang"}

/* 
    切片初始化完成后使用下标进行访问，
    访问超出长度的元素会产生越界错误

    直接访问某个下标，返回值是具体元素
*/
fmt.Printf("%v\n", allCourses[3])

/* 
    取多个元素,用到的语法和用数组创建切片时的语法基本相似 

    切片[start:end]

    ① start开始位置，end结束位置，左闭右开
    ② start省略，相当于开始到end
    ③	end省略，相当于从start到结束
    ④	都省略，相当于取全部
*/
courses1 := allCourses[1:2]
courses2 := allCourses[:3]
courses3 := allCourses[1:]
courses4 := allCourses[:]
fmt.Printf("%#v\n %#v\n %#v\n %#v\n",courses1,courses2,courses3,courses4)
println("-------------------------")

/* 
    courses1容量为4，修改值不触发容量变化
    底层被修改，故courses1和allCourses都被修改
*/
courses1[0] = "Java"
fmt.Printf("%#v\n %#v\n",courses1,allCourses)
println("-------------------------")
/* 
    courses1容量为5，元素长度为3,增加一个不触发容量变化
    返回的地址不变
    底层被修改，故courses2和allCourses都被修改
*/
courses2 = append(courses2, "PHP")
fmt.Printf("%#v\n %#v\n",courses2,allCourses)
```

>结果

```sh
python
[]string{"html"}
[]string{"javascript", "html", "css"}
[]string{"html", "css", "python", "golang"}
[]string{"javascript", "html", "css", "python", "golang"}
-------------------------
[]string{"Java"}
[]string{"javascript", "Java", "css", "python", "golang"}
-------------------------
[]string{"javascript", "Java", "css", "PHP"}
[]string{"javascript", "Java", "css", "PHP", "golang"}
```

#### 切片数据的添加

切片数据的添加主要使用`append`方法

执行添加操作时,新增的元素必须和切片要求的元素类型一致

##### 基本语法

```sh
append(切片A,元素B)
#将元素B添加到切片A的末尾,返回值是一个切片

根据打印结果可以得知:
这个操作不改变原始切片,返回的是一个全新的切片 

我们可以采取把全新切片赋值给原切片的方式实现对原切片的添加
```

>举例

```go
allcourse := []string{"python","java","PHP","RUST"}
newCourse := append(allcourse,"Ruby")
fmt.Printf("新切片%#v\n 原始切片%#v\n",newCourse,allcourse)
```

>结果

```sh
新切片[]string{"python", "java", "PHP", "RUST", "Ruby"}
原始切片[]string{"python", "java", "PHP", "RUST"}
```

##### 增加多个元素

>语法

```sh
append(切片A,元素B,元素C,...)
#可以一次性添加多个元素,
```

>举例

```go
var sliceA []string
sliceA = append(sliceA,"vue","react")
fmt.Printf("一次性添加多个值%#v\n",sliceA)
```

>结果

```sh
一次性添加多个值[]string{"vue", "react"}
```

##### 扩展运算符

>语法

```sh
append(切片A,切片B...)
#将切片A,切片B进行合并
#切片B被展开,使用其内部元素
在append方法中使用 切片... 相当于js中的拓展运算符,展开其所有子元素
```

>举例

```go
sliceB := []string{"go","grpc"}
sliceC := []string{"gin","beego"}
sliceD := []string{"mysql","redis","es"}

sliceE := append(sliceB,sliceC...)
// 只想把silceD的一部分放进去,可以使用 silceD[x:y]对silceD进行截取
sliceF := append(sliceB,sliceD[0:1]...)

// 只能扩展一组
// sliceE := append(sliceB,sliceC...,sliceD...)

// 扩展不能结合其他插入一起用
// sliceE := append(sliceB,sliceC...,"abc")
// sliceE := append(sliceB,"grom",sliceD...)

fmt.Printf("%#v\n %#v\n",sliceE,sliceF)
```

>结果

```sh
[]string{"go", "grpc", "gin", "beego"}
[]string{"go", "grpc", "mysql"}
```

#### 切片元素的删除

Golang没有提供删除切片元素的api,如果要删除元素的通用思路是

```sh
从切片中提取出要保留的元素,拼接成新切片
```

```go
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
```

#### 切片元素的拷贝

```go
package main

import "fmt"

func main()  {
	// 复制slice


	/* 
		方案1:
		直接赋值
	*/
	var sliceA = []string{"A","B"}
	sliceB := sliceA

	fmt.Printf("%#v\n",sliceB)
	//[]string{"A", "B"}





	/* 
		方案2:
		截取全部 
	*/
	var sliceC =[]string{"C","D"}
	sliceD := sliceC[:]
	fmt.Printf("%#v\n",sliceD)


	println("------------------------------------------------")
	/* 
		方案3:
		copy(①，②)
		① 拷贝到的切片 
		② 被拷贝的切片 
		把切片②拷贝给切片①,覆盖对应位置的元素
	*/
	var sliceE = []string{"E","F"}

	var sliceF []string
	copy(sliceF,sliceE)
	fmt.Printf("%#v\n",sliceF)
	//[]string(nil)
	// 原因,拷贝并不会自动扩容，sliceF没有空间，自然无法接收拷贝的元素


	var sliceG = make([]string,5)
	copy(sliceG,sliceE)
	fmt.Printf("%#v\n",sliceG)
	//[]string{"E", "F", "", "", ""}


	var sliceH = []string{"GO","GOD","GOOD"}
	copy(sliceH,sliceE)
	fmt.Printf("%#v\n",sliceH)
	// []string{"E", "F", "GOOD"}
}
```

### map

 map是一个key(索引)和value(值)的无序集合 ,主要查询方便，查询性能比切片高 ,时间复杂度为o(1)

#### map的定义

>语法

```sh
var ① map[②]③
①map的变量名 
②key的类型,
key的类型不是任意的，例如切片就不行，因为不定，但是数组是可以的
③value的类型
可以是任何类型
```

>举例

```go
var course map[string]string

fmt.Printf("%#v",course)
```

>结果

```sh
map[string]string(nil)
```

>注意

```go
// 报错，map的nil不能赋值(切片的nil是可以的)
//即map类型想要设值，必须要先进行初始化
// course["level"] = "hight"
```

#### map的初始化

##### {}赋值

>语法

```sh
map的初始化方式1:
var 变量名 map[keyType]valueType = map[keyType]valueType{ key1:value1,...}
简化 
var 变量名 = map[keyType]valueType{ key1:value1,...}
```

>初始化为空

```go
/* 
    初始化也可以不放值
*/
var animal = map[string]string{}
fmt.Printf("%#v\n",animal)
//map[string]string{}
```

>最后一对一定要以逗号结尾

```go
var course = map[string]string{
    // 最后一对值一定要以逗号结尾
    "name":"zs",
    "gender":"man",
}
fmt.Printf("%#v\n",course)
// map[string]string{"gender":"man", "name":"zs"}
```

##### make初始化

>语法

```sh
make初始化 
make(类型,长度)
map也有长度,底层的结构需要,不过我们用不到，会自动扩展
```

>举例

```go
books := make(map[string]string, 3)
```

#### map的值操作

>语法

```sh
map必须初始化后才能赋值，赋值方式： 
①写在大括号里 {key:value}
②变量名[key] = value 
```

>举例

```go
/* 
  取值:
        变量名[key]
    赋值和修改： 
        变量名[key] = value

    给不存在的key赋值就相当于增加对应的key，扩展方便
*/
course["level"] = "hight" 
fmt.Printf("%#v\n",course)
// map[string]string{"gender":"man", "level":"hight", "name":"zs"}
```

#### map的遍历

```sh
map是键值对的形式，所以只要循环他的key就可以实现遍历
```

>举例

```go
var course = map[string]string{
    "name":"zs",
    "gender":"man",
}

for key,value := range course{
    fmt.Printf("%#v---%#v\n",key,value)
}
// "gender"---"man"
// "name"---"zs"


// 如果只有一个，则是key
for val := range course{
    fmt.Printf("键为%#v\n",val)
}
// 键为"name"
// 键为"gender"


/* 
    注意：map是无序的，而且每次执行可能都是不相同的顺序
*/
```

>注意

```sh
map是无序的，而且每次执行可能都是不相同的顺序
```

#### map元素的获取和删除

##### 元素是否存在

>背景

```go
var course = map[string]string{
    "name":"zs",
    "gender":"man",
}

/* 
    我们不能使用 变量名[key]来判断map是否有对应的键 
*/
fmt.Printf("%#v\n",course["java"])
// ""
//如果key不存在，会返回对应的零值，但是如果有一个key的值就是零值就无法判断，所以不能
```

>判断key是否存在

```sh
实际上 变量名[key] 是有两个返回值的， 
①,② := 变量名[key]
① key对应的值 
② 布尔值,true代表存在，false代表不存在
```

> 举例

```go
d,ok := course["sex"]
if ok {
    println("存在",d)
}else{
    println("不存在",d)
}
```

##### 元素的删除

>语法

```sh
delete(map变量,key)

删除不存在的元素不会报错
```

> 举例

```go
delete(course,"age")
```

#### map的注意事项

```sh
很重要的提示,
map不是线程安全的，当你有多个线程同时对一个map进行操作，会报错

并发时使用的map必须使用sync包下的Map，即syncMap
```

### list

#### list与slice的比较

list又叫双向链表

```sh
①slice每次扩容都会先对之前的进行拷贝，再次分配空间，频繁扩容，就要频繁的开辟空间 
②slice要求底层必须是连续的存储空间，如果底层没有这样的条件就会分配对应大小的slice
```

```sh
list的每个元素除了存储值外，还存储了指向下一个元素的指针 
list每次增加一个元素就会申请一块空间，所以底层可以不连续

优势：
①空间可以不连续，所以分配能力比slice强
②插入方便，删除很快
劣势：
①多存一个元素的指针，造成空间浪费
②性能比slice低很多，查询很慢
```

#### 使用举例

##### list的定义

```go
package main

import (
	"container/list"
	"fmt"
)


func main()  {
	// list不是个golang全局内置关键字，需要引入包 
	// container/list引入后就可以使用list了 

	// 定义
	/* 
		方式①：var 变量名 list.List
		方式②: 变量名 := list.New()
	*/
	var myList list.List
	// myList := list.New()
}
```

##### list的一些使用

```go
// list方法非常多 
/* 
    PushBack 尾部放数据 
    PushFront 头部放数据 
*/
myList.PushBack("golang") //在最后放一个 go 
myList.PushBack("mySql")

fmt.Println(myList)//直接打印不到值，是地址 
// {{0xc0000a0180 0xc0000a01b0 <nil> <nil>} 2}

/* 
    遍历打印值
    用法需要记
*/
// 正向遍历
for i:=myList.Front();i!=nil;i=i.Next(){
    fmt.Println(i.Value)
}
// golang
// mySql


//倒序遍历 
for i:=myList.Back();i!=nil;i=i.Prev(){
    fmt.Println(i.Value)
}
// mySql
// golang

/* 
    list的很多操作需要用到element元素，一种他自己返回的元素
*/
// 得到一个 list的element元素 
ele:=myList.Front()
// 通过循环遍历拿到golang对应的element元素
for ;ele!=nil;ele=ele.Next(){
    if ele.Next().Value.(string)=="golang"{
        break
    }
}
// 在golang前面插入 gin
myList.InsertBefore("gin",ele)

//更多用法需要的时候看源码
```

>总结

```sh
list的API很多，但是性能很低，所以用的少，需要的时候查一下就可以
```

### 集合类型小结

集合类型四种

1.数组

```sh
不同长度的数组类型不一样
```

2.切片

```sh
动态数组，用起来很方便，而且性能高，尽量使用slice
```

3.map

```sh
和其余三种不一样，并发编程时存在一些问题，可以用包解决
```

4.list

```sh
操作方法多，和slice竞争，slice能完成的list也能，但是list性能低，用得少
```

## 函数

函数是golang中的一等公民，因为go语言采用函数式编程。

### 函数的定义

>基本语法

```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

函数定义解析：

- func：函数由 func 开始声明
- function_name：函数名称，参数列表和返回值类型构成了函数签名。
- parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
- return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
- 函数体：函数定义的代码集合

### 函数的闭包

#### 什么是闭包

闭包就是指某个函数有权访问另一个函数内部的数据或变量。

经常通过函数嵌套的方式实现闭包，因为内嵌函数能够访问祖先的数据和变量。

#### 为什么需要闭包

>需求

```sh
有个函数每次调用一次返回的结果值都是增加一次之后的值
```

>传统解决方案

```go
package main

import "fmt"


var local int

func autoIncrement() int {
	local += 1
	return local
}


func main()  {
	
	for i := 0; i < 5; i++ {
		fmt.Println(autoIncrement())
	}
	
}
```

>痛点

```sh
①这样虽然实现了需求，但是污染了外部变量
②且由于变量在外部，所以会被其他方式修改，不安全
③想再从0开始，只能重置外部的local，就会影响到其他在使用这个变量的
```

#### 闭包的解决

>举例

```go
package main

import "fmt"

/*
	将函数需要的变量在内部声明
	返回值是一个函数的调用结果，
	用这个被调用的函数来操作变量
*/
func autoIncrement() int{
	local := 0
	
	//由于内部的匿名函数访问到了其他函数中的局部变量，所以称为闭包
	// 这个函数调用，获得对应类型的返回值
	return func() int{
		local += 1
		return local
	}()
}

/* 
	和autoIncrement的区别，前者每次调用都是整体调用，会执行autoIncrement 
	就会使得local重置为0
	而autoIncrement01是将内部闭包函数抛出,只执行闭包函数，所以不会重置local

	想要重置local，只需要再调用一次autoIncrement01即可
*/
func autoIncrement01() func() int{
	local := 0
	
	//由于内部的匿名函数访问到了其他函数中的局部变量，所以称为闭包
	// 这个函数调用，获得对应类型的返回值
	return func() int{
		local += 1
		return local
	}
}
func main()  {
	/* 
		autoIncrement的使用，返回值是个int
		每次调用执行，local重置，计算完后抛出值，所以一直为1
	*/
	for i := 0; i < 3; i++ {
		fmt.Printf("我是autoIncrement的值%#v\n", autoIncrement())
	}

	/* 
		autoIncrement01的返回值是个函数
	*/
	//赋值的时候，调用autoIncrement01，loacl初始化为0
	next := autoIncrement01()
	for i := 0; i < 3; i++ {
		fmt.Printf("我是autoIncrement01的值%#v\n", next())
	}

	//想要重置loacl，只需要再次调用autoIncrement01
	next = autoIncrement01()
	fmt.Printf("我是local重置后再开始的值%#v\n", next())
}
```

>执行结果

```sh
我是autoIncrement的值1
我是autoIncrement的值1
我是autoIncrement的值1
------------------------------------
我是autoIncrement01的值1
我是autoIncrement01的值2
我是autoIncrement01的值3
------------------------------------
我是local重置后再开始的值1
```

>注意

```sh
分清楚暴露的是函数还是函数调用的结果。
```

#### 闭包的优势

>防止了变量的污染

定义在函数内部，生命周期和函数一起。

>可以随时重置

可以对变量进行重置，不影响到其他操作，在协程时不产生bug

### 函数中的defer

相当于java和python中的finally

>常见场景

```sh
连接数据库、打开文件、开始锁等这些操作，
无论成功与否，最后都需要去关闭数据库、关闭文件、解锁， 
```

>结合java理解

```java
try {
    开始锁的逻辑
}finally{
    无论try中的语句结果是什么，这个最后都会在之后执行 
    在这里解锁
}
/*
    这种格式看似可读性高，但是有个问题，就是try里面的逻辑可以写的特别多，最后我们忘记了该在finally中写什么 
    隔得太远了，实际可读性反而变低了
*/
```

>go的设计理念

```sh
于是golang在设计的时候，就将finally的操作写在defer后面，defer后面的语句在函数return之前执行，相当于就是位置移动
但是执行时机不变
```

>举例

```go
package main
import "sync"

func main()  {
	
	var mu sync.Mutex

	// 开始锁 
	mu.Lock()
	

	// 执行时机实际上在return之前
	//写在这是为了提高可读性和防止忘记关闭
	defer mu.Unlock()

	// 锁的逻辑 
	/* 
		....
	*/

	// defer实际在这执行，之后再return
	return 
}
```

#### defer的执行时机

>书写

```go
func main(){
    逻辑①
    defer 逻辑②
    
    逻辑③
    逻辑④
    return
}
```

>执行

```sh
逻辑①
逻辑③
逻辑④
逻辑②
```

>小结

```sh
defer书写在前面只是为了防止忘记关闭某些操作，实际执行时机在return前
```

#### 多个defer

>举例

```go
package main

import "fmt"

func main()  {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")

	return
}
```

>结果

```sh
4
3
2
1
```

>原理

```sh
defer都是在函数return之前执行 
defer是一个栈的概念，先写的先进栈，所以后写的在上面 
出栈时由于从上到下出栈，所以后写的先执行
```

#### defer的使用

defer由于他在ruturn之前执行，所以有能力修改返回值

```go
package main

import "fmt"

/*
	defer的执行时机在return之前，所以他有能力改变返回值
*/
func deferReturn() (ret int) {
	ret = 7

	defer func ()  {
		ret ++
	}()

	ret = 10

	// 在return前执行defer后的逻辑，所以 
	//	ret++ ，也就是11
	return
}


func main()  {
	res:=deferReturn()
	fmt.Printf("最终返回值为%d", res)

}
```

>结果

```sh
最终返回值为11
```

### 错误处理

#### error的设计理念

go语言的错误处理饱受诟病，他的错误处理包含三个

```sh
error、panic、recover
```

实际上是由于golang错误处理的理念不一样

```sh
一个函数可能出错，其他语言是用try...catch包住函数，防止整个代码崩溃
golang则抛弃了这种思路，而是认为，如果出错了，就抛出一个错误对象，通过判断error是否为nil，来知道函数是否出错了

开发函数的人需要有一个返回值去告诉调用者是否成功，返回值的处理更灵活，因为你可以决定是否使用或者接收他，也可以决定是否对外继续暴露
不过按照开发规范，golang设计者要求我们必须要处理这个error,所以代码中会大量出现 if err !=nil，这种就叫做防御性编程，程序健壮性好
```

>举例

```go
package main

import (
	"errors"
	"fmt"
)

func fn01() (int,error) {
	// 我们可以自定意任何错误为error
	//但是在golang中实际上有个专门的包生成error对象 
	/* 
		errors.New(字符串)
		这个包还有好几种用法，来返回不同的error
	*/
	return 0,errors.New("这是一个错误")
}

func main()  {

	if _,err := fn01();err != nil{
		fmt.Printf("出错原因是:%#v", err)
	}
}
```

>结果

```sh
出错原因是:&errors.errorString{s:"这是一个错误"}
```

error有很多实际的用法，随着后续知识的学习会了解到

#### panic

```sh
panic 这个函数会导致程序退出
相当于javascript中的throw
golang中不推荐随便使用panic
```

>举例

```go
func testA(){
	a := 10

	// 从这抛出异常，结束程序，后面的都不执行
	panic("这是一个错误")

	fmt.Println(a)
}
```

##### 使用场景

```sh
一般用到panic的场景： 
服务启动的过程中，
要启动服务必须某些依赖服务要存在，例如日志文件、mysql等等
如果任一依赖服务不满足，就主动调用panic

但是你的服务一旦启动，如果某行代码中不小心使用panic，可能会导致该核心程序挂掉，整个服务崩溃 --重大事故
所以不推荐
```

#### recover

>背景

```go
/* 
	但是有些时候，我们的代码也会被动触发panic
	例如 我们使用map之前没有初始化
*/
func testB(){
	var name map[string]string
	// map没进行初始化就被赋值了 
	//会触发panic
	name["go"] = "go工程师"
}
```

>用法

```go
/* 
	recover 用于捕捉panic，我们可以在捕捉到后做任何操作 
	常见的是捕捉到后进行日志打印 
	语法： r:= recover()
	如果有panic被捕获，recover()就会返回一个错误对象 
	处理时机 一般配合defer，因为defer的执行时机在return前
*/
func testC()  {
	defer func ()  {
		if r:=recover();r!=nil{
			fmt.Printf("捕获到的错误为%#v",r)
		}
	}()

	var name map[string]string
	name["go"] = "go工程师"
}
func main()  {
	testC()
	// 捕获到的错误为"assignment to entry in nil map
}
```

#### 注意事项

```sh
①尽量使用error对象来处理错误
②defer需要放在panic之前定义，放在panic后就不会执行了，所以写在最开始
③recover只有在defer调用的函数中才会生效
④recover处理异常后，逻辑并不会恢复到panic的那个点继续执行
⑤多个defer会形成栈，后定义的defer会先执行
```

## 结构体struct

  golang的结构体和C类似，是其他语言中轻量级的class

### type关键字的用法

>常见用法总结

```sh
1.定义结构体
2.定义接口
3.共享底层，类型定义
4.定义类型别名
5.类型判断
```

#### 定义类型别名

>语法

```go
type 类型 = 类型别名
```

>举例

```go
/* 

    别名实际上是为了更好的理解代码

    底层类型是一样的，只是取了一个别名，
    通过打印可知，根本类型没变，赋值也可以直接使用

    用途：
    增加代码可读性
*/
type MyInt = int
var i MyInt
fmt.Printf("%T\r\n", i)//int
var a  int = 8
i = a
fmt.Printf("%v\r\n",i)//8
```

>结果

```sh
int
8
```

>原理

```sh
在编译的时候，类型别名会直接替换为底层类型 
即编译时，MyInt会直接变成int
```

#### 自定义类型

>语法

```go
type 类型A 类型B
```

>举例

```go
/* 
    少了等号，这个实际上是定义了一个新类型，
    基于已有的类型B定义一个新的类型A
    和底层类型B可以接收完全一样的值，但是二者属于不同的类型 
    变量之间需要类型转换

    用途： 
    既希望使用原类型的特性，又需要在其上扩展方法特性
    内置类型没办法扩展，但是自定义类型可以使用接收者函数扩展，结构体可以理解为一种特殊的自定义类型
*/
type newInt int 
var b newInt
fmt.Printf("%T\r\n", b)
//main.newInt  包名.类型
var c int = 9
// 类型转换
b = newInt(c)
fmt.Printf("%v\r\n",b)
```

>结果

```sh
# 包名.类型名
main.newInt 
9
```

#### 类型断言

>语法

```go
对于interface{}类型的变量，使用
变量.(type) 可以拿到他的实际类型,配合switch使用

变量.(类型) 进行类型断言，直接强转
```

>举例

```go
var f interface{} = "abc"
switch f.(type) {
    case string:
        println("是字符串类型")
    default:
        println("不是字符串")
}

// 类型断言，interface{}不能使用其他类型的方法 
//经过类型断言为string，只要成立就可以使用string的方法了 
m := f.(string)
fmt.Println("m的长度为",len(m))
```

>结果

```sh
是字符串类型
m的长度为 3
```

其他用法等结构体和接口时讲解

#### 约定

```sh
type 后接的我们自定义的类型采用大写字母开头，驼峰命名

小写也是可以的，但是约定采用大写
```

### 为什么需要结构体

```go
package main

func main()  {
	/* 
		为什么需要结构体，
		例如我们要存储人的信息，包含姓名，年龄，身高 
		之前我们只能采取多维数组，类型单一，使用时要进行类型转换 
	*/
	var persons [][]string
	persons = append(persons, []string{"zs","18","180"})

	/* 
		类型麻烦，可以使用interface{}，存储的时候可以任意类型，但是使用的时候需要类型断言 
	*/
	var persons01 [][]interface{}
	persons01 = append(persons01, []interface{}{"lisi",18,1.8})
}
```

>解决痛点

```sh
①定义一个类型的集合，解决类型的痛点
②提高可读性和可操作性
```

>区别

```sh
golang没有面向对象的一些class的概念，但是代码的封装，面向对象的思想是有的
面向对象的三大特性：封装、继承和多态(接口)
```

>提示

```go
空结构体不占空间
```

### 结构体的定义

>语法

```sh
#核心：结构体本质上也是一种自定义类型
#注意：每行对应一个字段，不需要逗号结尾

type 类型名 struct {
    字段A 类型
    字段B 类型
    ...
}
```

>举例

```go
type Persons struct{
	name string 
	age  int 
	height float32
}
```

### 结构体的初始化

#### 方式①

> 语法

```sh
结构体名{值A,值B,...}
#注意：顺序和个数要与定义时一一对应
```

>举例

```go
p1 := Persons{"zs",18,1.8}
fmt.Printf("%#v\n",p1)
```

>结果

```sh
main.Persons{name:"zs", age:18, height:1.8}
```

#### 方式②

>语法

```sh
结构体名{
        字段A:值A,
        字段B:值B,
    }
    
"注意" 
最后一个必须逗号结尾，字段需要在结构体中 
没有顺序，没赋值的会采用类型零值
```

>举例

```go
p2 := Persons{
    name:"lisi",
}
fmt.Printf("%#v\n",p2)
```

>结果

```sh
main.Persons{name:"lisi", age:0, height:0}
```

> 特殊

初始化为空

```go
p9 := Persons{}
```



#### 方式③

```sh
切片形式，一次性初始化多个
```

>举例

```go
p2 := Persons{
    name:"lisi",
}

var p3 []Persons
// 因为分离了，所以需要指明元素是Persons的实例
p3 = append(p3,Persons{"zs",18,1.8})
p3 = append(p3,p2)
fmt.Printf("%#v\n",p3)
```

>结果

```sh
[]main.Persons{
    main.Persons{name:"zs", age:18, height:1.8}, 
    main.Persons{name:"lisi", age:0, height:0}
}
```

#### 方式④

方式③的优化，经常在实际中使用

>注意

```sh
注意：
1.这是省略了Persons
2.最后一项一定要以逗号结尾
```

>举例

```go
var p4 = []Persons{
    {name:"lisi"},
    {name:"wangwu",age:18},
    // 可以嵌套着混合写，他会自动识别
    {"zs",18,1.8},
}
fmt.Printf("%#v\n",p4)
```

>结果

```sh
[]main.Persons{
    main.Persons{name:"lisi", age:0, height:0}, 
    main.Persons{name:"wangwu", age:18, height:0},
    main.Persons{name:"zs", age:18, height:1.8}
}
```

### 结构体的赋值与取值

先准备一个用于演示的结构体

```go
type Persons struct{
    name string
    age  int
    height float32
}
```

#### 赋值

>语法

```sh
结构体实例.字段 = 值

这也是一种初始化的方式,其他字段会采用零值 
```

>举例

```go
p.age = 18
fmt.Printf("%#v\n", p)
```

>结果

```sh
main.Persons{name:"", age:18, height:0}
```

#### 取值

>语法

```sh
结构体实例.字段 
```

>举例

```go
fmt.Println(p.height)
// 无法取不存在的字段
// fmt.Println(p.abc)
```

>结果

```sh
0
```

### 匿名结构体 

#### 结构体定义的同时进行初始化

>语法

```sh
①
变量名 := struct{
    字段A 类型A 
    字段B 类型B
}{
    值A,
    值B
}

②
var 变量名 = struct{
    字段A 类型A 
    字段B 类型B
}{
    值A,
    值B
}
```

这种实际是在定义结构体时同时赋值

```sh
第一个大括号是定义，第二个大括号是赋值 
然后可以通过 结构体名.字段 来赋值或修改 

定义的语法和赋值的语法是和之前一样的，不过是把之前的分离写法，写到了一起
```

>举例

```go
address := struct {
    province string
    city string
    age int
}{
    age:200,
}
fmt.Printf("%#v\r\n",address)
```

>结果

```sh
struct { province string; city string; age int }{province:"", city:"", age:200}
```

#### 使用匿名结构体

有了上面的引子我们可以知道，匿名结构体只要去掉左边的变量名即可，那样就变成一次性的了

>举例

```go
fmt.Printf("%#v\r\n",struct {
    province string
    city string
    age int
}{
    city:"Rome",
})
```

>结果

```sh
struct { province string; city string; age int }{province:"", city:"Rome", age:0}
```

#### 小结

也有人将第一种称为广义的匿名结构体，主要是使用的使用定义，无需全局注册。结构体是验证类型，快速赋值的好东西。

### 嵌套结构体

为了方便演示，我们先定义两个需要嵌套的结构体

```go
type Persons struct{
	name string
	age  int
	height float32
}

type Persons01 struct{
	name string
	age  int
	height float32
}
```

#### 具名嵌套

```sh
给定一个字段，让该字段是要嵌套的结构体类型
取值赋值都是嵌套的形式,通过该字段去实现,这个字段就相当于桥梁

"优势"
避免了结构体的字段冲突 
	
"劣势" 
书写繁琐
```

>举例

```go
type Student struct{
	p1 Persons01
	score float32
	p Persons
}
s := Student{
    p: Persons{
        name:"张三",
    },
    score: 95.6,
    p1:Persons{}
}

fmt.Printf("学生姓名为%#v\r\n", s.p.name)
```

>结果

```sh
学生姓名为"张三"
```

#### 匿名嵌套

##### 基本情形

```sh
直接给要嵌套的结构体名
"优势"
书写方便，访问直接 
不需要再通过中间桥梁 
```

>举例

```go
type Teacher struct{
    Persons
	salary int
} 

t := Teacher{}
t.name = "王二"
t.salary = 10000
fmt.Printf("老师信息为%#v\r\n", t)
```

>结果

```sh
老师信息为main.Teacher{Persons:main.Persons{name:"王二", age:0, height:0}, salary:10000}
```

##### 字段冲突

```sh
当我们要使用某个字段，这个字段不唯一时会报错
```

>举例

```go
type Teacher struct{
    Persons
	salary int
    Persons01
} 

t := Teacher{}
//无法确定取的是哪个name,故报错
t.name = "王二"
t.salary = 10000
fmt.Printf("老师信息为%#v\r\n", t)
```

>结果

```sh
报错，因为name字段的归属不确定
```

##### 字段优先级

```sh
本地的字段优先级最高，可以避免嵌套的字段冲突
```

>举例

```go
type Teacher struct{
    Persons
	salary int
    name string
    Persons01
} 

t := Teacher{}
t.name = "王二"
t.salary = 10000
fmt.Printf("老师信息为%#v\r\n", t)
```

>结果

```sh
老师信息为main.Teacher{
    Persons:main.Persons{name:"", age:0, height:0}, 
    Persons01:main.Persons01{name:"", age:0, height:0}, 
    salary:10000,
    name:"王二"
}
```

##### 原理

以上只是匿名函数的语法糖，初始化时就可以看到底层实现

```go
type Persons struct{
	name string
	age  int
	height float32
}

type Teacher struct{
    Persons
	salary int
} 
```

想要实现

```go
t := Teacher{}
t.name = "王二"
t.salary = 10000
```

底层实现

```go
t := Teacher{
    Persons{
        name:"王二",
    },
    salary:10000
}
```

这就解释了字段的冲突，因为我们没有显式指出该给谁

我们也可以写这种进行初始化

#### 一种不常用的嵌套

```go
type Teacher struct{
	salary int
    adress struct{
        city string
        age int
    }
} 
```

不过这样的嵌套可读性太差，所以采用提取出来

### 结构体绑定方法

```sh
参考下面的"接收者方法"
```

### 结构体的比较

#### 原则

```sh
①不同结构体的实例无法比较
②同一结构体的实例，按照字段进行比较，字段值都相同就相等
```

#### 示例

>准备好的结构体

```go
type Test01 struct{
	name string
	age int
}

type Test02 struct{
	name string
	age int
}
```

>不同结构体的实例无法比较

```go
p := Test01{}
p1 := Test02{}

// ①不同结构体的实例之间不能进行比较
if p == p1 {
	fmt.Println("yes")
}

//报错：无法比较
```

>同一结构体按照字段进行比较

```go
p := Test01{}
p2 := Test01{
    age:10,
}
if p == p2 {
    fmt.Println("yes")
}	else{
    fmt.Println("No")
}
//age的值不一样，p的age为零值0，name为零值空字符串
```



## 指针pointer

变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

一个指针变量指向了一个值的内存地址。

### 指针的底层

#### 变量是什么

变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

>语法

```go
//获取变量在内存中的地址
&变量名
```

#### 指针是什么

>语法

```go
//声明变量
var 变量 类型

//声明变量对应的指针
var 变量 *类型
```

>理解

```sh
声明变量，即开辟了一块空间，存储的内容是变量的值

声明变量对应的指针，即申请了一块空间，存储的是变量在内存中的地址

指针类型的变量也是变量，他存储的是其他变量的地址，该地址指向目标变量的值空间
```

> 举例

```go
var a int= 100

//存储变量a 的地址
//b 是个指针类型的变量，他存储的是一个地址，这个地址指向变量a的值所在的空间
var b *int = &a

fmt.Printf("%#v\n", a)
fmt.Printf("%#v\n", b)
```

>执行结果

```sh
100
(*int)(0xc00001a130)
```

### 指针的定义和初始化

#### 指针的定义

```go
/* 
    指针的定义 

    var 变量 *类型 
*/
var a *string 
```

##### 空指针

```go
/* 
    指针需要初始化，不初始化是nil
    即空指针
*/
var b *int
fmt.Printf("%#v\n",b)
//空指针的错误常见于结构体 
type Persons struct{
    name string
}

var p *Persons
// 报错，nil上面没有这个
//因为指针不初始化是nil，和map类似
p.name = "张三"
println(p)
```

>打印结果

```sh
(*int)(nil) # fmt.Printf("%#v\n",b)
0x0  #println(p)
```

>小结

```sh
指针定义后不初始化是nil，和map一样
```

#### 基本类型的初始化

>注意

```sh
①我们不能在基本类型上面取地址

a = &"你好"

这样是不被允许的

②但是可以在变量上面取地址 
var d = "王五"
a = &d 
```

>举例

```go
var a *string
var d = "你好"
a = &d
fmt.Printf("%#v\n", a)
```

>结果

```sh
(*string)(0xc000014270)
```

#### 复杂类型的初始化

>复杂类型可以直接在值上取地址

```go
var f1 *map[string]string
var f2 *Persons
var f3 *[]string
f1 = &map[string]string{}

f2 = &Persons{}

f3 = &[]string{}
fmt.Printf("%#v\r\n %#v\r\n %#v\r\n",f1,f2,f3)
```

>结果

```sh
&map[string]string{}
&main.Persons{name:""}
&[]string{}
```

#### new函数

>语法

```go
new(Type)
new函数，参数传入一个类型，会开辟一块空间，然后返回这个类型对应的指针，
这样就同时实现了指针类型的定义和初始化 
```

>举例

```go
var kk = new(int)
fmt.Printf("%#v\r\n",kk)
```

>结果

```sh
(*int)(0xc00001a150)
```

#### 总结

```sh
初始化关键字有两个 
①
map、channel、slice定义兼初始化推荐使用make函数
指针定义兼初始化推荐使用 new函数

②
指针需要初始化，否则会出现nil空指针的问题
map也需要初始化，否则会报nil问题
```

### 理解指针的传递

>需求

```sh
定义一个函数来交换变量的值
```

#### 错误示范

```go
func swap(a,b *int){
    a,b = b,a
}
func main(){
    a := 7
    b := 5
    swap(&a,&b)
    fmt.Printf("%#v\n %#v\n",a,b)
}
```

>结果

```sh
7
5
```

>理解

从最终结果来看，没有发生交换，原因

```sh
本质上来说，都是值传递，我们传入指针变量，就是把指针变量拷贝一份进行传递
指针变量在内部进行交换，无法影响到外部

因为外部指针变量没变
```

#### 正确做法

```go
func swap(a,b *int){
    *a,*b = *b,*a
}
func main(){
    a := 7
    b := 5
    swap(&a,&b)
    fmt.Printf("%#v\n %#v\n",a,b)
}
```

>结果

```sh
5
7
```

>理解传递的本质

```sh
传参都是将变量拷贝一份进行传入，不同的是传入的指针变量，可以访问到外部变量

指针类型的变量是一座桥梁，所以为了区分他与普通传递的不同，把他叫做引用传递，但本质上还是值传递，传递指针类型变量的值
```

### go中指针的特性

#### 值和指针的转换

>地址与值的转换

```go
var a int = 10

var b *int = &a 

/*
	指针类型的变量B存储的是变量A的地址     &变量A
	想要通过变量B获取变量A的值  		  *变量B
*/
```

>语法

```sh
# 获取变量的地址
&变量名

# 通过指针变量获取值
*指针变量
```

>举例

```go
c := &map[string]string{
    "name":"张三",
}
(*c)["name"]= "李四"
fmt.Printf("%#v\n",c)
fmt.Printf("%#v\n",*c)
```

>执行结果

```sh
&map[string]string{"name":"李四"}
map[string]string{"name":"李四"}
```

#### 不安全的计算

golang的指针是不能和`C/C++`一样直接参与计算的，但是也可以使用`unsafe`包实现

### 指针对结构体的优化

其他数据类型指针变量要修改指向的值，都需要先拿到值，即通过`*指针变量`的方式，但是对于结构体，做了隐式转换

```go
/* 
    结构体实例的指针变量和结构体变量之间可以互换
*/
type Test struct{
    name string
}

//这是右边结构体的指针变量
p1 := &Test{
    name:"王五",
}

//常规的值修改 
(*p1).name = "李四"
fmt.Printf("%#v\n",p1)

//结构体的特例，结构体的指针变量可以直接使用和值的用法一样 
//精简结构体的使用
p1.name = "二虎"
fmt.Printf("%#v\n",p1)
```

>执行结果

```sh
&main.Test{name:"李四"}
&main.Test{name:"二虎"}
```

>注意

```sh
结构体的值类型和引用类型用法一样只适用于以下几种情形。
① 接收者方法传参
② 通过点语法修改结构体的值

不适用的情况：
①函数传参 
函数传参要求是什么类型就必须是什么类型，但是传递后通过点语法修改结构体时，不需要显示转换，因为触发了上面的情形
```

>需求

在结构体传值的时候，在内部的修改能反馈到外部变量中

#### 结构体与方法的实现

##### 值方法

```go
type Persons struct{
	name string
	age int
}

func (c Persons) AddAge() int{
	c.age ++
	return c.age
}

func main()  {
	p := Persons{}

	p.age = 20

	fmt.Printf("调用方法得到的年龄%v\n",p.AddAge())
	fmt.Printf("age变量对应的年龄%v\n",p.age)
}
```

>执行结果

```sh
调用方法得到的年龄21
age变量对应的年龄20
```

>讲解

```go
/*
	接收的是值类型 也就是值传递
	所谓的值传递，就是变量将自己复制一份，然后将复制的传入
*/

func (c Persons) AddAge() int{
    //修改的是变量复制的内容
    //如果复制后地址发生了变化，那么就无法修改到变量
	c.age ++
	return c.age
}
```

##### 指针方法

```go
type Persons struct{
	name string
	age int
}


func (c *Persons) changeName() string{
	c.name ="李四"
	return c.name
}

func main()  {
	p := Persons{}
	p.name = "张三"
    
	fmt.Printf("调用方法得到的名字%v\n",p.changeName())
	fmt.Printf("name变量对应的名字%v\n",p.name)
}
```

>执行结果

```sh
调用方法得到的名字李四
name变量对应的名字李四
```

>讲解

```go
/*
	接收的是指针类型，也就是说传参要传地址
	接收者方法对结构体做了优化，结构体的值和地址在传入时会根据需要自动转化
*/

func (c *Persons) changeName() string{
    //引用传递，传递的是地址
   	//对地址的修改，即修改了变量底层
	c.name ="李四"
	return c.name
}
```

#### 结构体与函数

```go
package main

import "fmt"

type Student struct{
	name string
	age int
}

// 使用函数
func changeName(c Student){
	c.name = "张三"
}

func changeAge(c *Student){
	// 结构体的指针和值再使用点语法时不需要转换
	c.age = 20
}

func main()  {
	s:=&Student{
		name:"王五",
		age:15,
	}

	// 函数传参时，参数要求是什么类型就必须什么类型
	changeName(*s)
	fmt.Printf("改名%#v\n",s)

	changeAge(s)
	fmt.Printf("改岁%#v\n",s)

}
```

>结果

```sh
改名&main.Student{name:"王五", age:15}
改岁&main.Student{name:"王五", age:20}
```



## 方法

>参考

```sh
https://www.jianshu.com/p/474314f956b5
```



>语法

```sh
func (接收器) 方法 {
	逻辑体
}
```

### 接收器

```go
两种形态：
"值传递形态"
(变量名① 类型②)
"引用传递形态"
(变量名① *类型②)


变量名①  方法调用者本身
类型②    方法调用者本身的类型

规范：变量名 一般用后面类型的首字母的小写来命名
```

>举例

```go
func (c *myInt) getName(){
    return c.name
}
/*
	go对接收器做了优化，无论传入的是值还是指针，用法都一样
	无需把指针转为值再进行使用
*/
```

>注意

```go
//不能给两种形态绑定同名方法
func (c *myInt) getName(){
    return c.name
}

func (c myInt) getName(){
    return c.lastName
}

//报错
```

### 方法

方法是我们对内置功能的扩展，帮助类型添加新的功能。

用法：

```sh
类型实例.方法名()
```

>使用语法

```go
func (接收器) 方法 {
	逻辑体
}
```

方法本身就是一个精简版的函数，所谓的接收者方法，就是在函数的前加了一个接收器,函数的入参变成了接收器传入的变量

>方法

```go
方法名(参数类型) 返回值类型{
    逻辑体
}
//入参在接收器
```

也就是调用方法的实例对象

>语法详述

```go
func (变量① 变量类型②) 方法名() 返回值类型{
    方法逻辑
}
```

```sh
和函数的区别，在函数前加了一组括号，表名接收这个函数的对象 

(变量名① 变量数据类型②) 叫做接收器,默认使用值传递，就把值复制一份传进来 
如果我们想对调用者进行修改，或者值太大了，可以考虑引用传递，就是把目标的指针传进来

变量名① 代表调用者，方法的调用者会将自身传入 
变量的数据类型② 将方法挂载到②上，成为类型②的一个内置方法  
```

>方法的调用

```go
type Student struct {
	name string
}

// 接收者函数不能作为嵌套函数
func (c Student) getName() string {
	return c.name
}


func main(){
    s:= Student{
		name:"张三",
	}

	fmt.Printf("%#v\r\n",s.getName())
}
```

使用`实例.方法名()`进行调用

由于他本质上是个函数，所以也可以作为一个变量使用

### 值传递与引用传递

#### 值传递

```sh
golang中无论是函数还是接收方法，默认使用值传递，就是在传参时先把变量拷贝一份，然后将拷贝的内容传入。

切片之所以会被误认为引用传递，是因为拷贝后切片地址不变，还能被改到底层。
而对于其他绝大多数，值传递进入函数或接收者方法的变量，是独立的。
对他们的修改不会影响到外部。
```

>举例

```go
type Student struct {
	name string
    age int
}
func (c Student) getAge() int {
    c.age += 1
	return c.age
}
func main(){
    s := Student{
        age:18
    }
    
    fmt.Prinf("%#v\n",s.getAge()) //19
    
    fmt.Print("%#v\n",s.age) //18
}
```

>结语

```sh
值传递传递的是原变量的拷贝，所以修改是分离的。
```

#### 引用传递

```sh
引用传递，传递的是变量的指针，所以可以通过他对外部变量进行修改

可以粗暴的认为把原变量传进来了，实际上是原变量的地址，但我们可以通过地址访问原变量
```

>举例

```go
type Student struct {
	name string
    age int
}
func (c *Student) getAge() int {
    c.age += 1
	return c.age
}

func main(){
    s := Student{
        age:18
    }
    
    fmt.Prinf("%#v\n",s.getAge()) //19
    
    fmt.Print("%#v\n",s.age) //19
}
```

#### 使用的感悟

```sh
接收者方法与普通函数的不同点。
① 无论实例是值还是指针，他们能调用的方法一样。
② 无论实例是值还是指针，调用方法实现的效果一样。

这是因为底层做了兼容处理，减少了指针与值转换的步骤

"区别"
① 实例是值传入的值类型，实例是指针，传入的是指针类型
② 传入的是指针，通过对实例参数的修改可以影响到外部，而值是一份拷贝，所以修改一般不会影响到外部，切片要单独考虑

"理解"
最终传入的是什么，实际上和调用者无关，看接收者方法的封装
```

>举例

```go
type Student struct {
	name string
}

//指针形态
func (c *Student) getAge() int {
    //接收到实例地址
    c.age += 1
	return c.age
}

//值形态
func (c Student) getName() string {
    //接收到实例的值
	return c.name
}

func main(){
    //值类型
    s:= Student{
		name:"张三",
	}

    //指针类型
    s1 := &Student{
        name:"李四",
    }
    
    //都能用，方法都正常执行
    s.getAge()
    s.getName()
    
    s1.getAge()
    s1.getName()
}
```

## go语言中的nil

### 类型的零值

```sh
不同类型的数据零值不一样 
bool      	false
numbers  	0 
string    	 ""
pointer    	nil
slice     	nil
map        	nil
channel    	nil
interface  	nil
function   	nil

struct 默认值不是nil，默认值是具体字段的默认值
```

>判断是否等于nil

```sh
判断一个值是否为 nil

变量名 == nil

等于就为true，反之false

注意： 
除了基本数据类型外都可以使用此方法判断
基本数据类型(bool\数字\字符串)不能和nil做判断
```

>举例

```go
var a map[string]string
if a==nil {
    println("yes")
}else{
    println("No")
}
```

### make制造的空与nil的区别

```sh
make初始化后不是nil，而是空的 

所以make和new函数都是定义加初始化 

例如我们会被切片分为 nil slice 和empty slice
map分为 nil map 和empty map
其他同理 
```

>举例

```go
var b = make([]string,0,0)
if b == nil {
    println("yes")
}else{
    println("No")
}
//No

//即空的不等于nil
```

#### 空与nil的异同

```sh
nil的无和empty的空的区别 
取值的时候都一样，nil的不能赋值
```

>举例

```go
c := make(map[string]string,0)
var d map[string]string
// 取值都一样 
println(c["name"])
println(d["name"])

//赋值有差别 
c["name"] = "李四" //make的进行了初始化，所以可以赋值
d["name"] = "王五" //nil不可以赋值，会panic
```

## 接口interface

### 什么是鸭子类型

>注意

```sh
go语言中处处都是interface，到处都是鸭子类型 duck typing
```

>俗语

```sh
当一只鸟走起来像鸭子，游泳起来像鸭子，叫起来也像鸭子，那么这只鸟就是鸭子
#以上都是动词，也就是方法 
```

>golang的鸭子

```sh
也就是说一个类型，只要他具备了方法集的所有方法，那么他就是这个方法集的实例，或者说他实现了这个方法集合 

golang中把这个方法集叫做interface接口 
通常会用结构体去实现它 
```

>小结

```sh
鸭子类型强调的是事物的外部行为，也就是对象暴露的方法 
而不关心内部的结构，例如对象的属性
```

### 如何定义接口

>语法

```go
type 接口名 interface{
    方法的声明
}

①只声明不写具体逻辑

②方法的声明格式
方法名(参数类型) 返回值类型
```

>举例

```go
type Cat interface{
	miaomiao()

	Walk()

	say() string
}
```

### 接口的实现

>举例

```go
/* 
	接口的实现，没有关键字

	只要某种类型实现了接口全部的方法，golang就会自动认定他实现了接口

	接口的实现没有强绑定关系，是自动判定是否实现了，即看起来像
*/
type HuaCat struct{
	name string
	age int
}

func (h HuaCat) miaomiao(){
	println("喵喵叫")
}

func (h HuaCat) Walk(){
	println("走猫步")
}

func (h HuaCat) say() string{
	return "猫言猫语"
}


type HeiCat struct {
	name string
}

func (h HeiCat) miaomiao(){
	println("喵喵叫")
}

func (h *HeiCat) Walk(){
	println("走猫步")
}

func (h HeiCat) say() string{
	return "猫言猫语"
}
```

>接口实现的使用

```go
func main()  {
	
	/* 
		HuaCat结构体由于具备了所有Cat接口的方法 
		所以可以用HuaCat类型的实例来实现接口 

		注意事项： 
		①如果接收器全是值类型,结构体实例可以写值形式  
		例如 HuaCat

		②如果接收器有一个指针类型，结构体就必须用引用形式
		例如HeiCat
	*/
	var d Cat = HuaCat{}
	d.Walk()


	// 由于有个接收器是指针类型，所以和结构体需要用引用形式
	var f Cat = &HeiCat{}
	f.miaomiao()


	/* 
		接口的妙用： 
		①快速生成方法 我们用结构体去实现接口时，可以通过idea等开发工具快速生成接口定义的方法模板，提升开发效率 

		②如果接口里面什么方法都没有,则什么类型都可以用来实现他，这就是any类型的实现原理 

		any的底层源码 
		type any = interface{}
	*/
}
```

### 多接口的实现

>理解

```sh
由于接口的实现不是强绑定，所以所谓多接口只要类型具备了多个接口的方法，就实现了
```

>举例

```go
package main

/*
	在实际开发中，我们不会在一个接口中声明多个方法
	因为方法过多了，不好维护，其使用效果差
*/

type Inter01 interface{
	Writer(string)
}

type Inter02 interface{
	say() string
	Walk()
}

type Test struct{

}

func (i *Test) Writer(string){
	println("天天向上")
}

func (i *Test) say() string{
	return "说"
}

func (i *Test) Walk(){
	println("走起！")
}

func main()  {
    var t1 Inter01 = &Test{}
    var t2 Inter02 = &Test{}

    t1.Writer("好好好")
    t2.Walk()
}
```

### 结构体嵌套接口

接口也是一个类型，所以类型能放的地方接口就能放，常见的应用就是结构体里面放接口

>嵌套方式

```go
//接口
type Inter01 interface{
	Writer(string)
}
//接口
type Inter02 interface{
	say() string
	Walk()
}

//具名嵌套
type Test01 struct {
	it Inter01
}
//匿名嵌套
type Test02 struct {
	Inter01
	Inter02
}
```

#### 嵌套结构体的用处

通常用另一个结构体去实现接口，然后根据需求，使用不同的结构体实现接口，达到不同的功能

>举例

```go
//接口
type Inter01 interface{
	Writer(string)
}

//具名嵌套
type Test01 struct {
	it Inter01
}

// 实现Inter01接口，书写汉字
type Chinese struct{
	word string
}

func (it *Chinese) Writer(string){
	println("写字")
}

//实现Inter01接口，书写英文 
type English struct{
	word string 
}

func (it *English) Writer(string){
	println("this is English")
}
```

>使用

根据不同的场景，用不同的结构体实例化，实现不同的功能

```go
func main()  {
		/* 
			在我们实现例化结构体的时候可以根据传参来实现接口 

			通过传不同的用于实现接口的结构体来达到不同功能
		*/
		t1:= &Test01{
				it:&Chinese{},
		}

		t2 := Test01{
			it:&English{},
		}

		// 结构体实例在实现时对接口采用不同的实现即可
		t1.it.Writer("XX")
		t2.it.Writer("xxx")
}
```

>好处

```sh
方便代码的解耦
```

### 通过interface解决动态类型传参

>背景

```sh
有时候我们会需要动态类型传参，例如封装一个加法函数，由于int类型繁多，所以要封装很多个
```

```go
我们只要一个简单的加法功能，但由于数字类型繁多 
需要定义非常多的函数(因为函数不能重名)

在其他语言中会用泛型解决这个问题，但go的泛型是后期加入的 
	
所以通常都用空接口来解决这种问题，因为空接口可以是任意类型
```

```go
func add(a int,b int) int {
	return a+b 
}

func addInt32(a int32,b int32) int32 {
	return a+b 
}

func addInt64(a int64,b int64) int64 {
	return a+b 
}

//空接口 
func addAll(a,b interface{}) int {
	return a+b
}
```

空接口解决了类型传递的问题，但是出现了新问题

```sh
interface{} 是任意类型，同时也就不是任何类型
即他可以替代任何类型，但由于他不是具体类型，本身不具备任何方法，所以谁的方法他都不能用
```

### 类型断言

>背景

```sh
为了解决传参时的类型限制，我们使用了空接口类型 
当我们接收到参数需要使用时，需要变回具体类型 
这时候就需要用到空接口类型独有的类型转换 -- 类型断言 
```

>语法

```go
空接口类型变量.(具体类型)
返回值有两个：
① 断言的结果 
② 布尔值，是否成功

//举例
var a interface{}

b,ok := a.(int)
```

>举例

```go
package main

import "fmt"
//空接口 
func addAll(a,b interface{}) int {
	/* 
		为了解决传参时的类型限制，我们使用了空接口类型 
		当我们接收到参数需要使用时，需要变回具体类型 
		这时候就需要用到空接口类型独有的类型转换 -- 类型断言 

		语法： 
			空接口类型变量.(具体类型)
			返回值有两个：
			①	断言的结果 
			② 布尔值，是否成功 
	*/
	ai,_ := a.(int)

	bi,_ := b.(int)
	return ai+bi
}

func main()  {
	a := 1
	b := 2
	fmt.Printf("值为%#v\n", addAll(a,b))
}
```

>隐患

```sh
上述我们使用"_"忽略了可能的错误，因为我们知道目前传入的一定是整数

但是如果我们传入字符串或者浮点数就不会成功，
注意断言如果失败了，返回的是"断言类型的零值"和"false"
```

对于这种情况，常用switch处理多类型的情况。

### 通过switch进行类型判断

>需求

```sh
现在需要支持 任意类型的传入和返回
# 即接收的是interface{}
# 返回的也是 interface{},所以接到返回值后需要断言才能使用
```

>语法

```go
switch 空接口变量.(type) {
    case 类型①:
        逻辑
    case 类型②:
        逻辑
    default:
        逻辑
}

switch 变量:=空接口变量.(type) {
		case 判断①:
			逻辑
		case 判断①:
			逻辑
		default:
			逻辑
}
```

>举例

```go
package main

import (
	"fmt"
	"strings"
)
func addAll(a,b interface{}) interface{} {
	defer func ()  {
		if r:=recover();r!=nil {
			fmt.Println("这是不支持的运算")
		}
		
	}()


	// 由于我们需求中 a,b 是一样的类型，这里的写法会被简化 
	//正常情况处理这种，使用泛型最好 
	switch a.(type) {
		case int:
			ai,_ := a.(int)
			bi,_ := b.(int)
			return ai+bi
		case float64:
			ai,_ := a.(float64)
			bi,_ := b.(float64)
			return ai+bi
		case string:
			ai,_ := a.(string)
			bi,_ := b.(string)
			return ai+bi
		default:
			panic("not supported")
	}
}

func main()  {
	 a := "hello"
	 b := "world"
	 fmt.Printf("运行结果为%#v\n",addAll(a,b))

	//  虽然打印结果是个字符串，但是实际上addAll()的返回值是个interface{}
	// 是因为打印时做了隐式的断言
	// 所以我们要使用的时候需要先进行类型断言
	re := addAll(a,b)
	res,_ :=re.(string)
	// 只有断言后是字符串类型才可以用字符串的方法
	k := strings.Split(res,"l")
	fmt.Printf("%#v\n",k)
}
```

>执行结果

```sh
运行结果为"helloworld"
[]string{"he", "", "owor", "d"}
```

>牢记

```sh
拿到空接口类型后要断言后才能按照具体类型使用
```



### 接口的嵌套

>背景

```sh
是为了能达到一个复用的效果
作用：
①继承另一个接口的方法
②扩展自己的方法

注意：
①不能有同名方法,当A继承了B，B已经有X方法，则A不能再有X方法
```

>举例

```go
package main
type Father interface{
	Writer()
	Walk()
}

type Son interface{
	Father
	Read()

	//不能有同名方法
	// Walk() string
}

type ReadWriter struct {

}
func (r *ReadWriter) Walk() {
	println("走路")
}
func (r *ReadWriter) Writer() {
	println("写作")
}
func (r *ReadWriter) Read() {
	println("阅读")
}
func main()  {
	var mrw Son = &ReadWriter{}
	mrw.Read()
}
```

### 接口遇到了slice的常见错误

#### 问题描述

众所周知，在golang中，我们可以将任意类型的变量赋值给 `interface{}`，通常大家会下意识写出类似代码：

```go
var slice []int = []int{1, 2, 3}
var sliceI []interface{} = slice
```

编译后报错

```sh
cannot use slice (type []int) as type []interface {} in assignment
```

疑问：

```sh
既然我可以将任意类型的变量赋值给 "interface{}"，
为什么就不能把任意类型的切片赋值给 "[]interface{}"
```

#### 问题的原因

首先需要明白，`[]interface{} `不是接口，而是一个切片，其元素类型为`interface{}`，即该切片中的元素实际可为任意类型。

其次，`[]MyType`切片与 `[]interface{}` 切片的内存布局是完全不同的：

* 每个` interface{} `接口占用两个字（一个字表示所包含内容的类型，另一个字表示所包含的数据或指向数据的指针），于是长度为 N 的 `[]interface{}` 切片，由 N * 2 个字长的数据块支持；
* 而同等长度的 `[]MyType` 切片，其数据块的长度为 N * `sizeof(MyType)` 个字。

​	由此而知，两者的数据结构不一致，便无法赋值了。
#### 解决方案

创建等长的 `[]interface{}` 切片，并逐一赋值。代码如下：

```go
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	sliceI := make([]interface{}, len(slice))
	for i, val := range slice {
		sliceI[i] = val
	}
	fmt.Println(sliceI)
}
```

#### 函数传参

可变类型参数时，我们将`...type`理解为`[]type`,但实际上是允许传入多个`type`类型的值，然后用`[]type`的形式去操作，

```go
package main

import "fmt"
/*
	...interface{}允许传递0~多个interface{}的值
	他形式上是[]interface{}，可是实际上就是多个interface{}
	所以传入任意值都可以
	
	不会出现上述的切片问题，例如我传入一个 []string
	这里也只会把它作为 interface{}
	因为他是把所有参数组合为 []interface{},而不是指datas的类型是 []interface{}，datas的实际形成在逻辑处理时出现，并不是在某个参数传递时确定
*/
func mPrint(datas ...interface{}){
	for _,val := range datas{
		fmt.Println(val)
	}
}


func main()  {
	/* 
		[]interface{}本质上是一个切片
		和interface{}不一样 

		所以 []type ≠ []interface{}
	*/

	var data = []interface{}{
		"body",18,1.80,
	}
	// var a = "你好"
	mPrint(data)

	println("---------------------")
	var name = []string{
		"zs","lisi",
	}
	mPrint(name)

}
```

>结果

```sh
[body 18 1.8]
---------------------
[zs lisi]
```

## Go工程管理

### 关于引包时找不到文件的解决方案

由于go的工程路径在`GOPATH`下，他会从`GOPATH`下的src开始寻找。

所以我们需要把某个目录定义成`GOPATH`，然后在这个目录的src文件夹下建立项目。

>图01

>图02

>图03

### package的定义

>package的作用

```sh
①用来组织源码，是多个go源码的集合,代码复用的基础,也就是常说的包
```

```sh
②每个源码文件开始都必须要申明所属的package,即源码是属于哪个包下面的
```

```sh
③在golang中，package的名称可以与源码文件名、文件夹名不一样，在Python中包和文件名绑定
```

>约定

```sh
①package的命名采用小驼峰命名，例如 firstTest
```

```sh
②同一个文件夹下的直属包名必须一致，这是强制约定的，idea会检查，但是vscode不会
```

>举例

`src01/first/01_包的定义.go`

```go
package firstTest
/*
	1. package 用来组织源码，是多个go源码的集合,代码复用的基础,常说的包

	2.每个源码go文件开始都必须要申明所属的package，即源码是属于哪个包下面的

	3.在golang中，package的名称可以与源码文件名、文件夹名不一样,在Python中包和文件名绑定

	4.同一个文件夹下的直属go文件，包名必须一致，这是强制约定的,idea会检查，但是vscode不会

	4.1 包只与本级文件夹，也就是本级目录关联。例如A文件夹下包名为a，其子文件夹包名任意，即使也为a也和A文件夹的无关
	4.2 A目录下包名为a，B目录下包名也为a，他们两个没有任何关系，因为包是以目录区分的，从包的导入就可以得知

	5.同一个目录下，包名一致，go文件之间是相通的，就相当于把一份大文件拆了N多份，所以不能申明同名函数，同名全局变量等等...
	5.1 同一个目录下，属于同一个包，go文件之间不需要相互引用，应为他们本质上写在一起
*/
import "fmt"

func Hello(){
	fmt.Println("欢迎使用")
}

func main(){
	//同一个目录下的直属go文件是相同的，不需要互相引用
	c := add(3,4)
	fmt.Println(c)
}
```

`src01/first/02_和01同包.go`

```go
package firstTest
/*
	同一个目录下的包名要一致
	
	本质上写在一起，所以可以互相引用
*/
func add(a,b int) int {
	return a+b
}
```

>理解：直属、同包、`本级目录`

`直属`

```sh
文件的直属go文件，指的是这个文件目录下的go文件，子文件夹里的不算，也就是是最近的一级目录，嫡系。
```

`同包`

```sh
同一个目录下的go文件在同一个包内，以文件夹为最小单位。

即使B是A的子文件夹，他也与A是不同的包。

不同目录下的包名可以一样，因为在引用包时不会具体到go文件，而是到文件夹也就是目录层。
```

`本级目录`

```sh
仅仅指的是本级文件夹，任何子级、父级、兄弟都属于其他目录。

所以可以说，本级目录下的go属于同一个包。
```

### 导包

>图03

#### 基本导包

```sh
# 单个导入
import 路径

①路径是从GOPATH的src目录下开始的，由文件夹名组成，和go.mod无关
②只需要到文件夹，不需要到具体的go文件，因为文件夹下的go文件包名一致

# 多个导入
import (
	包1
	包2
)
```

>举例

```go
package second

import "go_project/src01/first"
/*
	按路径导入
	import 路径

	①路径是从GOPATH的src目录下开始的，由文件夹名组成，和go.mod无关
	②只需要到文件夹，不需要到具体的go文件，因为文件夹下的go文件包名一致

	引入的包的使用
	包名.方法()
	包名.变量

	注意：
	①只有首字母大写的方法或变量才能在包外访问
	②但可以通过包暴露的方法间接访问到包内部的首字母小写的方法或变量

*/
func main()  {
	firstTest.Hello()
}
```

#### 包的别名

>语法

```sh
import 别名 路径

import (
	别名 路径
	
	路径
)
```

>使用

由于别名代替了引入的包，所以就变成了

```sh
别名.方法()
别名.变量
```

>举例

```go
package second

/*
	因为不同的目录可以存在同名的包，
	所以有时候我们需要通过给包起别名来区分引入的包

	语法
	import 别名 路径
*/
import firstOne "go_project/src01/first"
func sayHello()  {
	firstOne.Hello()
}
```

#### 包名使用的优化

>场景

```sh
每次都需要通过
包名/别名.方法()
包名/别名.变量

会让代码看起来很长,
```

>语法

```sh
#用一个英文的原点在引入的包之前

import . 路径

import (
	包1
	. 包2
)

#相当于把包里面所有暴露的方法变量都加入了本包
```

>小结

```sh
不推荐使用,原因：
① 因为可能和本包的方法或变量名冲突
② 不好溯源，不知道方法是哪个包的
```

>举例

```go
package second

/*
	每次都需要通过
	包名/别名.方法()
	包名/别名.变量

	会让代码看起来很长,

	语法：用一个英文的原点在引入的包之前

	import . 路径

	相当于把包里面所有暴露的方法变量都加入了本包
	不推荐使用,原因：
	① 因为可能和本包的方法或变量名冲突
	② 不好溯源，不知道方法是哪个包的
*/
import . "go_project/src01/first"
func sayNo()  {
	Hello()
}
```

#### 匿名导入

>场景

```sh
我们引入一个包暂时不使用，这样编译器会报错，并且会把这个包移除

可是我们想要保留这个包，这时候就可以使用匿名导入，就用下划线为别名
```

>语法

```go
import _ 路径

import (
	包1
    _ 包2
)
```

>举例

```go
package second

/*
	有如下的场景：
	我们引入一个包暂时不使用，这样编译器会报错，并且会把这个包移除
	可是我们想要保留这个包，这时候就可以使用匿名导入，就用下划线为别名

	语法：
	import _ 路径
*/
/*
	匿名导入的包虽然没有被我们显式使用，可以会执行该包的init方法
	这通常被用于初始化的场景
*/
import _ "go_project/src01/first"
```

>匿名导入的另一个使用场景

```sh
匿名导入的包虽然没被我们直接使用，但是会执行该包的`init`方法
这在很多初始化的时候使用很频繁
```

### mod文件的下载方式

`go.mod`文件是自动维护的，会记录我们项目使用的第三方包，相当于前端中的`package.json`

>初始形态

```go
module go_home

// go的版本
go 1.20
```

>使用第三方包

需要允许go modules初始化，ide会有提示

需要我们配置go的sdk和GoProxy

```sh
package three

/*
	①引入第三方包，观察项目结构和go.mod文件的变化
*/

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getData() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

将包名拷贝过来后由于还没有引入所以会飘红，引入包，来在最外层目录的终端，输入

```go
go get "github.com/gin-gonic/gin"
```

安装完成后,

`go.mod`

```go
module go_home

// go的版本
go 1.20
//依赖，可以有多个require
require (
	github.com/bytedance/sonic v1.9.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.14.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

同时还多了一个`go.sum`文件，是对文件做校验用的，自动维护

在`External Libraries`目录下的`Go Modules`下可以找到安装的依赖，相当于前端的`node_modules`

### 设置GOProxy国内镜像

>终端输入

```sh
go env
```

>关键点

```sh
# 111是因为这是go 1.11版本后引入
set GO111MODULE = on

# 要将他设置为on
go env -w GO111MODULE=on
```

```sh
set GOPROXY=https://goproxy.cn,direct

#设置为国内镜像源
go env -w GOPROXY=https://goproxy.cn,direct
```

最后

```sh
# 验证
go env 
```

>为什么要开启go modules

```sh
① 自动添加依赖，他会自己分析并下载依赖
② 删除未使用的依赖项
```

一些命令

```sh
go list -m -versions 包名
这样就会列出包有哪些可用版本
```

```sh
# 下载并替换现有的包
go get 包名@版本
```

### goget\gomod相关命令

go找包的两个途径

```sh
① pkg网站 相当于npm官网
② github直接搜
```

#### go get命令

> 用于下载包

```sh
go get 包名@版本
# 下载某个版本的包
# 用下载下来的版本替换当前项目版本
```

>举例

```go
go get github.com/go-redis/redis/v8
```

>升级某一个包

```sh
# 升级到最新的版本
go get -U  

# 升级到最新的修订版本
go get -U=patch

# 更换到某个版本，前提有
go get 包名@版本
```

>注意

```sh
go get 会修改go.mod文件
```

#### go mod

```sh
# 查看有哪些go mod命令
go mod help

# 把包下载到本地缓存
go mod download

#手动编辑go.mod文件和脚本
go mod edit

#看一下依赖图
go mod graph

#初始化一个mod文件
go mod init 模块名

# 下载缺少的依赖并删除不用的依赖
go mod tidy
```

```sh
go mod tidy最常用
他会同步修改 go.mod文件
```

#### go install

```sh
#全局安装
go install 包名@版本
```

#### go mod replace的应用场景

>场景

```sh
一开始，写了一个A项目，仓库是Project-A，但是代码仓库的go.mod中设置的是github.com/uni/A
```

`go.mod`

```go
module github.com/uni/A
```

即`go.mod`里面设置的module和项目文件夹名称不一致

```sh
B项目由于依赖了A项目。import的是 Project-A，因为是按照文件夹名引用的，
go get命令的时候，由于package和代码仓库的名称不一样，所以会失败
此时就需要用到replace命令
```

```sh
go mod edit -replace `go.mod的module`=`真实文件路径@版本号`

最后在go.mod会生成一行

replace go.mod的module => 真实文件路径 版本号

注意，要看你真的写了什么，最后会生成这种格式
```

## Go代码规范

### 为什么要代码规范

```sh
1.代码规范并不是强制的，但是不同的语言一些细微的规范还是要遵循
2.代码规范主要是为方便团队内部形成一个统一的代码风格，提高代码的可读性、统一性
```

### 代码规范

#### 命名规范

>包名

```sh
1.包名尽量和目录名保持一致，除非想要标明版本号，如果是版本号，则可以和版本号上一级目录保持一致

2.尽量采取有意义的包名，包名尽量简短

3.包名不要和标准库名冲突

4.包名全部小写，不要采取驼峰命名和下划线蛇形命名
```

>文件名

```sh
1.如果有多个单词可以采用蛇形命名法
例如： user_name.go
```

>变量名

常见的命名法有两种：蛇形和驼峰

```sh
1.蛇形：下划线分割 python和php常用
```

```sh
2.驼峰：
小驼峰：首字母小写，后面单词首字母大写，java、C、javaScript常用
大驼峰：所有单词首字母大写
```

```sh
go推荐采用驼峰规范
```

一些规范(随意)

```sh
① bool类型用has is can allow等开头
② 专有词，例如API URL，不要改成Api Url

# 不强制
```



>结构体

```sh
推荐采用驼峰命名法
```

>接口命名

```sh
推荐采用驼峰命名法,

#不强制
有些公司使用er结尾或I开头，根据公司规范
例如：
type IRead interface{}
type Reader interface{}
```

>常量命名

```sh
全部字母大写，如果有多个单词使用蛇形命名法
```

>目录命名

```sh
这个没有规范，驼峰、蛇形都行，甚至有些还会采用 点型
```

#### 注释规范

go提供两种注释

```sh
1. //   适合单行注释
2. /**/ 适合大段的注释
```

>变量名加注释

位置在变量后边或者变量上面

```go
//age 年龄
age := 19

//Course 对课程信息的包装
type Course struct {
    Name string
}
```

```go
c:= 19 //c 年龄

type Course struct {
    Name string
}//Course 对课程信息的包装
```

>包的注释

```sh
① 包的整体说明
② 包的作者
③ 创建时间
```

```go
//user 包，关于用户的包
//author: dsj
//datetime: 20231105
package user
```

>函数的注释

```sh
① 函数总体的介绍
② 参数说明
③ 返回值
```

```go
//GetCourse 获取课程信息
//参数：
// c: Course对象
//返回值
// 课程的名字
func GetCourse(c Course) string{
    
    //函数内逻辑注释，可以采用单行或多行
    
    //返回课程名称
    return c.name
}
```

>import的注释

在导入的上方或者后面

```go
//导入打印的包
import "fmt"

import "net" //用于http的包
```

#### import规范

>包的分组

```sh
① go自带的包
② 第三方包
③ 自己内部的包
```

>引入的时候分组来写即可，中间隔开。例如

```go
import (
    //系统内部的包
	"os"
    "fmt"
    
    //第三方包
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    
    
    //自己的包
    "myWork/firstTest"
    
)
```

## Go单元测试

### 如何写单元测试用例

大型的项目中为了保证项目的健壮性都会编写单元测试

go中的单元测试，需要运行命令

```sh
go test
```

go单元测试的特点

* `go test`命令是一个按照一定约定和组织的测试代码的驱动程序
* 在包目录中，所有以`_test.go`为后缀的源码文件都会被`go test`运行
* 我们写的`_test.go`源码文件不用担心内容过多，因为`go build`命令不会将这些测试文件打包到最后的可执行文件

>test文件分类

|     开头      |   作用   |  参数类型  |
| :-----------: | :------: | :--------: |
|   Test开头    | 功能测试 | *testing.T |
| Benchmark开头 | 性能测试 | *testing.B |
|  Example开头  | 示例函数 | *testing.E |
|   Fuzz开头    | 模糊测试 | *testing.F |

>基本分类的意义

* Test测试函数

  ```sh
  函数前缀为Test 主要用于测试程序的一些逻辑行为是否正确
  ```

* Benchmark又叫基准测试

  ```sh
  函数名前缀为Benchmark 主要测试函数的性能
  
  基准函数会自定义一个时间段用于执行代码,如果代码简洁,被测函数的执行次数需要成倍增加（直到达到基准测试函数给的预期,然后统计一共执行了多少轮,，每轮平均用时多少）
  ```

* Example测试

  ```sh
  函数的前缀名为 Example 为文档提示示例文档
  ```

* Fuzz模糊测试

  ```sh
  模糊测试（Fuzz Testing）是一种随机化的测试方式，它通过生成大量随机数据来模拟真实情况下的输入数据，以此检测代码中的各种漏洞和错误。在Golang的单元测试中，模糊测试可以帮助开发人员找出代码中可能存在的缺陷，并尽早修复这些问题。
  
  模糊测试的作用:
  模糊测试主要有两个作用。第一，它可以检测代码的健壮性，验证代码是否能够正确处理各种异常情况。第二，它可以帮助开发人员找出代码中的漏洞和错误，从而提高代码的质量和稳定性。
  
  模糊测试的使用方法：
  要使用Golang进行模糊测试，需要引入相应的测试框架。其中，GoFuzz是一个流行的Golang模糊测试框架，它可以自动生成大量的随机数据，并通过执行覆盖率分析等技术来检测代码中的漏洞和错误。
  ```

### Test功能测试

> 待测函数

`add.go`

```go
package goTest

func add(a, b int) int {
	return a + b
}
```

>Test测试

`add_test.go`

```go
package goTest

import "testing"

/*
	在go中test测试要和被测目标处于同一个包，
	因为需要使用到一些小写的不对外暴露的变量等
*/

// 测试函数
//  1. TestAdd 以Test开头，add方法的测试函数
//  2. 参数：
//     t 类型*testing.T,用于反馈测试结果
func TestAdd(t *testing.T) {
	//在测试函数中调用需要测试的方法
	//查看执行结果是否正确
	re := add(1, 3)

	//通过假定的测试结果来进行判断
	if re != 4 {
		//如果错误了如何报错 t.Errorf方法
		t.Errorf("expect %d,actual %d", 4, re)
	}
}
```

IDE和golang会将`_test.go`文件识别为测试文件

>注意事项

```sh
1. add.go与add_test.go 应该在同一个包下，因为一些小写开头的变量或方法，属于本包私有，保外无法访问不好测试

2.在 func TestAdd 旁边会有一个运行测试的小三角，可以直接运行测试，可以测试单个方法，运行add_test.go会运行里面所有的测试函数

3. t是*testing.T类型的变量，在测试函数里面执行需要测试的方法，看执行结果是不是理想结果，如果不是预期结果则进行错误捕获
```

>执行成功的案例

```sh
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
PASS
```

>错误捕获的案例

```sh
=== RUN   TestAdd
    add_test.go:22: expect 5,actual 4
--- FAIL: TestAdd (0.00s)

FAIL
```

### 如何跳过耗时的单元测试用例

有时候我们在测试文件中写的测试函数很多，例如十多个，我们自然是想将这些测试函数都运行，但是存在一个问题。

```sh
可能有些单元测试非常的耗时，如何去跳过比较耗时的单元测试
```

#### short模式

>举例

```go
package goTest

import (
	"fmt"
	"testing"
)

/*
	在go中test测试要和被测目标处于同一个包，
	因为需要使用到一些小写的不对外暴露的变量等
*/

// 测试函数
//  1. TestAdd 以Test开头，add方法的测试函数
//  2. 参数：
//     t 类型*testing.T,用于反馈测试结果
func TestAdd(t *testing.T) {
	//在测试函数中调用需要测试的方法
	//查看执行结果是否正确
	re := add(1, 3)

	//通过假定的测试结果来进行判断
	if re != 4 {
		//如果错误了如何报错 t.Errorf方法
		t.Errorf("expect %d,actual %d", 4, re)
	}
}

func TestAdd2(t *testing.T) {
	//在short之前，会运行
	fmt.Println("pass")

	//当我们 以short模式进行测试时
	//下面的不会运行
	if testing.Short() {
		t.Skip("short 模式下跳过")
	}

	//在short之后，不会运行
	fmt.Println("no!!")
}
```

>不使用short模式，正常测试

```sh
go test add_test.go
# 命令运行需要main包，无法执行

#推荐直接在add_test.go上执行右键执行 Run 'add_test.go' 

#也可以直接在控制台输入
go test
#他会自动取执行目录下所有的test文件
```

>执行结果

```sh
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestAdd2
pass
no!!
--- PASS: TestAdd2 (0.00s)
PASS
```

>使用short模式

```sh
go test -short
```

>执行结果

```sh
pass
PASS
ok      gotest  0.137s
```

>总结

```sh
如果开启了short模式，那么测试函数中short之前的部分会被执行，之后的会被跳过，类似return的约束
```

### 基于表格驱动测试

测试函数测试时，测试的数据不止一组，如果每组都写一个测试函数显然不合适

>举例

```go
//基于表格的单元测试
//解决我们想要测试多组数据的情形
func TestAdd3(t *testing.T) {
	//使用匿名结构体
	//构筑多组值
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{12, 12, 24},
		{-9, 8, -1},
	}

	//遍历结构体实例 快速写一堆测试用例
	for _, val := range dataset {
		re := add(val.a, val.b)
		if re != val.out {
			t.Errorf("期待值为%d,实际值为%d", val.out, re)
		}
	}
}
```

虽然称做表格，但实际上只是多组值而已，本质上就是遍历多组值来检测方法是否都能正确执行

### benchmark性能测试

在写一些核心的函数的时候，我们希望其性能比较高

>举例

```go
// 性能测试
// 参数：
// 性能测试的参数类型为 *testing.B
func BenchmarkAdd(be *testing.B) {
	//重置时间，当有多个性能测试函数在同一个测试文件时使用
	be.ResetTimer()
	
	var a, b, c int
	a = 123
	b = 456
	c = 579

	//be.N 测试运行的轮数，可以通过go test命令传入
	for i := 0; i < be.N; i++ {
		if actual := add(a, b); actual != c {
			fmt.Printf("期待值为%d,实际值为%d", c, actual)
		}
	}

	//关闭
	be.StopTimer()
}
```

>运行

```sh
① 在func BenchmarkAdd 旁边运行，执行单个

② 在add_test.go上运行，Run会多一个选项 gobench add_test.go

③ go test -bench=".*"  终端执行所有的性能测试函数
```

>结果

```sh
goos: windows
goarch: amd64
pkg: gotest
cpu: Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz
BenchmarkAdd
BenchmarkAdd-8          1000000000               0.2836 ns/op
PASS
```

### 小结

```sh
这里只是对单元测试的入门讲解，后续将根据实际情况进行补充
```

## Go并发编程

### go并发编程初体验

>背景

```sh
Python、Java、PHP都是使用多线程编程、多进程编程，多线程和多进程存在的问题主要是耗费内存

这些语言的多线程基本上都会被操作系统调度，例如启动一个python的多线程，本质上是在操作系统中开辟了一个新线程。主要的缺点有以下两个：
1.内存占用很大
2.线程切换成本高
```

随着`web2.0`的发展，对并发的要求越来越高，后来就出现了用户级线程，也被称为绿程、轻量级线程、协程。NodeJs是有协程的，其他语言为了适应发展也支持了协程。

>协程的优势

```sh
1.内存占用小
2.切换快，因为他并不是通过操作系统切换的
```

#### go语言的协程

go语言诞生在web2.0之后，故只提供了协程，其协程叫做`goroutine`,这也就决定了go只有协程的生态，二其他语言由于线程生态较为完善，开发协程就比较缓慢。

>语法

go开启协程非常简单

```sh
go 函数名
# 这样就开启了协程，该函数就被异步执行了
```

>举例

```go
package main

import "fmt"

func myPrint() {
	fmt.Println("我的打印")
}

func main() {
	/*
		go的协程非常易于书写

		go 函数名

		这样就开启了协程，该函数就被异步执行了
	*/
	go myPrint()
}
```

>现象

```sh
但是我们发现，我们执行后,myPrint并没有执行打印，
原因：
main 函数是主协程，主协程一旦执行完毕，或者说挂了，内部调用的子协程也会跟着完结

所以有一种说法叫 主死随从
就是主协程死掉了，子协程也会死，主要原因是主协程挂掉，程序就退出了，自然不会执行
```

>演示方案

```go
package main

import (
	"fmt"
	"time"
)

func myPrint() {
	fmt.Println("我的打印")
}

func main() {
	go myPrint()
	
	
	//验证goroutine是异步的，如果是同步的，应该在它之前打印
	fmt.Println("我来自主协程")
	
	
	//主协程沉睡2秒钟，等待子协程执行
	time.Sleep(2 * time.Second)
}
```

>执行结果

```sh
我来自主协程
我的打印

```

>小结

```sh
1. 主死随从 主协程一旦结束，子协程也会退出
2. 子协程是异步的
```

#### 一些案例

>协程执行for无限循环，等到主进程退出

```go
package main

import (
	"fmt"
	"time"
)

// 子协程 无限循环等到主协程结束退出
func loopPrint() {
	for {
		//加入时间限制，否则执行的太快了
		//因为我们的执行逻辑太少了
		time.Sleep(1 * time.Second)
		
		fmt.Println("无限打印")
	}
}

func main() {

	go loopPrint()

	//验证goroutine是异步的，如果是同步的，应该在它之前打印
	fmt.Println("我来自主协程")

	//主协程沉睡2秒钟，等待子协程执行
	time.Sleep(2 * time.Second)
}
```

>很多时候我们不需要在外面单独定义一个函数，而是使用匿名函数

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//匿名函数启动goroutine
	go func() {
		fmt.Println("匿名打印")
	}()

	//验证goroutine是异步的，如果是同步的，应该在它之前打印
	fmt.Println("我来自主协程")

	//主协程沉睡2秒钟，等待子协程执行
	time.Sleep(4 * time.Second)
}
```

>启动100个goroutine

由于goroutine协程占用内存极少，所以可以同时启动上万个goroutine

```go
package main

import (
	"fmt"
	"time"
)


func main() {
	//启动100个goroutine
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	//验证goroutine是异步的，如果是同步的，应该在它之前打印
	fmt.Println("我来自主协程")

	//主协程沉睡2秒钟，等待子协程执行
	time.Sleep(4 * time.Second)
}
```

问题:

```sh
①没有按顺序打印，甚至打印很多重复的出现
1.for循环的时候，每次变量都会被重用，就是说每次进入循环体的`i`都是新的
2.goroutine的执行是个异步的过程，他和for循环不同步

② 没有执行一百次
因为我们设置的主协程退出执行的时间不够执行，所以应该采取另一种方案，等到子协程完成
```

解决重复，采用值传递

方案1：

```go
for i := 0; i < 100; i++ {
    temp := i
    go func() {
        fmt.Println(temp)
    }()
}
```

方案二：

```go
for i := 0; i < 100; i++ {
    go func(k int) {
        fmt.Println(k)
    }(i)
}
```

顺序无法保证，因为goroutine是异步的

#### 拓展

想要深入了解goroutine的底层原理，可以去了解`GMP`调度原理

### 通过waitgroup等待协程的执行

>背景

```sh
由于协程"主死随从"的特性，主协程一旦结束子协程就会终止
之前为了让子协程执行，我们采用了让主协程沉睡的方式，但是实际需要沉睡多久，完全无法预期
所以需要采取另一种思路，让子协程运行完成后再结束
```

>问题

```sh
1.子的goroutine如何通知到主的goroutine自己结束了
2.主的goroutine如何知道子的goroutine已经结束了
```

>介绍

```sh
waitGroup主要用于goroutine的执行等待
Add方法要与Done方法配套，Add有多少个Done就需要多少个
所以为了防止忘记goroutine的关闭，常用defer配合Done
```

>步骤

```go
//1. 制造一个waitGroup实例
var wg sync.WaitGroup
```

```go
//2. wg.Add(整数)
// 要监控多少个goroutine执行结束
//在知道具体数量的情况下我们可以写具体数值
wg.Add(100)

//不知道具体数目时，每执行一次就调一次
wg.Add(1)
```

```go
//3. wg.Done()
//与Add成对出现，出现的次数等于Add里面的数值
wg.Done()
```

```go
//4.wg.Wait()
//写在主协程末尾，或者用defer写在最前，表示监听所有子协程完成后才允许终止主协程
wg.wait()
```

>举例

```go
package main

import (
	"fmt"
	"sync"
)

//子的goroutine如何通知到主的goroutine自己结束了
//主的goroutine如何知道子的goroutine已经结束了

func main() {

	//① 实例化WaitGroup
	var wg sync.WaitGroup

	//② 要监控多少个goroutine执行结束
	//由于我们知道有100个，所以就写100个
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			// Add与Done要同时使用，每次调用Done就自动减1
			// 如果不调用Done，这个协程就不会结束，最后会形成deadlock(死锁)
			wg.Done()
		}(i)
	}

	//等到所有goroutine结束再结束主协程
	wg.Wait()
}

/*
	小结:
	waitgroup主要用于goroutine的执行等待
	Add方法要与Done方法配套，Add有多少个Done就需要多少个
	所以为了防止忘记goroutine的关闭，常用defer配合Done

	结合defer的特性，
	① 我们可以把 defer wg.Wait() 写在主协程的开头
	② 把defer wg.Done()写在每次goroutine的开头

*/
```

> 如果不确定次数，更推荐Add和Done以`1:1`出现

```go
package main

import (
	"fmt"
	"sync"
)

//子的goroutine如何通知到主的goroutine自己结束了
//主的goroutine如何知道子的goroutine已经结束了

func main() {

	//① 实例化WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
        //② 要监控多少个goroutine执行结束
        //每次循环加一个，Done一个
        wg.Add(1)
        
		go func(i int) {
			fmt.Println(i)
			// Add与Done要同时使用，每次调用Done就自动减1
			// 如果不调用Done，这个协程就不会结束，最后会形成deadlock(死锁)
			wg.Done()
		}(i)
	}

	//等到所有goroutine结束再结束主协程
	wg.Wait()
}
```

>defer优化Done，如果goroutine的逻辑很长，我们常常会忘记写Done，或是将Done写错位置
>
>Done之后的goroutine代码是不会执行的，相当于阻断
>
>所以 用defer将Done写在最开始，由于defer的特性，它将最后执行

```go
package main

import (
	"fmt"
	"sync"
)

//子的goroutine如何通知到主的goroutine自己结束了
//主的goroutine如何知道子的goroutine已经结束了

func main() {

	//① 实例化WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
        //② 要监控多少个goroutine执行结束
        //每次循环加一个，Done一个
        wg.Add(1)
        
		go func(i int) {
            //防止遗忘
            defer wg.Done()
			fmt.Println(i)	
		}(i)
	}

	//等到所有goroutine结束再结束主协程
	wg.Wait()
}
```

>优化wait，wait的位置要在所有子协程的最后，也可以使用defer将它写在主协程靠前的位置

```go
package main

import (
	"fmt"
	"sync"
)

//子的goroutine如何通知到主的goroutine自己结束了
//主的goroutine如何知道子的goroutine已经结束了

func main() {
	//等到所有goroutine结束再结束主协程
	defer wg.Wait()
    
	//① 实例化WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
        //② 要监控多少个goroutine执行结束
        //每次循环加一个，Done一个
        wg.Add(1)
        
		go func(i int) {
            //防止遗忘
            defer wg.Done()
			fmt.Println(i)	
		}(i)
	}


}
```

### 通过mutex和atomic完成全局变量的原子操作

也就是goroutine当中的锁









