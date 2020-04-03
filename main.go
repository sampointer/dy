package main

import (
	"fmt"
	divvy "github.com/sampointer/dy/divvyyaml"
	"os"
)

func main() {
	var dy divvy.DivvyYaml
	var multiDoc bool

	// Do the most basic argument parsing possible
	if len(os.Args) < 2 {
		os.Stderr.WriteString("you must pass at least one path as an argument\n")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		multiDoc = true
	}

	// Allow multiple arguments and emit each as a different document
	for i, dir := range os.Args {
		if i == 0 {
			continue
		}

		err := dy.Parse(dir)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(2)
		} else {
			if multiDoc == true {
				fmt.Println("---")
			}
			fmt.Println(dy.Doc)
		}
	}
}
