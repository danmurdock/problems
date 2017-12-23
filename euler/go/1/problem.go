package main
import "fmt"
func main() {
		
	var sum int = 0
    for j:=1; j < 1000; j++ {
		if j%3 == 0 || j%5 == 0 {
			sum += j
		}
    }
	
    fmt.Println(sum)
}
