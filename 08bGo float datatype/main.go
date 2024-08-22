package main
import (
	"fmt" 
	"math"
)
func main() {
	var a float32 = 1.2e+12
	var b float64 = 1.5e+78
	fmt.Printf("value = %f and type = %T\n",a,a)
	fmt.Printf("value = %f and type = %T\n",b,b)

	fmt.Printf("\n\n")

	/*max and min value of float32 and float64*/

	fmt.Printf("min float64: %.50e\n", math.SmallestNonzeroFloat64)
    fmt.Printf("max float64: %.50e\n", math.MaxFloat64)

    fmt.Printf("min float32: %.50e\n", math.SmallestNonzeroFloat32)
    fmt.Printf("max float32: %.50e\n", math.MaxFloat32)
}