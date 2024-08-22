package main
import "fmt"
func main() {
	
	var i,j,k int 
	fmt.Print("Enter first number : ")
	fmt.Scan(&i)
	fmt.Println("Your number is : ", i)

	fmt.Println("Enter second and third number : ")
	fmt.Scan(&j, &k)
	fmt.Println("Your number is : ", j," and ",k)
	
	
	var a, b int
	
	fmt.Scanln(&a, &b)
	fmt.Println("Your number is : ", a," and ",b)

	var m, n string
	fmt.Println("Enter the name : ")
	fmt.Scanf("%v %v", &m,&n)
	fmt.Println("My name is : ", m, " ", n)

}