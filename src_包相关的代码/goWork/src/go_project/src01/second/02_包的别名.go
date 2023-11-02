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