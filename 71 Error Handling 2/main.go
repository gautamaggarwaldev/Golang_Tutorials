package main

import (
	"errors"
	"fmt"
)

func main() {
	ans, err := sumOf(11, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}


	ans, err = sumOf(1, 100)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}

func sumOf(start int, end int) (int, error) {
	total := 0
	if start > end {
		return 0, errors.New("ERROR: Not Possible start should always less than end")
	} else {
		for i := start; i <= end; i++ {
			total += i
		}
		return total, nil
	}

}
