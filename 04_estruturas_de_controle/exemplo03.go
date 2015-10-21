package main

import (
	"fmt"
)

func main() {
	a := 9
	b := 2

	switch {
	case a > b:
		fmt.Printf("%d é maior que %d", a, b)
	case a == b:
		fmt.Printf("%d é igual a %d", a, b)
	default:
		fmt.Printf("%d é menor ou igual que %d", a, b)
	}

	fmt.Println()

	a = 9
	b = 7

	switch {
	case b == 4:
		fmt.Println("b igual a 4")
	case a == 0:
		fmt.Println("a igual a 0")
	case a == 2:
		fmt.Println("a igual a 2")
	default:
		fmt.Println("sei la")
	}

	s := "vamos ver quantas vogais existem nesta string"
	vogais, consoantes := 0, 0 // inicialização multipla
	for _, v := range s {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			vogais++
		default:
			consoantes++
		}
	}

	fmt.Printf("existem %d vogais\n", vogais)
	fmt.Printf("existem %d consoantes\n", consoantes)

}
