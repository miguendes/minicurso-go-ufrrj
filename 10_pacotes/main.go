package main

import (
	"fmt"
	"minicurso-go-ufrrj/10_pacotes/calclib"
)

func main() {
	resultado := calclib.Soma(1, 2)
	
	fmt.Println(resultado)
	
	resultado = calclib.Multiplica(4, 2)
	fmt.Println(resultado)
	
		resultado = calclib.Divide(4, 2)
	fmt.Println(resultado)
}


