package main

import (
	"fmt"
	divvy "github.com/sampointer/dy/divvyyaml"
	"os"
)

func main() {
	var dy divvy.DivvyYaml

	// Do the most basic argument parsing possible
	if len(os.Args) < 2 {
		os.Stderr.WriteString("you must pass a path as an argument\n")
		os.Exit(1)
	}

	err := dy.Parse(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	} else {
		fmt.Println(dy.Doc)
	}
}
