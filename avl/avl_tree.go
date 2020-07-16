package avl

import (
	"math"
)

func max(a, b int) int { return int(math.Max(float64(a), float64(b))) }

func abs(a int) int { return int(math.Abs(float64(a))) }

type TreeNode struct {
	Val      int
	Height   int
	ChildCnt int // Count the number of root of the sub-tree which the current root is root
	Left     *TreeNode
	Right    *TreeNode
}

// calculate height of left subtree
func (root *TreeNode) leftSubTreeHeightAndNodeCount() []int {
	if root.Left == nil {
		return []int{-1, 0}
	}
	return []int{root.Left.Height, root.Left.ChildCnt}
}

// calculate height of right subtree
func (root *TreeNode) rightSubTreeHeightAndNodeCount() []int {
	if root.Right == nil {
		return []int{-1, 0}
	}
	return []int{root.Right.Height, root.Right.ChildCnt}
}

// update height of current root and check whether it is balanced or not
func (root *TreeNode) reComputeHeight() bool {
	ls := root.leftSubTreeHeightAndNodeCount()
	rs := root.rightSubTreeHeightAndNodeCount()
	root.Height = max(ls[0], rs[0]) + 1
	root.ChildCnt = ls[1] + rs[1] + 1
	return abs(ls[0]-rs[0]) > 1
}

// left rotate the unbalance subtree
func (root *TreeNode) rotateLeft() *TreeNode {
	newRoot := root.Right
	tmp := root.Right.Left
	root.Right = tmp
	root.reComputeHeight()
	newRoot.Left = root
	newRoot.reComputeHeight()
	return newRoot
}

// right rotate the unbalance subtree
func (root *TreeNode) rotateRight() *TreeNode {
	newRoot := root.Left
	tmp := root.Left.Right
	root.Left = tmp
	root.reComputeHeight()
	newRoot.Right = root
	newRoot.reComputeHeight()
	return newRoot
}

// re-balance the tree
func (root *TreeNode) reBalance() *TreeNode {
	ls := root.leftSubTreeHeightAndNodeCount()
	rs := root.rightSubTreeHeightAndNodeCount()
	if rs[0] > ls[0] {
		rLS := root.Right.leftSubTreeHeightAndNodeCount()  // left sub-tree of right sub-tree
		rRS := root.Right.rightSubTreeHeightAndNodeCount() // right sub-tree of right sub-tree
		if rLS[0] > rRS[0] {
			root.Right = root.Right.rotateRight()
		}
		return root.rotateLeft()
	} else {
		lLS := root.Left.leftSubTreeHeightAndNodeCount()  // left sub-tree of left sub-tree
		lRS := root.Left.rightSubTreeHeightAndNodeCount() // right sub-tree of left sub-tree
		if lRS[0] > lLS[0] {
			root.Left = root.Left.rotateLeft()
		}
		return root.rotateRight()
	}
}

// find root with min val of the subtree
func (root *TreeNode) findMinValue() *TreeNode {
	tmp := root
	for tmp.Left != nil {
		tmp = tmp.Left
	}
	return tmp
}

// find root with max val of the subtree
func (root *TreeNode) findMaxValue() *TreeNode {
	tmp := root
	for tmp.Right != nil {
		tmp = tmp.Right
	}
	return tmp
}

// remove the root with min val in the subtree
func (root *TreeNode) removeMinValue() *TreeNode {
	if root.Left == nil {
		return root.Right
	}
	root.Left = root.Left.removeMinValue()
	if root.reComputeHeight() {
		return root.reBalance()
	}
	return root
}

// Find the root with the maximum val smaller than the given val
// If not found, return the root with the minimum val
func (root *TreeNode) findMaximumSmallerOrEqual(val int) *TreeNode {
	if val == root.Val {
		return root
	} else if val > root.Val {
		if root.Right == nil {
			return root
		}
		rs := root.Right.findMaximumSmallerOrEqual(val)
		if rs.Val > val {
			return root
		}
		return rs
	} else {
		if root.Left == nil {
			return root
		}
		return root.Left.findMaximumSmallerOrEqual(val)

	}
}

// Find the root with minimum val larger than or equal given val
// If not found, return the root with the maximum val
func (root *TreeNode) findMinimumLargerOrEqual(val int) *TreeNode {
	if val == root.Val {
		return root
	} else if val < root.Val {
		if root.Left == nil {
			return root
		}
		ls := root.Left.findMinimumLargerOrEqual(val)
		if ls.Val < val {
			return root
		}
		return ls
	} else {
		if root.Right == nil {
			return root
		}
		return root.Right.findMinimumLargerOrEqual(val)
	}
}

// Export function
// find root with given val
func (root *TreeNode) Find(val int) *TreeNode {
	if val == root.Val {
		return root
	}
	if val < root.Val {
		if root.Left != nil {
			return root.Left.Find(val)
		}
		return nil
	} else {
		if root.Right != nil {
			return root.Right.Find(val)
		}
		return nil
	}
}

// Export function
// Insert new root with val and value
func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, 0, 1, nil, nil}
	}
	if val == root.Val {
		return root
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{val, 0, 1, nil, nil}
		} else {
			root.Left = root.Left.Insert(val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{val, 0, 1, nil, nil}
		} else {
			root.Right = root.Right.Insert(val)
		}
	}
	if root.reComputeHeight() {
		return root.reBalance()
	}
	return root
}

// Export function
// Remove a root from the sub-tree
func (root *TreeNode) Remove(val int) *TreeNode {
	if val < root.Val && root.Left != nil {
		root.Left = root.Left.Remove(val)
	} else if val > root.Val && root.Right != nil {
		root.Right = root.Right.Remove(val)
	} else if val == root.Val {
		if root.Left != nil && root.Right != nil {
			newRoot := root.Right.findMinValue()
			root.Val = newRoot.Val
			root.Right = root.Right.removeMinValue()
		} else {
			if root.Left != nil {
				return root.Left
			}
			return root.Right
		}
	}
	if root.reComputeHeight() {
		return root.reBalance()
	}
	return root
}

// Export function
// Find the root with minimum val larger than or equal given val
func (root *TreeNode) Ceiling(val int) *TreeNode {
	res := root.findMinimumLargerOrEqual(val)
	if res.Val < val {
		return nil
	}
	return res
}

// Export function
// Find the root with the maximum val smaller than the given val
func (root *TreeNode) Floor(val int) *TreeNode {
	res := root.findMaximumSmallerOrEqual(val)
	if res.Val > val {
		return nil
	}
	return res
}

// Export function
// Find kth smallest node
func (root *TreeNode) FindKthSmallestValueNode(k int) *TreeNode {
	if root == nil || k <= 0 {
		return nil
	}
	if root.ChildCnt < k {
		return nil
	} else if root.ChildCnt == k {
		return root.findMaxValue()
	} else if k == 1 {
		return root.findMinValue()
	} else {
		if root.Left != nil && root.Left.ChildCnt == k-1 {
			return root
		}
		ls := root.Left.FindKthSmallestValueNode(k)
		if ls != nil {
			return ls
		}
		if root.Left == nil {
			return root.Right.FindKthSmallestValueNode(k - 1)
		}
		return root.Right.FindKthSmallestValueNode(k - root.Left.ChildCnt - 1)
	}
}

// Export function
// Find kth largest node
func (root *TreeNode) FindKthLargestValueNode(k int) *TreeNode {
	if root == nil || k <= 0 {
		return nil
	}
	if root.ChildCnt < k {
		return nil
	} else if root.ChildCnt == k {
		return root.findMinValue()
	} else if k == 1 {
		return root.findMaxValue()
	} else {
		if root.Right != nil && root.Right.ChildCnt == k-1 {
			return root
		}
		rs := root.Right.FindKthLargestValueNode(k)
		if rs != nil {
			return rs
		}
		if root.Right == nil {
			return root.Left.FindKthLargestValueNode(k - 1)
		}
		return root.Left.FindKthLargestValueNode(k - root.Right.ChildCnt - 1)
	}
}
