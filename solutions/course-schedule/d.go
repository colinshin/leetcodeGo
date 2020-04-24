/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package course_schedule

import (
	"container/list"
)

/*
207. 课程表 https://leetcode-cn.com/problems/course-schedule
你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
在选修某些课程之前需要一些先修课程。
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？

示例 1:
输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。

示例 2:
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。

提示：
输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。
1 <= numCourses <= 10^5
*/

/*
这是图处理中比较经典拓扑排序问题，可以用BFS或DFS

方法1：BFS（广度优先遍历）

先修不依赖任何课程的课程C1s（可能多个），再修直接依赖C1s的课程C2s...;
为了方便，可以事先统计一把各门课依赖的前置课程的数量，称为入度；C1s就是入度为0的课程。
将入度为0的课程加入一个集合（如队列），一一出队，每出一个课程相当于修了这门课，
这时候需要遍历一下入度统计结果，将依赖这门课的课入度减1，如果发现入度为0了则入队

循环执行，最终队列空后，如果还有课程的入度不为0， 说明这些课程有循环依赖，如a依赖b，b依赖c，c依赖a
显然，队列里每出队一门课程，相当于修了该门课，这时候可以把总课程数减1；当队列为空后判断课程总是是不是0即可

时间复杂度
O(N+M)，遍历一个图需要访问所有节点和所有临边，N 和 M 分别为节点数量和临边数量；
空间复杂度O(N)，为建立入度表所需额外空间。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses) // 记录每门课程前置应修的课程数
	for _, req := range prerequisites {
		indegrees[req[0]]++
	}
	queue := list.New() // 装入度为0的课程，即没有依赖可直接修的课程
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		course := queue.Remove(queue.Front()).(int)
		numCourses-- // 修course这门课
		for _, req := range prerequisites {
			if req[1] != course {
				continue
			}
			// course修过了，依赖course的课程也可以修了
			indegrees[req[0]]--
			// req[0]这门课的前置课程都修完了
			if indegrees[req[0]] == 0 {
				queue.PushBack(req[0])
			}
		}
	}
	return numCourses == 0
}

/*
内层循环稍稍有点浪费，可以事先统计各个课程的后修课程
*/
func canFinish1(numCourses int, prerequisites [][]int) bool {
	// 记录每门课程前置应修的课程数
	indegrees := make([]int, numCourses)
	// 记录每门课的后修课程
	nexts := make([][]int, numCourses)
	for _, req := range prerequisites {
		indegrees[req[0]] += 1
		nexts[req[1]] = append(nexts[req[1]], req[0])
	}
	queue := list.New()
	// 装入度为0的课程，即没有依赖可直接修的课程
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			queue.PushBack(i)
		}
	}
	for queue.Len() > 0 {
		course := queue.Remove(queue.Front()).(int)
		numCourses-- // 修course这门课
		for _, next := range nexts[course] {
			// course修过了，依赖course的课程也可以修了
			indegrees[next]--
			if indegrees[next] == 0 { // 前置课程都修完了
				queue.PushBack(next)
			}
		}
	}
	return numCourses == 0
}

/*
方法2：深度优先遍历

针对每门课做dfs；借助一个标志列表 flags，用于判断每个节点 i （课程）的状态：

未被 DFS 访问：i == 0；
已被其他节点启动的DFS访问：i == -1；
已被当前节点启动的DFS访问：i == 1。
对所有课程执行 DFS，判断每个课程的起步 DFS 是否存在环，若存在环直接返回False。

DFS 流程：
终止条件：
当 flag[i] == -1，说明当前访问节点已被其他节点启动的 DFS 访问，无需再重复搜索，直接返回True。
当 flag[i] == 1，说明在本轮 DFS 搜索中节点 i 被第2 次访问，即 课程安排图有环，直接返回False。
将当前访问节点 i 对应 flag[i] 置1，即标记其被本轮 DFS 访问过；
递归访问当前节点 i 的所有邻接节点 j，当发现环直接返回False；
当前节点所有邻接节点已被遍历，并没有发现环，则将当前节点 flag 置为−1 并返回True。
若整个图 DFS 结束并未发现环，返回True。

复杂度同方法1
*/
func canFinish10(numCourses int, prerequisites [][]int) bool {
	flags := make([]int, numCourses)
	var dfs func(course int) bool
	dfs = func(course int) bool {
		if flags[course] == 1 {
			return false
		}
		if flags[course] == -1 {
			return true
		}
		flags[course] = 1
		for _, req := range prerequisites {
			if req[0] == course && !dfs(req[1]) {
				return false
			}
			/* 以下写法也是对的；两种写法都是对course邻居做判断
			if req[1] == course && !dfs(req[0], prerequisites, flags) {
				return false
			}*/
		}
		flags[course] = -1
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}

/*
dfs里的循环稍微有点浪费，可以事先统计好每个节点的邻居
*/
func canFinish11(numCourses int, prerequisites [][]int) bool {
	neighbors := make([][]int, numCourses)
	for _, req := range prerequisites {
		// 写成 neighbors[req[1]] = append(neighbors[req[1]], req[0]) 也对，都是统计邻居
		neighbors[req[0]] = append(neighbors[req[0]], req[1])
	}
	flags := make([]int, numCourses)
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
		return true
	}
	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}
	return true
}
