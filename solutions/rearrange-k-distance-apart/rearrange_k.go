/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package rearrange_k_distance_apart

import (
	"container/heap"
	"container/list"
	"sort"
)

func rearrangeString(s string, k int) string {
	if k <= 1 {
		return s
	}
	result := []byte(s)
	pairs := count(result)
	cmp := func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].char < pairs[j].char
		}
		return pairs[i].count > pairs[j].count
	}
	sort.Slice(pairs, cmp)
	j := 0
	for pairs[0].count > 0 {
		for i := 0; i < k; i++ {
			if pairs[0].count == 0 {
				break
			}
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

func rearrangeString1(s string, k int) string {
	result := []byte(s)
	pairs := count(result)
	h := &PairHeap{}
	for _, pair := range pairs {
		if pair.count > 0 {
			heap.Push(h, pair)
		}
	}
	j := 0
	queue := list.New()
	for h.Len() > 0 {
		pair := heap.Pop(h).(Pair)
		result[j] = pair.char
		j++
		pair.count--
		queue.PushBack(pair)
		if queue.Len() >= k {
			pair = queue.Remove(queue.Front()).(Pair)
			if pair.count > 0 {
				heap.Push(h, pair)
			}
		}
	}
	if j != len(s) {
		return ""
	}
	return string(result)
}

func count(s []byte) []Pair {
	pairs := make([]Pair, 26)
	for _, b := range s {
		pairs[b-'a'].char = b
		pairs[b-'a'].count++
	}
	return pairs
}

type Pair struct {
	count int
	char  byte
}

type PairHeap []Pair

func (h PairHeap) Len() int { return len(h) }
func (h PairHeap) Less(i, j int) bool {
	if h[i].count == h[j].count {
		return h[i].char < h[j].char
	}
	return h[i].count > h[j].count
}
func (h PairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *PairHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}
