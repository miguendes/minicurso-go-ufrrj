package main

import (
	"fmt"
)

func main() {

	// operadores condicionais
	// >, <, !=, ==

	// &&, ||

	fmt.Println(false && false)
	fmt.Println(true && false)
	fmt.Println(true && true)

	fmt.Println(false || false)
	fmt.Println(true || false)
	fmt.Println(true || true)

	a := 9
	b := 2

	if (a > b) {
			fmt.Printf("%d é maior que %d", a, b)
	} else if (a == b) {
		fmt.Printf("%d é igual a %d", a, b)
	} else {
		fmt.Printf("%d é menor ou igual que %d", a, b)
	}
	fmt.Println()

}
