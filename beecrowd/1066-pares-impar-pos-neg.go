package main

import "fmt"

func main() {

	values := [5]int{}
	for i := 0; i < 5; i++ {
		fmt.Scan(&values[i])
	}

	var pairs, odd, positives, negatives int
	for _, value := range values {
		if value % 2 == 0 {
			pairs += 1
		} else {
			odd += 1
		}

		if value < 0 {
			negatives += 1
		} else if value > 0{
			positives += 1
		}
	}

	fmt.Printf("%d valor(es) par(es)\n", pairs)
	fmt.Printf("%d valor(es) impar(es)\n", odd)
	fmt.Printf("%d valor(es) positivo(s)\n", positives)
	fmt.Printf("%d valor(es) negativo(s)\n", negatives)
}