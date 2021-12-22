package main

import (
	"fmt"

	"rsc.io/quote"
)

func Hello() string {
	return quote.Go()
}

func main() {
	fmt.Printf("%s", Hello())
}
