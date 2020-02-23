/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package point_game

import (
	"container/list"
	"math"
)

/*
你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。

示例 1:

输入: [4, 1, 8, 7]
输出: True
解释: (8-4) * (7-1) = 24
示例 2:

输入: [1, 2, 1, 2]
输出: False
注意:

除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允许的。
你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/24-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
递归回溯
只有四张牌，四种运算， 总共的可能组合为：
先选两张牌，4*3，四种运算，可能有4*3*4种
已选2张牌运算后的结果与剩下两张牌组合，现在相当于有3张牌；选出其中两张的情况有3*2*4种
现在相当于有2张牌，有2*4种运算情况
综上，总共可能有4*3*4*3*2*4*2*4=9216种组合

时空复杂度都是O(1)
*/
func judgePoint24(nums []int) bool {
	floats := make([]float64, len(nums))
	for i, v := range nums {
		floats[i] = float64(v)
	}
	return judge(floats)
}

func judge(nums []float64) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return math.Abs(float64(nums[0])-24.0) < 1e-6
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			var tmp []float64
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					tmp = append(tmp, nums[k])
				}
			}
			added := false
			for k := 0; k < 4; k++ {
				iVal, jVal := float64(nums[i]), float64(nums[j])
				if k < 2 && j > i || k == 3 && jVal == 0 {
					continue
				}
				var r float64
				if k == 0 {
					r = iVal + jVal
				} else if k == 1 {
					r = iVal * jVal
				} else if k == 2 {
					r = iVal - jVal
				} else if k == 3 {
					r = iVal / jVal
				}
				if added && len(tmp) > 0 {
					tmp[len(tmp)-1] = r
				} else {
					tmp = append(tmp, r)
				}
				added = true
				if judge(tmp) {
					return true
				}
			}
		}
	}
	return false
}

// 参考官方题解，用了list的解法如下：
func judgePoint24_0(nums []int) bool {
	floats := list.New()
	for _, v := range nums {
		floats.PushBack(float64(v))
	}
	return solve(floats)
}

func solve(nums *list.List) bool {
	if nums.Len() == 0 {
		return false
	}
	if nums.Len() == 1 {
		return math.Abs(nums.Front().Value.(float64)-24.0) < 1e-6
	}
	elementI := nums.Front()
	for i := 0; i < nums.Len(); i++ {
		elementJ := nums.Front()
		for j := 0; j < nums.Len(); j++ {
			if i == j {
				elementJ = elementJ.Next()
				continue
			}
			tmp := list.New()
			e := nums.Front()
			for k := 0; k < nums.Len(); k++ {
				if k != i && k != j {
					tmp.PushBack(e.Value)
				}
				e = e.Next()
			}
			for k := 0; k < 4; k++ {
				iVal, jVal := elementI.Value.(float64), elementJ.Value.(float64)
				if k < 2 && j > i || k == 3 && jVal == 0 {
					continue
				}
				if k == 0 {
					tmp.PushBack(iVal + jVal)
				} else if k == 1 {
					tmp.PushBack(iVal * jVal)
				} else if k == 2 {
					tmp.PushBack(iVal - jVal)
				} else if k == 3 {
					tmp.PushBack(iVal / jVal)
				}
				if solve(tmp) {
					return true
				}
				tmp.Remove(tmp.Back())
			}
			elementJ = elementJ.Next()
		}
		elementI = elementI.Next()
	}
	return false
}

/*
另一个解法
*/
func judgePoint24_1(nums []int) bool {
	if len(nums) != 4 {
		return false
	}
	return is24for4nums(float64(nums[0]), float64(nums[1]), float64(nums[2]), float64(nums[3]))
}

func is24for4nums(a, b, c, d float64) bool {
	return is24for3nums(a+b, c, d) ||
		is24for3nums(a-b, c, d) ||
		is24for3nums(a*b, c, d) ||
		is24for3nums(a/b, c, d) ||
		is24for3nums(b-a, c, d) ||
		is24for3nums(b/a, c, d) ||
		is24for3nums(a+c, b, d) ||
		is24for3nums(a-c, b, d) ||
		is24for3nums(a*c, b, d) ||
		is24for3nums(a/c, b, d) ||
		is24for3nums(c-a, b, d) ||
		is24for3nums(c/a, b, d) ||
		is24for3nums(a+d, b, c) ||
		is24for3nums(a-d, b, c) ||
		is24for3nums(a*d, b, c) ||
		is24for3nums(a/d, b, c) ||
		is24for3nums(d-a, b, c) ||
		is24for3nums(d/a, b, c) ||
		is24for3nums(b+c, a, d) ||
		is24for3nums(b-c, a, d) ||
		is24for3nums(b*c, a, d) ||
		is24for3nums(b/c, a, d) ||
		is24for3nums(c-b, a, d) ||
		is24for3nums(c/b, a, d) ||
		is24for3nums(b+d, a, c) ||
		is24for3nums(b-d, a, c) ||
		is24for3nums(b*d, a, c) ||
		is24for3nums(b/d, a, c) ||
		is24for3nums(d-b, a, c) ||
		is24for3nums(d/b, a, c) ||
		is24for3nums(c+d, a, b) ||
		is24for3nums(c-d, a, b) ||
		is24for3nums(c*d, a, b) ||
		is24for3nums(c/d, a, b) ||
		is24for3nums(d-c, a, b) ||
		is24for3nums(d/c, a, b)
}

func is24for3nums(a, b, c float64) bool {
	return is24for2nums(a+b, c) ||
		is24for2nums(a-b, c) ||
		is24for2nums(a*b, c) ||
		is24for2nums(a/b, c) ||
		is24for2nums(b-a, c) ||
		is24for2nums(b/a, c) ||
		is24for2nums(a+c, b) ||
		is24for2nums(a-c, b) ||
		is24for2nums(a*c, b) ||
		is24for2nums(a/c, b) ||
		is24for2nums(c-a, b) ||
		is24for2nums(c/a, b) ||
		is24for2nums(b+c, a) ||
		is24for2nums(b-c, a) ||
		is24for2nums(b*c, a) ||
		is24for2nums(b/c, a) ||
		is24for2nums(c-b, a) ||
		is24for2nums(c/b, a)
}

func is24for2nums(a, b float64) bool {
	return is24(a+b) ||
		is24(a-b) ||
		is24(a*b) ||
		is24(a/b) ||
		is24(b/a)
}

func is24(a float64) bool {
	return math.Abs(a-24) < 1e-6
}
