package main

import "fmt"

func main() {
	var inputs, sortedInputs [3]int
	fmt.Scan(&inputs[0], &inputs[1], &inputs[2])

	sortedInputs = inputs
	for i := 0; i < len(inputs); i++ {
		for j := 0; j < len(inputs)-1; j++ {
			if sortedInputs[j] > sortedInputs[j+1] {
				sortedInputs[j], sortedInputs[j+1] = sortedInputs[j+1], sortedInputs[j]
			}
		}
	}

	for _, value := range sortedInputs {
		fmt.Println(value)
	}
	fmt.Print("\n")
	for _, value := range inputs {
		fmt.Println(value)
	}
}
