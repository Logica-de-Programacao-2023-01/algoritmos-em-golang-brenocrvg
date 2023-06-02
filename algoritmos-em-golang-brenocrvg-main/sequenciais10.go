package main

import "fmt"

func main() {
	var peso float64
	fmt.Println("digite o seu peso em quilos: ")
	fmt.Scan(&peso)
	fmt.Println("o seu peso em libras é igual a: ", peso*2.20462)

}
