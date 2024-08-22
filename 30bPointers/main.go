package main

import "fmt"

func main() {
	x := 56
	fmt.Println(x);
	fmt.Println(&x);

	y := &x;
	// y ---> var y *int
	fmt.Printf("%v\t%T\n", y, y);
	fmt.Println(*y); //56
	fmt.Println(*&x); //56

	*y = 10;
	fmt.Println(x); //10
	fmt.Println(*y);//10
	fmt.Println(y); //memory addrress


	var a int;
	var z *int = &a;
	fmt.Println(z);

}