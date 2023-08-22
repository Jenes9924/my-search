package list

import "errors"

// List 链表实现
type List[T any] struct {
	Head *listNode[T]
	Tail *listNode[T]
	size int
}

type listNode[T any] struct {
	Next     *listNode[T]
	Value    any
	Previous *listNode[T]
}

func newListNode[T any](value any) *listNode[T] {
	return &listNode[T]{Value: value}
}

func (l *List[T]) Set(index int, v any) {
	if index > l.size {
		panic("index out of length")
	}
	var n = l.Head
	for c := 1; c < index; c++ {
		if n == nil {
			panic("it's nil")
		}
		n = n.Next
	}
	node := &listNode[T]{Value: v}
	node.Next = n.Next
	n.Next = node

}
func (l *List[T]) Add(v any) {
	node := &listNode[T]{Value: v}
	if l.Head == nil {
		if l.Tail != nil {
			panic("error")
		}
		l.Head, l.Tail = node, node
		l.size++
		return
	}
	var n *listNode[T] = l.Head
	for n != nil {
		if n.Next == nil {
			n.Next = node
			l.size++
			break
		}
		n = n.Next
	}
}

func (l *List[T]) Get(index int) (v any, err error) {
	if l.size <= index {
		return nil, errors.New("数组下标越界")
	}
	var res *listNode[T] = l.Head
	for i := 1; i < index; i++ {
		res = res.Next
	}
	return res.Value, nil
}

func (l *List[T]) GetType(index int) (v any, err error) {
	if l.size <= index {
		return nil, errors.New("数组下标越界")
	}
	var res *listNode[T] = l.Head
	for i := 1; i < index; i++ {
		res = res.Next
	}
	return res.Value, nil
}

func (l *List[T]) Size() int {
	return l.size
}

// Put 替换
func (l *List[T]) Put(v any, ix int) int {
	if ix >= l.size {
		panic("index out of length")
	}
	var n = l.Head
	for c := 1; c < ix; c++ {
		if n == nil {
			panic("it's nil")
		}
		n = n.Next
	}
	node := &listNode[T]{Value: v}
	if n.Next != nil {
		node.Next = n.Next.Next
	}
	n.Next = node
	return l.size
}
