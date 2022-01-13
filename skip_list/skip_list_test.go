package skip_list

import (
	"fmt"
	"my-search/util"
	"testing"
	"time"
)

func TestSkipList(t *testing.T) {
	fmt.Printf("random start time is %s\n", time.Now().Format("2006-01-02 15:04:05"))
	randomNumberArray := util.GenerateRandomNumber(1, 10000000000, 6224000)
	//randomNumberArray := generateRandomNumber(1, 10000000000, 4000)
	skl := NewSkipList(32)
	starTime := time.Now()
	fmt.Printf("start time is %s\n", starTime.Format("2006-01-02 15:04:05"))
	count := 0
	//randomNumberArray = []int{9014411797, 3244934821, 3244934821, 2350684784, 4731529698, 6597469165, 908390875, 1055099731}
	st := time.Now()
	for _, v := range randomNumberArray {
		skl.Add(v)
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
		n := skl.Find(v)
		equals := v == n.Data
		fmt.Printf(" find data time is : %fS , result is %t \n", time.Now().Sub(t).Seconds(), equals)
		if !equals {
			fmt.Printf("")
		}
	}
	fmt.Println("end")
}

func TestRandomLevel(t *testing.T) {
	skl := NewSkipList(32)
	m := map[int]int{}
	ts := 300000000
	for i := 0; i < ts; i++ {
		k := skl.randomLevel()
		skl.level = k
		if v, ok := m[k]; ok {
			m[k] = v + 1
		} else {
			m[k] = 1
		}
	}
	for k, v := range m {
		var f float64 = float64(v) / float64(ts)
		fmt.Printf(" %d The number of occurrences : %9f \n ", k, f)
	}
}

func TestSlice(t *testing.T) {
	var k, res, high = 4, 0, 16
	for i := 0; i < high; i++ {
		r := 1
		for j := 0; j < i; j++ {
			r = r * k
		}
		res = res + r
	}
	fmt.Println(5726623061 - res)
}
