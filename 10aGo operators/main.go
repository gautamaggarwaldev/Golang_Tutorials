package main
import ("fmt")
func main() {
	a := 15 + 15

	var (
		sum1 = 125 + 63
		sum2 = sum1 + 489
		sum3 = sum2 + sum1
	)

	fmt.Println(a)

	fmt.Println(sum1)
	fmt.Println(sum2)
	fmt.Println(sum3)

}