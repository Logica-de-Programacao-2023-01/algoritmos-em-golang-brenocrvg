package main

import "fmt"

func main() {
	//Faça um algoritmo que imprima os números de 1 a 100, substituindo os múltiplos de 3 pela palavra "Fizz" e os múltiplos de 5 pela palavra "Buzz". Para os números que são múltiplos de ambos, utilize a palavra "FizzBuzz".
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}

	}

}
