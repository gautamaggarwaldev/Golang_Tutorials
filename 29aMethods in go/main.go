/*
	Go does not have classes. However, you can define methods on types.
	A method is a function with a special receiver argument.
	The receiver appears in its own argument list between the func keyword and the method name.

	SYNTAX :---

	func(reciver_name Type) method_name(parameter_list)(return_type) {

	}
*/

package main

import (
	"fmt"
	"math"
)
type person struct {
	name string
	age int
}

type vertex struct {
	x,y float64
}

func(p person) showdata() {
	// fmt.Println("Name:",p.name)
	// fmt.Println("Age:",p.age)
}

func(v vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
/*
	var pp person
	pp.name = "Gautam"
	pp.age = 65
	pp.showdata()
*/
	pp := person {
		name: "gautam",
		age: 165,
	}
	fmt.Printf("%+v\n",pp)

	v := vertex{3,7}

	pp.showdata()
	fmt.Println(v.abs())

}