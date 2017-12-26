package main

import "fmt"

const Divisor = 20

func main() {

	var divides = false
	var number = 20

	for !divides {
		number++
		divides = true
		for j := 2; j <= Divisor; j++ {
			if number%j != 0 {
				divides = false
				break
			}

		}
	}

	fmt.Printf("The smallest integer divisible by all numbers from 1 to %v is %v.\n", Divisor, number)
}
