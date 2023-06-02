package main

import "fmt"

func main() {
	//Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles, se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.
	var x, y int
	fmt.Println("digite o primeiro número: ")
	fmt.Scan(&x)
	fmt.Println("digite o segundo número: ")
	fmt.Scan(&y)
	if x < 0 || y < 0 {
		fmt.Println("a soma deles é: ", x+y)
	} else if x > 0 && y > 0 {
		fmt.Println("o resultado da multiplicação dos números é: ", x*y)
	}

}
