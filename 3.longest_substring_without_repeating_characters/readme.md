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
2. 我们记寻找的子串开始的索引为left，结束的索引为right，right不断增加，用一个全局变量max记录已经发现的没有重复字符的子串长度；<br>
用一个map，键为字符中的字符，值为其索引；如果发现m中能找到right所在的元素，说明字符出现了重复；<br>
因为求的是没有重复字符的最长子串，此时left只需要移动到第一次出现该字符的下一位，right不变，继续增加即可。<br>
时间复杂度O(n), 空间复杂度O(n)
```
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
```
