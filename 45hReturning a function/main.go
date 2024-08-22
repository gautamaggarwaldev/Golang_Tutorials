package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x);

	y := bar();
	fmt.Println(y());

	fmt.Printf("%T\n",foo); //func() int
	fmt.Printf("%T\n",bar); //func() func() int
	fmt.Printf("%T\n",y); //func() int

}

func foo() int {
	return 42;
}

func bar() func() int {
	return func() int {
		return 74;
	}
}

