// package main
package stringsMethods

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
