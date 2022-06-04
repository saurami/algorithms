/*
There are two functions (recursive and standalone) instead of one because
the recursive function calls itself with values based on conditions and
the standalone function calls the recursive function with proper starting
values.
*/

package tree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (tree *BinarySearchTree) insert(value int) *BinarySearchTree {
	if tree.root == nil {
		tree.root = &Node{data: value, left: nil, right: nil}
	} else {
		tree.root.insert(value)
	}
	return tree
}

func (node *Node) insert(value int) {
	if value < node.data {
		if node.left == nil {
			node.left = &Node{data: value, left: nil, right: nil}
		} else {
			node.left.insert(value)
		}
	} else {
		if node.right == nil {
			node.right = &Node{data: value, left: nil, right: nil}
		} else {
			node.right.insert(value)
		}
	}
}

func (tree *BinarySearchTree) search(value int) bool {
	if tree.root == nil {
		return false
	} else {
		return tree.root.search(value)
	}
}

func (node *Node) search(value int) bool {
	if value < node.data {
		if node.left == nil {
			return false
		} else {
			return node.left.search(value)
		}
	} else if value > node.data {
		if node.right == nil {
			return false
		} else {
			return node.right.search(value)
		}
	} else if value == node.data {
		return true
	}
	return false
}

func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Print(root.data, " ")
	inOrder(root.right)
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.data, " ")
	preOrder(root.left)
	preOrder(root.right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Print(root.data, " ")
}

/*func main() {
	tree := &BinarySearchTree{}
	tree.insert(10)
	tree.insert(5)
	tree.insert(15)
	tree.insert(4)
	tree.insert(16)
	tree.insert(6)
	tree.insert(13)
	tree.insert(1)
	tree.insert(12)
	tree.insert(2)
	tree.insert(20)
	inOrder(tree.root)    // 1 2 4 5 6 10 12 13 15 16 20
	fmt.Println()
	postOrder(tree.root)  // 2 1 4 6 5 12 13 20 16 15 10
	fmt.Println()
	preOrder(tree.root)   // 10 5 4 1 2 6 15 13 12 16 20
	fmt.Println()
	fmt.Println(tree.search(12))
	fmt.Println(tree.search(20))
	fmt.Println(tree.search(21))
}
*/
