package main

import "fmt"

/* 	Go 语言数据类型包含基础类型和复合类型两大类。
基础数据类型包括：布尔型、整型、浮点型、复数型、字符型、字符串型、错误类型。
复合数据类型包括：指针、数组、切片、字典、通道、结构体、接口。
*/

func main() {
	// 布尔型
	fmt.Println("一.布尔类型")
	var bool1, bool2 bool = false, true // 布尔型的值只可以是常量 true 或者 false
	fmt.Printf("  a1=%s 	a2=%s \n\n ", bool1, bool2)

	fmt.Println("二.数字类型")
	fmt.Println("无符号整形")
	var uint_8 uint8 = 255     // uint8,无符号8位整形,存储空间8bit，值范围0 ~ 255.
	var uint_16 uint16 = 65535 // uint16,无符号16位整形,存储空间16bit,值范围0 ~65535
	var uint_32 uint32 = 4294967295
	var uint_64 uint64 = 18446744073709551615
	fmt.Printf("  uint_8=%d\n  uint_16=%d\n  uint_32=%d\n  uint_64=%d\n\n", uint_8, uint_16, uint_32, uint_64)

	fmt.Println("整形")
	var int_ int = -9999999999999
	var int_8 int8 = -128
	var int_16 int16 = -32768
	var int_32 int32 = -2147483648
	var int_64 int64 = -9223372036854775808
	fmt.Printf("  int_=%d\n  int_8=%d\n  int_16=%d\n  int_32=%d\n  int_64=%d\n\n", int_, int_8, int_16, int_32, int_64)

	fmt.Println("浮点数")
	var float_32 float32 = 1.12345678
	var float_64 float64 = 1.12345678901234567
	fmt.Printf("  float_32=%f\n  float_64=%v \n\n", float_32, float_64)

	fmt.Println("复数形")
	var complex_64 complex64 = 1.12345678 + 1.12345678i
	var complex_128 complex128 = 2.1234567890123456 + 2.1234567890123456i
	fmt.Printf("  complex_64=%v\n  complex_128=%v \n\n", complex_64, complex_128)

	fmt.Println("字符型")
	var byte_ byte = 'a'
	var rune_ rune = '符'
	fmt.Printf("  byte_=%v\n  rune_=%v\n\n", byte_, rune_)

	fmt.Println("指针地址")
	var uintptr_ uintptr = 18446744073709551615
	fmt.Printf("  uintptr_=%v\n\n", uintptr_)

}
