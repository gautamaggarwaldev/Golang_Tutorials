/*Go programming provides a pretty simple error handling framework with inbuilt error interface type of the following declaration −

type error interface {
   Error() string
}

*/

package main

import (
	"errors"
	"fmt"
	"math"
)
func main() {

	res, err := Sqrt(-1);

	if(err != nil) {
		fmt.Println(err);
	} else {
		fmt.Println(res);
	}


	res, err = Sqrt(50);

	if(err != nil) {
		fmt.Println(err);
	} else {
		fmt.Println(res);
	}

} 

func Sqrt(num float64) (float64, error) {
	if(num < 0) {
		return 0, errors.New("Error: negative number sqrt passed");
	}
	return math.Sqrt(num), nil;
}