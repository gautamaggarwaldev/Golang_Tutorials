package main

import "fmt"

const pi float32 = 3.14 // Typed Constant
const x = 5 // Untyped constant

const (
	a int = 1
	b = 32.5
	c = true
	d = "hello ji"
)

func main() {
	var radius float32 = 5.0
	fmt.Println(pi*radius*radius);
	fmt.Println(x); // Unchanged constant value

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}