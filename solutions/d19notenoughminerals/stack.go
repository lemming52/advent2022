package d19notenoughminerals

import (
	"container/list"
	"log"
)

type BuildStack struct {
	stack *list.List
}

func newBuildStack() *BuildStack {
	return &BuildStack{stack: list.New()}
}

func (s *BuildStack) Push(b *BuildOrder) {
	s.stack.PushFront(b)
}

func (s *BuildStack) End(b *BuildOrder) {
	s.stack.PushBack(b)
}

func (s *BuildStack) Pop() interface{} {
	if s.stack.Len() > 0 {
		elem := s.stack.Front()
		return s.stack.Remove(elem)
	}
	log.Fatal("stack error")
	return nil
}

func (b *BuildStack) Size() int {
	return b.stack.Len()
}

func (b *BuildStack) Empty() bool {
	return b.stack.Len() == 0
}
