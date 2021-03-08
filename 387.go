func firstUniqChar(s string) int {
	var m = map[rune]int{} // char rune
	for idx, ch := range s {
		if _, ok := m[ch]; ok {
			m[ch] = -1
		} else {
			m[ch] = idx
		}
	}

	var vv int = 1e10
	for _, v := range m {
		if v < 0 {
			continue
		} else {
			if v < vv {
				vv = v
			}
		}
	}
	if vv == 1e10 {
		return -1
	}
	return vv
}

