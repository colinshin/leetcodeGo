/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rearrange_k_distance_apart

import (
	"sort"
)

type Pair struct {
	count int
	char  byte
}

func rearrangeString(s string, k int) string {
	if k <= 1 {
		return s
	}
	result := []byte(s)
	pairs := make([]Pair, 26)
	for _, b := range result {
		pairs[b-'a'].char = b
		pairs[b-'a'].count++
	}
	cmp := func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].char < pairs[j].char
		}
		return pairs[i].count > pairs[j].count
	}
	sort.Slice(pairs, cmp)

	j := 0
	for pairs[0].count > 0 {
		for i := 0; check(i, k, pairs); i++ {
			if i >= len(pairs) || pairs[i].count == 0 {
				return ""
			}
			result[j] = pairs[i].char
			j++
			pairs[i].count--
		}
		sort.Slice(pairs, cmp)
	}
	return string(result)
}

func check(i, k int, pairs []Pair) bool {
	return i < k && pairs[0].count > 0
}
