package main
import "fmt"
func main() {
	var num int = 5
	fmt.Println("Value :",num)//5
	fmt.Println("Memory Address :",&num) //0xc00000a0b8
	
	/*Create Pointer*/
	var ptr *int = &num
	fmt.Println("Memory Address :",ptr) //0xc00000a0b8
	fmt.Println("Memory Address :",*ptr) //5

	val := 20
	fmt.Printf("Value of val valriable : %v\n",val)
	fmt.Printf("Address of val valriable : %v\n",&val)
	var ptr1 *int = &val
	fmt.Printf("Value of *ptr1 valriable OR Address of val : %v\n",ptr1)
	fmt.Printf("Address of ptr1 : %v\n",&ptr1)
	fmt.Printf("value of val : %v\n",*ptr1)

}