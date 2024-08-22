package main
import "fmt"
func main() {
	/*Access map elements*/
	var map1 = make(map[string]string)
	map1["Company"] = "BMW"
	map1["Model"] = "Maybatch"
	map1["Engine"] = "15000CC"

	fmt.Println(map1["Company"])
	fmt.Println(map1["Model"])
	fmt.Println(map1["Engine"])

	/*Update and Add Map Elements*/
	var map2 = make(map[string]string)
	map2["Company"] = "Audi"
	map2["Model"] = "Oracel 360"
	map2["Engine"] = "12000CC"
	fmt.Println(map2)

	map2["Engine"] = "16000CC"
	map2["Company"] = "Faraari"
	map2["color"] = "black" // new pair add
	fmt.Println(map2)

	/*Remove element from map*/
	fmt.Println(map1) //map[Company:BMW Engine:15000CC Model:Maybatch]
	delete(map1,"Model")
	fmt.Println(map1) //map[Company:BMW Engine:15000CC]


	/*Check For Specific Elements in a Map*/
	var map3 = map[string]int{"A":65,"B":66,"C":67,"D":68,"E":69,"F":70,"G":71,"H":72}
	val1, ok1 := map3["A"]//Checking for existing key and its value
	val2, ok2 := map3["B"]//Checking for existing key and its value
	val3, ok3 := map3["C"]//Checking for existing key and its value
	val4, ok4 := map3["I"]//Checking for non-existing key and its value
	val5, ok5 := map3["E"]//Checking for existing key and its value
	_, ok6 := map3["H"]// Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(val4, ok4)
	fmt.Println(val5, ok5)
	fmt.Println(ok6)

}