package main

import "log"

func main() {
	defer posSusto()
	log.Panic("panico")
}

func posSusto() {
	log.Println("que susto")
}
