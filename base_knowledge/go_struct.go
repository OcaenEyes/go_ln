package main

import "fmt"

//结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
//
//结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
//
//Title ：标题
//Author ： 作者
//Subject：学科
//ID：书籍ID
//定义结构体
//结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：
//
//type struct_variable_type struct {
//   member definition
//   member definition
//   ...
//   member definition
//}
//一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
//
//variable_name := structure_variable_type {value1, value2...valuen}
//或
//variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func struct1() {
	var book1 Books

	book1.title = "gzy 你好"
	book1.author = "gzy"
	book1.subject = "welcome"
	book1.bookId = 100001

	book2 := Books{"你好 gzy", "gzy", "welcome", 100002}

	fmt.Println(book1)

	fmt.Println(book2)
}
func main() {
	struct1()
}
