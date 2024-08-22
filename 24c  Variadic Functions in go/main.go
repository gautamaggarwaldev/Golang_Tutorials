/*
	Variadic functions can be called with any number of trailing arguments. For example, 
	fmt.Println is a common variadic function.
*/
package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, i:= range nums {
		total += i
	}
	fmt.Println(total)

}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
	sum(1,3,6,5,4)
	sum(1,3,9,7,4,1,30,1,10,100)
}