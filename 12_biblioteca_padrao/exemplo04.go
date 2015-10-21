package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arq, err := os.Create("teste_bufio.txt")

	if err != nil {
		fmt.Println("Erro ao criar o arquivo!")
		os.Exit(1)
	}

	writer := bufio.NewWriter(arq)

	n1, err := writer.WriteString("Escrevendo a primeira linha\n")
	fmt.Println(n1, " bytes escritos com sucesso!")

	n1, err = writer.WriteString("Escrevendo a segunda linha\n")
	fmt.Println(n1, " bytes escritos com sucesso!")

	n1, err = writer.WriteString("Escrevendo a terceira linha\n")
	fmt.Println(n1, " bytes escritos com sucesso!")

	writer.Flush()

	arq.Close()

	arq, err = os.Open("teste_bufio.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo!")
		os.Exit(1)
	}
	defer arq.Close()

	scanner := bufio.NewScanner(arq)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao scannear o arquivo!")
		os.Exit(1)
	}
}
