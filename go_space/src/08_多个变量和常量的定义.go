package main

import "fmt"

func main() {
	//不同类型变量的声明（定义）

	//var a int
	//var b float64

	//可以自动推导类型
	var (
		a int
		b float64
	)

	a, b = 10, 3.14

	fmt.Println("a = ", a)

	fmt.Println("b = ", b)

	//选中代码，按ctrl+\,注释和取消注释代码块的快捷键
	// const i int = 10
	// const j float64 =3.14

	//可以自动推导类型
	/*const (
		i int = 10
		j float64 = 3.14

	)*/

	fmt.Println("i = ", i)
	fmt.Println("j = ", j)

}
