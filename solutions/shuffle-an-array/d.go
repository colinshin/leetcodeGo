/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package shuffle_an_array

import "math/rand"

/*
打乱一个没有重复元素的数组。

示例:

// 以数字集合 1, 2 和 3 初始化数组。
int[] nums = {1,2,3};
Solution solution = new Solution(nums);

// 打乱数组 [1,2,3] 并返回结果。任何 [1,2,3]的排列返回的概率应该相同。
solution.shuffle();

// 重设数组到它的初始状态[1,2,3]。
solution.reset();

// 随机返回数组[1,2,3]打乱后的结果。
solution.shuffle();

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shuffle-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type Solution struct {
	original []int
	helper   []int
}

func Constructor(nums []int) Solution {
	r := Solution{}
	r.original = nums
	r.helper = make([]int, len(nums))
	_ = copy(r.helper, nums)
	return r
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	_ = copy(s.helper, s.original)
	return s.original
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	for i := len(s.helper) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		s.helper[i], s.helper[j] = s.helper[j], s.helper[i]
	}
	return s.helper
}
