package main

import "fmt"

func main() {
	//Faça um algoritmo que imprima a tabuada de multiplicação de 1 a 10 para um número fornecido pelo usuário.
	var numero int

	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numero)

	fmt.Println("Tabuada de multiplicação para o número", numero)
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}
