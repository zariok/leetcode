package main

// recursive search to find max depth
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1

}

// used for testing
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// used for testing
func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for i < len(vals) {
		current := queue[0]
		queue = queue[1:]
		if vals[i] != nil {
			current.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, current.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			current.Right = &TreeNode{Val: vals[i].(int)}

			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}
