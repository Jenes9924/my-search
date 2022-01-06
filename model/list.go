package model

import "errors"

type List struct {
	Head *listNode
	size int
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
		l.size++
		return
	}
	var n *listNode = l.Head
	for n != nil {
		if n.Next == nil {
			n.Next = node
			l.size++
			break
		}
		n = n.Next
	}
}

func (l *List) Get(index int) (v interface{}, err error) {
	if l.size <= index {
		return nil, errors.New("数组下标越界")
	}
	var res *listNode = l.Head
	for i := 1; i < index; i++ {
		res = res.Next
	}
	return res.Value, nil
}

func (l *List) Size() int {
	return l.size
}
