package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Duration(rand.Intn(15)) * time.Millisecond)
	}
}

func main() {

	f("simples")

	fmt.Println()

	go f("goroutine1")
	go f("goroutine2")

	var input string
	fmt.Scanln(&input)
	fmt.Println("pronto")
}
