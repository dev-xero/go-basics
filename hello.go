package main

import "fmt"
import "rsc.io/quote"

var x int = 10
var y int = 90

func main() {
	fmt.Println(quote.Go())
	fmt.Println(x + y)
}
