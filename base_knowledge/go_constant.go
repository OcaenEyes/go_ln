package main

import "fmt"
import "unsafe"

/*
常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：

const identifier [type] = value
你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。

显式类型定义： const b string = "abc"
隐式类型定义： const b = "abc"
多个相同类型的声明可以简写为：

const c_name1, c_name2 = value1, value2

常量还可以用作枚举：

const (
    Unknown = 0
    Female = 1
    Male = 2
)
数字 0、1 和 2 分别代表未知性别、女性和男性。

常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
*/

/*
iota
iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

iota 可以被用作枚举值：

const (
    a = iota
    b = iota
    c = iota
)
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

const (
    a = iota
    b
    c
)
*/
const (
	a1 = "gzy"
	b1 = len(a1)
	c1 = unsafe.Sizeof(a1)
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH

	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)

	println(a1, b1, c1)

	const (
		a2 = iota //0
		b2        //1
		c2        //2
		d2 = "ha" //独立值，iota += 1
		e2        //"ha"   iota += 1
		f2 = 100  //iota +=1
		g2        //100  iota +=1
		h2 = iota //7,恢复计数
		i2        //8
	)
	fmt.Println(a2, b2, c2, d2, e2, f2, g2, h2, i2)

}
