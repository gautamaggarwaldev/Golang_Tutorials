package main
import "fmt"
func main() {
	x := 10
	y := 7

	fmt.Println(x>y) //true
	fmt.Println(x<y) //false

	z := 10

	fmt.Println(x==z)//true
	fmt.Println(x!=y)//true
	fmt.Println(x!=z)//false

	fmt.Println(z>=y)//true
	fmt.Println(z<=y)//false

}