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
	//s := []int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}
	//k := s[0:3]
	//k1 := s[3:5]
	//k2:= s[5:9]
	var one, two, three, _ = 1, 2, 3, 4
	k := []*int{&one, nil, &two, nil, nil}
	//k = append(k, nil)
	copy(k[3:], k[2:])
	k[2] = &three
	k1, k2 := k[0:2], k[2:]
	k3 := append(k1, &three)
	k3 = append(k3, k2...)
	k = k3
	k4 := append(append(k1, &three), k2...)
	fmt.Printf("%v,%v,%v", k, k1, k4)
}
