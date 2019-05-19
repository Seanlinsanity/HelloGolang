package main

import (
	"LearnGo/tree"
	"fmt"
)

//MyTreeNode binary tree
type MyTreeNode struct {
	node *tree.Node
}

func (myNode *MyTreeNode) postOrderTraverse() {
	if myNode == nil || myNode.node == nil {
		return
	}
	right := MyTreeNode{myNode.node.Right}
	right.postOrderTraverse()
	myNode.node.Print()
	left := MyTreeNode{myNode.node.Left}
	left.postOrderTraverse()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{Value: 4}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.
		Node)
	root.Left.Right = tree.CreateNode(2)
	//       3
	//    /     \
	//   4       5
	//    \     /
	//     2   0
	root.Traverse()

	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{Value: 6, Left: nil, Right: &root},
	}

	fmt.Printf("nodes: %v", nodes)
	fmt.Println()
	root.Print()

	root.Right.Left.SetValue(100)
	root.Right.Left.Print()
	fmt.Println()

	pRoot := &root
	pRoot.SetValue(30)
	root.Print()
	fmt.Println()

	root.Traverse()

	// myNode := MyTreeNode{&root}
	// myNode.postOrderTraverse()
	// fmt.Println()

}
