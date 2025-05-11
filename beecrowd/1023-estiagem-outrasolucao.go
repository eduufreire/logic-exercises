package main

import "fmt"

// solução usada com base
// https://github.com/xTecna/solucoes-da-beecrowd/blob/main/problemas/estruturas-e-bibliotecas/1023/README.md
// o programador usou um método onde ele usa o indice do vetor para conseguir ordenar o consumo em ordem crescente
// para evitar mais um for de ordenação
// utiliza a media retirada de consumo/morades, onde é gerado um valor inteiro e é usado como indice, colocando os moradores como valor
// isso foi feito porque no exercicio pede para exibir a quantidadeMorades-consumoMedio em ordem crescente
// como o enunciado determina o valor maximo de consumo é 201, permite que esse método seja utilizado


func main() {
	var cidade, imoveis int = 0, -1
	var consumos [201]int

	for {
		moradores, consumo, totalMoradores, totalConsumo := 0, 0, 0, 0
		fmt.Scan(&imoveis)
		if imoveis == 0 {
			break
		}

		cidade += 1
		for i := 0; i < imoveis; i++ {
			fmt.Scan(&moradores, &consumo)
			totalMoradores += moradores
			totalConsumo += consumo
			consumos[consumo/moradores] = moradores
		}

		consumoMedioCidade := float64(totalConsumo) / float64(totalMoradores)

		fmt.Printf("Cidade# %d:\n", cidade)
		for key, value := range consumos {
			if value > 0 {
				fmt.Printf("%d-%d ", value, key)
			}
		}
		fmt.Printf("\nConsumo medio: %.2f m3.\n", consumoMedioCidade)
	}

}
