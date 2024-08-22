package main
import "fmt"
func main() {
/*
	Maps Are References :--
	* Maps are references to hash tables.
	* If two map variables refer to the same hash table, changing the content of one variable affect the content
	of the other.
*/

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Printf("\nAfter change to b:\n\n")

	fmt.Println(a)
	fmt.Println(b)
}