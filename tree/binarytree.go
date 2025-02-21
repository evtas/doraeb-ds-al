package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func BuildTree(arr []int) *Node {
	root := &Node{val: arr[0]}
	for i := 1; i < len(arr); i++ {
		Insert(root, arr[i])
	}
	return root
}

func Insert(root *Node, val int) {
	if val < root.val {
		if root.left != nil {
			Insert(root.left, val)
		} else {
			root.left = &Node{val: val}
		}
	} else {
		if root.right != nil {
			Insert(root.right, val)
		} else {
			root.right = &Node{val: val}
		}
	}
}

func PreOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	PreOrder(root.left)
	PreOrder(root.right)
}

func InOrder(root *Node) {
	if root == nil {
		return
	}
	InOrder(root.left)
	fmt.Println(root.val)
	InOrder(root.right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Println(root.val)
}

func LevelOrder(root *Node) {
	q := make([]*Node, 0)
	q = append(q, root)
	for {
		if len(q) == 0 {
			break
		}
		curNode := q[0]
		q = q[1:]
		if curNode.left != nil {
			q = append(q, curNode.left)
		}
		if curNode.right != nil {
			q = append(q, curNode.right)
		}
		fmt.Println(curNode.val)
	}
}

func main() {
	array := []int{4, 3, 2, 5, 1, 6, 8, 9}
	root := BuildTree(array)
	LevelOrder(root)
}
