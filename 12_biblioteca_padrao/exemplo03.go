package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arq, err := os.Open("teste.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo!")
	}
	defer arq.Close()

	scanner := bufio.NewScanner(arq)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo!")
	}
}
