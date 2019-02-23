package main

import "fmt"

/* A struct is a collection of fields.
 * (And a type declaration does what you'd expect.)
 */
type Vertex struct {
	X int
	Y int
}

/* A struct literal denotes a newly allocated struct value by listing the
 * values of its fields. You can list just a subset of fields by using the
 * Name: syntax. (And the order of named fields is irrelevant.)
 * The special prefix & constructs a pointer to a struct literal.
 */
var (
	a = Vertex{8, 5}		// has type Vertex
	b = &Vertex{4, 8}		// has type *Vertex
	c = Vertex{X: 1}		// Y: 0 is implicit
	d = Vertex{}			// X: 0 and Y: 0
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{4, 5}
	// Struct fields are accessed using a dot.
	v.X = 4
	fmt.Println(v.X)

	/* Go has pointers, but no pointer arithmetic. Struct fields can be
	 * accessed through a struct pointer. The indirection through the
	 * pointer is transparent. 最后一句话的意思应该是指通过指针来访问
	 * 结构体成员和通过结构体变量来访问结构体成员的方法是一样的,都是使用
	 * 点号(.),不像C语言那样,指针访问结构体成员需要使用->操作符.
	 */
	q := &v
	q.X = 1e9
	fmt.Println(v)

	fmt.Println(a, b, c, d)
}
