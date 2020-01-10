package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%f %f %f\n", &a, &b, &c)
	if formaTriangulo(a, b, c) {
		perimetro(a, b, c)
	} else {
		area(a, b, c)
	}
}

func formaTriangulo(a, b, c float64) bool {
	return (math.Abs(b-c) < a && a < (b+c)) &&
		(math.Abs(a-c) < b && b < (a+c)) &&
		(math.Abs(a-b) < c && c < (a+b))
}

func perimetro(a, b, c float64) {
	fmt.Printf("Perimetro = %.1f\n", a+b+c)
}

func area(a, b, c float64) {
	fmt.Printf("Area = %.1f\n", ((a+b)/2)*c)
}
