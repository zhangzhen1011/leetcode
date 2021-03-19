// 最长回文子串
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	var i, j = 0, len(s) - 1
	var max = 1
	var ret = string(s[0])
	for i <= j {
		if i == j {
			i++
			j = len(s) - 1
			continue
		}
		if !check(s[i : j+1]) {
			j--
		} else {
			if j-i+1 > max {
				max = j - i + 1
				ret = s[i : j+1]
			}
			i++
			j = len(s) - 1
		}
	}
	return string(ret)
}

func check(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
