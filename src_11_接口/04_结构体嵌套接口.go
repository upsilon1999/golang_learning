// package main
package structHaveInterface

type Inter01 interface{
	Writer(string)
}

type Inter02 interface{
	say() string
	Walk()
}

type Test01 struct {
	it Inter01
	name string
}

type Test02 struct {
	name string
	Inter01
	Inter02
}

// 实现Inter01接口，书写汉字
type Chinese struct{
	word string
}

func (it *Chinese) Writer(string){
	println("写字")
}

//实现Inter01接口，书写英文 
type English struct{
	word string 
}

func (it *English) Writer(string){
	println("this is English")
}
func main()  {
		/* 
			在我们实现例化结构体的时候可以根据传参来实现接口 

			通过传不同的用于实现接口的结构体来达到不同功能
		*/
		t1:= &Test01{
				it:&Chinese{},
		}

		t2 := Test01{
			it:&English{},
		}

		// 结构体实例在实现时对接口采用不同的实现即可
		t1.it.Writer("XX")
		t2.it.Writer("xxx")
}