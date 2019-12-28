package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Tries to open a file
	f, err := os.Open("thisfiledoesnotexist.txt")

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
	fmt.Println(bs)

}
