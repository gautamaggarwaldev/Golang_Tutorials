package main
import "fmt"
func main() {

	/*
	array declaration :---
	
	var array_name = [5]int{} 
	array1_name := [5]int{} 

	*/

	/*array declaration with fixed size*/
	
	var arr1 = [3] int {1,2,3}
	arr2 := [5] int {1,2,3,4,5}
	fmt.Println(arr1)
	fmt.Println(arr2)

	/*array declaration with variable size*/
	var arr3 = [...] int {1,2,3,4,5,6,7,8}
	fmt.Println(arr3)

	var arr4 = [...] string {"gautam", "hello", "garima", "govil", "aggarwal"}
	fmt.Println(arr4)
}