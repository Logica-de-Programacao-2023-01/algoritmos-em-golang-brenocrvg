package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("digite o primeiro número: ")
	fmt.Scan(&x)
	fmt.Println("digite o segundo número: ")
	fmt.Scan(&y)
	if x > y {
		fmt.Println("o maior número é", x)
	} else {
		fmt.Println("o maior número é", y)
	}

}
