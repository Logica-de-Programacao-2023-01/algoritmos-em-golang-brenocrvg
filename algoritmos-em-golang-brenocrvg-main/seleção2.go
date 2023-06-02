package main

import "fmt"

func main() {
	//Faça um algoritmo que leia três números inteiros e mostre o menor deles.
	var x, y, z int
	fmt.Println("digite o primeiro número: ")
	fmt.Scan(&x)
	fmt.Println("digite o segundo número: ")
	fmt.Scan(&y)
	fmt.Println("digite o terceiro número: ")
	fmt.Scan(&z)
	if x < y && x < z {
		fmt.Println("o menor é: ", x)
	} else if y < x && y < z {
		fmt.Println("o menor é: ", y)
	} else if z < x && z < y {
		fmt.Println("o menor é: ", z)
	}
}
