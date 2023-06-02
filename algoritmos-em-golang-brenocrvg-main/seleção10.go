package main

import "fmt"

func main() {
	//Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação, de acordo com a seguinte tabela:
	//até 9 anos: mirim
	//10 a 13 anos: infantil
	//14 a 17 anos: juvenil
	//maiores de 18 anos: adulto
	var idade int
	fmt.Println("digite sua idade: ")
	fmt.Scan(&idade)
	if idade < 10 {
		fmt.Println("mirim")
	} else if idade >= 10 && idade < 14 {
		fmt.Println("infantil")
	} else if idade >= 14 && idade < 18 {
		fmt.Println("juvenil")
	} else {
		fmt.Println("adulto")
	}

}
