// package main
package interfaceNesting

/*
	接口嵌套：是为了能达到一个复用的效果
	作用：
	①继承另一个接口的方法
	②扩展自己的方法

	注意：
	①不能有同名方法,当A继承了B，B已经有X方法，则A不能再有X方法
*/
type Father interface{
	Writer()
	Walk()
}

type Son interface{
	Father
	Read()

	//不能有同名方法
	// Walk() string
}

type ReadWriter struct {

}
func (r *ReadWriter) Walk() {
	println("走路")
}
func (r *ReadWriter) Writer() {
	println("写作")
}
func (r *ReadWriter) Read() {
	println("阅读")
}
func main()  {
	var mrw Son = &ReadWriter{}
	mrw.Read()
}