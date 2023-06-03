package main

import "fmt"
import "rsc.io/quote"

var x int = 10
var y int = 90
const pi float64 = 3.14
var isGolang bool = true

var testArray = [3]int {1, 2, 3}
var varLengthTestArray = [...]int {9, 8, 7, 6}
var testSlice = []int {1, 2, 3, 4, 5}

func main() {
	fmt.Println(quote.Go())
	fmt.Println(x + y)
	fmt.Printf("The value of PI is: %v\n", pi)

	if (isGolang) {
		fmt.Println("Currently writing in Golang")
	}

	fmt.Println(testArray[0])
	fmt.Println()

	// loops have a weirdly nice syntax
	for i := 0; i < len(varLengthTestArray); i++ {
		fmt.Println(varLengthTestArray[i])
	}

	fmt.Println()
	for j := 0; j < len(testSlice); j++ {
		fmt.Print(testSlice[j], " ")
	}
}
