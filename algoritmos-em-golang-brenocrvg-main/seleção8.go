package main

import "fmt"

func main() {
	var numero int
	fmt.Println("digite um número de 1 a 7: ")
	fmt.Scan(&numero)
	if numero == 1 {
		fmt.Println("domingo")
	} else if numero == 2 {
		fmt.Println("segunda-feira")
	} else if numero == 3 {
		fmt.Println("terça-feira")
	} else if numero == 4 {
		fmt.Println("quarta-feira")
	} else if numero == 5 {
		fmt.Println("quinta-feira")
	} else if numero == 6 {
		fmt.Println("sexta-feira")
	} else if numero == 7 {
		fmt.Println("sábado")
	} else {
		fmt.Println("ta de sacanagem ne")
	}

}
