package main

import (
	"fmt"
	"math/rand"
)
func main() {
	// ODD - EVEN
	var x int
	fmt.Println("Enter the number : ")
	fmt.Scan(&x)

	if x%2==0 {
		fmt.Print("Even number")
	} else {
		fmt.Print("Odd number")
	}
	


/*

	//Number is 3-digit or not

	var y int
	fmt.Println("Enter the number : ")
	fmt.Scan(&y)

	if(y>99 && y<1000) {
		fmt.Printf("%v is three digit number", y)
	} else {
		fmt.Printf("%v is not 3-digit number",y)
	}
*/

/*
	time := 12
	if(time>10) {
		fmt.Print("ok")
	} 
*/

fmt.Println();

z := 42

if a:=2 * rand.Intn(z); a>=z {
	fmt.Printf("a is %v i.e. greater than or equal to z which is %v", a,z);
} else {
	fmt.Printf("a is %v i.e. smaller than z which is %v", a,z)
}



}