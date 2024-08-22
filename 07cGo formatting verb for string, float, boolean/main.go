package main
import "fmt"
func main() {
	str := "hello world!"
	flag := true
	i := 85.6957896
	/*String*/
	fmt.Printf("%s\n", str)
	fmt.Printf("%q\n", str)
	fmt.Printf("%32s\n", str)
	fmt.Printf("%-32s\n", str)
	fmt.Printf("%x\n", str)
	fmt.Printf("% x\n", str)

	/*boolean*/
	fmt.Printf("%t\n", flag)

	/*Float*/
	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%20.4f\n", i) // width 20, precision 4
	fmt.Printf("%g\n", i)
}