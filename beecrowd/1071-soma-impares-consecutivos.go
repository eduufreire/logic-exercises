package main

import (
	"fmt"
)

func main() {

	var x, y int
	fmt.Scan(&x, &y)

	if y < x {
		x, y = y, x
	}

	impares := make([]int, 0)
	for i := x + 1; i < y; i++ {
		if i % 2 != 0 {
			impares = append(impares, i)
		}
	}

	sum := 0
	for _, value := range impares {
		sum = value + sum 
	}

	fmt.Println(sum)
}
