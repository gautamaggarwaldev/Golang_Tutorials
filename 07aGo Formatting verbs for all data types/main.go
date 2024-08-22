package main
import "fmt"
func main() {
	var i int = 10
	var str = "hello world"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%v%%\n", i) //use to pritn % sign

	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Printf("%T\n", str)  // print the type of data
}