package datastructures

// TODO should receivers be pointers or values
// TODO serialize
// TODO print
// TODO make Root private

type BST struct {
	Root *BSTNode
}

type BSTNode struct {
	Val   int
	Left  *BSTNode
	Right *BSTNode
}

func (this *BST) Insert(val int) {
	newNode := &BSTNode{val, nil, nil}
	if this.Root == nil {
		this.Root = newNode
	} else {
		insertNode(this.Root, newNode)
	}
}

func insertNode(node, newNode *BSTNode) {
	if newNode.Val < node.Val {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

func (this *BST) Delete(key int) {
	this.Root = deleteNode(this.Root, key)
}

func deleteNode(node *BSTNode, key int) *BSTNode {
	if node == nil {
		return nil
	}
	if key != node.Val {
		if key < node.Val {
			node.Left = deleteNode(node.Left, key)
		} else {
			node.Right = deleteNode(node.Right, key)
		}
		return node
	}

	// key == node.Val
	if node.Left == nil && node.Right == nil {
		return nil
	}
	if node.Left != nil && node.Right != nil {
		leftMostNodeInRightSubtree := *node.Right
		for leftMostNodeInRightSubtree.Left != nil {
			leftMostNodeInRightSubtree = *leftMostNodeInRightSubtree.Left
		}

		// copy leftMostNodeInRightSubtree value and attach current node children to it. This way we delete current node
		copy := BSTNode{leftMostNodeInRightSubtree.Val, node.Left, node.Right}

		// now we need to delete leftMostNodeInRightSubtree from the right subtree
		copy.Right = deleteNode(copy.Right, leftMostNodeInRightSubtree.Val)

		return &copy
	}

	if node.Left != nil {
		return node.Left
	}
	return node.Right
}

func (this *BST) Search(val int) *BSTNode {
	return search(this.Root, val)
}

func search(node *BSTNode, val int) *BSTNode {
	if node == nil {
		return nil
	}
	if val == node.Val {
		return node
	}
	if val < node.Val {
		return search(node.Left, val)
	}

	return search(node.Right, val)
}

// The method returns the greatest element in this set less than or equal to the given element, or nil if there is no such element.
func (this *BST) Floor(val int) *BSTNode {
	return floor(this.Root, val)
}

func floor(node *BSTNode, val int) *BSTNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	if node.Val > val {
		return floor(node.Left, val)
	}
	r := floor(node.Right, val)
	if r != nil && r.Val <= val {
		return r
	}
	return node
}

// The method returns the least element in this set greater than or equal to the given element, or nil if there is no such element.
func (this *BST) Ceil(val int) *BSTNode {
	return ceil(this.Root, val)
}

func ceil(node *BSTNode, val int) *BSTNode {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	if node.Val < val {
		return ceil(node.Right, val)
	}
	r := ceil(node.Left, val)
	if r != nil && r.Val >= val {
		return r
	}
	return node
}
