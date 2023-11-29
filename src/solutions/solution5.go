package solutions

import (
	"fmt"
	"unicode/utf8"
)

type Solution5 struct {
}

func (s Solution5) Solve() {
	fmt.Println("\n===== Solution 5 =====")

	fmt.Printf("Is 'źaź' a palindrome: %v \n", IsPalindrome("źaź"))
	fmt.Printf("Is 'ąćą' a palindrome: %v \n", IsPalindrome("ąćą"))
	fmt.Printf("Is 'test' a palindrome: %v \n", IsPalindrome("test"))
	fmt.Printf("Is '' a palindrome: %v \n", IsPalindrome(""))
}

func IsPalindrome(s string) bool {
	reversedString := reverseString(s)

	if s != reversedString {
		return false
	}

	return true
}

func reverseString(s string) string {
	// try this runes := []rune(s)
	// now you have slice of runs to play with
	size := len(s)
	reversedStringBuffer := make([]byte, size)

	for i := 0; i < size; {
		r, n := utf8.DecodeRuneInString(s[i:])
		i += n
		utf8.EncodeRune(reversedStringBuffer[size-i:], r)
	}

	return string(reversedStringBuffer)
}
