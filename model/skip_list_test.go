package model

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSkipList(t *testing.T) {
	fmt.Printf("random start time is %s\n", time.Now().Format("2006-01-02 15:04:05"))
	//randomNumberArray := generateRandomNumber(1,10000000000,6224000)
	randomNumberArray := generateRandomNumber(1, 10000000000, 4000)
	skl := NewSkipList(32)
	starTime := time.Now()
	fmt.Printf("start time is %s\n", starTime.Format("2006-01-02 15:04:05"))
	count := 0
	randomNumberArray = []int{9014411797, 3244934821, 3244934821, 2350684784, 4731529698, 6597469165, 908390875, 1055099731}
	for _, v := range randomNumberArray {
		skl.Add(v)
		count++
		if count%300 == 0 {
			fmt.Printf("add %d number\n", count)
		}
	}
	fmt.Printf(" insert 1KW number consume time is : %fS\n", time.Now().Sub(starTime).Seconds())
	for _, v := range randomNumberArray {
		t := time.Now()
		n := skl.Find(v)
		fmt.Printf(" find data time is : %fS , result is %t \n", time.Now().Sub(t).Seconds(), v == n.Data)
	}
	fmt.Println("end")
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	m := map[int]string{}
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		tl := len(m)
		m[num] = ""
		if (len(m) - tl) != 0 {
			nums = append(nums, num)
		}
	}
	return nums
}

func TestRandomLevel(t *testing.T) {
	skl := NewSkipList(32)
	m := map[int]int{}
	ts := 300000000
	for i := 0; i < ts; i++ {
		k := skl.randomLevel()
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
