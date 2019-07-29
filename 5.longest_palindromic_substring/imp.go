package __longest_palindromic_substring

func longestPalindrome(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			if isPalindromic(&sub) && len(r) < j-i+1 {
				r = sub
			}
		}
	}
	return r
}

func isPalindromic(s *string) bool {
	for i, j := 0, len(*s)-1; i < j; {
		if (*s)[i] != (*s)[j] {
			return false
		}
		i++
		j--
	}
	return true
}
