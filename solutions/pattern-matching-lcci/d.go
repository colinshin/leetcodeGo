/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package pattern_matching_lcci

/*
面试题 16.18. 模式匹配
你有两个字符串，即pattern和value。
pattern字符串由字母"a"和"b"组成，用于描述字符串中的模式。
例如，字符串"catcatgocatgo"匹配模式"aabab"（其中"cat"是"a"，"go"是"b"），该字符串也匹配像"a"、"ab"和"b"这样的模式。
但需注意"a"和"b"不能同时表示相同的字符串。编写一个方法判断value字符串是否匹配pattern字符串。

示例 1：
输入： pattern = "abba", value = "dogcatcatdog"
输出： true

示例 2：
输入： pattern = "abba", value = "dogcatcatfish"
输出： false

示例 3：
输入： pattern = "aaaa", value = "dogcatcatdog"
输出： false

示例 4：
输入： pattern = "abba", value = "dogdogdogdog"
输出： true
解释： "a"="dogdog",b=""，反之也符合规则

提示：
0 <= len(pattern) <= 1000
0 <= len(value) <= 1000
你可以假设pattern只包含字母"a"和"b"，value仅包含小写字母。
*/

/*
边界情况较多；主要思路是统计pattern里'a'、'b'的数量，然后穷举实际匹配到的字符串的长度
*/
func patternMatching(pattern string, value string) bool {
	if len(pattern) == 0 {
		return len(value) == 0
	}
	if len(value) == 0 {
		return len(pattern) == 1
	}
	return check(pattern, value)
}

func check(pattern, value string) bool {
	nA, nB := countAB(pattern)
	if nA == 0 {
		return canCut(value, nB)
	}
	if nB == 0 {
		return canCut(value, nA)
	}
	if canCut(value, nA) || canCut(value, nB) {
		return true
	}
	return rangeABLensToCheck(pattern, value, nA, nB)
}

func countAB(pattern string) (int, int) {
	nA, nB := 0, 0
	for _, v := range pattern {
		if v == 'a' {
			nA++
		} else {
			nB++
		}
	}
	return nA, nB
}

func canCut(value string, k int) bool {
	m := len(value)
	if m%k != 0 {
		return false
	}
	subSize := m / k
	sub := value[:subSize]
	for i := subSize; i < m; i += subSize {
		if value[i:i+subSize] != sub {
			return false
		}
	}
	return true
}

func rangeABLensToCheck(pattern, value string, nA, nB int) bool {
	m := len(value)
	for lenA := 1; lenA*nA+nB <= m; lenA++ {
		if (m-lenA*nA)%nB != 0 {
			continue
		}
		lenB := (m - lenA*nA) / nB
		if canMatch(pattern, value, lenA, lenB) {
			return true
		}
	}
	return false
}

func canMatch(pattern, value string, lenA, lenB int) bool {
	matchedA, matchedB := "", ""
	j := 0
	for _, v := range pattern {
		if v == 'a' {
			sub := value[j : j+lenA]
			if matchedA == "" {
				matchedA = sub
			} else if matchedA != sub {
				return false
			}
			j += lenA
		} else {
			sub := value[j : j+lenB]
			if matchedB == "" {
				matchedB = sub
			} else if matchedB != sub {
				return false
			}
			j += lenB
		}
	}
	return true
}
