package main

import "fmt"

func main() {
	//Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.
	var x int
	fmt.Println("digite um número inteiro positivo: ")
	fmt.Scan(&x)

	fmt.Println("divisores de ", x, ":")

	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Println(i)
		}
	}

}
