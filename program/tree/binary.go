package tree

// BinaryTree is a Node of a binary tree.
type BinaryTree struct {
	left, right *BinaryTree
	value       int
	root        *BinaryTree
	count       int
}

// NewBinary creates a new binary tree object with a given
// root and a given value
func NewBinary(root *BinaryTree, value int) *BinaryTree {
	return &BinaryTree{root: root, value: value, count: 1}
}

// Root returns the root of this BinaryTree.
func (t *BinaryTree) Root() *BinaryTree {
	return t
}

// Value returns the Value of this node.
func (t *BinaryTree) Value() int {
	return t.value
}

// Children returns all the existing children of this node.
func (t *BinaryTree) Children() []*BinaryTree {
	children := make([]*BinaryTree, 0, 2)
	if t.left != nil {
		children = append(children, t.left)
	}
	if t.right != nil {
		children = append(children, t.right)
	}
	return children
}

// Size returns the count of Nodes in this tree.
func (t *BinaryTree) Size() int {
	return t.count
}
