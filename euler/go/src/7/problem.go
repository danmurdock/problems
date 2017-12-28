package main

import "fmt"
import "../utils/prime"

const Max = 100

func main() {

	var cap = 1000000
	var primes = prime.Seive(cap)

	for len(primes) < 10001 {
		cap *= 2
		primes = prime.Seive(cap)
	}

	fmt.Printf("%v.\n", primes[10000])
}
