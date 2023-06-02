package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Println("digite o primeiro número: ")
	fmt.Scan(&x)
	fmt.Println("digite o segundo número: ")
	fmt.Scan(&y)
	fmt.Println("digite o terceiro número: ")
	fmt.Scan(&z)
	fmt.Println("a soma dos três números é igual a: ", x+y+z)

}
