// 前缀和, 枚举数组所有子区间的和
func numSubarraysWithSum(A []int, S int) int {
	p := make([]int, len(A)+1)
	total := 0
	for i := 0; i < len(A); i++ {
		total += A[i]
		p[i+1] = total
	}

	count := map[int]int{}
	var ret int
	for _, val := range p {
		ret += count[val]
		count[val+S] += 1
	}

	return ret
}
