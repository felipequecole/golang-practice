package main

import "fmt"

func main() {
	var entrada float64
	fmt.Scanf("%f\n", &entrada)
	if entrada < 0 || entrada > 100 {
		fmt.Println("Fora de intervalo")
	} else if entrada <= 25 {
		fmt.Println("Intervalo [0,25]")
	} else if entrada <= 50 {
		fmt.Println("Intervalo (25,50]")
	} else if entrada <= 75 {
		fmt.Println("Intervalo (50,75]")
	} else {
		fmt.Println("Intervalo (75,100]")
	}
}
