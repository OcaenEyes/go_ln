package main

import "fmt"

/*
循环类型	描述
for 循环	重复执行语句块
循环嵌套	在 for 循环中嵌套一个或多个 for 循环
循环控制语句
循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：
控制语句	描述
break 语句	经常用于中断当前 for 循环或跳出 switch 语句
continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
goto 语句	将控制转移到被标记的语句。
*/

/*
无限循环
for true  {
        fmt.Printf("这是无限循环。\n");
    }
*/

/*
Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

for init; condition; post { }
和 C 的 while 一样：

for condition { }
和 C 的 for(;;) 一样：

for { }
init： 一般为赋值表达式，给控制变量赋初值；
condition： 关系表达式或逻辑表达式，循环控制条件；
post： 一般为赋值表达式，给控制变量增量或减量。
for语句执行过程如下：

1、先对表达式 1 赋初值；

2、判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句，然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

for key, value := range oldMap {
    newMap[key] = value
}
*/

func main() {
	forSum()

	forSum2()

	forEachTest()

	forInclude()

	forBreak()

	forBreakSignLabel()

	forContinue()

	forContinueSignLabel()

	forGoto()
}

func forSum() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forSum2() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)

	if sum == 1 {
		fmt.Printf("sum是1\n")

	} else {
		fmt.Printf("sum 不是1\n")
		fmt.Printf("sum 是 %d\n", sum)
	}

	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
}

func forEachTest() {
	/*
		For-each range 循环
		这种格式的循环可以对字符串、数组、切片等进行迭代输出元素
	*/

	strings := []string{"hello", "gzy"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第%d位 x的值是：%d \n", i, x)
	}
}

/*
 Go 语言嵌套循环的格式：
for [condition |  ( init; condition; increment ) | Range]
{
   for [condition |  ( init; condition; increment ) | Range]
   {
      statement(s);
   }
   statement(s);
}
*/

func forInclude() {
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			fmt.Println(i, j)
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

func forBreak() {
	/*
		用于循环语句中跳出循环，并开始执行循环之后的语句。
		break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。
		在多重循环中，可以用标号 label 标出想 break 的循环。
	*/
	/* 定义局部变量 */
	var a int = 10
	//在变量 a 大于 15 的时候跳出循环
	for a < 20 {
		fmt.Printf("a 的值为 : %d\n", a)
		a++
		if a > 15 {
			/* 使用 break 语句跳出循环 */
			break
		}
	}

}

//多重循环，演示了使用标记和不使用标记的区别
func forBreakSignLabel() {
	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}

//continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
//
//for 循环中，执行 continue 语句会触发 for 增量语句的执行。
//
//在多重循环中，可以用标号 label 标出想 continue 的循环。
//在变量 a 等于 15 的时候跳过本次循环执行下一次循环
func forContinue() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 2
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}

//使用标记和不使用标记的区别
func forContinueSignLabel() {
	// 不使用标记
	fmt.Println("---- continue ---- ")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- continue label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			continue re
		}
	}
}

/*
goto label;
..
.
label: statement;
*/
//在变量 a 等于 15 的时候跳过本次循环并回到循环的开始语句 LOOP 处
func forGoto() {
	/*定义局部变量*/
	var a int = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a = a + 2
			goto LOOP
		}

		fmt.Printf("a的值为： %d\n", a)
		a++
	}
}
