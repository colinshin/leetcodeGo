/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package course_schedule

import (
	"container/list"
)

/*
207. 课程表 https://leetcode-cn.com/problems/course-schedule
现在你总共有 n 门课需要选，记为 0 到 n-1。
在选修某些课程之前需要一些先修课程。
例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]

给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？

示例 1:
输入: 2, [[1,0]]
输出: true
解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。

示例 2:
输入: 2, [[1,0],[0,1]]
输出: false
解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。

说明:
输入的先决条件是由边缘列表表示的图形，而不是邻接矩阵。详情请参见图的表示法。
你可以假定输入的先决条件中没有重复的边。

提示:
这个问题相当于查找一个循环是否存在于有向图中。如果存在循环，则不存在拓扑排序，因此不可能选取所有课程进行学习。
通过 DFS 进行拓扑排序 - 一个关于Coursera的精彩视频教程（21分钟），介绍拓扑排序的基本概念。
拓扑排序也可以通过 BFS 完成。
*/

/*
方法1：入度表（广度优先遍历）

算法流程：
统计课程安排图中每个节点的入度，生成 入度表 indegrees。
借助一个队列 queue，将所有入度为0的节点入队。
当 queue 非空时，依次将队首节点出队，在课程安排图中删除此节点 pre：
并不是真正从邻接表中删除此节点 pre，而是将此节点对应所有邻接节点 cur 的入度−1，即 indegrees[cur] -= 1。
当入度−1后邻接节点 cur 的入度为0，说明 cur 所有的前驱节点已经被 “删除”，此时将 cur 入队。
在每次 pre 出队时，执行 numCourses--；
若整个课程安排图是有向无环图（即可以安排），则所有节点一定都入队并出队过，即完成拓扑排序。
换个角度说，若课程安排图中存在环，一定有节点的入度始终不为0。
因此，拓扑排序出队次数等于课程个数，返回 numCourses == 0 判断课程是否可以成功安排。
复杂度分析：
时间复杂度
O(N+M)，遍历一个图需要访问所有节点和所有临边，N 和 M 分别为节点数量和临边数量；
空间复杂度O(N)，为建立邻接矩阵所需额外空间。

作者：jyd
链接：https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses) // 记录每门课程前置应修的课程数
	for _, req := range prerequisites {
		indegrees[req[0]] += 1
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
			indegrees[req[0]]--         // course修过了，依赖course的课程也可以修了
			if indegrees[req[0]] == 0 { // 前置课程都修完了
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
	indegrees := make([]int, numCourses) // 记录每门课程前置应修的课程数
	nexts := make([][]int, numCourses)   //记录每门课的后修课程
	for _, req := range prerequisites {
		indegrees[req[0]] += 1
		nexts[req[1]] = append(nexts[req[1]], req[0])
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
		for _, next := range nexts[course] {
			indegrees[next]--         // course修过了，依赖course的课程也可以修了
			if indegrees[next] == 0 { // 前置课程都修完了
				queue.PushBack(next)
			}
		}
	}
	return numCourses == 0
}

/*
方法2：深度优先遍历

算法流程（思路是通过 DFS 判断图中是否有环）：
借助一个标志列表 flags，用于判断每个节点 i （课程）的状态：
未被 DFS 访问：i == 0；
已被其他节点启动的DFS访问：i == -1；
已被当前节点启动的DFS访问：i == 1。
对 numCourses 个节点依次执行 DFS，判断每个节点起步 DFS 是否存在环，若存在环直接返回False。
DFS 流程；
终止条件：
当 flag[i] == -1，说明当前访问节点已被其他节点启动的 DFS 访问，无需再重复搜索，直接返回True。
当 flag[i] == 1，说明在本轮 DFS 搜索中节点 i 被第2 次访问，即 课程安排图有环，直接返回False。
将当前访问节点 i 对应 flag[i] 置1，即标记其被本轮 DFS 访问过；
递归访问当前节点 i 的所有邻接节点 j，当发现环直接返回False；
当前节点所有邻接节点已被遍历，并没有发现环，则将当前节点 flag 置为−1 并返回True。
若整个图 DFS 结束并未发现环，返回True。

复杂度同方法1

作者：jyd
链接：https://leetcode-cn.com/problems/course-schedule/solution/course-schedule-tuo-bu-pai-xu-bfsdfsliang-chong-fa/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
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
