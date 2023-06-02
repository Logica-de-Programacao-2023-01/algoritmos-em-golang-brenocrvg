package main

import "fmt"

func main() {
	var peso, altura, IMC float64

	fmt.Println("digite sua altura: ")
	fmt.Scan(&altura)
	fmt.Println("digite seu peso: ")
	fmt.Scan(&peso)
	fmt.Println("seu IMC é igual a: ", peso/(altura*altura))

	if IMC < 18.5 {
		fmt.Println("magreza")
	} else if IMC >= 18.5 && IMC <= 24.9 {
		fmt.Println("normal")
	} else {
		fmt.Println("obesidade")
	}

}
