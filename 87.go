// 题目给出了递归方案
// 超时
func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	for i := 1; i < len(s1); i++ {
		// 未交换
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		// 交换
		if isScramble(s1[0:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[0:len(s2)-i]) {
			return true
		}
	}
	return false
}

// 增加剪枝,  有两个用例超时;  再添加记忆化， 通过
var memo = make(map[string]bool)

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if val, ok := memo[s1+"-"+s2]; ok {
		if val {
			return true
		} else {
			return false
		}
	}
	if !countEqual(s1, s2) {
		memo[s1+"-"+s2] = false
		return false
	}
	for i := 1; i < len(s1); i++ {
		// 未交换
		if isScramble(s1[0:i], s2[0:i]) && isScramble(s1[i:], s2[i:]) {
			memo[s1+"-"+s2] = true
			return true
		}
		// 交换
		if isScramble(s1[0:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[0:len(s2)-i]) {
			memo[s1+"-"+s2] = true
			return true
		}
	}
	memo[s1+"-"+s2] = false
	return false
}

func countEqual(s1, s2 string) bool {
	tmp := [256]int{}
	for i := 0; i < len(s1); i++ {
		tmp[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		tmp[s2[i]]--
	}
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != 0 {
			return false
		}
	}
	return true
}

// --------------- dp TODO
// dp[i][j][len]: s1[i:i+len], s2[j:j+len], 是否true
func isScramble(s1 string, s2 string) bool {
}
