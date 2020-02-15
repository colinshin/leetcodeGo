# [5. Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/)

## 题目
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example:
```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```
Example:
```
Input: "cbbd"
Output: "bb"
```
## 分析
1. 朴素实现 O(n^3)
```
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

```
2.考虑回文的特点<br>
```
```
