package main

type BSTIterator struct {
	stack Stack
}

type Stack []*TreeNode

func (s *Stack) len() int {
	return len(*s)
}

func (s *Stack) peek() *TreeNode {
	return (*s)[s.len()-1]
}

func (s *Stack) pop() (node *TreeNode) {
	node = s.peek()
	*s = (*s)[:s.len()-1]
	return
}

func (s *Stack) push(node *TreeNode) {
	*s = append(*s, node)
}

func Constructor(root *TreeNode) BSTIterator {
	var stack = make(Stack, 0)

	node := root
	for node != nil {
		stack.push(node)
		node = node.Left
	}

	return BSTIterator{stack}
}

/** @return the next smallest number */
func (it *BSTIterator) Next() (r int) {
	node := it.stack.pop()
	r = node.Val
	if node.Right != nil {
		node = node.Right
		for node != nil {
			it.stack.push(node)
			node = node.Left
		}
	}
	return r
}

/** @return whether we have a next smallest number */
func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

func main() {

}
