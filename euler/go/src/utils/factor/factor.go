package factor

import (
	"math"

	"../prime"
)

func Factor(value int) []int {

	var factors []int

	primes := prime.Seive(int(math.Ceil(math.Sqrt(float64(value)))))

	accumulate := value

	for _, p := range primes {

		for accumulate%p == 0 {
			accumulate = accumulate / p
			factors = append(factors, p)
		}

		if accumulate == 1 {
			break
		}
	}

	return factors
}
