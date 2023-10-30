// package main
package useDefer

import "sync"

func main()  {
	/* 
		相当于java和python中的finally

		连接数据库、打开文件、开始锁等这些操作，
		无论成功与否，最后都需要去关闭数据库、关闭文件、解锁， 
		例如java中 

		try {
			开始锁的逻辑
		}finally{
			无论try中的语句结果是什么，这个最后都会在之后执行 
			在这里解锁
		}

		这种格式看似可读性高，但是有个问题，就是try里面的逻辑可以写的特别多，最后我们忘记了该在finally中写什么 
		隔得太远了，实际可读性反而变低了
	*/

	/* 
		于是golang在设计的时候，就将finally的操作写在defer后面，defer后面的语句在函数return之前执行，相当于就是位置移动
		但是执行时机不变
	*/

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