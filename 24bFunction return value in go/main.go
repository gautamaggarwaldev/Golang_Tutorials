package main
import "fmt"
func add(x int, y int) int{
	return x+y
}

func sub(a int, b int) (result int) {
	result = a-b
	return result
}

func myfun(x int, y string)(ans int, str string){
	ans = x + x
	str = y + "world"
	return
}

func val() (int, int){
	return 3,4
}

func main() {
	fmt.Println(add(2,3))
	fmt.Println(sub(10,4))
	fmt.Println(myfun(10,"hello"))
	a,b := myfun(30,"good")
	fmt.Println(a,b)
	_,u := myfun(30,"good")
	w,_ := myfun(30,"good")
	fmt.Println(u)
	fmt.Println(w)
	fmt.Println(val())
}