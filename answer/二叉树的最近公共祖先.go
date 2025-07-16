package answer

import "fmt"

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func BuildTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		// 左孩子
		if i < len(nums) && nums[i] != nil {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// 右孩子
		if i < len(nums) && nums[i] != nil {
			node.Right = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func FindTree(root *TreeNode) (*TreeNode, *TreeNode) {
	// 辅助函数：在树中查找值为 target 的节点
	var findNode func(*TreeNode, int) *TreeNode
	findNode = func(node *TreeNode, target int) *TreeNode {
		if node == nil {
			return nil
		}
		if node.Val == target {
			return node
		}
		left := findNode(node.Left, target)
		if left != nil {
			return left
		}
		return findNode(node.Right, target)
	}

	p := findNode(root, 7)
	q := findNode(root, 4)

	if p == nil || q == nil {
		fmt.Println("节点不存在")
		return nil, nil
	}
	return p, q
}
