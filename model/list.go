package model

type List struct {
	Head *listNode
	Tail *listNode
}

type listNode struct {
	Next     *listNode
	Value    interface{}
	Previous *listNode
}

func newListNode(value interface{}) *listNode {
	return &listNode{Value: value}
}

func (l *List) Add(v interface{}) {
	node := newListNode(v)
	if l.Head == nil {
		l.Head = node
	}
}
