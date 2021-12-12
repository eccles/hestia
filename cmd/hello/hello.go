package main

import "fmt"
import "rsc.io/quote"

func Hello() string {
    return quote.Go()
}

func main() {
    fmt.Println(Hello())
}
