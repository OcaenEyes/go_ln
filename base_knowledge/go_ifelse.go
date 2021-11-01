package main

import "fmt"

//if 布尔表达式 {
//   /* 在布尔表达式为 true 时执行 */
//} else {
//  /* 在布尔表达式为 false 时执行 */
//}

// if 语句嵌套
//if 布尔表达式 1 {
//   /* 在布尔表达式 1 为 true 时执行 */
//   if 布尔表达式 2 {
//      /* 在布尔表达式 2 为 true 时执行 */
//   }
//}

/*
语句	描述
if 语句	if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
if 嵌套语句	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
switch 语句	switch 语句用于基于不同条件执行不同动作。
select 语句	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
*/

/*
switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。

switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
*/
func main() {
	/*定义局部变量*/
	var a int = 100

	if a < 20 {
		fmt.Printf("a小于20\n")
	} else {
		fmt.Printf("a大于20\n")
	}

	if a > 12 {
		fmt.Printf("a大于12\n")
	}

	fmt.Printf("a值为：%d\n", a)

	switchTest()

	typeSwitchTest()

	fallthroughSwitchTest()
}

func switchTest() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

//Type Switch
//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
//
//Type Switch 语法格式如下：
//
//switch x.(type){
//    case type:
//       statement(s);
//    case type:
//       statement(s);
//    /* 你可以定义任意个数的case */
//    default: /* 可选 */
//       statement(s);
//}

func typeSwitchTest() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T \n", i)
	case int:
		fmt.Printf("x 是 int 型 \n")
	case float64:
		fmt.Printf("x 是 float64 型 \n")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型 \n")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型 \n")
	default:
		fmt.Printf("未知型 \n")
	}
}

//fallthrough
//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
func fallthroughSwitchTest() {

	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
}
