package main

import "fmt"

func main() {

	treeAnimals := map[string]map[string]map[string]string{
		"vertebrado": {
			"ave": {
				"carnivoro": "aguia",
				"onivoro":   "pomba",
			},
			"mamifero": {
				"onivoro":   "homem",
				"herbivoro": "vaca",
			},
		},
		"invertebrado": {
			"inseto": {
				"hematofago": "pulga",
				"herbivoro":  "lagarta",
			},
			"anelideo": {
				"hematofago": "sanguessuga",
				"onivoro":    "minhoca",
			},
		},
	}

	var primeira, segunda, terceira string
	fmt.Scan(&primeira, &segunda, &terceira)

	result1 := treeAnimals[primeira]
	result2 := result1[segunda]
	result3 := result2[terceira]

	fmt.Println(result3)
}
