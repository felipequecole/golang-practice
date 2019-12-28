package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// Tries to open a file
	f, err := os.Open("/home/felipequecole/go/src/github.com/felipequecole/curso/error-handling/thisfileexists.txt")

	// if there's an error, halt.
	if err != nil {
		fmt.Println(err)
		return
	}
	//Defer means it will run at the end of function execution
	defer f.Close()

	// Tries to read from file
	bs, err := ioutil.ReadAll(f)
	// Again, if there's any error, halt.
	if err != nil {
		fmt.Println(err)
		return
	}

	// print lines from file
	fmt.Println(string(bs))

}
