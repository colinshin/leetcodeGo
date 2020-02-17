package longest_substring_without_repeating_characters

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
滑动窗口
时间复杂度：O(2n)=O(n)，在最糟糕的情况下，每个字符将被i 和j 访问两次。
空间复杂度：O(k)，Set 的大小。取决于字符串的大小n 以及字符集 / 字母的大小m 。
*/
func lengthOfLongestSubstring(s string) int {
	result := 0
	set := make(map[byte]struct{}, 0)
	for l, r := 0, 0; r < len(s); {
		c := s[r]
		if _, found := set[c]; !found {
			set[c] = struct{}{}
			r++
			if r-l > result {
				result = r - l
			}
		} else {
			delete(set, s[l])
			l++
		}
	}
	return result
}

/*
优化的滑动窗口
上述的方法最多需要执行 2n 个步骤。事实上，它可以被进一步优化为仅需要 n 个步骤。
可以定义字符到索引的映射，而不是使用集合来判断一个字符是否存在。 当找到重复的字符时，可以立即跳过该窗口。
*/
func lengthOfLongestSubstring1(s string) int {
	m := make(map[byte]int, 0)
	result := 0
	for l, r := 0, 0; r < len(s); r++ {
		c := s[r]
		if pos, found := m[c]; found && pos > l {
			l = pos
		}
		if r-l+1 > result {
			result = r - l + 1
		}
		m[c] = r + 1 // 下一个位置的索引，第一个元素和最后一个元素的索引不用加入map
	}
	return result
}

/*
另一个修改版，map里记录字符的个数，通用的滑动窗口解法
*/
func lengthOfLongestSubstring10(s string) int {
	found := make(map[byte]int, 0)
	result := 0
	for l, r := 0, 0; r < len(s); {
		c := s[r]
		found[c]++
		r++
		for found[c] > 1 {
			found[s[l]]--
			l++
		}
		if r-l > result {
			result = r - l
		}
	}
	return result
}

/* 不通过，因假设前提不成立；但是思路可以看一下

假设字符集为 ASCII 128,以上方法里的map可以改成数组
以前的我们都没有对字符串 s 所使用的字符集进行假设
当我们知道该字符集比较小的时侯，我们可以用一个整数数组作为直接访问表来替换 Map
常用的表如下所示：
int [26] 用于字母 ‘a’ - ‘z’ 或 ‘A’ - ‘Z’
int [128] 用于ASCII码
int [256] 用于扩展ASCII码
*/
func lengthOfLongestSubstring2(s string) int {
	const charNum = 128
	index := make([]int, charNum)
	result := 0
	for l, r := 0, 0; r < len(s); r++ {
		c := s[r]
		if index[c] > l {
			l = index[c]
		}
		if r-l+1 > result {
			result = r - 1 + l
		}
		index[c] = r + 1
	}
	return result
}
