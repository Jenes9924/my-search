package BTree

import (
	"fmt"
	"my-search/util"
	"testing"
	"time"
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
	k := []*int{&one, nil, &two, nil, &three}
	//k = append(k, nil)
	copy(k[6:], k[5:])
	k[2] = &three
	k1, k2 := k[0:2], k[2:]
	k3 := append(k1, &three)
	k3 = append(k3, k2...)
	k = k3
	k4 := append(append(k1, &three), k2...)
	fmt.Printf("%v,%v,%v", k, k1, k4)
}

func TestBtree(t *testing.T) {
	tree := NewBPlusTree(4)
	fmt.Printf("random start time is %s\n", time.Now().Format("2006-01-02 15:04:05"))
	randomNumberArray := util.GenerateRandomNumber(1, 10000000000, 6224000)
	starTime := time.Now()
	fmt.Printf("start time is %s\n", starTime.Format("2006-01-02 15:04:05"))
	count := 0
	st := time.Now()
	for _, v := range randomNumberArray {
		tree.Insert(v)
		count++
		if count%400000 == 0 {
			n := time.Now()
			fmt.Printf("add %d number, spend time is : %9fS \n", count, n.Sub(st).Seconds())
			st = n
		}
	}
	fmt.Printf(" insert 1KW number consume time is : %fS\n", time.Now().Sub(starTime).Seconds())
	for _, v := range randomNumberArray {
		t := time.Now()
		n := tree.Search(v)
		equals := v == n
		fmt.Printf(" find data time is : %fS , result is %t \n", time.Now().Sub(t).Seconds(), equals)
		if !equals {
			fmt.Printf("")
		}
	}
	fmt.Println("end")
}
