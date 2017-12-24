package main

import "fmt"
import "../prime"

func main() {

	primes := prime.Seive(2000000)

	var sum int

	for _, prime := range primes {
		sum += prime
	}

	fmt.Println(sum)

}
