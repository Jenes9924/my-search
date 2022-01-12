package BTree

type IndexNodeList struct {
	// 索引
	index []*Integer
	// 下层索引
	nextLevel []*IndexNodeList
}

func (i *IndexNodeList) IsNil() bool {
	return i.index == nil || i.index[0] == nil
}

func (i *IndexNodeList) Data() *interface{} {
	return nil
}

func (i *IndexNodeList) Type() int {
	return 1
}

func (i *IndexNodeList) Index() []*Integer {
	return i.index
}

type DataNode struct {
	index int
	data  *interface{}
}
