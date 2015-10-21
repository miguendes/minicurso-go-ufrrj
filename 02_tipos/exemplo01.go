package main

import "fmt"
import "reflect"

func main() {

	// https://golang.org/ref/spec#Method_sets

	// Numbers: int, float, complex

	// Boolean: bool

	// String

	fmt.Println("1 + 1", 1+1)
	fmt.Println("1.5 + 1.7", 1.5+1.7)
	fmt.Println("isso é uma string")

	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(1.5))
	fmt.Println(reflect.TypeOf(1 + 6i))

	fmt.Println(reflect.TypeOf(true))

	fmt.Println(reflect.TypeOf("isso é uma string"))

	fmt.Println(reflect.TypeOf("isso é um char"[0]))

	var myInt int = 9

	fmt.Println(reflect.TypeOf(myInt))

	// ex array
	fmt.Println(reflect.TypeOf([4]int{1, 2, 3, 4}))

}
