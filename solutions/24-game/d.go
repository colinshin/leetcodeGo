/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package point_game

func judgePoint24(nums []int) bool {
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
