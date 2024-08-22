/* function is recursive if it calls itself and reaches a stop condition.*/
package main
import "fmt"
func myfunc(a int) int {
	if a == 14 {
		return 0
	}
	fmt.Println(a)
	return myfunc(a+1)
}


func fibo(x int) int {
	if (x==0) {
		return 0
	} else if (x==1) {
		return 1
	} else{
		return fibo(x-1) + fibo(x-2)
	} 
}

func factorial(y int) int {
	if(y==0 || y==1) {
		return 1
	} else {
		return y * factorial(y-1)
	}
}

/*calculate the sum of positive numbers*/

func calculate(z int) int {
	if z == 0 {
		return 0
	} else {
		return z + calculate(z-1)
	}
}

func main() {
	myfunc(1)
	
	fmt.Println("Fibonnacci")
	for i:=0; i<=10; i++ {
		fmt.Println(fibo(i))
	}

	fmt.Println("Factorial : ",factorial(6))
	fmt.Println("sum of first",50,"numbers are : ",calculate(50))
}