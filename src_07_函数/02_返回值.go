package main

/*
函数可以没有参数或者没有返回值
*/
func noWords() {
	println("一无所有")
}

/*
	返回值的变量加类型形式，如果使用变量加类型的形式
	那么所有返回值都需要使用，而不是某几个使用变量加类型，有几个只有类型
	即形式要统一
*/

//错误①，返回值不统一
//func fn01() (sum int,error) {
//	return 10,nil
//}

// 返回值的变量相当于在开头就定义了，不要重复定义
func fn02() (sum int, err error) {
	sum = 5
	return sum, nil
}

// return 不写返回值，相当于
// return sum,err
// 值是类型零值
func fn03() (sum int, err error) {
	return
}

func main() {

}
