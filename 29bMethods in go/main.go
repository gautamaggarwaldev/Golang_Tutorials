package main

import (
	"fmt"
	"math"
)
//You can declare a method on non-struct types, too.
type my_method float64

func (m my_method) abs() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}

// Pointer receivers
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods 
//often need to modify their receiver, pointer receivers are more common than value receivers.
func (v *Vertex) Scale (f float64) {  
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale1(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	m := my_method(-math.Sqrt2)

	v:= Vertex{6,7}
	fmt.Println(m.abs())
	v.Scale(2) //*10 
	Scale1(&v,10)
	// fmt.Println(v.Abs())


	p := &Vertex{6,7}
	p.Scale(3)
	Scale1(p, 7) 
	fmt.Println(v, p)
}