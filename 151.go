import "strings"

// 翻转字符串里的单词
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	var tmp = []string{}
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == "" {
			continue
		}
		tmp = append(tmp, ss[i])
	}

	return strings.Join(tmp, " ")
}

