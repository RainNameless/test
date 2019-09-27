package main /* 定义——main包，此包用于初始化,表示一个可独立执行的程序 */

import "fmt" /* 定义fmt包，此包用于（输入/输出） */

func main() { /* 定义main函数,main函数是每一个程序开始执行的函数,如果有 init() 函数则会先执行该函数 */
	fmt.Println("语言结构. ") /*  fmt.Println,此函数用于将字符串输出到控制台，并在最后自动增加换行字符 \n。 */
}
