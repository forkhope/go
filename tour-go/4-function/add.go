package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

/* When two or more consecutive named function parameters share a type,
 * you can omit the type from all but the last.
 * In this example, we shortened 
 *		x int, y int
 * to
 		x, y int
 */
func add2(x, y int) int {
	return x + y
}

/* A function can take zero or more arguments.
 * In this example, add() takes two parameters of type in.
 * Notice that the type comes after the variable name.
 * add() 函数最后面的那个 int 应该是表示函数返回值的类型.
 */
func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(543, 24))
}
