// package main
package deferCharacteristic

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