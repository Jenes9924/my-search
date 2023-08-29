package b_plus_tree

import "encoding/json"

type IndexNode struct {
	// 索引
	DataNodes []*DataNode `json:"idxs"`
	Values    []string    `json:"values"`
	// 下层索引
	NextLevel []*IndexNode `json:"next_level"`
	Father    *IndexNode   `json:"-"`
	Depth     int          `json:"-"`
}

func (ixn *IndexNode) MarshalJSON() ([]byte, error) {
	var last = struct {
		Index []int        `json:"ix"`
		Next  []*IndexNode `json:"next,omitempty"`
	}{}
	var res []int
	for i := range ixn.DataNodes {
		if ixn.DataNodes[i] != nil {
			res = append(res, ixn.DataNodes[i].Idx)
		}
	}
	last.Index = res
	//if ixn.NextLevel[0] != nil {
	//	last.Next = ixn.NextLevel
	//}
	for i := range ixn.NextLevel {
		if ixn.NextLevel[i] == nil {
			break
		}
		last.Next = append(last.Next, ixn.NextLevel[i])
	}
	return json.Marshal(last)
}

func (ixn *IndexNode) ToString() string {
	var res []*int
	for _, idx := range ixn.DataNodes {
		if idx != nil {
			res = append(res, &idx.Idx)
		} else {
			res = append(res, nil)
		}
	}
	bs, _ := json.Marshal(res)
	return string(bs)
}

func (b *BPlusTree) newIndexNode(idxs []*DataNode, nextLevel []*IndexNode) *IndexNode {
	t := &IndexNode{DataNodes: idxs, NextLevel: nextLevel}
	if nextLevel != nil || nextLevel[0] != nil {
		for _, node := range nextLevel {
			if node != nil {
				node.Father = t
			}
		}
	}
	return t
}

type DataNode struct {
	// 数据 id
	Idx  int         `json:"idx"`
	Data interface{} `json:"-"`
	Next *DataNode   `json:"-"`
	Prev *DataNode   `json:"-"`
}

func (b *BPlusTree) newDataNode(idx int, data interface{}) *DataNode {
	return &DataNode{Idx: idx, Data: data}
}
