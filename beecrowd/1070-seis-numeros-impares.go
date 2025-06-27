package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	count := 0
	for count < 6 {
		if x % 2 != 0 {
			fmt.Println(x)
			count += 1
		}
		x += 1
	}
}