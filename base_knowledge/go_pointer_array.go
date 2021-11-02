package main

import "fmt"

func main() {
	ptrArray1()
}

//需要保存数组，这样我们就需要使用到指针。
//
//以下声明了整型指针数组：
//
//var ptr [MAX]*int;
//ptr 为整型指针数组。因此每个元素都指向了一个值

func ptrArray1() {
	const Max int = 3
	a := []int{10, 100, 200}
	var i int
	var ptr [Max]*int

	for i = 0; i < Max; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < Max; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}
