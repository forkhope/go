package main

import (
	"fmt"
	"math"
)

/* Every Go program is made up of packages.
 * Progarms start running in package main.
 * 使用 import 引入其他包,以供程序调用.
 * By convention, the package name is the same as the last
 * element of the import path.
 * 这句话的意思是说,按照惯例,导入的包名就是导入的包路径的最后一个元素.
 * 程序通过包名来调用包的功能.如import "fmt",然后 fmt.Println()调用了
 * fmt包里面的 Println() 方法.
 */
func main() {
	fmt.Println("Happy", math.Pi, "DAY")
	fmt.Println("Now you have %g problems.", math.Nextafter(2,3))
}
