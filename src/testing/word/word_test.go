package word

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("abcdcba") {
		t.Error("IsPalindrome(\"abcdcba\") = true")
	}
	if !IsPalindrome("kayak") {
		t.Error("IsPalindrome(\"abcdbcba\") = false")
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error("IsPalindrome(\"palindrome\") = true")
	}
}

// func TestFrenchPalindrome(t *testing.T) {
// 	input := `été`
// 	if !IsPalindrome(input) {
// 		t.Errorf("IsPalindrome(%q) = false", input)
// 	}
// }

// func TestCannalPalindrome(t *testing.T) {
// 	input := "A man, a plan, a canal: Panama"
// 	if !IsPalindrome(input) {
// 		t.Errorf("IsPalindrome(%q) = false", input)
// 	}
// }

func TestIsPalindromeIgnoreSpaceLetter(t *testing.T) {
	// 测试驱动表格
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if rc := IsPalindromeIgnoreSpaceLetter(test.input); rc != test.want {
			t.Errorf("IsPalindromeIgnoreSpaceLetter(%q) = %v", test.input, rc)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if IsPalindrome("yvm") != false {
			b.Errorf("IsPalindrome(\"yvm\") = true")
		}
	}
}
