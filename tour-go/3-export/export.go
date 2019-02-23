package main

import (
	"fmt"
	"math"
)

/* After importing a package, you can refer to the names it exports.  
 * In Go, a name is exported if it begins with a capitcal letter.
 * Foo is an exported name, as is FOO. The name foo is not exported.
 * 注意,这里是说 Foo 和 FOO 都是从包里导出的名字,不是说Foo 和 FOO一样,
 * 即它们其实是区分大小写的.
 */
func main() {
	/* 如果下面的 math.Pi 写为 math.pi 会报错,如下:
	 * ./export.go:18: cannot refer to unexported name math.pi
	 * ./export.go:18: undefined: math.pi
	 *
	 * 如果写成 math.PI 也会报错,从这里可以看到 math.Pi 和 math.PI 不一样.
	 * ./export.go:22: undefined: math.PI
	 *
	 * 另外,可以看到写成 math.pi 报的错有两个,第一个是说 math.pi 不是一个
	 * 有效的导出名字,第二个是指出程序中具体的错误位置.
	 * 而写成 math.PI 报的错只有一个,就是指出程序中错误的那个,可以看出来,
	 * math.PI 被认为是有效的导出名字,只是这个名字没有被定义而已,而 math.pi
	 * 直接就被认为不是有效的导出名字.
	 */
	fmt.Println(math.Pi)
}
