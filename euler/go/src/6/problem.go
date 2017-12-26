package main

import "fmt"

const Max = 100

func main() {

	var sumsquares = 0
	var squaresum = 0

	for i := 0; i <= Max; i++ {
		sumsquares += i * i
		squaresum += i
	}
	squaresum = squaresum * squaresum

	fmt.Printf("%v.\n", squaresum-sumsquares)
}
