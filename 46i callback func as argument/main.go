package main

import "fmt"

func main() {
	x := math(56, 40, add)
	y := math(56, 40, subtract)

	fmt.Println(x);
	fmt.Println(y);

	fmt.Printf("%T\n",add);
	fmt.Printf("%T\n",subtract);
	fmt.Printf("%T\n",math);
}

func math(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}