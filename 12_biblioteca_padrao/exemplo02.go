package main

import (
	"fmt"
	"os"
)

func main() {
	arq, err := os.Create("teste.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo!")
		os.Exit(1)
	}
	defer arq.Close()

	arq.WriteString("Escrevendo a primeira linha\n")
	arq.WriteString("Escrevendo a segunda linha\n")
	arq.WriteString("Escrevendo a terceira linha\n")

	// arq.Close()
}
