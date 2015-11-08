package main

import (
	"fmt"
)

func main() {
	//for

	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	for i := 0; ; i++ {
		if i == 10 {
			break
		}
		fmt.Print(i, " ")
	}

	fmt.Println()

	i := 0
	for ; i < 10; i++ {
		fmt.Print(i, " ")
	}

	// while
	fmt.Println()

	i = 0
	for i < 10 {
		fmt.Print(i, " ")
		i++
	}

	// loop infinito
	fmt.Println()

	i = 0
	for false { // for true {
		fmt.Print(i, " ")
		i++
	}

}
