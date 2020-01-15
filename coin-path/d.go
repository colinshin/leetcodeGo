package coin_path

import "math"

func cheapestJump(A []int, B int) []int {
	n := len(A)
	if n == 0 || B < 1 || A[0] == -1 {
		return nil
	}
	if B == 1 {
		var r []int
		for i := 0; i < n; i++ {
			if A[i] == -1 {
				return nil
			}
			r = append(r, i+1)
		}
		return r
	}

	type item struct {
		path    []int
		minCost int
	}
	dp := make([]item, n+1)
	dp[1] = item{path: []int{1}, minCost: A[0]}
	for i := 1; i < n; i++ {
		if A[i] == -1 {
			continue
		}

		dpIndex := i + 1
		minCost := math.MaxInt64
		lastIndex := -1
		for j := min(1, dpIndex-B); j < dpIndex; j++ {
			if dp[j].path != nil && dp[j].minCost < minCost {
				minCost = dp[j].minCost
				lastIndex = j
			}
		}
		if lastIndex == -1 {
			continue
		}
		dp[dpIndex].minCost = minCost + A[i]
		dp[dpIndex].path = append(dp[lastIndex].path, dpIndex)
	}

	return dp[n].path
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
