package main

import (
	"container/list"
	"fmt"
)

type SortedStack struct {
	s *list.List
}

func (ss *SortedStack) reverseStack() {
	if ss.s.Len() > 0 {
		x := ss.s.Back().Value.(int)
		ss.s.Remove(ss.s.Back())
		ss.reverseStack()
		ss.s.PushFront(x)
	}
}

func main() {
	ss := &SortedStack{}
	ss.s = list.New()

	ss.s.PushBack(11)
	ss.s.PushBack(2)
	ss.s.PushBack(32)
	ss.s.PushBack(3)
	ss.s.PushBack(41)

	ss.reverseStack()

	for ss.s.Len() > 0 {
		fmt.Println(ss.s.Front().Value.(int))
		ss.s.Remove(ss.s.Front())
	}
}
