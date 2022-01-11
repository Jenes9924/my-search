package model

import (
	"errors"
	"fmt"
	"math/rand"
)

//var SKIPLIST_P = 1 / 2

type SkipList struct {
	// 存储每个索引的头部节点
	Head     *skipListNode
	level    int
	maxLevel int
	size     int
}

func NewSkipList(indexCount int) *SkipList {
	return &SkipList{maxLevel: indexCount}
}

// 跳表节点
type skipListNode struct {
	Data int
	Next []*skipListNode
	//Previous *skipListNode
}

// 索引
type index struct {
}

// Add insert
func (sl *SkipList) Add(data int) bool {

	//node := sl.Find(data)
	err := sl.add(data)
	if err == nil {
		sl.size++
		fmt.Printf("add data is %d \n", data)
	}
	return err == nil
}

func (sl *SkipList) randomLevel() int {
	level := 1
	for true {
		//a,b := rand.Intn(sl.maxLevel) & 0xFFFF,sl.maxLevel*0xFFFF
		t := rand.Float64()
		if t < 0.5 && level < sl.maxLevel {
			level += 1
		} else {
			break
		}
	}
	//t := sl.level + 1
	//if level > t{
	//	level = t
	//}
	return level
}

func (sl *SkipList) add(data int) error {
	if sl.Head == nil {
		sl.Head = &skipListNode{
			Data: data,
			Next: []*skipListNode{},
		}
		sl.level = 1
		return nil
	}
	n := sl.get(data)
	if n.Data == data {
		return errors.New("数据已存在")
	}
	//随机决定层级
	level := sl.randomLevel()
	// 构建新增node
	insertNode := &skipListNode{Data: data, Next: []*skipListNode{}}
	// 全部承接前面节点
	isFirst := false
	if n.Data > data {
		isFirst = true
		if level < len(sl.Head.Next) {
			level = len(sl.Head.Next)
		}
	}
	var (
		idx = level - 1
		//t = sl.Head[idx]
		t = sl.Head
	)
	tx := len(t.Next)
	for tidx := level; tidx > tx; tidx-- {
		t.Next = append(t.Next, nil)
	}
	sl.level = level
	if isFirst {
		for i := 0; i < level; i++ {
			insertNode.Next = append(insertNode.Next, t)
		}
		sl.Head = insertNode
	} else {
		t = sl.getByLevel(data, level)
		for ; idx >= 0; idx-- {
			if t == nil {
				fmt.Printf("")
			}
			node := sl.search(data, t, idx)
			if len(node.Next) <= idx {
				node.Next = append(node.Next, insertNode)
			} else {
				insertNode.Next = append(insertNode.Next, node.Next[idx])
				node.Next[idx] = insertNode
			}

			t = node
			//}
		}
		insertNode.Next = sl.reverse(insertNode.Next)
	}
	if sl.level < level {
		sl.level = level
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

func (sl SkipList) getByLevel(data, level int) *skipListNode {
	if level > sl.level {
		return nil
	}
	var (
		ix = sl.level - 1
		n  = sl.Head
	)
	for ; ix >= (level - 1); ix-- {
		node := sl.search(data, n, ix)
		if len(node.Next) <= ix || node.Next[ix] == nil {
			n = node
		} else if node.Next[ix].Data == data {
			n = node.Next[ix]
		} else {
			n = node
		}
		if n.Data == data {
			break
		}
	}
	return n
}
func (sl SkipList) get(data int) *skipListNode {
	return sl.getByLevel(data, 1)
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
	if node.Data == data || len(node.Next) <= level || node.Next[level] == nil {
		return node
	}
	if node.Next[level].Data < data {
		return sl.search(data, node.Next[level], level)
	} else {
		return node
	}
}

func (sl *SkipList) Remove(data int) bool {

	var (
		ix = sl.level - 1
		n  = sl.Head
	)
	for ; ix >= 0; ix-- {
		node := sl.search(data, n, ix)
		if node == nil {
			continue
		} else {
			n = node.Next[ix]
		}
		if n.Data == data {
			node.Next[ix] = n.Next[ix]
		}
		n = node
	}
	return true
}

func (sl *SkipList) reverse(ns []*skipListNode) []*skipListNode {
	if len(ns) == 1 {
		return ns
	}
	var t []*skipListNode
	for ix := len(ns) - 1; ix >= 0; ix-- {
		t = append(t, ns[ix])
	}
	ns = t
	return t
}
