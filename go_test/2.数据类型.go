package main

import "fmt"

/* 	Go 语言数据类型包含基础类型和复合类型两大类。
	基础数据类型包括：布尔型、整型、浮点型、复数型、字符型、字符串型、错误类型。
	复合数据类型包括：指针、数组、切片、字典、通道、结构体、接口。
*/

func main(){

	var a bool = true		 /* 布尔型的值只可以是常量 true 或者 false */
	fmt.Println("bool=",a)

	var b uint8 = 255 		/* uint8,无符号8位整形,存储空间8bit，值范围0 ~ 255. */ /*  */
	var c uint16=65535		/* uint16,无符号16位整形,存储空间16bit,值范围0 ~65535 */
	var d



	fmt.Println("uint8=",b,c)
	fmt.Println()
	fmt.Println()
}