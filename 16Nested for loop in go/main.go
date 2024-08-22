package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	
	fmt.Println()
	fmt.Println()

	for k:=5; k>=0; k-- {
		for l:=1; l<=k; l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
