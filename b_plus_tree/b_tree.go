package b_plus_tree

import "fmt"

type BTree struct {
	length   int          `json:"length"`
	depth    int          `json:"depth"`
	maxIndex int          `json:"max_index"`
	Root     *IndexNode   `json:"Root"`
	nilInS   []*IndexNode `json:"-"`
	nilDS    []*DataNode  `json:"-"`
}

type Node interface {
	Type() int
	Index() []*Integer
	Next() []Node
	Data() interface{}
	IsNil() bool
}

type Integer = int

func NewBPlusTree(length int) *BTree {
	if length < 3 {
		length = 3
	}
	return &BTree{
		depth:    0,
		maxIndex: length,
		Root:     nil,
		length:   length,
	}
}

// todo 插入数据需要计算高度
func (b *BTree) Insert(id int, data interface{}) {
	// 原本是向在 B+ 树中产生自增ID,但是后来发现不可行，毕竟索引树可能有多棵
	//idx := b.maxIndex + 1
	datanode := b.newDataNode(id, data)
	ixs, ix := b.get(datanode.Idx, b.Root)
	b.insert(datanode, ixs, ix)
	//b.maxIndex = idx
	//return idx
}

/*
*
node: 新增的叶子节点
*/
func (b *BTree) insert(dataNd *DataNode, ixnd *IndexNode, ix int) {
	if ixnd == nil {
		b.Root = b.NewIndexNode(nil, 1)
		//b.Root.DataNodes[0] = dataNd
		b.Root.DataNodes = append(b.Root.DataNodes, dataNd)
		b.depth++
		return
	}
	//b.get()

	// 插入 dataNd，然后重建索引

	// 底层双向链表的插入
	var tmp *DataNode

	if ix == -1 {
		tmp = ixnd.DataNodes[0]
		prev := tmp.Prev
		dataNd.Next = tmp
		tmp.Prev = dataNd
		if prev != nil {
			prev.Next = dataNd
			dataNd.Prev = prev
		}
	} else {
		tmp = ixnd.DataNodes[ix]
		nt := tmp.Next
		tmp.Next = dataNd
		dataNd.Prev = tmp
		if nt != nil {
			dataNd.Next = nt
			nt.Prev = dataNd
		}
	}
	fmt.Println("")

	// 索引节点的插入
	//length := len(ixnd.DataNodes)
	//// 按照正确的逻辑来说，此时返回的 ix 应该是小于或者等于最大长度的
	////ixnd.DataNodes[ix+1] = dataNd
	//var dataNodes []*DataNode
	//var p *DataNode = dataNd
	//for t := ix; t > -1; t-- {
	//	p = p.Prev
	//}
	//for i := 0; i < length; i++ {
	//	dataNodes = append(dataNodes, p.Next)
	//	p = p.Next
	//}
	//ixnd.DataNodes = dataNodes
	// 索引节点插入的简化操作
	// 切片底层是共享数组，有大坑
	var dns = []*DataNode{}
	for _, dataNode := range ixnd.DataNodes {
		dns = append(dns, dataNode)
	}
	dns = append(dns[:ix+1], append([]*DataNode{dataNd}, dns[ix+1:]...)...)
	ixnd.DataNodes = dns
	b.rebuildIndex(ixnd)
}

func (b *BTree) riseNode(nd *DataNode, father *IndexNode, leftIxn, rightIxn *IndexNode) {
	ix := 0
	if father == nil {
		f := b.newIndexNode([]*DataNode{nd}, []*IndexNode{leftIxn, rightIxn})
		b.Root = f
		f.Depth = 1
		leftIxn.Father, rightIxn.Father = f, f
		b.depth++
		return
	}
	// 获取父节点 DataNode 插入位置
	var tmp []*DataNode
	for _, v := range father.DataNodes {
		tmp = append(tmp, v)
	}
	for i, v := range tmp {
		if v != nil && v.Idx <= nd.Idx {
			ix = i + 1
		}
	}
	// 中间插入dataNode
	father.DataNodes = append(tmp[:ix], append([]*DataNode{nd}, tmp[ix:]...)...)

	// 下面插入 indexNode
	var nl []*IndexNode
	for _, v := range father.NextLevel {
		nl = append(nl, v)
	}
	father.NextLevel = append(nl[:ix+1], append([]*IndexNode{rightIxn}, nl[ix+1:]...)...)
	father.NextLevel[ix] = leftIxn

	leftIxn.Father, rightIxn.Father = father, father

	b.rebuildIndex(father)
}

// 无需整体构建，只需要局部重建索引，一直逐级上升修改，冒泡到 Root 节点
//
//	func (b *BTree) rebuild(ixs *IndexNode) {
//		b.rebuildIndex(ixs)
//	}
//
// 分裂步骤
func (b *BTree) rebuildIndex(ixn *IndexNode) {
	if len(ixn.DataNodes) < b.maxIndex {
		return
	}
	interceptIx := b.maxIndex / 2
	//if b.maxIndex%2 == 0 {
	//	interceptIx++
	//}

	// 先切割 indexNode,左边节点和右边节点
	var (
		dns       []*DataNode
		nextLevel []*IndexNode
	)
	for _, dataNode := range ixn.DataNodes {
		dns = append(dns, dataNode)
	}

	var leftIxn, rightIxn *IndexNode
	// 如果是最底层的数据节点，那么就不需要切割 nextlevel 下一层
	if len(ixn.NextLevel) == 0 {
		leftIxn = b.newIndexNode(dns[0:interceptIx], ixn.NextLevel)
		rightIxn = b.newIndexNode(dns[interceptIx:], ixn.NextLevel)
	} else {
		for _, ntl := range ixn.NextLevel {
			nextLevel = append(nextLevel, ntl)
		}
		leftIxn = b.newIndexNode(dns[0:interceptIx], nextLevel[0:interceptIx+1])
		rightIxn = b.newIndexNode(dns[interceptIx+1:], nextLevel[interceptIx+1:])
	}
	b.riseNode(ixn.DataNodes[interceptIx], ixn.Father, leftIxn, rightIxn)

}

func (b *BTree) Search(index int) (interface{}, bool) {
	idxn, ix := b.search(index)
	if ix == -1 {
		return nil, false
	}
	//if ix == len(idxn.DataNodes) {
	//	ix = ix - 1
	//}
	dn := idxn.DataNodes[ix]
	if dn == nil || dn.Idx != index {
		return nil, false
	}
	return dn.Data, true
}

//func (b *BTree) recursiveSearch(index, depth int, n *IndexNode) (*IndexNode, int) {
//	if n == nil {
//		return nil, -1
//	}
//	idxs := n.DataNodes
//	maxIx := b.length - 1
//	depth++
//	for i, dn := range idxs {
//		//当遍历到最后一个非空的时候，就开始到下一层
//		if dn == nil {
//			if len(n.NextLevel) != 0 || n.NextLevel[i] != nil {
//				return b.recursiveSearch(index, depth, n.NextLevel[i])
//			}
//			if depth != b.depth {
//				fmt.Printf("error!")
//			}
//			return n, i - 1
//		}
//		if i == maxIx && dn.Idx < index {
//			if n.NextLevel[i+1] != nil {
//				return b.recursiveSearch(index, depth, n.NextLevel[i+1])
//			}
//			if depth != b.depth {
//				fmt.Printf("error!")
//			}
//			return n, i
//		}
//		if dn.Idx == index {
//			if depth == b.depth {
//				return n, i
//			}
//			return b.recursiveSearch(index, depth, n.NextLevel[i+1])
//		}
//		if dn.Idx > index {
//			return b.recursiveSearch(index, depth, n.NextLevel[i])
//		}
//	}
//	return n, maxIx
//}

func (b *BTree) search(index int) (*IndexNode, int) {
	idxn, ix := b.recursiveSearch(index, b.Root)
	return idxn, ix
}

// 递归搜索
func (b *BTree) recursiveSearch(index int, n *IndexNode) (*IndexNode, int) {

	if n == nil {
		return nil, -1
	}
	dns := n.DataNodes
	if len(n.NextLevel) == 0 {
		for i, v := range dns {
			if v.Idx == index {
				return n, i
			}
		}
	} else {
		for i, v := range dns {
			if v.Idx > index {
				return b.recursiveSearch(index, n.NextLevel[i])
			}
		}
		return b.recursiveSearch(index, n.NextLevel[len(n.NextLevel)-1])
	}
	return nil, -1
}

// 返回与id最相近地树节点
// 索引节点不断向下查找
// 返回负数的时候表示不存在相同的值
func (b *BTree) get(index int, nd *IndexNode) (*IndexNode, int) {
	if nd == nil {
		return nil, -2
	}

	dataNodes := nd.DataNodes
	for i, dn := range dataNodes {
		if dn == nil || dn.Idx > index {
			// 如果该索引没有下级单元，直接返回该索引节点和要插入的节点位置
			if len(nd.NextLevel) == 0 || nd.NextLevel[0] == nil {
				return nd, i - 1
			}
			return b.get(index, nd.NextLevel[i])
		}
		//if dn.Idx == index  {
		//	if len(nd.NextLevel) == 0 || nd.NextLevel[0] == nil {
		//		return nd, i
		//	}
		//	return b.get(index, nd.NextLevel[i])
		//}
	}
	if len(nd.NextLevel) > 0 {
		return b.get(index, nd.NextLevel[len(nd.NextLevel)-1])
	} else {
		return nd, len(nd.DataNodes) - 1
	}
}

func (b *BTree) NewIndexNode(father *IndexNode, depth int) *IndexNode {
	return &IndexNode{
		//DataNodes: make([]*DataNode, b.length, b.length),
		DataNodes: []*DataNode{},
		//NextLevel: make([]*IndexNode, b.length+1, b.length+1),
		NextLevel: []*IndexNode{},
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

func (b *BTree) Delete(index int) bool {

	return true
}

// 删除需要同步更改索引，所以不能单一搜索底部的索引node,还需要把途中所有符合的索引node都包含出来
func (b *BTree) delete(index int) bool {
	ixnd, ix := b.get(index, b.Root)
	// 判断是否存在
	if ix < 0 || ixnd.DataNodes[ix].Idx != index {
		return false
	}
	// 先在双向链表中删除
	nd := ixnd.DataNodes[ix]
	if nd.Prev != nil {
		nd.Prev.Next = nd.Next
	}
	if nd.Next != nil {
		nd.Next.Prev = nd.Prev
	}

	// 修改索引

	b.deleteRaiseNode(index, ixnd)
	return false
}
func (b *BTree) deleteRaiseNode(index int, ixnd *IndexNode) {

	var dns []*DataNode
	for _, dn := range ixnd.DataNodes {
		if dn.Idx == index {
			continue
		}
		dns = append(dns, dn)
	}
	ixnd.DataNodes = dns

	if ixnd.Father == nil {
		return
	}

	if len(dns) >= b.maxIndex/2 {
		b.deleteRaiseNode(index, ixnd.Father)
		return
	}
	// 先向 father 借调
	// 统一向后借调，没有则向前借
	// pixn 和 nixn 的数量一样多的时候，优先向前借

	ft := ixnd.Father
	ftIndex := -1
	for i, v := range ft.DataNodes {
		if v.Idx > index {
			ftIndex = i
			break
		}
	}
	// 说明需要向前借，后面没有节点，father也没有可以借调的
	if ftIndex == -1 {
		// todo 向前借先不实现
		pixn := ft.NextLevel[len(ft.DataNodes)-1]
		// 判断是否需要合并
		if len(pixn.DataNodes)-1 < b.maxIndex/2 {

		}
	} else {
		nixn := ft.NextLevel[ftIndex+1]
		// 先判断是否需要合并
		if len(nixn.DataNodes)-1 < b.maxIndex/2 {
			// 直接合并，不需要先借过来再判断是否合并

		}

		// 先从father 那边借过去
		dns = append(dns, ft.DataNodes[ftIndex])
		// 在从另一边剔除

		var nixnDns []*DataNode
		// 尽量不用 slice
		for _, v := range nixn.DataNodes {
			if v.Idx == ft.DataNodes[ftIndex].Idx {
				continue
			}
			nixnDns = append(nixnDns, v)
		}
		nixn.DataNodes = nixnDns
		// 合并
		if len(nixnDns) < b.maxIndex/2 {

		}

	}

}