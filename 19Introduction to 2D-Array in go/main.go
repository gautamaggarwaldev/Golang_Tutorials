package main
import "fmt"
func main() {
	arr := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	fmt.Println(arr)

	arr1 := [5][2]int{{0,0},{1,0},{2,0},{3,0},{4,0}}
	fmt.Println(arr1)

	var arr3 [3][3]int
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			fmt.Scanf("%d", &arr3[i][j])
		}
	}
	fmt.Println(arr3)
}