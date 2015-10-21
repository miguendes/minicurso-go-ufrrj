package main

import (
	"fmt"
)

func main() {
	arr := [4]int{1, 2, 3, 4}

	fmt.Println(arr)

	arr2 := [4]int{} // var arr2 [4]int

	fmt.Println(arr2)

	arr2[0], arr2[1], arr2[2], arr2[3] = 1, 2, 3, 4

	fmt.Println(arr2)

	arr3 := [...]int{4, 3, 2, 1, 0}
	fmt.Println(arr3)

	fmt.Printf("arr3 tem len(%d) e cap(%d)\n", len(arr3), cap(arr3))

	// matriz em Go
	var mat [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			mat[i][j] = i + j
		}
	}
	fmt.Println("mat: ", mat)

}
