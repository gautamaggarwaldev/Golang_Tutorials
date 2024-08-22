package main

import "fmt"

func main() {
	
	a := 5.23
	fmt.Printf("%X \n",a) //uppercase hexadecimal notation
	fmt.Printf("%x \n",a) //lowercase hexadecimal notation

	b := 15

	fmt.Printf("%#b \n", b)  //binary 

	c := 15
	fmt.Printf("%#o \n", c) //octal


	s := "Hello"

	fmt.Printf("%q\n", s)

	y := 647

	fmt.Printf("%#x\n", y)
	fmt.Printf("%#X\n", y)

}
