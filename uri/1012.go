package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Scanf("%f %f %f\n", &a, &b, &c)
	triangulo(a, c)
	circulo(c)
	trapezio(a, b, c)
	quadrado(b)
	retangulo(a, b)
}

func triangulo(base float64, altura float64) {
	area := (base * altura) / 2
	fmt.Printf("TRIANGULO: %.3f\n", area)
}

func circulo(raio float64) {
	pi := 3.14159
	area := pi * (raio * raio)
	fmt.Printf("CIRCULO: %.3f\n", area)
}

func trapezio(base1 float64, base2 float64, altura float64) {
	area := ((base1 + base2) * altura) / 2
	fmt.Printf("TRAPEZIO: %.3f\n", area)
}

func quadrado(lado float64) {
	fmt.Printf("QUADRADO: %.3f\n", lado*lado)
}

func retangulo(lado1 float64, lado2 float64) {
	fmt.Printf("RETANGULO: %.3f\n", lado1*lado2)
}
