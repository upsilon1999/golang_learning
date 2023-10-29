package main

/*
参数 ...类型
代表接受零个或多个该类型的参数,在函数中这是一个切片，所以可以用切片的方式使用
这种写法必须在参数列表的最后，且只能有一个
*/
func fn01(desc string, items ...int) (sum int, err error) {
	//items接受零个或多个int类型的参数组成切片
	for _, val := range items {
		sum += val
	}
	return sum, nil
}

func main() {
	println(fn01("hello", 1, 2, 3, 4, 5, 6))
}
