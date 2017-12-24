package prime

import "fmt"
import "testing"

func TestSeive(t *testing.T) {
	fmt.Println(Seive(1000))
	fmt.Println(Seive(4000000))
}
