package BTree

import "fmt"

type BTree struct {
	length   int
	depth    int
	maxIndex int
	root     *IndexNode
}

type Node interface {
	Type() int
	Index() []*Integer
	Next() []Node
	Data() *interface{}
	IsNil() bool
}

type Integer = int

func NewBPlusTree(length int) *BTree {
	return &BTree{
		depth:  0,
		root:   nil,
		length: length,
	}
}

// todo 插入数据需要计算高度
func (b *BTree) Insert(data *interface{}) {
	datanode := b.newDataNode(b.maxIndex+1, data)
	ixs, ix := b.search(datanode.Idx, 0, b.root)
	b.insert(datanode, ixs, ix)

}

func (b *BTree) insert(node *DataNode, ixs *IndexNode, ix int) {
	if ixs == nil {
		b.root = b.NewIndexNode(nil, 1)
		b.root.Idxs = append(b.root.Idxs, node)
		b.depth++
		return
	}
	// 插入 node，然后重建索引
	ixs.Idxs[ix+1] = node
	b.rebuildIndex(ixs)

}

func (b *BTree) riseNode(node *DataNode, father *IndexNode, t1, t2 *IndexNode) {
	ix := 0
	if father == nil {
		b.root = b.NewIndexNode(nil, 1)
		b.root.Idxs = append(b.root.Idxs, node)
		b.root.NextLevel[ix] = t1
		b.root.NextLevel[ix+1] = t2
		b.depth++
		return
	}
	// 获取父节点插入位置
	tmp := father.Idxs
	for i, v := range tmp {
		if v.Idx <= node.Idx {
			ix = i
		}
	}
	// 中间插入 node，然后重建索引
	it, it2 := tmp[0:ix+1], tmp[ix+1:len(tmp)]
	t := append(it, node)
	t = append(t, it2...)
	copy(tmp[ix+2:], tmp[ix+1:])
	tmp[ix+1] = node
	// 中间插入 分裂的 t2

	b.rebuildIndex(father)

}

// 无需整体构建，只需要局部重建索引，一直冒泡到 root 节点
//func (b *BTree) rebuild(ixs *IndexNode) {
//	b.rebuildIndex(ixs)
//}

func (b *BTree) rebuildIndex(ixs *IndexNode) {
	if ixs == nil || len(ixs.Idxs) < b.length {
		return
	}
	interceptIx := (b.length / 2) + 1
	//插入父节点的 dataNode
	dn := ixs.Idxs[interceptIx]

	// 当前 indexNode 分裂 以及 是否需要 删除
	t1 := b.newIndexNode(ixs.Idxs[0:interceptIx], ixs.NextLevel)
	t2 := b.newIndexNode(ixs.Idxs[interceptIx:b.length], ixs.NextLevel)
	if ixs.Depth != b.depth {
		t2 = b.newIndexNode(ixs.Idxs[interceptIx+1:b.length], ixs.NextLevel)
	}
	b.riseNode(dn, ixs.Father, t1, t2)
}

func (b *BTree) Search(index int) *interface{} {
	idxn, ix := b.search(index, 0, b.root)
	if ix == -1 {
		return nil
	}
	dn := idxn.Idxs[ix]
	if dn == nil || dn.Idx != index {
		return nil
	}
	return dn.Data
}

func (b *BTree) search(index, depth int, n *IndexNode) (*IndexNode, int) {
	if n == nil {
		return nil, -1
	}
	idxs := n.Idxs
	maxIx := len(idxs) - 1
	depth++
	for i, dn := range idxs {
		//当遍历到最后一个非空的时候，就开始到下一层
		if dn == nil {
			if n.NextLevel[i] != nil {
				return b.search(index, depth, n.NextLevel[i])
			}
			if depth != b.depth {
				fmt.Printf("error!")
			}
			return n, i - 1
		}
		if i == maxIx && dn.Idx < index {
			if n.NextLevel[i+1] != nil {
				return b.search(index, depth, n.NextLevel[i+1])
			}
			if depth != b.depth {
				fmt.Printf("error!")
			}
			return n, i
		}
		if dn.Idx == index {
			if depth == b.depth {
				return n, i
			}
			return b.search(index, depth, n.NextLevel[i+1])
		}
		if dn.Idx > index {
			return b.search(index, depth, n.NextLevel[i])
		}
	}
	return n, maxIx
}

func (b *BTree) NewIndexNode(father *IndexNode, depth int) *IndexNode {
	return &IndexNode{
		Idxs:      make([]*DataNode, b.length, b.length),
		NextLevel: make([]*IndexNode, b.length+1, b.length+1),
		Father:    father,
		Depth:     depth,
	}
}

func (b *BTree) IsNil(n Node) bool {
	if n == nil {
		return true
	}
	return false
}
