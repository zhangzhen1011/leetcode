import "strings"

// 超时
func longestSubstring(s string, k int) int {
	var count = map[uint8]int{}
	var m int
	var i, j = 0, 0
	for i = 0; i < len(s); i++ {
		count = map[uint8]int{}
		for j = i; j < len(s); j++ {
			count[s[j]] += 1
			if check(count, k) {
				m = max(j-i+1, m)
			}
		}
	}
	return m
}

func check(m map[uint8]int, k int) bool {
	for _, v := range m {
		if v == 0 {
			continue
		}
		if v < k {
			return false
		}
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 滑动窗口
func longestSubstringSlidingWin(s string, k int) int {

}

// 分治
func longestSubstring(s string, k int) int {
	if len(s) == 0 { // 递归出口1
		return 0
	}
	var count = map[byte]int{}
	for _, v := range []byte(s) {
		count[v] += 1
	}

	var split byte
	for a, v := range count {
		if v < k {
			split = a
			break
		}
	}
	if split == 0 { // 递归出口2
		return len(s)
	}

	var result int
	for _, ss := range strings.Split(s, string([]byte{split})) {
		result = max(result, longestSubstring(ss, k))
	}

	return result
}
