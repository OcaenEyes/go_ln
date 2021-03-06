package main

import "fmt"

/*
变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
*/

//指针声明格式如下：
//
//var var_name *var-type
//var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：
//
//var ip *int        /* 指向整型*/
//var fp *float32    /* 指向浮点型 */

func main() {
	pointerTest1()
	usePointer()
	nilPointer()
}

func pointerTest1() {
	var a int = 10
	fmt.Printf("变量的地址:%x\n", &a)
}

/*
如何使用指针
指针使用流程：

定义指针变量。
为指针变量赋值。
访问指针变量中指向地址的值。
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容
*/

func usePointer() {
	var a int = 20 //声明实际变量
	var ip *int    //声明指针变量

	ip = &a //指针变量的存储地址

	fmt.Printf("a变量的地址是：%x \n", &a)

	/*指针变量的存储地址*/
	fmt.Printf("ip变量储存的指针地址：%x\n", ip)

	/*使用指针访问值*/
	fmt.Printf("*ip变量的值：%d\n", *ip)
}

/*
Go 空指针
当一个指针被定义后没有分配到任何变量时，它的值为 nil。

nil 指针也称为空指针。

nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。

*/

//空指针判断：
//
//if(ptr != nil)     /* ptr 不是空指针 */
//if(ptr == nil)    /* ptr 是空指针 */

func nilPointer() {
	var ptr *int

	if ptr == nil {
		fmt.Printf("ptr的值为: %x\n", ptr)
	}
}
