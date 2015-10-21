package main

import (
	"fmt"
)

func swap(i, j *int) {
	*i, *j = *j, *i
}
func main() {
	x := 9

	fmt.Println("x = ", x)

	y := &x

	*y = 8

	fmt.Println("y = ", *y, " x = ", x)

	a, b := 1, 2
	fmt.Println("a = ", a, " b = ", b)

	swap(&a, &b)
	fmt.Println("troca ")
	fmt.Println("a = ", a, " b = ", b)

}
