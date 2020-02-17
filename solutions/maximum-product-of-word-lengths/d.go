package maximum_product_of_word_lengths

/*
给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。
你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-product-of-word-lengths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxProduct(words []string) int {
	max := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if !haveCommonChar(words[i], words[j]) {

			}
		}
	}

}

func haveCommonChar(s, t string) bool {

}
