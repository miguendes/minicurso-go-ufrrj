package main

import (
	"fmt"
	"math"
)

type Figura interface {
	calculaArea() float64
}

type Quadrado struct {
	lado float64
}

func (q Quadrado) calculaArea() float64 {
	return q.lado * q.lado
}

type Retangulo struct {
	x, y float64
}

func (r Retangulo) calculaArea() float64 {
	return r.x * r.y
}

type Circulo struct {
	raio float64
}

func (c Circulo) calculaArea() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	c := Circulo{5.0}
	//q := Quadrado{3.0}
	//r := Retangulo{2.0, 3.0}

	figuras := []Figura{c}

	for _, figura := range figuras {
		fmt.Println("figura = ", figura.calculaArea())
	}
}
