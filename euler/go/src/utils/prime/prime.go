package prime

func Seive(value int) []int {

	max := value
	numbers := make([]int, max)
	var primes []int

	for j := 1; j < max; j++ {
		numbers[j] = 1
	}

	for j := 2; j < max; j++ {

		if numbers[j] == 1 {
			primes = append(primes, j)
		}

		var multiple = j + j
		for multiple < max {
			numbers[multiple] = 0
			multiple += j

		}
	}

	return primes
}
