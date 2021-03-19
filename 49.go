import "sort"

// 思路一样
// 写法一
func groupAnagrams(strs []string) [][]string {
	var m = map[string][]string{}
	for _, s := range strs {
		ns := sortString(s)
		m[ns] = append(m[ns], s)
	}
	var ret [][]string
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func sortString(s string) string {
	var tmp []int
	for _, i := range s {
		tmp = append(tmp, int(i))
	}
	sort.Ints(tmp)
	var ret []byte
	for _, i := range tmp {
		ret = append(ret, byte(i))
	}
	return string(ret)
}

// 写法二
func groupAnagrams(strs []string) [][]string {
	var m, res = map[string][]string{}, [][]string{}
	for _, s := range strs {
		runes := []rune(s)
		sort.Sort(sortRunes(runes))
		m[string(runes)] = append(m[string(runes)], s)
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

