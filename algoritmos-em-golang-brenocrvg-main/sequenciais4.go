package main

import "fmt"

func main() {
	var x, y, z float64
	fmt.Println("digite o primeiro número: ")
	fmt.Scan(&x)
	fmt.Println("digite o segundo número: ")
	fmt.Scan(&y)
	fmt.Println("digite  terceiro número: ")
	fmt.Scan(&z)
	fmt.Println("a média ponderada é igual a: ", (x*2+y*3+z*5)/10)
}
