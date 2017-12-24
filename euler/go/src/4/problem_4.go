package main

import "fmt"
import "strconv"
import "../utils/palindrome"

func main() {

	var product int
	var a int
	var b int

	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			p := i * j
			if palindrome.IsPalindrome(strconv.Itoa(p)) {
				if p > product {
					product = p
					a = i
					b = j
				}
			}
		}
	}

	fmt.Println("The largest Palindromic product of 2 3 digit numbers.")
	fmt.Printf("%v * %v = %v\n", a, b, product)
}
