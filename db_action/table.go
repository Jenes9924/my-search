package db_action

import "my-search/b_plus_tree"

type Table struct {
	name string
	tree *b_plus_tree.BPlusTree
}

func NewTable(name string) *Table {
	return &Table{name: name, tree: b_plus_tree.NewBPlusTree(3)}
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) SetName(name string) {
	t.name = name
}

func (t *Table) Tree() *b_plus_tree.BPlusTree {
	return t.tree
}

func (t *Table) SetTree(tree *b_plus_tree.BPlusTree) {
	t.tree = tree
}

func (t *Table) Insert(data interface{}) {
	//t.tree.Insert(data)
}

func (t *Table) Select(condition string) interface{} {

	return nil
}
