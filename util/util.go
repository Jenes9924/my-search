package util

import (
	"math/rand"
	"time"
)

//生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
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
