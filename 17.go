func letterCombinations(digits string) []string {
	ret = []string{}
	if len(digits) == 0 {
		return nil
	}
	helper(digits, 0, []byte{})
	return ret
}

var ret = []string{}

func helper(digits string, idx int, path []byte) {
	if idx == len(digits) {
		ret = append(ret, string(path))
		return
	}
	ch := digits[idx]
	for _, val := range digitsMap[ch] {
		helper(digits, idx+1, append(path, val))
	}
}

var digitsMap = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

