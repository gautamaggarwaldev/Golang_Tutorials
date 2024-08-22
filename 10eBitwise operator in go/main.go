package main
import "fmt"
func main() {
	/*
	1. AND(&) ---> Sets each bit to 1 if both bits are 1
	2. OR(|) ---> Sets each bit to 1 if one of two bits is 1
	3. XOR(^) ---> Sets each bit to 1 if only one of two bits is 1
	4. left shift(<<) ---> Shift left by pushing zeros in from the right
	5. right shift(>>) ---> Shift right by pushing copies of the leftmost bit in from the left, and let the 
	rightmost bits fall off
	*/

	a := 5	
	b := 3
	fmt.Printf("a & b = %d\n", a & b)
	fmt.Printf("a | b = %d\n", a | b)
	fmt.Printf("a ^ b = %d\n", a ^ b)
	fmt.Printf("a left shift by 2 = %d\n", a << 2)
	fmt.Printf("a right shift by 2 = %d\n", a >> 2)
}