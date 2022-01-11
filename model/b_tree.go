package model

type BTree struct {
	Length int
	Height int
	Root   *IndexNode
}

type IndexNode struct {
	Length    int
	Index     []int
	NextLevel []*IndexNode
}
type DataNode struct {
	Index int
	data  interface{}
}

func NewBPlusTree(length int) *BTree {
	return &BTree{
		Height: 0,
		Root:   nil,
		Length: length,
	}
}
