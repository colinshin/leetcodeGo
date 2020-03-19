/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package longest_palindromic_substring

func longestPalindrome1(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			if isPalindromic(sub) && len(result) < j-i+1 {
				result = sub
			}
		}
	}
	return result
}

func isPalindromic(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome2(s string) string {
	result := ""
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func longestPalindrome3(s string) string {
	result := ""
	dp := make([]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(s) - 1; j >= 0; j-- {
			dp[j] = s[i] == s[j] && (j-i < 3 || dp[j-1])
			if dp[j] && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		l, r := expandAroundCenter(s, i)
		if r-l > right-left {
			left, right = l, r
		}
	}
	return s[left : right+1]
}

func expandAroundCenter(s string, i int) (int, int) {
	l1, r1 := expandAround2Centers(s, i, i)   // 以索引i为中心向两边扩展得到的最长回文串的左右边界
	l2, r2 := expandAround2Centers(s, i, i+1) // 以索引i即其下一个位置为中心向两边扩展得到的最长回文串的左右边界
	if r2-l2 > r1-l1 {
		return l2, r2
	}
	return l1, r1
}

func expandAround2Centers(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
	}
	return i + 1, j - 1
}
