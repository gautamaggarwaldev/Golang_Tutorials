package main

import "fmt"

func main() {

	f := incremeanter()
	fmt.Println(f());
	fmt.Println(f());
	fmt.Println(f());
	fmt.Println(f());
	fmt.Println(f());
	fmt.Println(f());
	fmt.Println(f());
	
	g := incremeanter()
	fmt.Println(g());
	fmt.Println(g());
	fmt.Println(g());
	fmt.Println(g());
	fmt.Println(g());
	fmt.Println(g());
	fmt.Println(g());


}

func incremeanter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}