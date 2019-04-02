package tree

// Tree is a recursive datastructure defined by
// a key or a value and having pointers to children
// which are also trees.
type Tree interface {
	Root() Tree
	Value() interface{}
	Children() []Tree
	Size() int
}

const (
	// INORDER is the value to get an in-order traversal of a tree.
	INORDER = 0
	// POSTORDER is the value to get a post-order traversal of a tree.
	POSTORDER = 1
	// PREORDER is the value to get a pre-order traversal of a tree.
	PREORDER = 2
	// OUTORDER is the value to get an out-of-order traversal of a tree.
	OUTORDER = 3
)

// Traverse returns all the keys of this tree according
// to a certain traversal order.
func Traverse(t Tree, mode int) []interface{} {
	c := t.Children()
	traversals := make([][]interface{}, len(c))
	for i := range c {
		traversals[i] = Traverse(c[i], mode)
	}
	res := make([]interface{}, 0, t.Size())
	switch mode % 4 {
	case INORDER:
		// INORDER is only defined for binary trees
		if len(traversals) == 2 {
			res = append(res, traversals[0]...)
			res = append(res, t.Value())
			res = append(res, traversals[1]...)
			break
		}
	case OUTORDER:
		// OUTORDER is only defined for binary trees
		if len(traversals) == 2 {
			res = append(res, traversals[1]...)
			res = append(res, t.Value())
			res = append(res, traversals[0]...)
			break
		}
	case POSTORDER:
		for i := range traversals {
			res = append(res, traversals[i]...)
		}
		res = append(res, t.Value())
		break
	case PREORDER:
		res = append(res, t.Value())
		for i := range traversals {
			res = append(res, traversals[i]...)
		}
		break
	}
	return res
}

// Leaves returns all the values of the nodes
// that don't have children.
func Leaves(t Tree) []Tree {
	if IsLeaf(t) {
		return []Tree{t}
	}
	res := make([]Tree, 0)
	c := t.Children()
	for i := range c {
		res = append(res, Leaves(c[i])...)
	}
	return res
}

// Level is the distance of this leaf to the trees root.
func Level(t Tree) int {
	return distanceToChild(t.Root(), t)
}

// Height returns the height of the heighest subtree plus one.
// If there are no subtrees the hight is 0.
func Height(t Tree) int {
	c := t.Children()
	h := make([]int, len(c))
	for i := range c {
		h[i] = Height(c[i])
	}
	return max(h) + 1
}

func max(values []int) int {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

/** distanceToChild returns the count of pointers you have to
 *  follow to get to a specific child. Therefore 0 means it's
 *  the node itself and -1 meas this node isn't in the same
 *  subtree.
 */
func distanceToChild(t, n Tree) int {
	if t == n {
		return 0
	}
	if IsLeaf(t) {
		return -1
	}
	c := t.Children()
	d := make([]int, len(c))
	for i := range c {
		d[i] = distanceToChild(c[i], n)
	}
	return max(d) + 1
}

// IsLeaf returns true if the tree doesn't
// have children.
func IsLeaf(t Tree) bool {
	return len(t.Children()) == 0
}
