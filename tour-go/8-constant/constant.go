package main

import "fmt"

const Pi = 3.14

/* Numeric constants are high-precision values.
 * An untyped constant takes the type needed by its context.
 */
const (
	Big = 1<<100
	Small = Big>>99
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64 {
	return x*0.1
}

/* Constants are declared like vairable, but with the const keyword.
 * Constants can be charater, string, boolean, or numeric values.
 * 常量不需要指定数据类型,它的数据类型是隐式的(untyped),使用常量为变量赋值
 * 会发生隐式的类型转换,如果常量的值和变量的类型不符,会报错.
 */
func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	/* 不可以为常量重新赋值,否则会报错:
	 * ./constant.go:11: cannot assign to Pi
	 * ./constant.go:11: cannot use 5.89 (type float64) 
	 *		as type ideal in assignment
	 */
	//Pi = 5.89

	/* 常量Pi的值为5.89,属于浮点类型,使用Pi为int型变量赋值,会报错:
	 * ./constant.go:26: constant 3.14 truncated to integer
	 */
	//var a int = Pi

	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	/* 使用 needInt(Big) 会报错,提示溢出,这里是说数字常量的取值范围很大:
	 * ./constant.go:51: constant 1267650600228229401496703205376 
	 *		overflows int
	 */
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
