package main

import "fmt"
func main() {
	a := 12
	b := 42.0

	fmt.Printf("Type : %T and Value : %v\n", a,a)
	fmt.Printf("Type : %T and Value : %v\n", b,b)

	var c float32 = 15.32
	fmt.Printf("Type : %T and Value : %v\n", c,c)
/*
	this doesn't work
	
	b = c
    fmt.Printf("Type : %T and Value : %v\n", b,b)

*/

	// type conversion
	b = float64(c)
    fmt.Printf("Type : %T and Value : %v\n", b,b)


	var x int = 15
	var y float64 = 16.32

	ans := float64(x)*y // explicit conversion 
	fmt.Println(ans)


/*  
    throws error
	var m int = 69
	var n float32 = m

*/
    
}