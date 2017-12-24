package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {

	if !IsPalindrome("abccba") {
		t.Fail()
	}

	if !IsPalindrome("abcba") {
		t.Fail()
	}

	if IsPalindrome("abcbb") {
		t.Fail()
	}

}
