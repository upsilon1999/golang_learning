// package main
package manyInterface

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