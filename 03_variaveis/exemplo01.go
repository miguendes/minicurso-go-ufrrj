package main

import "fmt"

func main() {
	// declaração de variaveis em Go

	var myInt int
	var myFloat32 float32 // existem float32 (float em java, c) e float64 (double em java)
	var myString string

	myInt = 5
	myFloat32 = 5.32
	myString = "string"

	fmt.Println(myInt)
	fmt.Println(myFloat32)
	fmt.Println(myString)

	var myInt64 int64 // int8, int16, int32
	var myUint uint32 // uint8, uint16, ...

	myInt64 = -563255666556555455
	fmt.Println(myInt64)

	myUint = 855
	fmt.Println(myUint)
	//	myUint = -2   -- erro
	// fmt.Println(myUint)

	// operadores aritimeticos
	fmt.Println(1 + 2)
	fmt.Println(1 - 2)
	fmt.Println(2 * 2)
	fmt.Println(4 / 2)

	// constantes
	const a = 2 + 3.5
	fmt.Println(a)
	


}
