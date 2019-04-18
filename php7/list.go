package php7

import (
	"github.com/z7zmey/php-parser/syntaxtree/linkedtree"
)

type stackedNodeList struct {
	list  []linkedtree.NodeID
	stack []int
}

func (s *stackedNodeList) Reset() {
	s.list = s.list[:0]
	s.stack = s.stack[:0]
}

func (s *stackedNodeList) add(n linkedtree.NodeID) {
	s.list = append(s.list, n)
}

func (s *stackedNodeList) push() {
	s.stack = append(s.stack, len(s.list))
}

func (s *stackedNodeList) pop() []linkedtree.NodeID {
	p := 0
	if len(s.stack) > 0 {
		p = s.stack[len(s.stack)-1]
	}
	list := s.list[p:]

	s.list = s.list[:p]
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}

	return list
}
