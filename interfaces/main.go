package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("inicializando..")
	// retangulo := retangulo{
	// 	largura: 2,
	// 	altura:  1,
	// }
	// circulo := circulo{
	// 	radius: 3,
	// }

	//ExibeGeometria(retangulo)
	////ExibeGeometria(circulo)
	// p := ProblemaDeNetwork{
	// 	rede:     true,
	// 	hardware: false,
	// }
	// ExibeError(p)

	var lista []interface{}

	lista = append(lista, 10)
	lista = append(lista, true)
	lista = append(lista, "teste")
	for _, valor := range lista {
		if v, ok := valor.(string); ok {
			fmt.Println(v + "string")
		} else {
			fmt.Println(v + "não string")
		}
	}
}
func ExibeError(e error) {
	fmt.Println(e.Error())
}

type ProblemaDeNetwork struct {
	rede     bool
	hardware bool
}

func (p ProblemaDeNetwork) Error() string {
	if p.rede {
		return "problema de rede"
	} else if p.hardware {
		return "problema de hardware"
	} else {
		return "outro problema"
	}

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
