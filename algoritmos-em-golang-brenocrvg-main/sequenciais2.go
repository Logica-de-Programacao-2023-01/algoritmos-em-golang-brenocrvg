package main

import "fmt"

func main() {
	var x int
	fmt.Println("digite o seu número: ")
	fmt.Scan(&x)
	fmt.Println("o dobro, triplo e quádruplo do seu número são: ", x*2, x*3, x*4)

}
