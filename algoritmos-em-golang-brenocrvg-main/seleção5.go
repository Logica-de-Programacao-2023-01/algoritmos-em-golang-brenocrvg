package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.
	var x int
	fmt.Println("digite um número: ")
	fmt.Scan(&x)
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("o número é múltiplo de 5 e 3")
	} else {
		fmt.Println("não é múltiplo de 3 e 5")
	}

}
