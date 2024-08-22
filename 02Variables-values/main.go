package main
import "fmt"

func main() {
	a := 45      // := ---> short declaration operator
	fmt.Println(a)

	var b bool = true
	fmt.Println(b)

	d,e,f,_,g := 15,12.3,"hello",10,false
	fmt.Println(d,e,f,g)

	/*Default values of different data types*/

	var x int
	fmt.Println(x) //0
	var y bool
	fmt.Println(y) //false
	var z float64
	fmt.Println(z) //0
	var w string
	fmt.Println(w) //""


	/*This will not work because we initialize four varibles m,n,o,p but we use only 
	three variables i.e m,n,o it throws compile time error*/

	// m,n,o,p := 15,12.3,"hello",false
	// fmt.Println(m,n,o)
}