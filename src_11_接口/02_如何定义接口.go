// package main
package defineInterface

/*
	接口的定义：

	type 接口名 interface{
		方法的声明
	}

	①只声明不写具体逻辑
	②方法的声明格式
	方法名() 返回值类型
*/
type Cat interface{
	miaomiao()

	Walk()

	say() string
}

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