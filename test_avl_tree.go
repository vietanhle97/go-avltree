package main

import (
	"fmt"
	. "go-avltree/avl"
	"math"
	"time"
)

func max(a, b int) int { return int(math.Max(float64(a), float64(b))) }

func inOrder(root *TreeNode, res *[]interface{}) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

func levelOfAllSubTree(root *TreeNode, level int, res, check *map[interface{}]int) {
	if root == nil {
		return
	}
	levelOfAllSubTree(root.Left, level+1, res, check)
	levelOfAllSubTree(root.Right, level+1, res, check)
	(*res)[root.Val] = root.Height
	(*check)[root.Val] = level
}

func insertAVLTree(root *TreeNode, val []int) *TreeNode {
	for _, v := range val {
		if root == nil {
			root = &TreeNode{Val: v}
		} else {
			root = root.Insert(v)
		}
	}
	return root
}

func main() {
	//traverse := make([]interface{}, 0)
	//height := map[interface{}]int{}
	//check := map[interface{}]int{}
	val := []int{1, 2, 3, 4, 6, 7, 8, 11, 14, 15}
	start := time.Now()
	avl := insertAVLTree(nil, val)
	end := time.Now()
	//inOrder(avl, &traverse)
	//levelOfAllSubTree(avl, 0, &height, &check)
	ceiling, floor := avl.Ceiling(16), avl.Floor(10)
	node := avl.Find(4)
	if node == nil {
		fmt.Println(node)
	} else {
		fmt.Println(node.Val)
	}
	if ceiling != nil {
		fmt.Println(ceiling.Val)
	} else {
		fmt.Println(ceiling)
	}
	if floor != nil {
		fmt.Println(floor.Val)
	} else {
		fmt.Println(floor)
	}
	fmt.Println(avl.CountNodes(), end.Sub(start))
}
