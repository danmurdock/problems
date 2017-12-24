package main

import "fmt"
import "../utils/factor"

func main() {

	factors := factor.Factor(600851475143)

	var sum int

	for _, factor := range factors {
		sum += factor
	}

	fmt.Println(factors)
	fmt.Println(sum)

}
