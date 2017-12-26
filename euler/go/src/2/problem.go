package main

import "fmt"

func main() {

	var sum = 0
	var lfib = 1
	var fib = 1

	for fib < 4000000 {
		var nfib = lfib + fib
		if nfib%2 == 0 {
			sum += nfib
		}
		lfib = fib
		fib = nfib
	}

	fmt.Println(sum)
}
