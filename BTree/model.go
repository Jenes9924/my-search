package BTree

type IndexNode struct {
	// 索引
	Idxs []*DataNode
	// 下层索引
	NextLevel []*IndexNode
	Father    *IndexNode
	Depth     int
}

func (b *BTree) newIndexNode(idxs []*DataNode, nextLevel []*IndexNode) *IndexNode {
	return &IndexNode{Idxs: idxs, NextLevel: nextLevel}
}

type DataNode struct {
	Idx  int
	Data interface{}
	Next *DataNode
}

func (b *BTree) newDataNode(idx int, data interface{}) *DataNode {
	return &DataNode{Idx: idx, Data: data}
}
