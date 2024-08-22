package main

import "fmt"

func main() {
	var DayNum int
	fmt.Print("Enter the Day Number : ")
	fmt.Scan(&DayNum)

	/*All the case values should have the same type as the switch expression. Otherwise, the compiler will raise an error*/

	switch DayNum {
		case 1: fmt.Print("Monday")
		case 2: fmt.Print("Tuesday")
		case 3: fmt.Print("Wednesday")
		case 4: fmt.Print("Thrusday")
		case 5: fmt.Print("Friday")
		case 6: fmt.Print("Saturday")
		case 7: fmt.Print("Sunday")
		default : fmt.Print("Invalide Day Number !")
	}
}