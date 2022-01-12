package BTree

type BTree struct {
	Length int
	Height int
	Root   *IndexNodeList
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
		Height: 0,
		Root:   nil,
		Length: length,
	}
}

func (b BTree) Insert(data interface{}) {

}

func (b *BTree) Search(data interface{}) {
	//n := b.Root
	//length := 1
	//for _, index := range n.Index() {
	//
	//
	//}
}

func (b *BTree) search(index int, n *IndexNodeList) *IndexNodeList {
	idxs := n.Index()
	maxIx := len(idxs) - 1
	for i, idx := range idxs {
		if idx == nil || i == maxIx {
			return b.search(index, n.nextLevel[i+1])
		}
		if *idx > index {
			return b.search(index, n.nextLevel[i])
		}
		if *idx == index {
			return n
		}

	}
	return nil
}

func (b *BTree) NewIndexNode() *IndexNodeList {
	return &IndexNodeList{
		index:     make([]*Integer, b.Length, b.Length),
		nextLevel: make([]*IndexNodeList, b.Length, b.Length),
	}
}

func (b *BTree) IsNil(n Node) bool {
	if n == nil {
		return true
	}
	return false
}
