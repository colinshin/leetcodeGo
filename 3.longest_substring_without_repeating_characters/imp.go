package __longest_substring_without_repeating_characters

import "github.com/zrcoder/leetcodeGo/util/integer"

func lengthOfLongestSubstring(s string) int {
	m := map[interface{}]int{}
	max := 0
	for left, right := 0, 0; right < len(s); right ++ {
		char := s[right]
		if index, found := m[char]; found {
			left = integer.Max(left, index)
		}
		max = integer.Max(max, right-left+1)
		m[char] = right + 1 // add or update m[char]
	}
	return max
}
