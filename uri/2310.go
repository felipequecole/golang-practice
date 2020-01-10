package main

import "fmt"

func main() {
	var total int
	s, b, a := 0, 0, 0
	s1, b1, a1 := 0, 0, 0
	fmt.Scanf("%d\n", &total)
	for i := 0; i < total; i++ {
		var nome string
		var l1, l2, l3 int
		fmt.Scanf("%v\n", &nome)
		fmt.Scanf("%d %d %d\n", &l1, &l2, &l3)
		s += l1
		b += l2
		a += l3
		fmt.Scanf("%d %d %d\n", &l1, &l2, &l3)
		s1 += l1
		b1 += l2
		a1 += l3
	}

	fmt.Printf("Pontos de Saque: %.2f %%.\n", (float64(s1) / float64(s) * 100))
	fmt.Printf("Pontos de Bloqueio: %.2f %%.\n", (float64(b1) / float64(b) * 100))
	fmt.Printf("Pontos de Ataque: %.2f %%.\n", (float64(a1) / float64(a) * 100))
}
