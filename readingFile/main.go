package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		os.Exit(1)
	}
	fn := argsWithProg[1]

	dat, err := os.ReadFile(fn)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(dat))
}
