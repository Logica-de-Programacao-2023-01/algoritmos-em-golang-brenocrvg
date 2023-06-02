package main

import "fmt"

func main() {
	//Faça um algoritmo que leia vários números inteiros e mostre o maior deles. A leitura deve ser interrompida quando for lido o valor zero.
	var numero, maior int

	fmt.Println("digite os números (0 para): ")

	for {
		fmt.Println("numero: ")
		fmt.Scan(&numero)
		if numero == 0 {
			break
		}
		if numero > maior {
			maior = numero
		}

	}
	if maior != 0 {
		fmt.Println("o maior número é: ", maior)

	} else {
		fmt.Println("nada")
	}
}
