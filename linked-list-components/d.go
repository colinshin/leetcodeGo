package linked_list_components

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
给定一个链表（链表结点包含一个整型值）的头结点 head。

同时给定列表 G，该列表是上述链表中整型值的一个子集。

返回列表 G 中组件的个数，这里对组件的定义为：链表中一段最长连续结点的值（该值必须在列表 G 中）构成的集合。

示例 1：

输入:
head: 0->1->2->3
G = [0, 1, 3]
输出: 2
解释:
链表中,0 和 1 是相连接的，且 G 中不包含 2，所以 [0, 1] 是 G 的一个组件，同理 [3] 也是一个组件，故返回 2。
示例 2：

输入:
head: 0->1->2->3->4
G = [0, 3, 1, 4]
输出: 2
解释:
链表中，0 和 1 是相连接的，3 和 4 是相连接的，所以 [0, 1] 和 [3, 4] 是两个组件，故返回 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/linked-list-components
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numComponents(head *ListNode, G []int) int {
	if len(G) == 0 || head == nil {
		return 0
	}

	set := make(map[int]struct{}, len(G))
	for _, v := range G {
		set[v] = struct{}{}
	}

	count := 0
	for p := head; p != nil; p = p.Next {
		_, isCurrentNodeInSet := set[p.Val]
		isNextNodeInSet := false
		if p.Next != nil {
			_, isNextNodeInSet = set[p.Next.Val]
		}

		if isCurrentNodeInSet && !isNextNodeInSet {
			count++
		}
	}
	return count
}