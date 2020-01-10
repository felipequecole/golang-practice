package main

import "fmt"

func main() {
	var falta int
	var presente1 int
	var presente2 int
	fmt.Scanf("%d\n%d %d", &falta, &presente1, &presente2)
	if presente1+presente2 > falta {
		fmt.Println("Deixa para amanha!")
	} else {
		fmt.Println("Farei hoje!")
	}
}
