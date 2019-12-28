package main

import (
	"errors"
	"fmt"
)

func main() {
	defer recuperar()
	err := errors.New("deu ruim")
	panic(err)
}

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Me recuperei de: ", r)
	}
}
