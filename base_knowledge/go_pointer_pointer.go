package main

import "fmt"

/*
如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址
*/
//指向指针的指针变量声明格式如下：
//
//var ptr **int;

func main() {
	pointerPtr()
}

func pointerPtr() {
	var a int = 100

	var ptr *int

	var pptr **int

	// 指针ptr地址
	ptr = &a

	// 指向 指针ptr地址
	pptr = &ptr

	/* 获取ppt的值*/
	fmt.Printf("变量 a= %d\n", a)

	fmt.Printf("变量a的地址：%x\n", &a)

	fmt.Printf("指针ptr 指向的地址：%x\n", ptr)

	fmt.Printf("指针变量*ptr, ptr指针指向地址对应的值 =%d\n", *ptr)

	fmt.Printf("指针ptr 所在的地址：%x\n", &ptr)

	fmt.Printf("指针pptr 指向的地址:%x\n", pptr)

	fmt.Printf("指向指针的 指针变量 **pptr= %d\n", **pptr)

}
