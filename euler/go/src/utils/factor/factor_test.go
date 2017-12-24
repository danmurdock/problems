package factor

import "fmt"
import "testing"

func TestFactor(t *testing.T) {
	fmt.Println(Factor(12341234))
	fmt.Println(Factor(1030))
}

func TestFactorKAT(t *testing.T) {

	primes := Factor(12341234)

	var sum int

	for _, prime := range primes {
		sum += prime
	}

	if sum != 829 {

		fmt.Println(sum)
		t.Fail()
	}
}
