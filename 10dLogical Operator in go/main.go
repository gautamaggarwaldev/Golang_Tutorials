package main
import "fmt"
func main() {
	/* &&  ||  ! */

	fmt.Println((7<8)&&(9>3)) //true because both statement are true 
	fmt.Println((7<8)&&(9<3)) //false because one statement is true and other is false

	fmt.Println((7<8)||(9<3)) //true because one statement is true

	fmt.Println(!(7<8)) //false because the statement is true

	/* ! ---> this operator makes true to false and false to true*/
}