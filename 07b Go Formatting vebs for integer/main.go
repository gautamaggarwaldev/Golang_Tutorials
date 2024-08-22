package main
import "fmt"
func main() {
	i := 47469

	fmt.Printf("%b\n", i) //binary format
	fmt.Printf("%d\n", i) //decimal format
	fmt.Printf("%+d\n", i) //show sign with number
	fmt.Printf("%o\n", i) //octal format
	fmt.Printf("%O\n", i) //octal format with 0o
	fmt.Printf("%x\n", i) //hex format, lowercase
	fmt.Printf("%X\n", i) //hex format with uppercase
	fmt.Printf("%#x\n", i) //hex fromat with 0x
	fmt.Printf("%8d\n", i) // pad with spaces
	fmt.Printf("%-8d\n", i)
	fmt.Printf("%08d\n", i) //pad with zero
}