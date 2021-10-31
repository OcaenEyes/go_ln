package main

//	变量名由字母、数字、下划线组成，其中首个字符不能为数字
//	var identifier type
//	var identifier1, identifier2 type

import "fmt"

func main() {
	var a string = "gzy"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)

	//	变量声明
	//第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	//
	//var v_name v_type
	//v_name = value
	//零值就是变量没有做初始化时系统默认设置的值
	/*
		数值类型（包括complex64/128）为 0

		布尔类型为 false

		字符串为 ""（空字符串）

		以下几种类型为 nil：

		var a *int
		var a []int
		var a map[string] int
		var a chan int
		var a func(string) int
		var a error // error 是接口
	*/
	// 声明变量并初始化
	var a1 = "gzy"
	fmt.Println(a1)

	// 没有初始化就为零值
	var b1 int
	fmt.Println(b1)

	// bool 零值为 false
	var c1 bool
	fmt.Println(c1)

	var i2 int
	var f2 float64
	var b2 bool
	var s2 string
	fmt.Printf("%v %v %v %q\n", i2, f2, b2, s2)

	//第二种，根据值自行判定变量类型。
	//
	//var v_name = value

	var d2 = true
	fmt.Println(d2)

	// 第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：
	//
	//v_name := value
	//例如：
	//
	//var intVal int
	//intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
	//直接使用下面的语句即可：
	//
	//intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	//intVal := 1 相等于：
	//
	//var intVal int
	//intVal =1
	//可以将 var f string = "Runoob" 简写为 f := "Runoob"

	f3 := "gzy" // var f string = "Runoob"
	fmt.Println(f3)

	// 多变量声明
	/*
		多变量声明
		//类型相同多个变量, 非全局变量
		var vname1, vname2, vname3 type
		vname1, vname2, vname3 = v1, v2, v3

		var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

		vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


		// 这种因式分解关键字的写法一般用于声明全局变量
		var (
		    vname1 v_type1
		    vname2 v_type2
		)
	*/
	var x4, y4 int
	var ( // 这种因式分解关键字的写法一般用于声明全局变量
		a4 int
		b4 bool
	)

	var c4, d4 int = 1, 2
	var e4, f4 = 123, "hello"

	//这种不带声明格式的只能在函数体中出现
	//g, h := 123, "hello"
	g4, h4 := 123, "hello"
	println(x4, y4, a4, b4, c4, d4, e4, f4, g4, h4)
	/*
		使用 := 赋值操作符
		可以将它们简写为 a := 50 或 b := false。
		a 和 b 的类型（int 和 bool）将由编译器自动推断。
		这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
		使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明
	*/
}
