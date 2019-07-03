package main

type Tree struct {
	node *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (t *Tree) Insert(value int) *Tree {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.Insert(value)
	}
	return t
}

func (n *Node) Insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.Insert(value)
		}
	}
}

func printNode(n *Node) {
	if n == nil {
		return
	}
	println(n.value)
	printNode(n.left)
	printNode(n.right)
}

func (n *Node) exists(value int) bool {
	if n == nil {
		return false
	}
	if n.value == value {
		return true
	}
	println("value", n.value)
	if value <= n.value {
		return n.left.exists(value)
	} else {
		return n.right.exists(value)
	}

}
func main() {
	t := Tree{}
	t.Insert(10).
		Insert(8).
		Insert(20).
		Insert(9).
		Insert(0).
		Insert(15).
		Insert(25)

	printNode(t.node)

	println(t.node.exists(25))
	println(t.node.exists(26))
}
