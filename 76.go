//
func minWindow(s string, t string) string {
	cnt := map[byte]int{}
	for _, val := range []byte(t) {
		cnt[val] += 1
	}
	var ret string
	var i, j = 0, 0
	var min int = 1e8
	var count = 0
	for ; j < len(s); j++ {
		if cnt[s[j]] > 0 {
			count++
		}
		cnt[s[j]]--
		if count == len(t) {
			for cnt[s[i]] != 0 {
				cnt[s[i]]++
				i++
			}
			if j-i+1 < min {
				ret = s[i : j+1]
				min = j - i + 1
			}
			cnt[s[i]]++
			count--
			i++
		}
	}
	return ret
}

func checkAll(cnt map[byte]int) bool {
	for _, val := range cnt {
		if val > 0 {
			return false
		}
	}
	return true
}
