package main

import (
	"fmt"
	"math"
)

// https://blog.golang.org/go-slices-usage-and-internals
// https://golang.org/doc/effective_go.html
// https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm


// INTERFACES
// representam um conjunto de assinaturas de metodos
// some types(ex structs, ints ) implements the interface methods
//

type FormaGeometrica interface {
	area() float64
}

// definição para método Interface FormaGeometrica
func getArea(formaGeometrica FormaGeometrica) float64 {
	return formaGeometrica.area()
}

type Circulo struct {
	x,y,raio float64
}

type Retangulo struct {
	largura, altura float64
}

// implementacao de Circulo parao  metodo area de FormaGeometrica
func(circulo Circulo) area() float64 {
	return math.Pi * circulo.raio * circulo.raio
}

// implementacao de Retangulo para o  metodo area de FormaGeometrica
func(retangulo Retangulo) area() float64 {
	return retangulo.altura * retangulo.largura
}
	
func main() {
	circulo := Circulo{x:0,y:0,raio:5}
	retangulo := Retangulo {largura	:10, altura:5}

	fmt.Printf("Área do circulo: %f\n",getArea(circulo))
	fmt.Printf("Área do retangulo: %f\n",getArea(retangulo))
}