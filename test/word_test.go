package word

import "testing"

func TestISPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("ANNA") {
		t.Error(`IsPalindrome("ANNA") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

// func TestFrenchPalindrome(t *testing.T) {
// 	if !IsPalindrome("été") {
// 		t.Error(`IsPalindrome("été") = false`)
// 	}
// }

// func TestCanalPalindrome(t *testing.T) {
// 	input := "A man, a plan, a canal: Panama"
// 	if !IsPalindrome(input) {
// 		t.Errorf(`IsPalindrome(%q) = false`, input)
// 	}
// }
