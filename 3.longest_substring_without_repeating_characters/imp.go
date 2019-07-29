package __longest_substring_without_repeating_characters

import "github.com/zrcoder/leetcodeGo/util/integer"

func lengthOfLongestSubstring(s string) int {
	m := map[interface{}]int{}
	max := 0
	for left, right := 0, 0; right < len(s); right ++ {
		c := s[right]
		if index, found := m[c]; found {
			left = integer.Max(left, index+1)
		}
		max = integer.Max(max, right-left+1)
		m[c] = right
	}
	return max
}
