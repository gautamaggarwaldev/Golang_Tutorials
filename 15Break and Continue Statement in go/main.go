package main
import "fmt"
func main() {
	/*Continue Statement = The continue statement is used to skip one or more iterations in the loop. 
	It then continues with the next iteration in the loop.*/

	for i:=0; i<=20; i+=2 {
		if i == 16 {
			continue
		}
		fmt.Print(i," ")
	}

	fmt.Println()

	/*Break Statement = The break statement is used to break/terminate the loop execution.*/

	for j:=0; j<=10; j++ {
		if j == 8 {
			break
		}
		fmt.Print(j," ")
	}
}