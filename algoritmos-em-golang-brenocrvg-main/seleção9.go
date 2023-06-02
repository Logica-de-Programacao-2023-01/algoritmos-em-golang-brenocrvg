package main

import (
	"fmt"
	"sort"
)

func main() {
	var x, y, z float64
	fmt.Println("digite o primeiro número: ")
	fmt.Scan(&x)
	fmt.Println("digite o segundo número: ")
	fmt.Scan(&y)
	fmt.Println("digite o terceiro número: ")
	fmt.Scan(&z)
	numeros := []float64{x, y, z}
	sort.Float64s(numeros)

	fmt.Println("números em ordem crescente:", numeros)

}
