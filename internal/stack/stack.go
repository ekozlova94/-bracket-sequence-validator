package stack

type Stack struct {
	node *Node
	size int
}

type Node struct {
	Value byte
	Next  *Node
}

func (s *Stack) Push(v byte) {
	if s.node == nil {
		s.node = &Node{v, nil}
		s.size = 1
		return
	}
	temp := &Node{v, nil}
	temp.Next = s.node
	s.node = temp
	s.size++
}

func (s *Stack) Top() *Node {
	return s.node
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Pop() {
	if s.size == 0 {
		return
	}
	s.node = s.node.Next
	s.size--
	return
}
