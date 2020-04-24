/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package course_schedule_ii

import "container/list"

// bfs
func findOrder(numCourses int, prerequisites [][]int) []int {
	indegrees := make([]int, numCourses)
	nexts := make([][]int, numCourses)
	for _, req := range prerequisites {
		indegrees[req[0]]++
		nexts[req[1]] = append(nexts[req[1]], req[0])
	}
	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue.PushBack(i)
		}
	}
	var result []int
	for queue.Len() > 0 {
		course := queue.Remove(queue.Front()).(int)
		result = append(result, course)
		numCourses--
		for _, next := range nexts[course] {
			indegrees[next]--
			if indegrees[next] == 0 {
				queue.PushBack(next)
			}
		}
	}
	if numCourses == 0 {
		return result
	}
	return nil
}

// dfs
func findOrder11(numCourses int, prerequisites [][]int) []int {
	neighbors := make([][]int, numCourses)
	for _, req := range prerequisites {
		neighbors[req[0]] = append(neighbors[req[0]], req[1])
	}
	flags := make([]int, numCourses)
	var result []int
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if flags[course] == 1 {
			return false
		}
		if flags[course] == -1 {
			return true
		}
		flags[course] = 1
		for _, neighbor := range neighbors[course] {
			if !dfs(neighbor) {
				return false
			}
		}
		flags[course] = -1
		result = append(result, course)
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return nil
		}
	}
	return result
}
