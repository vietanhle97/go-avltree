package main

import (
	"fmt"
	. "go-avltree/avl"
	"math"
	"strconv"
	"time"
)

func max(a, b int) int { return int(math.Max(float64(a), float64(b))) }

func inOrder(root *TreeNode, res *[]int, child *[][]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res, child)
	*res = append(*res, root.Val)
	*child = append(*child, []int{root.Val, root.ChildCnt})
	inOrder(root.Right, res, child)
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

func printOrder(i int, n string) string {
	if 10 < i && i < 20 {
		return strconv.Itoa(i) + "th" + " " + n
	}
	if i%10 == 1 {
		return strconv.Itoa(i) + "st" + " " + n
	} else if i%10 == 2 {
		return strconv.Itoa(i) + "nd" + " " + n
	} else if i%10 == 3 {
		return strconv.Itoa(i) + "rd" + " " + n
	}
	return strconv.Itoa(i) + "th" + " " + n
}

func main() {
	traverse := make([]int, 0)
	child := make([][]int, 0)
	height := map[interface{}]int{}
	check := map[interface{}]int{}
	val := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384}
	start := time.Now()
	avl := insertAVLTree(nil, val)
	end := time.Now()
	inOrder(avl, &traverse, &child)
	levelOfAllSubTree(avl, 0, &height, &check)
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
	fmt.Println(traverse, height, avl.ChildCnt, end.Sub(start))
	fmt.Println(child)
	avl.Remove(16)
	avl.Remove(4)
	avl.Remove(8)
	traverse = []int{}
	height = map[interface{}]int{}
	child = [][]int{}

	inOrder(avl, &traverse, &child)
	levelOfAllSubTree(avl, 0, &height, &check)
	fmt.Println(traverse, height, avl.ChildCnt)
	fmt.Println(child)
	avl.Insert(16)
	avl.Insert(4)
	avl.Insert(8)
	traverse = []int{}
	height = map[interface{}]int{}
	child = [][]int{}

	inOrder(avl, &traverse, &child)
	levelOfAllSubTree(avl, 0, &height, &check)
	fmt.Println(traverse, height, avl.ChildCnt)
	fmt.Println(child)
	for i := 1; i < avl.ChildCnt+1; i++ {
		n := avl.FindKthSmallestValueNode(i)
		m := avl.FindKthLargestValueNode(i)
		if n != nil {
			fmt.Println(printOrder(i, "smallest"), n.Val, " ")
		} else {
			fmt.Print(n, " ")
		}
		if m != nil {
			fmt.Println(printOrder(i, "largest"), m.Val, " ")
		} else {
			fmt.Print(m, " ")
		}
	}
	fmt.Println("")
}
