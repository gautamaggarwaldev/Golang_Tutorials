package main

import (
	"fmt"
)
func main() {
	arr := [5]int{}
	fmt.Println("Enter the elements : ")
	for i:=0; i<5; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println(arr)
	
}