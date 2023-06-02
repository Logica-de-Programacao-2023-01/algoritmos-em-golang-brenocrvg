package main

import "fmt"

func main() {
	var idade int
	fmt.Println("digite sua idade em anos: ")
	fmt.Scan(&idade)
	fmt.Println("sua idade em dias é igual a: ", idade*365)

}
