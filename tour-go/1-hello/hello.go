package main

import "fmt"

/* 编译和运行这个程序的方法: go run hello.go.
 * 似乎没有生成可执行文件,直接解析执行.
 *
 * Go语言里面,左大括号需要跟在语句的后面,不能单独放在一行,否则会报错.
 * 例如写成
 *			func main()
 *			{
 *			}
 * 则在使用 go run hello.go 语句来运行程序时,报如下的错:
 * ./hello.go:9: syntax error: unexpected semicolon or newline before {
 */
func main() {
	fmt.Println("Hello, 世界")
}
