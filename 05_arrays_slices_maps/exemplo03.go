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


	v, prs := m["treze"]
	if  prs { //  prs == false
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

	for _, v := range m {
		fmt.Println(", ", v)
	}
	
	delete(m, "dez")
	
	fmt.Printf("tipo: %T - Elem: %v", m, m)
	
	
}
