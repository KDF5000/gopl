package ch11

import (
	"testing"
)

func TestPalindrom(t *testing.T) {
	if !IsPalindrome("abcdefedcba") {
		t.Error(`IsPalindrome("abcdefedcba") = false`)
	}

	if !IsPalindrome("a") {
		t.Error(`IsPalindrome("a") = false`)
	}

	if IsPalindrome("abbaa") {
		t.Error(`IsPalindrome("abbaa")=true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}
