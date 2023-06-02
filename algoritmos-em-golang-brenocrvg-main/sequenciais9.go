package main

import "fmt"

func main() {
	var valor float64
	fmt.Println("digite o preço do produto: ")
	fmt.Scan(&valor)
	fmt.Println("o preço com desconto de 10% é: ", valor-(valor*0.1))

}
