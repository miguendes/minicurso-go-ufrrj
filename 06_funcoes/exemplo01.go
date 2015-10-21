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
}

func retornaDoisValores(a, b int) (int, int) {
	return a + b, a - b
}
func main() {

	fmt.Printf("fat(5) = %d\n", fat(5))

	var arr [4]int
	fmt.Println(arr)
	insereNoArray(arr) // go passa valores e não referencia
	fmt.Println(arr)

	x, y := retornaDoisValores(8, 4)
	fmt.Println(x, " - ", y)

}
