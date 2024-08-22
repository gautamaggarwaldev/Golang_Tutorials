package main

import (
	"fmt"
	
	"golang.org/x/exp/constraints"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}


type myDataTypes interface {
	constraint.Integer | constraint.Float  // ********************* // 
}

type mygarima int;

func addT[T myDataTypes](a, b T) T {
	return a + b
}

func main() {

	var n mygarima = 45;

	fmt.Println(addI(5,2));
	fmt.Println(addF(5.3,2.5));

	fmt.Println(addT(n,2));
	fmt.Println(addT(5.3,2.5));

}