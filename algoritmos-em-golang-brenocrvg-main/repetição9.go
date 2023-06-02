package main

import "fmt"

func main() {
	//Faça um algoritmo que leia vários números inteiros e mostre a média aritmética entre eles. A leitura deve ser interrompida quando for lido o valor zero.
	var numero, soma, quantnum float64
	fmt.Println("digite os números (0 para): ")

	for {
		fmt.Println("numero: ")
		fmt.Scan(&numero)

		if numero == 0 {
			break
		}
		soma += numero
		quantnum++
	}
	if quantnum > 0 {
		media := float64(soma) / quantnum
		fmt.Printf("media aritmética: %.2f\n", media)

	} else {
		fmt.Println("nada")
	}

}
