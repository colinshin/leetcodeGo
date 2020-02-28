/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package countsort

import "sort"

// 计数排序s，假设s中元素值域为[min, max], 元素可以有负数，非原地的计数排序
func countSort(s []int) []int {
	if len(s) == 0 {
		return s
	}
	count := make(map[int]int, 0)
	for _, v := range s { // 统计s中每个元素出现的个数
		count[v]++
	}
	sum := 0
	rangeAsSorted(count, func(k, v int) {
		sum += v
		count[k] = sum
	})
	result := make([]int, len(s))
	for _, num := range s {
		index := count[num] - 1
		result[index] = num
		count[num]-- // 如果有重复的元素i，下一次插入的位置是当前插入位置的前一位
	}
	return result
}

func rangeAsSorted(m map[int]int, f func(k, v int)) {
	keys := make([]int, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	sort.Ints(keys)
	for _, key := range keys {
		f(key, m[key])
	}
}
