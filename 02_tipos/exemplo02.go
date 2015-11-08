package main

import (
	 "fmt"
	 "reflect"
 )

func main() {
	// declaração de variaveis em Go

	var myInt int         // uint8, uint16, uint32, uint64, int8, int16, int32 and int64.
	var myFloat32 float32 // existem float32 (float em java, c) e float64 (double em java)
	var myString string
	var myBoolean bool

	myInt = 11
	myFloat32 = 5.32
	myString = "string"
	myBoolean = true

	fmt.Println(myInt)
	fmt.Println(myFloat32)
	fmt.Println(myString)
	fmt.Println(myBoolean)

	var ( 
		myInt64 int64; myUint uint32
	 )

	myInt64 = -563255666556555455
	fmt.Println(myInt64)

	myUint = 855
	fmt.Println(myUint)
	//	myUint = -2   -- erro
	// fmt.Println(myUint)

	var myComplex complex128 = 5 + 6i

	fmt.Println(myComplex)

	// cast

	var float32toInt64 = int64(myFloat32)
	fmt.Println(float32toInt64)

	var intToFloat = float64(myInt)
	fmt.Println(reflect.TypeOf(intToFloat))
	
	fmt.Println(intToFloat)
					fmt.Println(float32toInt64)
	
}
