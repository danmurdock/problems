package main

import "fmt"

func main() {

	var a = 1
	var b = 1
	var c int

	for a < 998 {
		b = 1
		for b <= a {
			c = 1000 - a - b
			if (a*a + b*b) == c*c {
				fmt.Printf("%v^2  + %v^2 = %v^2, %v\n", a, b, c, a*b*c)
			}
			b++
		}
		a++
	}

	a = 1
	b = 1
	c = 1

	for a < 10 {
		b = 1
		for b < 10 {
			c = 1
			for c < 10 {
				if (a*a + b*b) == c*c {
					fmt.Printf("%v^2  + %v^2 = %v^2, %v\n", a, b, c, a*b*c)
				}
				c++
			}
			b++
		}
		a++
	}
}
