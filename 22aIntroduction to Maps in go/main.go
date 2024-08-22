package main
import (
	"fmt"
	
)
func main() {
/*
	1.Maps are used to store data values in key:value pairs.
	2.Each element in a map is a key:value pair.
	3.A map is an unordered and changeable collection that does not allow duplicates.
	4.The length of a map is the number of its elements. You can find it using the len() function.
	5.The default value of a map is nil.
	6.Maps hold references to an underlying hash table.

	* SYNTAX ----> var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
				   b := map[KeyType]ValueType{key1:value1, key2:value2,...}
*/

	var a = map[string]string{"brand" : "ford", "model" : "mustang", "year" : "1947"}   // map[Key]Value
	b := map[string]int {"Ice-cream" : 50, "Pizza" : 150, "Burger" : 70, "Soft-drink" : 40}
	fmt.Println(a)
	fmt.Println(b)


/*  Create map using make() function
	SYNTAX ---> var a = make(map[KeyType]ValueType)
				b := make(map[KeyType]ValueType)
*/

	var map1 = make(map[string]bool)

	map1["yes"] = true
	map1["no"] = false
	map1["wrong"] = false
	map1["right"] = true
	map1["correct"] = true
	map1["done"] = true

	fmt.Printf("%v\n", map1)


/*	Create an empty map :--
	Note: The make()function is the right way to create an empty map. If you make an empty map in a 
	different way and write to it, it will causes a runtime panic.
*/
	var map2 = make(map[string]string)
	var map3 map[string]int
	fmt.Println(map2) // map[]
	fmt.Println(map2 == nil) // false
	fmt.Println(map3) // map[]
	fmt.Println(map3 == nil) // true


/*
	Allowed Key Types in map :--

	The map key can be of any data type for which the equality operator (==) is defined. These include:

	* Booleans
	* Numbers
	* Strings
	* Arrays
	* Pointers
	* Structs
	* Interfaces (as long as the dynamic type supports equality)
	* Invalid key types are:
	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	* Slices
	* Maps
	* Functions
	These types are invalid because the equality operator (==) is not defined for them.
*/

	


}