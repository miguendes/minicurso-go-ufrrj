package main

import (
	"fmt"
)

type Ponto struct{}

func main() {
	var a []int

	a = append(a, 5)
	a = append(a, 4)
	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)

	s := fmt.Sprintf("%t", true)
	fmt.Printf("s: %s\n", s)

	s = fmt.Sprintf("%T", "true")
	fmt.Printf("s: %s\n", s)

	fmt.Printf("%T\n", a)

	fmt.Printf("%010.3f\n", 2.5662)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	p := Ponto{}

	fmt.Printf("%T\n", p)
}
