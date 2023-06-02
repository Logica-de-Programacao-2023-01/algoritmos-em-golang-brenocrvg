package main

import "fmt"

func main() {
	var dias, valor float64
	fmt.Println("digite a quantidade de dias trabalhados: ")
	fmt.Scan(&dias)
	fmt.Println("digite o valor da diária: ")
	fmt.Scan(&valor)
	fmt.Println("o salário total é igual a: ", dias*valor)

}
