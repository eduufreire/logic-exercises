package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scan(&n)

	if n < 5 || n > 2000 {
		fmt.Errorf("invalid")
	}

	for i := 1; i < n; i++ {
		if i % 2 == 0 {
			fmt.Printf("%d^2 = %d\n", i, int(math.Pow(float64(i), 2)))
		}
	}
}