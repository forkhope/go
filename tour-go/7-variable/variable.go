package main

import "fmt"

/* The var statement declares a list of variables; as in function
 * argument lists, the type is last.后面的变量不可少,否则报错:
 * ./variable.go:9: syntax error: unexpected semicolon or newline
 */
var x, y, z int
var c, python, java bool

/* 定义变量需要使用 var 关键字,写成下面的形式会报错:
 * ./variable.go:13: non-declaration statement outside function body
 * ./variable.go:13: syntax error: unexpected name, 
 *	expecting semicolon or newline
 */
//a int

/* A var declaration can include initializers, one per variable.
 * If an initializer is present, the type can be omitted; the variable
 * will take the type of the initializer.
 */
//var l, m, n int = 1, 2, 3
var l, m, n = 1, 2, 3 // 上面的声明是正确的,它也可以写成这个形式
// 初始化多个变量时,这些变量可以是不同的类型,其类型为初始值决定.
var you, qing, ren = true, false, "no!"

/* 为多个变量初始化时,必须为每个变量都赋值,不能省略任何一个,不存在省略
 * 就会使用默认值的说法.
 * 写为 val r, s, t int = 4,,6;或者 var r, s, t int = 4,6 会报错如下:
 * ./variable.go:29: non-declaration statement outside function body
 * ./variable.go:29: syntax error: unexpected name, 
 *		expecting semicolon or newline
 *
 * 写为 val r, s, t int = 4,6, 会报错如下:
 * ./variable.go:36: non-declaration statement outside function body
 * ./variable.go:36: syntax error: unexpected name, 
 * 		expecting semicolon or newline
 * ./variable.go:40: non-declaration statement outside function body
 * ./variable.go:41: syntax error: unexpected }
 */
//val r, s, t int = 4,6,

/* 在函数体外使用 := 操作符会报错:
 * ./variable.go:47: non-declaration statement outside function body
 */
//q := 6

func main() {
	fmt.Println(x, y, z, c, python, java)
	fmt.Println(l, m, n, you, qing, ren)

	/* Inside a function, the := short assignment statement can be used in
	 * place of a var declaration with implicit type.
	 * (Outside a function, every construct begins a keyword and the :=
	 * construct is not available.)
	 * 即在函数体外面必须使用 var 来定义变量,使用 := 会报错.
	 */
	var r, s, t int = 4, 5, 6
	var tian, xia = "tian", true
	fmt.Println(r, s, t, tian, xia)
}
