package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 洗牌算法
func shuffle(nums []int) {
	for i := 0; i < len(nums); i++ {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i] // 当前数和数组中其他数随机交换
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	shuffle(a)
	fmt.Println(a)
}
