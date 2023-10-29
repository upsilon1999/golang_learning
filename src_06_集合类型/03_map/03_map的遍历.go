// package main
package traverseMap

import "fmt"

func main()  {
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
}