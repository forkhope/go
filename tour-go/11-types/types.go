package main

import (
	"math/cmplx"
	"fmt"
)

var (
	ToBe bool = false;
	MaxInt uint64 = 1<<64 - 1

	// 包名是 import 的包路径的最后一个元素,下面使用cmplx来调用包的功能
	z complex128 = cmplx.Sqrt(-5+12i)
)

/* Go's basic types are:
 * 	bool
 *	string	
 *	int		int8	int16	int32	int64
 *	uint	uint8	uint16	uint32	uint64	uintptr
 *	byte	// alias for uint8
 *	rune	// alias for uint32, represents a Unicode code
 *	point
 *	float32			float64
 * 	complex64		complex128
 */
func main() {
	const f = "%T(%v)\n"	// %T: 打印变量的类型
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}
