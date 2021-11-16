package model

import (
	"errors"
	"math/rand"
)

type SkipList struct {
	// 存储每个索引的头部节点
	Head       *skipListNode
	level      int
	indexCount int
}

func NewSkipList(indexCount int) *SkipList {
	rand.Int()
	return &SkipList{indexCount: indexCount}
}

// 跳表节点
type skipListNode struct {
	Data int
	Next []*skipListNode
}

// 索引
type index struct {
}

func (sl *SkipList) Add(data int) {

}

func (sl *SkipList) add(data int) error {
	if sl.Head == nil {
		sl.Head = &skipListNode{
			Data: data,
			Next: nil,
		}
		return nil
	}
	var n *skipListNode
	for ix := 0; ix < sl.level; ix++ {
		node := sl.search(data, sl.Head.Next[ix], ix)
		if len(node.Next) < (ix + 1) {
			n = node
		}
		if node.Data == data || node.Next[ix].Data == data {
			return errors.New("已存在")
		}
		if node.Data < data && node.Next[ix].Data > data {
			//确认位置，进行插入

		}
	}
}

func (sl *SkipList) search(data int, node *skipListNode, level int) *skipListNode {
	if node == nil {
		return node
	}
	if len(node.Next) < level {
		return node
	}
	if node.Data < data {
		sl.search(data, node.Next[level], level)
	} else {
		return node
	}
}
