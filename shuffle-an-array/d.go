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
	for i, v := range nums {
		r.helper[i] = v
	}
	return r
}

/** Resets the array to its original configuration and return it. */
func (s *Solution) Reset() []int {
	for i, v := range s.original {
		s.helper[i] = v
	}
	return s.original
}

/** Returns a random shuffling of the array. */
func (s *Solution) Shuffle() []int {
	n := len(s.helper)
	for i := 0; i < n-1; i++ {
		j := i + rand.Intn(n-i)
		s.helper[i], s.helper[j] = s.helper[j], s.helper[i]
	}
	return s.helper
}
