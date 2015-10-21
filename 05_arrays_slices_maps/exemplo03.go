package main

import (
	"fmt"
)

func main() {

	m := make(map[string]int)

	m["sete"] = 7
	m["treze"] = 13

	fmt.Println(m)

	v1 := m["sete"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "treze")
	fmt.Println("map:", m)

	fmt.Println("len:", len(m))

	if v, prs := m["treze"]; prs { //  prs == true
		fmt.Println("v está no map:", v)
	} else {
		fmt.Println("13 não está no map")
	}

	if v, prs := m["sete"]; prs {
		fmt.Println("v está no map:", v)
	} else {
		fmt.Println("7 não está no map")
	}

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	m["oito"] = 8
	m["dez"] = 10
	m["onze"] = 11

	for k, v := range m {
		fmt.Println(k, ", ", v)
	}
}
