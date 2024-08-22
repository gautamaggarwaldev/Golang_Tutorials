package main

import "fmt"

func main() {
	i := 1
	for i<=5 {
		fmt.Print(i,"\t")
		i++
	}
	fmt.Printf("\n")
	for j := 1; j <= 10; j++ {
		fmt.Printf("%v\t", j)
	}
	fmt.Printf("\n")

	for i := range 10 {
		fmt.Println("range", i)
	}

/*
	infinite loop
	for {
		fmt.Print(2)
	}
*/
	
}