package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro e mostre se ele é par ou ímpar.
	var x int
	fmt.Println("digite um número: ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("o número é par")
	} else {
		fmt.Println("o número é ímpar")
	}

}
