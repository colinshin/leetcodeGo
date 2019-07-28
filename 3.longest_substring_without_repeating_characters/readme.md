# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

## 分析
1. 朴素实现 O(n^3)
```
import (
	"github.com/zrcoder/leetcodeGo/util/integer"
	"github.com/zrcoder/dsGo/base/set"
)

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i:=0; i<len(s); i++ {
		for j:=i+1; j<=len(s); j++ {
			subStr := s[i:j]
			if allUnique(subStr) {
				max = integer.Max(j-i, max)
			} else {
				break
			}
		}
	}
	return max
}

func allUnique(s string) bool {
	uniqChars := set.New()
	for _, c := range s {
		if uniqChars.Has(c) {
			return false
		}
		uniqChars.Add(c)
	}
	return true
}
```
2. 
```
import "code.huawei.com/interest/LetsGo/leetcode/util/integer"

func lengthOfLongestSubstring(s string) int {
	m := map[interface{}]int{}
	max := 0
	for left, right := 0, 0; right < len(s); right ++ {
		char := s[right]
		if index, found := m[char]; found {
			left = integer.Max(left, index)
		}
		max = integer.Max(max, right-left+1)
		m[char] = right + 1
	}
	return max
}
```
