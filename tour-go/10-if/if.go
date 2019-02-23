package main

import (
	"fmt"
	"math"
)

/* The if statement looks as it does in C or java, except that the ()
 * are gone (they are not even optional) and the { } are required.
 */
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/* Like for, the if statement can start with a short statement to
 * execute before the condition. Variables declared by the statement are
 * only in scope until the end of the if.
 */
func pow(x, n, lim float64) float64 {
	// Variables declared inside an if's short statement are also
	// available inside any of the else blocks.
	// 这里, else { 必须写在 if 的右大括号后面,不能单独写在一行,否则报错
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))

	/* 注意,如果右括号单独放在一行的话,第二个pow(3,3,20)后面一定要有逗号
	 * 否则会报错: ./if.go:37: syntax error: unexpected semicolon 
	 * or newline, expecting )
	 * 如果不想写第二个逗号的话,可以把右括号放在第二个pow()的后面,不单独
	 * 放在一行,即写为:
	 * fmt.Println(
	 *		pow(3, 2, 10),
	 *		pow(3, 3, 20))
	 */
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
