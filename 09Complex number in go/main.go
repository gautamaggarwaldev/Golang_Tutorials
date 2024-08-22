package main
import "fmt"
func main() {
	c1 := complex(10,11)
	fmt.Println(c1) // (10+11i)

	realPart := real(152)
	imagPart := imag(14i)

	fmt.Println(realPart,imagPart)

	c := complex(5,6)
	d := complex(7,3)
	e := c+d
	fmt.Println(e)

	var a complex64
	var realpart float32
	var imgpart float32
	a = 12 + 15i
	realpart = real(a)
	imgpart = imag(a)
	fmt.Println(a)
	fmt.Println(realpart)
	fmt.Println(imgpart)
}