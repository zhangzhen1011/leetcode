// 分割回文串 : 求所有可能？
// 如何高效切分字符串到组合?
///////////////// 单纯回溯
func partition(s string) [][]string {
	ret = make([][]string, 0)
	dfs(s, []string{})
	return ret
}

var ret [][]string

func dfs(s string, path []string) {
	if s == "" {
		tmp := make([]string, len(path))
		copy(tmp, path)
		ret = append(ret, tmp)
		return
	}
	for i := 1; i <= len(s); i++ {
		if isHuiWen(s[:i]) {
			dfs(s[i:], append(path, s[:i]))
		}
	}
}

// 双指针判断回文
func isHuiWen(s string) bool {
	if s == "" {
		return true
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 双指针判断回文 ,   增加map, 用于剪枝
var m = map[string]bool{}

func isHuiWenV2(s string) bool {
	if s == "" {
		return true
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			m[s] = false
			return false
		}
		i++
		j--
	}
	m[s] = true
	return true
}

// tip: 求“所有可能”的需求，一般都是会有递归或者回溯， 遍历所有方案

////////////////////////// 上述方式回溯+记忆化, 回文的判断
