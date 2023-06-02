package main

import "fmt"

func main() {
	var x int
	fmt.Println("digite um número qualquer: ")
	fmt.Scan(&x)
	fmt.Println("o antecessor e sucessor do seu número são: ", x-1, x+1)

}
