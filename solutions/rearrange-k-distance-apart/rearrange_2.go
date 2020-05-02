/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rearrange_k_distance_apart

import "sort"

func reorganizeString(s string) string {
	cnt := make([]int, 26)
	result := []byte(s)
	for _, b := range result {
		cnt[b-'a'] += 100
	}
	for i := 0; i < 26; i++ {
		cnt[i] += i
	}
	sort.Ints(cnt)
	j := 1

	for _, c := range cnt {
		k := c / 100
		b := 'a' + byte(c%100)
		if k > (len(s)+1)/2 {
			return ""
		}
		for i := 0; i < k; i++ {
			if j >= len(s) {
				j = 0
			}
			result[j] = b
			j += 2
		}
	}
	return string(result)
}

// TODO:
