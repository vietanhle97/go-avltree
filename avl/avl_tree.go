package avl

import (
	"math"
)

func max(a, b int) int { return int(math.Max(float64(a), float64(b))) }

func abs(a int) int { return int(math.Abs(float64(a))) }

type TreeNode struct {
	Val    int
	Height int
	Left   *TreeNode
	Right  *TreeNode
}

// calculate height of left subtree
func (node *TreeNode) leftSubTreeHeight() int {
	if node.Left == nil {
		return -1
	}
	return node.Left.Height
}

// calculate height of right subtree
func (node *TreeNode) rightSubTreeHeight() int {
	if node.Right == nil {
		return -1
	}
	return node.Right.Height
}

// update height of current root and check whether it is balanced or not
func (node *TreeNode) reComputeHeight() bool {
	ls := node.leftSubTreeHeight()
	rs := node.rightSubTreeHeight()
	node.Height = max(ls, rs) + 1
	return abs(ls-rs) > 1
}

// left rotate the unbalance subtree
func (node *TreeNode) rotateLeft() *TreeNode {
	newRoot := node.Right
	tmp := node.Right.Left
	node.Right = tmp
	node.reComputeHeight()
	newRoot.Left = node
	newRoot.reComputeHeight()
	return newRoot
}

// right rotate the unbalance subtree
func (node *TreeNode) rotateRight() *TreeNode {
	newRoot := node.Left
	tmp := node.Left.Right
	node.Left = tmp
	node.reComputeHeight()
	newRoot.Right = node
	newRoot.reComputeHeight()
	return newRoot
}

// rebalance the tree
func (node *TreeNode) reBalance() *TreeNode {
	if node.rightSubTreeHeight() > node.leftSubTreeHeight() {
		if node.Right.leftSubTreeHeight() > node.Right.leftSubTreeHeight() {
			node.Right = node.Right.rotateRight()
		}
		return node.rotateLeft()
	} else {
		if node.Left.rightSubTreeHeight() > node.leftSubTreeHeight() {
			node.Left = node.Left.rotateLeft()
		}
		return node.rotateRight()
	}
}

// find node with min val of the subtree
func (node *TreeNode) findMinValue() *TreeNode {
	tmp := node
	for tmp.Left != nil {
		tmp = tmp.Left
	}
	return tmp
}

// find node with max val of the subtree
func (node *TreeNode) findMaxValue() *TreeNode {
	tmp := node
	for tmp.Right != nil {
		tmp = tmp.Right
	}
	return tmp
}

// remove the node with min val in the subtree
func (node *TreeNode) removeMinValue() *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	node.Left = node.Left.removeMinValue()
	if node.reComputeHeight() {
		return node.reBalance()
	}
	return node
}

//traverse the tree with post order to count the total nodes
func (node *TreeNode) postOrder(cur []int, res *int) {
	if node == nil {
		return
	}
	ls, rs := []int{0}, []int{0}
	node.Left.postOrder(ls, res)
	node.Right.postOrder(rs, res)
	cur[0] = ls[0] + rs[0] + 1
	*res = max(*res, cur[0])
}

// Find the node with the maximum val smaller than the given val
func (node *TreeNode) findMaximumSmallerOrEqual(val int) *TreeNode {
	if val == node.Val {
		return node
	} else if val > node.Val {
		if node.Right == nil {
			return node
		}
		rs := node.Right.findMaximumSmallerOrEqual(val)
		if rs.Val > val {
			return node
		}
		return rs
	} else {
		if node.Left == nil {
			return node
		}
		return node.Left.findMaximumSmallerOrEqual(val)

	}
}

// Find the node with minimum val larger than or equal given val
func (node *TreeNode) findMinimumLargerOrEqual(val int) *TreeNode {
	if val == node.Val {
		return node
	} else if val < node.Val {
		if node.Left == nil {
			return node
		}
		ls := node.Left.findMinimumLargerOrEqual(val)
		if ls.Val < val {
			return node
		}
		return ls
	} else {
		if node.Right == nil {
			return node
		}
		return node.Right.findMinimumLargerOrEqual(val)
	}
}

// Export function
// find node with given val
func (node *TreeNode) Find(val int) *TreeNode {
	if val == node.Val {
		return node
	}
	if val < node.Val {
		if node.Left != nil {
			return node.Left.Find(val)
		}
		return nil
	} else {
		if node.Right != nil {
			return node.Right.Find(val)
		}
		return nil
	}
}

// Export function
// Insert new node with val and value
func (node *TreeNode) Insert(val int) *TreeNode {
	if node == nil {
		return &TreeNode{val, 0, nil, nil}
	}
	if val == node.Val {
		return node
	}
	if val < node.Val {
		if node.Left == nil {
			node.Left = &TreeNode{val, 0, nil, nil}
		} else {
			node.Left = node.Left.Insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{val, 0, nil, nil}
		} else {
			node.Right = node.Right.Insert(val)
		}
	}
	if node.reComputeHeight() {
		return node.reBalance()
	}
	return node
}

// Export function
// Remove a node from the sub-tree
func (node *TreeNode) Remove(val int) *TreeNode {
	if val < node.Val && node.Left != nil {
		node.Left = node.Left.Remove(val)
	} else if val > node.Val && node.Right != nil {
		node.Right = node.Right.Remove(val)
	} else if val == node.Val {
		if node.Left != nil && node.Right != nil {
			newRoot := node.Right.findMinValue()
			node.Val = newRoot.Val
			node.Right = node.Right.removeMinValue()
		} else {
			if node.Left != nil {
				return node.Left
			}
			return node.Right
		}
	}
	if node.reComputeHeight() {
		return node.reBalance()
	}
	return node
}

// Export function
// Count the total of nodes in the sub-tree
func (node *TreeNode) CountNodes() int {
	res := 0
	cur := []int{0}
	node.postOrder(cur, &res)
	return res
}

// Export function
// Find the node with minimum val larger than or equal given val
func (node *TreeNode) Ceiling(val int) *TreeNode {
	res := node.findMinimumLargerOrEqual(val)
	if res.Val < val {
		return nil
	}
	return res
}

// Export function
// Find the node with the maximum val smaller than the given val
func (node *TreeNode) Floor(val int) *TreeNode {
	res := node.findMaximumSmallerOrEqual(val)
	if res.Val > val {
		return nil
	}
	return res
}
