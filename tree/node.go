package tree

import "fmt"

//Node for Tree
type Node struct {
	Value       int
	Left, Right *Node
}

//CreateNode with value
func CreateNode(Value int) *Node {
	return &Node{Value: Value}
}

//Print node value
func (node Node) Print() {
	fmt.Println(node.Value)
}

//SetValue set value for node
func (node *Node) SetValue(Value int) {
	node.Value = Value
}
