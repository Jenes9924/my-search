package BTree

import (
	"fmt"
	"testing"
)

type NewInt int
type IntAlias = int

func TestSliceMake(t *testing.T) {
	k := make([]*NewInt, 10, 20)
	p := make([]*IntAlias, 10, 20)
	n := make([]*int, 10, 20)
	var i NewInt = 1
	var j IntAlias = 1
	var d = 1
	k[0] = &i
	p[0] = &j
	n[0] = &d
	fmt.Printf("%v , %v", k, p)
}

func TestDivision(t *testing.T) {
	var d = 11/2 + 1
	fmt.Printf("%v", d)
}
