package b_plus_tree

import (
	"encoding/json"
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
	k = append(k, nil)
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
	for ix, v := range randomNumberArray {
		tree.Insert(ix, v)
		count++
		//fmt.Printf("add %d number \n", count)
		//if count > 20 {
		//	Show(tree.Root,tree.depth,tree.length)
		//}
		if count%387 == 0 {
			n := time.Now()
			fmt.Printf("add %d number, spend time is : %9f S , now is %d \n", count, n.Sub(st).Seconds(), ix)
			st = n
			//bs, err := json.Marshal(tree)
			_, err := json.Marshal(tree)
			if err != nil {
				fmt.Printf("json error \n")
			} else {
				//fmt.Println(string(bs))
			}
		}
	}
	fmt.Printf(" insert 1KW number consume time is : %fS\n", time.Now().Sub(starTime).Seconds())
	for _, v := range randomNumberArray {
		t := time.Now()
		n, _ := tree.Search(v)
		equals := v == n
		fmt.Printf(" find data time is : %fS , result is %t \n", time.Now().Sub(t).Seconds(), equals)
		if !equals {
			fmt.Printf("")
		}
	}
	fmt.Println("end")
}

func WriteArray(node *IndexNode, rowIndex, columnIndex, treeDepth int, res [][]string) {
	if node == nil {
		return
	}
	res[rowIndex][columnIndex] = node.ToString()
	currLevel := (rowIndex + 1) / 2
	if currLevel == treeDepth {
		return
	}
	gap := treeDepth - currLevel - 1
	if node.NextLevel != nil {
		res[rowIndex+1][columnIndex-gap] = "/"
		for _, indexNode := range node.NextLevel {
			if indexNode != nil {
				WriteArray(indexNode, rowIndex+2, columnIndex-gap*2, treeDepth, res)
			}
		}
	}
}

func Show(node *IndexNode, depth, length int) {
	if node == nil {
		panic("error ! node is nil !")
	}
	arrayHeight := depth*length - 1
	arrayWidth := (2<<(depth-2))*3 + 1
	array := make([][]string, arrayHeight)
	for i := range array {
		array[i] = make([]string, arrayWidth)
	}
	for i := 0; i < arrayHeight; i++ {
		for j := 0; j < arrayWidth; j++ {
			array[i][j] = " "
		}
	}
	WriteArray(node, 0, arrayWidth/2, depth, array)
	for _, v := range array {
		s := ""
		for j := 0; j < len(v); j++ {
			s = s + v[j]
			if len(v[j]) > 1 && j <= len(v)-1 {
				a := len(v[j]) - 1
				if len(v[j]) > 4 {
					a = 2
				}
				j += a
			}
		}
		fmt.Println(s)
	}
}

func TestDo(t *testing.T) {
	// 简单情况下的B+树测试
	//fmt.Println("Simple B+ Tree Test:")
	//simpleBPlusTreeTest()

	// 复杂情况下的B+树测试
	fmt.Println("\nComplex B+ Tree Test:")
	complexBPlusTreeTest()
	//a := []int{1, 2, 3}
	//a = append(a[:1], append([]int{8, 9}, a[1:]...)...)
	//fmt.Printf("%v\n", a)
}

func main() {
	// 简单情况下的B+树测试
	fmt.Println("Simple B+ Tree Test:")
	simpleBPlusTreeTest()

	// 复杂情况下的B+树测试
	fmt.Println("\nComplex B+ Tree Test:")
	complexBPlusTreeTest()
}

func simpleBPlusTreeTest() {
	bpt := NewBPlusTree(3)

	// 插入键值对
	bpt.Insert(10, "Value 10")

	bpt.Insert(20, "Value 20")

	bpt.Insert(5, "Value 5")

	bpt.Insert(15, "Value 15")

	// 查找键值对
	value, found := bpt.Search(10)
	if found {
		fmt.Println("Key 10 found:", value)
	} else {
		fmt.Println("Key 10 not found")
	}

	value, found = bpt.Search(20)
	if found {
		fmt.Println("Key 20 found:", value)
	} else {
		fmt.Println("Key 20 not found")
	}

	value, found = bpt.Search(5)
	if found {
		fmt.Println("Key 5 found:", value)
	} else {
		fmt.Println("Key 5 not found")
	}

	value, found = bpt.Search(15)
	if found {
		fmt.Println("Key 15 found:", value)
	} else {
		fmt.Println("Key 15 not found")
	}
}

func complexBPlusTreeTest() {
	bpt := NewBPlusTree(3)

	// 插入键值对
	bpt.Insert(10, "Value 10")
	bpt.Insert(20, "Value 20")
	bpt.Insert(5, "Value 5")
	bpt.Insert(15, "Value 15")
	bpt.Insert(25, "Value 25")
	bpt.Insert(3, "Value 3")
	bpt.Insert(7, "Value 7")
	bpt.Insert(12, "Value 12")
	bpt.Insert(17, "Value 17")
	bpt.Insert(22, "Value 22")
	bpt.Insert(27, "Value 27")
	bpt.Insert(2, "Value 2")
	bpt.Insert(4, "Value 4")
	bpt.Insert(6, "Value 6")
	bpt.Insert(9, "Value 9")

	// 查找键值对
	value, found := bpt.Search(10)
	if found {
		fmt.Println("Key 10 found:", value)
	} else {
		fmt.Println("Key 10 not found")
	}

	value, found = bpt.Search(25)
	if found {
		fmt.Println("Key 25 found:", value)
	} else {
		fmt.Println("Key 25 not found")
	}

	value, found = bpt.Search(8)
	if found {
		fmt.Println("Key 8 found:", value)
	} else {
		fmt.Println("Key 8 not found")
	}

	// 删除键值对
	//bpt.Delete(10)

	value, found = bpt.Search(10)
	if found {
		fmt.Println("Key 10 found:", value)
	} else {
		fmt.Println("Key 10 not found")
	}
}

func TestBTree_Delete(t *testing.T) {
	bpt := NewBPlusTree(5)

	// 插入键值对
	bpt.Insert(100, "Value 10")
	bpt.Insert(200, "Value 20")
	bpt.Insert(300, "Value 5")
	bpt.Insert(400, "Value 15")
	bpt.Insert(500, "Value 25")
	bpt.Insert(600, "Value 3")
	bpt.Insert(700, "Value 7")
	bpt.Insert(800, "Value 12")
	bpt.Insert(900, "Value 17")
	bpt.Insert(550, "Value 22")
	bpt.Insert(650, "Value 27")
	bpt.Insert(680, "Value 2")
	bpt.Insert(750, "Value 4")
	bpt.Insert(820, "Value 6")
	//bpt.Insert(9, "Value 9")

	// 查找键值对
	value, found := bpt.Search(100)
	if found {
		fmt.Println("Key 10 found:", value)
	} else {
		fmt.Println("Key 10 not found")
	}

	value, found = bpt.Search(550)
	if found {
		fmt.Println("Key 25 found:", value)
	} else {
		fmt.Println("Key 25 not found")
	}

	value, found = bpt.Search(820)
	if found {
		fmt.Println("Key 8 found:", value)
	} else {
		fmt.Println("Key 8 not found")
	}

	// 删除键值对
	bpt.Delete(680)
	bpt.Delete(700)
	bpt.Delete(650)
	//bpt.Delete(700)

	value, found = bpt.Search(10)
	if found {
		fmt.Println("Key 10 found:", value)
	} else {
		fmt.Println("Key 10 not found")
	}
}
