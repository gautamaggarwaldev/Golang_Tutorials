/*
	Go Structures:---

	* A struct (short for structure) is used to create a collection of members of different data types, into a 
	single variable.
	* While arrays are used to store multiple values of the same data type into a single variable, structs are used 
	to store multiple values of different data types into a single variable.
	* A struct can be useful for grouping data together to create records.

	SYNTAX :---
	type struct_name struct {
		member1 datatype;
		member2 datatype;
		member3 datatype;
		so on...
	}
*/

package main
import "fmt"
type Employee struct {
	name string
	Id int
	salary int
}
func main() {
	var e1 Employee
	var e2 Employee

	e1.name = "Gautam"
	e1.Id = 45
	e1.salary = 74123698520

	e2.name = "Garima"
	e2.Id = 18
	e2.salary = 96321478520


	fmt.Printf("Name of Ist employe : %v\n",e1.name)
	fmt.Printf("ID of Ist employe : %v\n",e1.Id)
	fmt.Printf("Salary of Ist employe : %v\n",e1.salary)

	fmt.Printf("Name of 2nd employe : %v\n",e2.name)
	fmt.Printf("ID of 2nd employe : %v\n",e2.Id)
	fmt.Printf("Salary of 2nd employe : %v\n",e2.salary)
}