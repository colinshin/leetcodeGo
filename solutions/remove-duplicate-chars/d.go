/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package remove_duplicate_chars

/*
给定一个仅包含小写字母的字符串，去除字符串中重复的字母，使得每个字母只出现一次。
需保证返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

示例 1:

输入: "bcabc"
输出: "abc"
解释: 原字符串中 ‘b’, ‘c’ 均有多个，在不改变原有字符位置前提下，去掉最前面的 “bc”，使得输出字符串字典序最小
示例 2:

输入: "cbacdcbc"
输出: "acdb"
*/

/*
遍历s， 挑选合适的字母追加到result
对于遍历到的当前字母X，和已经放入result的尾部字母T：

X已经在result里，什么也不做
X不在result里：
	X<T，则看T在原字符串s中X之后是不是还有
		没有, 将X追加到result
		有, 可以从result中删除T，并接着对result尾部字母T'与X做相同的判断处理，最后将X追加到result
		——这里的玩法像是栈！！！
	X>t 追加X

借助两个map：
count首先记录每个字母在s中出现的次数；在修改result时根据情况增减字母个数
inResult记录字母是否已经在result中；在修改result时根据情况标记字母是否在result中
*/
func removeDuplicateLetters(s string) string {
	const letterNums = 26
	count := make(map[rune]int, letterNums) // count letters in s, we will change the numbers when make result later
	for _, c := range s {
		count[c]++
	}
	inResult := make(map[rune]bool, letterNums)
	var result []rune // use result as stack
	for _, c := range s {
		count[c]--
		if result == nil {
			result = append(result, c)
			inResult[c] = true
			continue
		}
		if inResult[c] {
			continue
		}
		for len(result) != 0 && c < result[len(result)-1] && count[result[len(result)-1]] > 0 {
			// pop all letters at the tail of result which are bigger than c
			last := result[len(result)-1]
			inResult[last] = false
			result = result[:len(result)-1]
		}
		result = append(result, c)
		inResult[c] = true
	}
	return string(result)
}
