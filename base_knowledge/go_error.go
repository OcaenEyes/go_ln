package main

import "fmt"

//通过内置的错误接口提供了非常简单的错误处理机制。
//
//error类型是一个接口类型，这是它的定义：
//
//type error interface {
//    Error() string
//}
//我们可以在编码中通过实现 error 接口类型来生成错误信息。
//
//函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
//
//func Sqrt(f float64) (float64, error) {
//    if f < 0 {
//        return 0, errors.New("math: square root of negative number")
//    }
//    // 实现
//}
//在下面的例子中，我们在调用Sqrt的时候传递的一个负数，然后就得到了non-nil的error对象，将此对象与nil比较，结果为true，所以fmt.Println(fmt包在处理error时会调用Error方法)被调用，以输出错误，请看下面调用的示例代码：
//
//result, err:= Sqrt(-1)
//
//if err != nil {
//   fmt.Println(err)
//}

func main() {
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}
}

// 定义一个 DivideError 结构
type DivideError struct {
	divide  int
	divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    divide: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.divide)
}

// 定义 `int` 类型除法运算的函数
func Divide(varDivide int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			divide:  varDivide,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDivide / varDivider, ""
	}
}
