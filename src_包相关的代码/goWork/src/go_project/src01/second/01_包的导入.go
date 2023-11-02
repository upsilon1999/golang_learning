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
