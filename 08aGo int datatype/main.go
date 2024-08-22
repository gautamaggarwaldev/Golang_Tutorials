package main
import (
	"fmt"
	"math"
)
func main() {	
	/*signed integer*/
	var a int = 500
	var b int = -8500

	fmt.Printf("value = %v and type = %T\n",a, a)
	fmt.Printf("value = %v and type = %T\n",b, b)

	var c int16 = 12365
	fmt.Printf("value = %v and type = %T\n",c,c)


	/*unsigned integer*/
	var x uint = 2500
	var y uint = 7500

	fmt.Printf("value = %v and type = %T\n",x,x)
	fmt.Printf("value = %v and type = %T\n",y,y)

	var z uint64 = 7418529633217469954
	fmt.Printf("value = %v and type = %T\n",z,z)

	fmt.Printf("\n\n\n\n")


	/*min-max value of int and uint*/

	fmt.Printf("int min and max value =  %d and %d\n", math.MinInt, math.MaxInt)  /*64-bit OS*/
	fmt.Printf("int8 min and max value =  %d and %d\n", math.MinInt8, math.MaxInt8)  
	fmt.Printf("int16 min and max value =  %d and %d\n", math.MinInt16, math.MaxInt16)  
	fmt.Printf("int32 min and max value =  %d and %d\n", math.MinInt32, math.MaxInt32)  
	fmt.Printf("int64 min and max value =  %d and %d\n", math.MinInt64, math.MaxInt64) 
	
	fmt.Printf("\n\n")

	fmt.Printf("uint min and max value =  %d and %d\n", uint(0), uint(math.MaxUint))  /*64-bit OS*/
	fmt.Printf("uint8 min and max value =  %d and %d\n", uint(0), uint(math.MaxUint8))  
	fmt.Printf("uint16 min and max value =  %d and %d\n", uint(0), uint(math.MaxUint16))  
	fmt.Printf("uint32 min and max value =  %d and %d\n", uint(0), uint(math.MaxUint32))  
	fmt.Printf("uint64 min and max value =  %d and %d\n", uint(0), uint(math.MaxUint64))  



}