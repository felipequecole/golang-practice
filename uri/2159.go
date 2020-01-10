package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scanf("%f\n", &n)
	p := n / math.Log(n)
	m := 1.25506 * p
	fmt.Printf("%.1f %.1f\n", p, m)
}
