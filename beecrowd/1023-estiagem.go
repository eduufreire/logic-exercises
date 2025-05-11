package main

import "fmt"

func main() {
	var cidade, imoveis int = 0, -1

	for {
		moradores, consumo, totalMoradores, totalConsumo :=  0, 0, 0, 0
		fmt.Scan(&imoveis)
		if imoveis == 0 {
			break
		}

		cidade += 1
		quantidadeMoradores, consumoMedioImovel := make([]int, imoveis), make([]int, imoveis)
		for i := 0; i < imoveis; i++ {
			fmt.Scan(&moradores, &consumo)

			mediaConsumo := consumo / moradores
			quantidadeMoradores[i], consumoMedioImovel[i] = moradores, mediaConsumo
			totalMoradores += moradores
			totalConsumo += consumo
		}

		for j := 0; j < imoveis; j++ {
			for k := 0; k < imoveis-1; k++ {
				if consumoMedioImovel[k] > consumoMedioImovel[k+1] {
					consumoMedioImovel[k], consumoMedioImovel[k+1] = consumoMedioImovel[k+1], consumoMedioImovel[k]
					quantidadeMoradores[k], quantidadeMoradores[k+1] = quantidadeMoradores[k+1], quantidadeMoradores[k]
				}
			}
		}

		consumoMedioCidade := float64(totalConsumo) / float64(totalMoradores)

		fmt.Printf("Cidade# %d:\n", cidade)
		for j := 0; j < imoveis; j++ {
			fmt.Printf("%d-%d ", quantidadeMoradores[j], consumoMedioImovel[j])
		}
		fmt.Printf("\nConsumo medio: %.2f m3.\n", consumoMedioCidade)
	}

}
