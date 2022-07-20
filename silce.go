package slice

import (
	"reflect"
	"sync"
)

type Node struct {
	parma []interface{}
	next  *Node
}

type Slice struct {
	Data  []interface{}
	First *Node
	Last  *Node
	size  int
	sync.RWMutex
}

func (e *Slice) InitQueue() *Slice {
	q := new(Slice)
	q.First = nil
	q.Last = nil
	q.size = 0
	return q
}

func (e *Slice) PushHead(parma []interface{}) {
	n := new(Node)
	n.parma = parma
	e.RLock()
	defer e.RUnlock()
	if e.First == nil {
		e.First = n
		e.Last = n
		n.next = nil
	} else {
		n.next = e.First
		e.First = n
	}
	e.size++
}

func (e *Slice) Pop() {
	e.RLock()
	defer e.RUnlock()
	f := e.First
	if f == nil {
		return
	} else if f.next == nil {
		e.First = nil
	} else {
		e.First = f.next
	}
	e.size--
}

func (e *Slice) InjectFilter(parmaNext []interface{}) interface{} {
	if e.First == nil {
		return "the first node is empty"
	} else {
		fp := e.First.parma
		e.Lock()
		defer e.Unlock()
		if reflect.TypeOf(fp) == reflect.TypeOf(parmaNext) {
			e.PushHead(parmaNext)
		}
	}
	return e
}

func (e *Slice) Top() *Node {
	return e.First
}

func (e *Slice) Length() int {
	return e.size
}
