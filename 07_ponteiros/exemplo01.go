package main

import (
	"fmt"
)

func swap(i, j *int) {
	*i, *j = *j, *i
	
	/*
		temp := *i
		*i = *j
		*j = temp
	
	 */
}
// main: &a = 0x001
// fakeSwap(i = a)   -- &i == &a == 0x001
// fakeSwap(i = 0x002) -- &i == 0x002
// main &a == 0x001
 func fakeSwap(i, j *int) {
	i, j = j, i
	
	//i = 0x002
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
	
	fmt.Println("antes ")
	fmt.Println("a = ", a, " b = ", b)
	swap(&a, &b)
	fmt.Println("depois ")
	fmt.Println("a = ", a, " b = ", b)
	

}
