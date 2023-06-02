package main

import "fmt"

func main() {
	//Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 10% se o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.
	var salario float64
	fmt.Println("digite o seu salário: ")
	fmt.Scan(&salario)
	if salario < 1000 {
		fmt.Println("seu salário com aumento de 10% é: ", salario+(salario*0.1))
	} else if salario >= 1000 {
		fmt.Println("seu salário com aumento de 5% é: ", salario+(salario*0.05))
	}
}
