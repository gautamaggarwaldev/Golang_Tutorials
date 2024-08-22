package main

import (
	"fmt"
	
)
func main() {
	nums := []int{1,2,3,5,6}
	sum := 0

	for i := range nums {
		sum += nums[i]
	}
	fmt.Println("Sum = ", sum)

	for i,num := range nums { //range on arrays and slices provides both the index and value for each entry.
		if num == 3 {
			fmt.Println("index : ", i)
		}
	}

	/*range on map iterates over key/value pairs.*/

	my_map := map[string]int{"one":1,"two":2,"three":3,"four":4}
	for k, v := range my_map {
		fmt.Printf("%s : %v", k, v)
	}
	fmt.Println()

	for k:= range my_map {
		fmt.Println(k)
	}

	/*range on strings iterates over Unicode code points. The first value is the starting 
	byte index of the rune and the second the rune itself*/
	for i,c := range "gautam" {
		fmt.Println(i,c)
	}
}