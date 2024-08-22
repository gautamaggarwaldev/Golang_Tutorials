package main
import "fmt"
func main() {
	var MonthNum int
	fmt.Print("Enter the Month Number : ")
	fmt.Scan(&MonthNum)

	switch MonthNum {
	case 1,3,5,7,10,12 : 
		fmt.Print("These months have 31 days")
	case 4,6,8,9,11 :
		fmt.Print("These months have 30 days")
	case 2 :
		fmt.Print("This months have 28 days")
	default : fmt.Print("Invalid Number")
	}
	


}