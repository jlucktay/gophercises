/*
https://gophercises.com/exercises/quiz
*/
package main

import (
	"flag"
	"fmt"
)

var csvFilename string

func init() {
	flag.StringVar(&csvFilename, "csv", "problems.csv", "the quiz inputs")
}

func main() {
	flag.Parse()

	fmt.Printf("The value of 'csvFilename': %v\n", csvFilename)
}
