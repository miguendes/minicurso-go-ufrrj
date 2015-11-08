package main

import (
	"fmt"
)

type Ponto struct {
	x, y int
}

func (p *Ponto) move(x, y int) {
	p.x = x
	p.y = y
}

func (p Ponto) soma(p1 Ponto) Ponto {
	return Ponto{p.x + p1.x, p.y + p1.y}
}

func (p Ponto) String() string {
	return fmt.Sprintf("x = %d; y = %d", p.x, p.y)
}

func main() {
	p1 := Ponto{1, 2}
	p2 := Ponto{y: 3, x: 4} // nesse caso pode ser qualquer ordem

	fmt.Println("p1 = ", p1)
	fmt.Println("p2 = ", p2)

	p1.move(5, 6)
	fmt.Println("p1 = ", p1)

	p3 := p1.soma(p2)
	fmt.Println("p3 = ", p3)
	fmt.Println("p1 = ", p1)

	fmt.Println(p3)

}
