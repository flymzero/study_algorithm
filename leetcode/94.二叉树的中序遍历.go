package leetcode

/*
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
给定一个二叉树，返回它的中序 遍历。
输入: [1,null,2,3]
1
\
2
/
3

输出: [1,3,2]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 测试用例
var exampleMode = &TreeNode{
	Val: 13,
	Left: &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   11,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 17,
		Left: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 25,
			Left: &TreeNode{
				Val:   22,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   27,
				Left:  nil,
				Right: nil,
			},
		},
	},
}

// 递归
func inorderTraversal1(root *TreeNode) []int {
	var result = []int{}
	if root != nil {
		// 左中右递归
		result = append(result, inorderTraversal1(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderTraversal1(root.Right)...)
	}
	return result
}

// 基于栈的遍历
func inorderTraversal2(root *TreeNode) []int {
	var res = []int{}
	//
	var stack = []*TreeNode{}
	var curNode = root
	// 栈空间不为空或者当前的节点不为空
	for len(stack) > 0 || curNode != nil {
		// 当前的节点不为空时，把它的左节点作为下一个根节点
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		} else { // 栈空间不为空, 说明上一次节点为空，取出栈中的节点，把它的右节点作为根节点
			lastNode := stack[len(stack)-1]
			res = append(res, lastNode.Val)
			stack = stack[:len(stack)-1]
			curNode = lastNode.Right
		}
	}
	return res
}
