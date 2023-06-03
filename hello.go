package main

import "fmt"
import "rsc.io/quote"

var x int = 10
var y int = 90
const pi float64 = 3.14
var isGolang bool = true

func main() {
	fmt.Println(quote.Go())
	fmt.Println(x + y)
	fmt.Printf("The value of PI is: %v\n", pi)

	if (isGolang) {
		fmt.Println("Currently writing in Golang")
	}
}
