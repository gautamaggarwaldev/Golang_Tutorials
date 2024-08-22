package main
import "fmt"
func main() {
	/*Access Elements of a Slice*/
	slice := []int{10,20,30,40,50,60,70,80}

	fmt.Println(slice[0])
	fmt.Println(slice[2])

	fmt.Println(slice[5])
	slice[5] = 100
	fmt.Println(slice[5])
	fmt.Println("Length = ",len(slice))
	fmt.Println("Capacity = ",cap(slice))

	/*Append Elements To a Slice*/

	slice = append(slice, 201, 501)
	fmt.Println(slice)
	fmt.Println("Length = ",len(slice))
	fmt.Println("Capacity = ",cap(slice))


	/*Append One Slice To Another Slice*/
	myslice1 := []int{1,2,3}
	myslice2 := []int{4,5,6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Println(myslice3)


	/*Change the length of a slice*/
	arr := [10]int{7,6,3,1,9,65,4,2,0,4}
	slice1 := arr[1:9]
	fmt.Println("slice = ", slice1)
	fmt.Println("capacity = ", cap(slice1))
	fmt.Println("length = ", len(slice1))

	slice1 = arr[1:5]
	fmt.Println("slice = ", slice1)
	fmt.Println("capacity = ", cap(slice1))
	fmt.Println("length = ", len(slice1))

	slice1 = append(slice, 6,3,0,2,8,9,4)
	fmt.Println("slice = ", slice1)
	fmt.Println("capacity = ", cap(slice1))
	fmt.Println("length = ", len(slice1))


/*
	When using slices, Go loads all the underlying elements into the memory.
	If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
	The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the 
	memory used for the program. 

	Syntax ---> copy(dest, src)
*/


	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
  
	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)
  
	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))

}