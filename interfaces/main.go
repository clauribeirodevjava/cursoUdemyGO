package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("inicializando..")
	retangulo := retangulo{
		largura: 2,
		altura:  1,
	}
	circulo := circulo{
		radius: 3,
	}

	ExibeGeometria(retangulo)
	ExibeGeometria(circulo)
}

type geometria interface {
	area() float64
}

func (r retangulo) area() float64 {
	var area float64
	area = r.altura * r.largura
	return area
}

func (c circulo) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

type retangulo struct {
	largura, altura float64
}
type circulo struct {
	radius float64
}

func ExibeGeometria(g geometria) {
	fmt.Println(g.area())
}
