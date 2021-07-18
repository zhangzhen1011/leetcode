// 单词拆分
// dfs+trie
func wordBreak(s string, wordDict []string) bool {

	// 建树
	root := &trieNode{}
	for _, a := range wordDict {
		node := root
		for _, b := range a {
			if node.ch[int(b-'a')] == nil {
				node.ch[int(b-'a')] = &trieNode{}
			}
			node = node.ch[int(b-'a')]
		}
		node.isEnd = true
	}

	var memo = make([]bool, len(s))
	return dfs(root, s, 0, memo)
}

type trieNode struct {
	ch    [26]*trieNode
	isEnd bool
}

func dfs(root *trieNode, s string, cur int, memo []bool) bool { // s[cur:]
	if cur == len(s) {
		return true
	}
	if memo[cur] {
		return false
	}
	node := root
	for i := cur; i < len(s); i++ {
		// 找前缀
		if node.ch[int(s[i]-'a')] != nil {
			node = node.ch[int(s[i]-'a')]
			if node.isEnd && dfs(root, s, i+1, memo) { // i+1
				return true
			}
		} else {
			break
		}
	}
	memo[cur] = true
	return false
}

