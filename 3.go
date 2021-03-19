//TODO, 优化
func lengthOfLongestSubstring(s string) int {
	var m = map[uint8]bool{}
	var max int
	var i, j = 0, 0
	for i <= j {
		for j < len(s) {
			if _, ok := m[s[j]]; ok {
				delete(m, s[i])
				i++
			} else {
				if j-i > max {
					max = j - i + 1
				}
				m[s[j]] = true
				j++
			}
		}
		if j == len(s) {
			break
		}
	}
	return max
}
