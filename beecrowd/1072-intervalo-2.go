package main

import "fmt"

func main() {
	var n, in, out, value int
	fmt.Scan(&n)

	for x := 0; x < n; x++ {
		fmt.Scan(&value)
		if value >= 10 && value <= 20{
			in++
		} else {
			out++
		}
	}

	fmt.Print(in, " in\n")
	fmt.Print(out, " out\n")
}
