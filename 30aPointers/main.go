package main

import "fmt"

func main() {
	x := 15
	fmt.Println(x);
	fmt.Println(&x);

	fmt.Printf("%v\t %T\n", &x, &x); // value --> 0xc000090068    type --> *int

	var s string = "hello World!";
	fmt.Printf("%v\t %T\n", &s, &s); // type ---> *string

	var b bool = true;
	fmt.Printf("%v\t %T\n", &b, &b);  //type ---> *bool
}