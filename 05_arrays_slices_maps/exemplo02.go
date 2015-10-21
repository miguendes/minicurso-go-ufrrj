package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}

	fmt.Println(slice1)
	slice1[2] = -9
	fmt.Println(slice1)

	// slice1[3] = 5 -- erro

	slice2 := make([]int, 0)
	// slice2[0] = 9 -- erro slice tam 0
	slice2 = append(slice2, 9)
	slice2 = append(slice2, 8)
	fmt.Println(slice2)
	fmt.Printf("slice2 tem len(%d) e cap(%d)\n", len(slice2), cap(slice2))

	slice2 = append(slice2, 7)
	fmt.Println(slice2)
	fmt.Printf("slice2 tem len(%d) e cap(%d)\n", len(slice2), cap(slice2)) // slices dobram de tamanho

	// com len(3) é possivel acessar pos: 0, 1, 2
	slice2[2] = 6
	fmt.Println(slice2)

	slice2 = append(slice2, 5)
	slice2 = append(slice2, 4)

	fmt.Println(slice2)
	fmt.Println(slice2[0:3])
	fmt.Println(slice2[2:4])

	fmt.Printf("slice2 tem len(%d) e cap(%d)\n", len(slice2), cap(slice2)) // slices dobram de tamanho

	fmt.Println(slice2[2:8]) // slice2[2:9] 9 está 'out of bounds'

}
