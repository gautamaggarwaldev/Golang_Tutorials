package main
import "fmt"
func main() {

	/* Access the particular element from array */

	age := [...]int{10,20,30,40,50,60}
	fmt.Println(age[0])
	fmt.Println(age[3])
	fmt.Println(age[4])
	fmt.Println(age[5])

	age[0] = 100
	fmt.Println(age[0])

	salary := [...]int{}
	fmt.Println(salary)  // []

	person := [5]int{}
	fmt.Println(person) //[0 0 0 0 0]

	array := [6]int{1,2,3}
	fmt.Println(array) // [1 2 3 0 0 0]


	/*Initialize Only Specific Elements in the Array*/
	garima := [...]int{5:50, 6:60, 1:10, 3:30, 10:100}
	fmt.Println(garima)  // [0 10 0 30 0 50 60 0 0 0 100]


	name := [6]int{1,2,3,4,5,6}
	cast := [...]int{12,6,3,459,5,2,6,97,41,6,20,6}

	fmt.Println(len(name)) //6
	fmt.Println(len(cast)) //12


	arr := []int{}
	fmt.Println(arr) // []
	fmt.Println(len(arr)) // 0
}