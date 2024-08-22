package main

import (
	"fmt"
	"maps"
)
func main() {
	var map1 = map[string]int{"a":1,"b":2,"c":3,"d":4}
	for k,v := range map1{
		fmt.Printf("%v : %v\n",k,v)
	}

/*	** Maps are unordered data structures. If you need to iterate over a map in a specific order, you must 
	have a separate data structure that specifies that order.
*/

	var my_map1 = map[string]int{"one":1,"two":2,"three":3,"four":4}
	var my_map2 []string
	my_map2 = append(my_map2, "one","two","three","four")

	for key,value := range my_map1{
		fmt.Printf("%v : %v, ",key,value)
	}

	fmt.Println()
	
	for _,ele := range my_map2 {
		fmt.Printf("%v : %v, ",ele,my_map1[ele])
	}
	
	fmt.Println()

	clear(my_map1)
	fmt.Println(my_map1)

	x := map[string]string{"hello":"world", "good":"morning"}
	fmt.Println(x)
	
	y := map[string]string{"hello":"world", "good":"morning"}
	if(maps.Equal(x,y)) {
		fmt.Println("x == y")
	}

}