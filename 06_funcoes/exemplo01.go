package main

import (
	"fmt"
)

func fat(n int) int {
	if n == 0 {
		return 1
	}
	return n * fat(n-1)
}
func insereNoArray(arr [4]int) {
	arr[0] = 9
	fmt.Println("Dentro da funcao: ", arr)
}

func retornaDoisValores(a, b int) (int, int) {
	return a + b, a - b
}
func divideDoisValores(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}
func main() {

	fmt.Printf("fat(5) = %d\n", fat(5))

	var arr [4]int
	fmt.Println(arr)
	
	insereNoArray(arr) // go passa valores e não referencia
	
	fmt.Println("Fora da funcão: ", arr)

	x, y := retornaDoisValores(8, 4)
	fmt.Println(x, " - ", y)
	
	fmt.Println(divideDoisValores(8, 4))
	
	if res, foiPossivel := divideDoisValores(8, 2); foiPossivel {
		fmt.Println("Divisao é igual: ", res)
	} else {
		fmt.Println("Erro ao dividir")
	}

}
