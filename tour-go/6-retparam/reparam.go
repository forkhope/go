package main

import "fmt"

func split(sum int) (x, y int) {
	/* Go语言可以自动推导变量类型,但是需要使用 := 操作符,而下面的 x 和 y
	 * 使用的都是 = 操作符,好像它们没有指定变量类型就直接使用了,其实不是
	 * 这样,而是x 和 y 的类型已经在函数的定义那里指定了: x, y int.
	 * 例如,下面定义了两个变量: z = sum; s := sum; 第一句会报错,第二句不会,
	 * ./reparam.go:14: undefined: z
	 * ./reparam.go:14: cannot assign to z
	 */
	x = sum * 4/9
	y = sum - x
	//z = sum	// 不能这样为 z 赋值, 此时 z 还没有定义.
	s := sum	// 可以直接为 s 赋值,因为 := 操作符可以自动推导变量类型

	/* 如果没有下面一句,使用 go run reparam.go 来执行程序时,会报如下的错:
	 * ./reparam.go:16: s declared and not used
	 */
	fmt.Println("split(), s = ", s)

	/* 这里,单独一条 return 语句会返回 x 和 y 的语句,但是一个函数写成
	 * func split(sum int) (x, y int) 的形式,并不代表,这个函数就只能通过
	 * 单独的 return 语句返回,也可以使用 return value1, value2 的形式来
	 * 返回其他的值,这里一定要返回两个值,只返回一个值会报错.
	 */
	if sum == 20 {
		return 20, 40 // 返回值中不需要有x,y,虽然在返回值定义那里,写了x,y
	}

	return
}

/* Functions tabke parameters. In Go, functions can return multiple
 * "result parameters", not just a single value. They can be named and
 * act just like variables.
 * If the result parameters and named, a return statement without arguments
 * returns the current values of the results.
 * 这里是说, Go语言的函数返回值可以带有名字,相当于变量,例如上面的split()
 * 函数写为 func split(sum int) (x, y int) {, 则定义了 x 和 y 两个变量,
 * 在函数里面可以直接 x 和 y 两个标识符.然后单独的一条 return 可以返回这
 * 两个变量的值,不需要在 return 语句里面再指定所要返回的变量.
 */
func main() {
	fmt.Println(split(17))
	fmt.Println(split(20))  // 测试返回其他值的情况
}
