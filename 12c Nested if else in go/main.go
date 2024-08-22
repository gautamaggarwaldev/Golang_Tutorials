package main

import "fmt"

func main() {
	var num = 53

	if num > 50 {
		fmt.Println("value of num is greater than 50")
		if num%2==0 {
			fmt.Println("num is also divided by 2")
		} else {
			fmt.Println("num is not divided by 2")
		}
	} else {
		fmt.Println("value is less than 50")
	}

}