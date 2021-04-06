package main

import (
	"flag"
	"fmt"
)

var age uint

func init() {
	flag.UintVar(&age, "age", 0, "Tell you my age object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, my age is: %d!\n", age)
}