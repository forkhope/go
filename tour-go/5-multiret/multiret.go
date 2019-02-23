package main

import "fmt"

/* A function can return any number of results.
 * This function returns two strings.
 */
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// Go语言可以自动推导变量的类型,但是好像需要使用 := 操作符
	a, b := swap("hello", "world")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
