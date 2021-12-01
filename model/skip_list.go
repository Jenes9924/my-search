package model

import (
	"errors"
	"math/rand"
)

//var SKIPLIST_P = 1 / 2

type SkipList struct {
	// 存储每个索引的头部节点
	Head     []*skipListNode
	level    int
	maxLevel int
	size     int
}

func NewSkipList(indexCount int) *SkipList {
	rand.Int()
	return &SkipList{maxLevel: indexCount}
}

// 跳表节点
type skipListNode struct {
	Data int
	Next []*skipListNode
}

// 索引
type index struct {
}

func (sl *SkipList) Add(data int) bool {

	//node := sl.Find(data)
	err := sl.add(data)
	if err == nil {
		sl.size++
	}
	return err == nil

}

func (sl *SkipList) randomLevel() int {
	level := 1
	for true {
		if rand.Float64() < 1/2 && level < sl.maxLevel {
			level += 1
		} else {
			break
		}
	}
	return level
}

func (sl *SkipList) add(data int) error {
	if sl.Head == nil {
		sl.Head = append([]*skipListNode{}, &skipListNode{
			Data: data,
			Next: nil,
		})
		sl.level = 1
		return nil
	}
	n := sl.get(data)
	if n.Data == data {
		return errors.New("数据已存在")
	}
	level := sl.randomLevel()

	insertNode := &skipListNode{Data: data, Next: []*skipListNode{}}
	// 全部承接前面节点
	max := 0
	if len(n.Next) > level {
		max = len(n.Next)
	} else {
		max = level
	}
	for idx := 0; idx < max; idx++ {
		insertNode.Next = append(insertNode.Next, n.Next[idx])
		n.Next[idx] = insertNode
	}
	// 超过前面节点的索引
	ix := level - 1
	for ; ix >= max; ix-- {
		node := sl.search(data, n, ix)
		if node == nil {
			// 还有一个情况，就是当head 的节点少一部分，而不是全部都少
			l := len(sl.Head)
			if l <= ix {
				for i := l - 1; i <= ix; i++ {
					sl.Head = append(sl.Head, insertNode)
				}
			}
			sl.Head[ix] = insertNode
		} else {
			insertNode.Next[ix] = node.Next[ix]
			node.Next[ix] = insertNode
		}
	}
	return nil
}

// Find /**通过搜索获取 data 是否存在
func (sl *SkipList) Find(data int) *skipListNode {
	n := sl.get(data)
	if n.Data != data {
		return nil
	}
	return n
}

func (sl SkipList) get(data int) *skipListNode {
	var (
		ix = sl.level - 1
		n  = sl.Head[ix]
	)
	for ; ix >= 0; ix-- {
		node := sl.search(data, n, ix)
		if node == nil {
			n = node
			break
		} else {
			n = node.Next[ix]
		}
		if n.Data == data {
			break
		}
	}
	return n
}

//func (sl *SkipList) searchBy(data, level int, node *skipListNode) *skipListNode {
//	if level < 1 {
//		return nil
//	}
//	n := node
//	for ; level >= 0; level-- {
//		node := sl.search(data, n, level)
//		if node == nil {
//			n = node
//			break
//		} else {
//			n = node.Next[level]
//		}
//		if n.Data == data {
//			break
//		}
//	}
//	return n
//}

/**
同一层级搜索
*/
func (sl *SkipList) search(data int, node *skipListNode, level int) *skipListNode {
	if node.Next == nil || node.Next[level] == nil {
		return node
	}
	//if len(node.Next) < level {
	//	return node
	//}
	if node.Next[level].Data < data {
		return sl.search(data, node.Next[level], level)
	} else {
		return node
	}
}
