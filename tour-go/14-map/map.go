package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/* A map maps keys to values. Maps must be created with make (not new) 
 * before use; the nil map is empty and connot be assigned to.
 * 根据代码来看, []里面的是键,[]后面的是值
 */
var m map[string]Vertex

// Map literals are like struct literals, but the keys are required.
var v = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

/* If the top-level type is just a type name, you can omit it
 * from the elements of the literal.
 */
 var s = map[string]Vertex{
	 "Tian": {58.40, 55.90},
	 "Xia": {24.90, 80.28},
 }

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,		// 最后面的逗号不可少
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(v)
	fmt.Println(s)

	// 下面描述了如何改变一个已有 map 的值
	m := make(map[string]int)
	// Insert or update an element in map m: m[key] = elem
	m["Answer"] = 42	// 如果"Answer"不存在,则插入一个"Answer"键
	// Retrieve an element: elem = m[key]
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48	// 如果"Answer"存在,则更新它的值
	fmt.Println("The value:", m["Answer"])

	// Test that a key is present with a two-value assignment:
	// elem, ok = m[key]
	// If key is in m, ok is true. If not, ok is false and elem is
	// the elem is the zero value for the map's element type.
	r, ok := m["Answer"]
	fmt.Println("The value:", r, "Present?", ok)
	// 上面说明的时候,是用ok来作说明,但这不代表保存判断结果的变量就只能
	// 是ok,这里使用其他任意合法的标识符都可以.
	r, any := m["Answer"]
	fmt.Println("The value:", r, "Present?", any)

	// Delete an element: delete(m, key)
	delete(m, "Answer")
	// Similarly, when reading from a map if the key is not present
	// the result is the zero value for the map's element type.
	// 下面由于m["Answer"]已经被删除,打印出来将是 0 值
	fmt.Println("The value:", m["Answer"])
}
