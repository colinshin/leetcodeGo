/*
 * Copyright (c) zrcoder 2019-2020. All rights reserved.
 */

package clone_graph

import "container/list"

/*
133. 克隆图 https://leetcode-cn.com/problems/clone-graph

给你无向 连通 图中一个节点的引用，请你返回该图的 深拷贝（克隆）。

图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。

class Node {
    public int val;
    public List<Node> neighbors;
}


测试用例格式：
简单起见，每个节点的值都和它的索引相同。例如，第一个节点值为 1（val = 1），第二个节点值为 2（val = 2），以此类推。该图在测试用例中使用邻接列表表示。
邻接列表 是用于表示有限图的无序列表的集合。每个列表都描述了图中节点的邻居集。
给定节点将始终是图中的第一个节点（值为 1）。你必须将 给定节点的拷贝 作为对克隆图的引用返回。

示例 1：
输入：adjList = [[2,4],[1,3],[2,4],[1,3]]
输出：[[2,4],[1,3],[2,4],[1,3]]
解释：
图中有 4 个节点。
节点 1 的值是 1，它有两个邻居：节点 2 和 4 。
节点 2 的值是 2，它有两个邻居：节点 1 和 3 。
节点 3 的值是 3，它有两个邻居：节点 2 和 4 。
节点 4 的值是 4，它有两个邻居：节点 1 和 3 。
示例 2：
输入：adjList = [[]]
输出：[[]]
解释：输入包含一个空列表。该图仅仅只有一个值为 1 的节点，它没有任何邻居。
示例 3：
输入：adjList = []
输出：[]
解释：这个图是空的，它不含任何节点。
示例 4：
输入：adjList = [[2],[1]]
输出：[[2],[1]]
提示：
节点数不超过 100 。
每个节点值 Node.val 都是唯一的，1 <= Node.val <= 100。
无向图是一个简单图，这意味着图中没有重复的边，也没有自环。
由于图是无向的，如果节点 p 是节点 q 的邻居，那么节点 q 也必须是节点 p 的邻居。
图是连通图，你可以从给定节点访问到所有节点。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/clone-graph
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type Node struct {
	Val       int
	Neighbors []*Node
}

/*
把图看成一棵树会比较好理解
要注意和树不同的是，孩子节点可能反过来成为父节点，递归或遍历会形成无限循环
需要用额外的数据结构如哈希表来记录节点是否访问过
直观的做法是哈希表定义为键为*Node值为bool，这样不太好处理；想想，改成值为原图中节点，键为新图中节点，会比较好。
*/
/*
递归DFS

假设所有节点个数为n，
时间复杂度O(n)，每个节点处理一次，栈调用时间复杂度O(H),H为图的最大深度，综合复杂度O(n)
空间复杂度O(n)，哈希表需要O(n)，栈需要O(H)
*/
func cloneGraph1(node *Node) *Node {
	visited := make(map[*Node]*Node, 0)
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if r, ok := visited[n]; ok {
			return r
		}
		r := &Node{Val: n.Val, Neighbors: make([]*Node, len(n.Neighbors))}
		visited[n] = r
		for i, v := range n.Neighbors {
			r.Neighbors[i] = dfs(v)
		}
		return r
	}
	return dfs(node)
}

/*
迭代BFS，存放临时节点的容器可随意选用，这里选list
时间复杂度O(n)，每个节点处理一次
空间复杂度O(n)，哈希表需要O(n), BFS使用的容器需要O(W)，其中W是图的宽度， 综合复杂度O(n)
*/
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(node)
	visited := make(map[*Node]*Node, 0)
	visited[node] = &Node{Val: node.Val, Neighbors: make([]*Node, len(node.Neighbors))}
	for queue.Len() > 0 {
		n := queue.Remove(queue.Front()).(*Node)
		for i, v := range n.Neighbors {
			if _, ok := visited[v]; !ok {
				queue.PushBack(v)
				visited[v] = &Node{Val: v.Val, Neighbors: make([]*Node, len(v.Neighbors))}
			}
			visited[n].Neighbors[i] = visited[v]
		}
	}
	return visited[node]
}
