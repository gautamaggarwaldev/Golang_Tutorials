package main
import "fmt"
func main() {
/*
	Slices are similar to arrays, but are more powerful and flexible.

	Like arrays, slices are also used to store multiple values of the same type in a single variable.

	However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	In Go, there are several ways to create a slice:

	1. Using the []datatype{values} format
	2. Create a slice from an array
	3. Using the make() function

*/

	myslice := []int{}
	fmt.Println(len(myslice)) //0
	fmt.Println(cap(myslice)) //0
	fmt.Println(myslice) //[]

	slice := [5]int{7,3,4,8,9}
	fmt.Println(len(slice)) //5
	fmt.Println(cap(slice)) //5
	fmt.Println(slice) //[7 3 4 8 9]



	/*Create a Slice From an Array
	
		Syntax  var myarray = [length]datatype{values} An array
		myslice := myarray[start:end] A slice made from the array
	*/

	arr := [7]int{7,5,3,6,9,1,0}
	slice1 := arr[2:5] // arr[n:m] ---> it print the value from nth index to (m-1)th index 
	fmt.Printf("value = %d\n", slice1)
	fmt.Printf("capacity = %v\n", cap(slice1)) //The slice can grow to the end of the array.
	fmt.Printf("length = %v\n", len(slice1))


/*
	Create a Slice With The make() Function
	slice_name := make([]type, length, capacity)
	Note: If the capacity parameter is not defined, it will be equal to length.
*/


myslice3 := make([]int, 5, 10)
fmt.Printf("myslice1 = %v\n", myslice3)
fmt.Printf("length = %d\n", len(myslice3))
fmt.Printf("capacity = %d\n", cap(myslice3))

// with omitted capacity
myslice4 := make([]int, 5)
fmt.Printf("myslice2 = %v\n", myslice4)
fmt.Printf("length = %d\n", len(myslice4))
fmt.Printf("capacity = %d\n", cap(myslice4))

}