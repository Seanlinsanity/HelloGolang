package tree

import "fmt"

//Traverse a node
func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
	// if node == nil {
	// 	return
	// }

	// node.Left.Traverse()
	// node.Print()
	// node.Right.Traverse()
}

//TraverseFunc :traverse and implement custom function
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
