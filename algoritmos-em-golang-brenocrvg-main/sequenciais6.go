package main

import "fmt"

func main() {
	var salario float64
	fmt.Println("digite o seu salário: ")
	fmt.Scan(&salario)
	fmt.Println("seu novo salário é: ", salario+(salario*0.15))
}
