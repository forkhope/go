package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

/* The expression new(T) allocates a zeroed T value and returns a pointer
 * to it. 它的两种使用方式如下:
 *		var t *T = new(T)
 * or
 *		t := new(T)
 */
func main() {
	v := new(Vertex)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)

	var t *Vertex = new(Vertex)
	fmt.Println(t)
	t.X, t.Y = 45, 2
	fmt.Println(t)
}
