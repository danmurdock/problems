package palindrome

// IsPalindrome returns true if a string is a palindrome (XYYX or XYZYX)
func IsPalindrome(value string) bool {

	length := len(value)

	for i := 0; i < length/2; i++ {
		if value[i] != value[length-i-1] {
			return false
		}
	}

	return true
}
