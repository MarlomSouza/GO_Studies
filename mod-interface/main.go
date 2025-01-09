package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
}

type retangulo struct {
	largura, altura float64
}

func (circulo circulo) area()  float64 {
	return math.Pi *  circulo.radius * circulo.radius
}

type circulo struct{
	radius float64
}

func (retangulo retangulo) area() float64 {
	return  retangulo.largura * retangulo.altura
}

func ExibiArea(g geometria){
	area := g.area()	
	fmt.Println(area)
}


func main() {
	fmt.Println("Inicializando...")

	retangulo := retangulo {
		largura: 1,
		 altura: 2,
	}

	circulo := circulo{
		radius: 3,
	}

	ExibiArea(retangulo)
	ExibiArea(circulo)
}