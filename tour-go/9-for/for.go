package main

import "fmt"

/* Go has only on looping construct, the for loop.
 * The basic for loop looks as it does in C or java, except that the ()
 * are gone (they are not even optional) and the { } are required.
 */
func main() {
	sum := 0
	/* for 循环不能再写小括号(),而大括号{}是必须的,即使是单行语句.
	 * 注意,下面的 i++ 是合法语句,但是 ++i 却不合法,写 ++i 会报错.
	 */
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	/* 下面的写法会报错:
	 * ./for.go:23: syntax error: unexpected :=, expecting )
	 * ./for.go:26: non-declaration statement outside function body
	 * ./for.go:27: syntax error: unexpected }
	 */
	//for (i := 0; i < 10; i++) {
	//	sum +=i
	//}

	/* 下面的写法会报错:
	 * ./for.go:30: syntax error: unexpected semicolon or newline before {
	 */
	//for i := 0; i < 10; i++
		//sum +=i

	/* 在函数体内,:=操作符可以看到一种简写的定义语句,由于 sum 在上面已经
	 * 定义,下面不能再写为 sum := 1,否则会报错:
	 * ./for.go:40: no new variables on left side of :=
	 * 应写为 sum = 1
	 */
	//sum := 1
	sum = 1
	// As in C or java, you can leave the pre and post statements empty
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	/* As that point you can drop the semicolons: C's while is spelled for
	 * in Go. 即写成下面的形式,就相当于 C语言里面的 while 循环.
	 */
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	// If you omit the loop condition, it loops forever.
	//for ;; {
	//}

	/* And with no clauses at all, the semicolons can be omitted, so an
	 * infinite loop is compactly expressed.
	 */
	//for {
	//}
}
