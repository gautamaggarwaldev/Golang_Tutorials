package main
import "fmt"
func main() {
	a := 150
	b := 18

	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)  /*if denomenator is greater than numerator then answer is zero 0*/
	
	x := 78
	y := 5
	fmt.Println(x%y)   /*if x < y then answer is numerator value*/

	q := 1
	q++ /*q += 1*/
	fmt.Println(q)

	w := 2
	w--  /*w -= 1*/
	fmt.Println(w)

	/*
	GO does not support post increament operator
	--x
	++x 
	it doesn't support 
	*/

	c := float64(b)
	g:= 456.32
	fmt.Println(g*c)

}