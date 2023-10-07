package search_in_a_binary_search_tree

// topic 👉 https://leetcode.com/problems/search-in-a-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val == val:
		return root
	case root.Val > val:
		return searchBST(root.Left, val)
	case root.Val < val:
		return searchBST(root.Right, val)
	default:
		return nil
	}
}
