package main

import (
	"container/list"
	"fmt"
)

type SortedStack struct {
	s *list.List
}

func (ss *SortedStack) sortedInsert(x int) {
	if ss.s.Len() == 0 || x > ss.s.Back().Value.(int) {
		ss.s.PushBack(x)
		return
	}

	temp := ss.s.Back().Value.(int)
	ss.s.Remove(ss.s.Back())
	ss.sortedInsert(x)
	ss.s.PushBack(temp)
}

func (ss *SortedStack) sort() {
	if ss.s.Len() > 0 {
		x := ss.s.Back().Value.(int)
		ss.s.Remove(ss.s.Back())
		ss.sort()
		ss.sortedInsert(x)
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

	ss.sort()
	for ss.s.Len() > 0 {
		fmt.Println(ss.s.Back().Value.(int))
		ss.s.Remove(ss.s.Back())
	}
}
