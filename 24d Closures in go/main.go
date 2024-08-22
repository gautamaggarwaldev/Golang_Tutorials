package main
import "fmt"

func new_my_fun() func() int {
	i := 0
	return func() int {
		i+=1
		return i
	}
}

func main() {

	/*Method 2 to create a closure*/
	my_fun := new_my_fun()
	
	fmt.Println(my_fun())
	fmt.Println(my_fun())


	/*Method 1 to create a closure*/
	i := 0

	counter := func() int {
		i += 1
		return i
	}
	fmt.Println(counter()) //1
	fmt.Println(counter()) //2

}