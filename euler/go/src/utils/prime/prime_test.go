package prime

import "fmt"
import "testing"

func TestSeive(t *testing.T) {
	fmt.Println(Seive(1000))
	fmt.Println(Seive(4000000))
}

func TestSeiveKAT(t *testing.T) {

	primes := Seive(2000000)

	var sum int

	for _, prime := range primes {
		sum += prime
	}

	if sum != 142913828922 {

		fmt.Println(sum)
		t.Fail()
	}
}
